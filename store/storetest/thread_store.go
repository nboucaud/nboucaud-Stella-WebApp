// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package storetest

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/store"
)

func TestThreadStore(t *testing.T, ss store.Store, s SqlStore) {
	t.Run("CRTMigrationFixQuery", func(t *testing.T) { testCRTMigrationFixQuery(t, ss, s) })
	t.Run("ThreadSQLOperations", func(t *testing.T) { testThreadSQLOperations(t, ss, s) })
	t.Run("ThreadStorePopulation", func(t *testing.T) { testThreadStorePopulation(t, ss) })
}

func testThreadStorePopulation(t *testing.T, ss store.Store) {
	makeSomePosts := func() []*model.Post {

		u1 := model.User{
			Email:    MakeEmail(),
			Username: model.NewId(),
		}

		u, err := ss.User().Save(&u1)
		require.NoError(t, err)

		c, err2 := ss.Channel().Save(&model.Channel{
			DisplayName: model.NewId(),
			Type:        model.CHANNEL_OPEN,
			Name:        model.NewId(),
		}, 999)
		require.NoError(t, err2)

		_, err44 := ss.Channel().SaveMember(&model.ChannelMember{
			ChannelId:   c.Id,
			UserId:      u1.Id,
			NotifyProps: model.GetDefaultChannelNotifyProps(),
			MsgCount:    0,
		})
		require.NoError(t, err44)
		o := model.Post{}
		o.ChannelId = c.Id
		o.UserId = u.Id
		o.Message = "zz" + model.NewId() + "b"

		otmp, err3 := ss.Post().Save(&o)
		require.NoError(t, err3)
		o2 := model.Post{}
		o2.ChannelId = c.Id
		o2.UserId = model.NewId()
		o2.RootId = otmp.Id
		o2.Message = "zz" + model.NewId() + "b"

		o3 := model.Post{}
		o3.ChannelId = c.Id
		o3.UserId = u.Id
		o3.RootId = otmp.Id
		o3.Message = "zz" + model.NewId() + "b"

		o4 := model.Post{}
		o4.ChannelId = c.Id
		o4.UserId = model.NewId()
		o4.Message = "zz" + model.NewId() + "b"

		newPosts, errIdx, err3 := ss.Post().SaveMultiple([]*model.Post{&o2, &o3, &o4})

		olist, _ := ss.Post().Get(context.Background(), otmp.Id, true, false, false, "")
		o1 := olist.Posts[olist.Order[0]]

		newPosts = append([]*model.Post{o1}, newPosts...)
		require.NoError(t, err3, "couldn't save item")
		require.Equal(t, -1, errIdx)
		require.Len(t, newPosts, 4)
		require.Equal(t, int64(2), newPosts[0].ReplyCount)
		require.Equal(t, int64(2), newPosts[1].ReplyCount)
		require.Equal(t, int64(2), newPosts[2].ReplyCount)
		require.Equal(t, int64(0), newPosts[3].ReplyCount)

		return newPosts
	}
	t.Run("Save replies creates a thread", func(t *testing.T) {
		newPosts := makeSomePosts()
		thread, err := ss.Thread().Get(newPosts[0].Id)
		require.NoError(t, err, "couldn't get thread")
		require.NotNil(t, thread)
		require.Equal(t, int64(2), thread.ReplyCount)
		require.ElementsMatch(t, model.StringArray{newPosts[0].UserId, newPosts[1].UserId}, thread.Participants)

		o5 := model.Post{}
		o5.ChannelId = model.NewId()
		o5.UserId = model.NewId()
		o5.RootId = newPosts[0].Id
		o5.Message = "zz" + model.NewId() + "b"

		_, _, err = ss.Post().SaveMultiple([]*model.Post{&o5})
		require.NoError(t, err, "couldn't save item")

		thread, err = ss.Thread().Get(newPosts[0].Id)
		require.NoError(t, err, "couldn't get thread")
		require.NotNil(t, thread)
		require.Equal(t, int64(3), thread.ReplyCount)
		require.ElementsMatch(t, model.StringArray{newPosts[0].UserId, newPosts[1].UserId, o5.UserId}, thread.Participants)
	})

	t.Run("Delete a reply updates count on a thread", func(t *testing.T) {
		newPosts := makeSomePosts()
		thread, err := ss.Thread().Get(newPosts[0].Id)
		require.NoError(t, err, "couldn't get thread")
		require.NotNil(t, thread)
		require.Equal(t, int64(2), thread.ReplyCount)
		require.ElementsMatch(t, model.StringArray{newPosts[0].UserId, newPosts[1].UserId}, thread.Participants)

		err = ss.Post().Delete(newPosts[1].Id, 1234, model.NewId())
		require.NoError(t, err, "couldn't delete post")

		thread, err = ss.Thread().Get(newPosts[0].Id)
		require.NoError(t, err, "couldn't get thread")
		require.NotNil(t, thread)
		require.Equal(t, int64(1), thread.ReplyCount)
		require.ElementsMatch(t, model.StringArray{newPosts[0].UserId, newPosts[1].UserId}, thread.Participants)
	})

	t.Run("Update reply should update the UpdateAt of the thread", func(t *testing.T) {
		rootPost := model.Post{}
		rootPost.RootId = model.NewId()
		rootPost.ChannelId = model.NewId()
		rootPost.UserId = model.NewId()
		rootPost.Message = "zz" + model.NewId() + "b"

		replyPost := model.Post{}
		replyPost.ChannelId = rootPost.ChannelId
		replyPost.UserId = model.NewId()
		replyPost.Message = "zz" + model.NewId() + "b"
		replyPost.RootId = rootPost.RootId

		newPosts, _, err := ss.Post().SaveMultiple([]*model.Post{&rootPost, &replyPost})
		require.NoError(t, err)

		thread1, err := ss.Thread().Get(newPosts[0].RootId)
		require.NoError(t, err)

		rrootPost, err := ss.Post().GetSingle(rootPost.Id, false)
		require.NoError(t, err)
		require.Equal(t, rrootPost.UpdateAt, rootPost.UpdateAt)

		replyPost2 := model.Post{}
		replyPost2.ChannelId = rootPost.ChannelId
		replyPost2.UserId = model.NewId()
		replyPost2.Message = "zz" + model.NewId() + "b"
		replyPost2.RootId = rootPost.Id

		replyPost3 := model.Post{}
		replyPost3.ChannelId = rootPost.ChannelId
		replyPost3.UserId = model.NewId()
		replyPost3.Message = "zz" + model.NewId() + "b"
		replyPost3.RootId = rootPost.Id

		_, _, err = ss.Post().SaveMultiple([]*model.Post{&replyPost2, &replyPost3})
		require.NoError(t, err)

		rrootPost2, err := ss.Post().GetSingle(rootPost.Id, false)
		require.NoError(t, err)
		require.Greater(t, rrootPost2.UpdateAt, rrootPost.UpdateAt)

		thread2, err := ss.Thread().Get(rootPost.Id)
		require.NoError(t, err)
		require.Greater(t, thread2.LastReplyAt, thread1.LastReplyAt)
	})

	t.Run("Deleting reply should update the thread", func(t *testing.T) {
		rootPost := model.Post{}
		rootPost.RootId = model.NewId()
		rootPost.ChannelId = model.NewId()
		rootPost.UserId = model.NewId()
		rootPost.Message = "zz" + model.NewId() + "b"

		replyPost := model.Post{}
		replyPost.ChannelId = rootPost.ChannelId
		replyPost.UserId = model.NewId()
		replyPost.Message = "zz" + model.NewId() + "b"
		replyPost.RootId = rootPost.RootId

		newPosts, _, err := ss.Post().SaveMultiple([]*model.Post{&rootPost, &replyPost})
		require.NoError(t, err)

		thread1, err := ss.Thread().Get(newPosts[0].RootId)
		require.NoError(t, err)
		require.EqualValues(t, thread1.ReplyCount, 2)
		require.Len(t, thread1.Participants, 2)

		err = ss.Post().Delete(replyPost.Id, 123, model.NewId())
		require.NoError(t, err)

		thread2, err := ss.Thread().Get(rootPost.RootId)
		require.NoError(t, err)
		require.EqualValues(t, thread2.ReplyCount, 1)
		require.Len(t, thread2.Participants, 2)
	})

	t.Run("Deleting root post should delete the thread", func(t *testing.T) {
		rootPost := model.Post{}
		rootPost.ChannelId = model.NewId()
		rootPost.UserId = model.NewId()
		rootPost.Message = "zz" + model.NewId() + "b"

		newPosts1, _, err := ss.Post().SaveMultiple([]*model.Post{&rootPost})
		require.NoError(t, err)

		replyPost := model.Post{}
		replyPost.ChannelId = rootPost.ChannelId
		replyPost.UserId = model.NewId()
		replyPost.Message = "zz" + model.NewId() + "b"
		replyPost.RootId = newPosts1[0].Id

		_, _, err = ss.Post().SaveMultiple([]*model.Post{&replyPost})
		require.NoError(t, err)

		thread1, err := ss.Thread().Get(newPosts1[0].Id)
		require.NoError(t, err)
		require.EqualValues(t, thread1.ReplyCount, 1)
		require.Len(t, thread1.Participants, 2)

		err = ss.Post().PermanentDeleteByUser(rootPost.UserId)
		require.NoError(t, err)

		thread2, _ := ss.Thread().Get(rootPost.Id)
		require.Nil(t, thread2)
	})

	t.Run("Thread last updated is changed when channel is updated after UpdateLastViewedAtPost", func(t *testing.T) {
		newPosts := makeSomePosts()
		opts := store.ThreadMembershipOpts{
			Following:             true,
			IncrementMentions:     false,
			UpdateFollowing:       true,
			UpdateViewedTimestamp: false,
			UpdateParticipants:    false,
		}
		_, e := ss.Thread().MaintainMembership(newPosts[0].UserId, newPosts[0].Id, opts)
		require.NoError(t, e)
		m, err1 := ss.Thread().GetMembershipForUser(newPosts[0].UserId, newPosts[0].Id)
		require.NoError(t, err1)
		m.LastUpdated -= 1000
		_, err := ss.Thread().UpdateMembership(m)
		require.NoError(t, err)

		_, err = ss.Channel().UpdateLastViewedAtPost(newPosts[0], newPosts[0].UserId, 0, 0, true, true)
		require.NoError(t, err)

		assert.Eventually(t, func() bool {
			m2, err2 := ss.Thread().GetMembershipForUser(newPosts[0].UserId, newPosts[0].Id)
			require.NoError(t, err2)
			return m2.LastUpdated > m.LastUpdated
		}, time.Second, 10*time.Millisecond)
	})

	t.Run("Thread last updated is changed when channel is updated after IncrementMentionCount", func(t *testing.T) {
		newPosts := makeSomePosts()

		opts := store.ThreadMembershipOpts{
			Following:             true,
			IncrementMentions:     false,
			UpdateFollowing:       true,
			UpdateViewedTimestamp: false,
			UpdateParticipants:    false,
		}
		_, e := ss.Thread().MaintainMembership(newPosts[0].UserId, newPosts[0].Id, opts)
		require.NoError(t, e)
		m, err1 := ss.Thread().GetMembershipForUser(newPosts[0].UserId, newPosts[0].Id)
		require.NoError(t, err1)
		m.LastUpdated -= 1000
		_, err := ss.Thread().UpdateMembership(m)
		require.NoError(t, err)

		err = ss.Channel().IncrementMentionCount(newPosts[0].ChannelId, newPosts[0].UserId, true, false)
		require.NoError(t, err)

		assert.Eventually(t, func() bool {
			m2, err2 := ss.Thread().GetMembershipForUser(newPosts[0].UserId, newPosts[0].Id)
			require.NoError(t, err2)
			return m2.LastUpdated > m.LastUpdated
		}, time.Second, 10*time.Millisecond)
	})

	t.Run("Thread last updated is changed when channel is updated after UpdateLastViewedAt", func(t *testing.T) {
		newPosts := makeSomePosts()
		opts := store.ThreadMembershipOpts{
			Following:             true,
			IncrementMentions:     false,
			UpdateFollowing:       true,
			UpdateViewedTimestamp: false,
			UpdateParticipants:    false,
		}
		_, e := ss.Thread().MaintainMembership(newPosts[0].UserId, newPosts[0].Id, opts)
		require.NoError(t, e)
		m, err1 := ss.Thread().GetMembershipForUser(newPosts[0].UserId, newPosts[0].Id)
		require.NoError(t, err1)
		m.LastUpdated -= 1000
		_, err := ss.Thread().UpdateMembership(m)
		require.NoError(t, err)

		_, err = ss.Channel().UpdateLastViewedAt([]string{newPosts[0].ChannelId}, newPosts[0].UserId, true)
		require.NoError(t, err)

		assert.Eventually(t, func() bool {
			m2, err2 := ss.Thread().GetMembershipForUser(newPosts[0].UserId, newPosts[0].Id)
			require.NoError(t, err2)
			return m2.LastUpdated > m.LastUpdated
		}, time.Second, 10*time.Millisecond)
	})

	t.Run("Thread membership 'viewed' timestamp is updated properly", func(t *testing.T) {
		newPosts := makeSomePosts()

		opts := store.ThreadMembershipOpts{
			Following:             true,
			IncrementMentions:     false,
			UpdateFollowing:       true,
			UpdateViewedTimestamp: false,
			UpdateParticipants:    false,
		}
		tm, e := ss.Thread().MaintainMembership(newPosts[0].UserId, newPosts[0].Id, opts)
		require.NoError(t, e)
		require.Equal(t, int64(0), tm.LastViewed)

		opts.UpdateViewedTimestamp = true
		_, e = ss.Thread().MaintainMembership(newPosts[0].UserId, newPosts[0].Id, opts)
		require.NoError(t, e)
		m2, err2 := ss.Thread().GetMembershipForUser(newPosts[0].UserId, newPosts[0].Id)
		require.NoError(t, err2)
		require.Greater(t, m2.LastViewed, int64(0))
	})

	t.Run("Thread membership 'viewed' timestamp is updated properly for new membership", func(t *testing.T) {
		newPosts := makeSomePosts()

		opts := store.ThreadMembershipOpts{
			Following:             true,
			IncrementMentions:     false,
			UpdateFollowing:       false,
			UpdateViewedTimestamp: true,
			UpdateParticipants:    false,
		}
		tm, e := ss.Thread().MaintainMembership(newPosts[0].UserId, newPosts[0].Id, opts)
		require.NoError(t, e)
		require.NotEqual(t, int64(0), tm.LastViewed)
	})

	t.Run("Thread last updated is changed when channel is updated after UpdateLastViewedAtPost for mark unread", func(t *testing.T) {
		newPosts := makeSomePosts()
		opts := store.ThreadMembershipOpts{
			Following:             true,
			IncrementMentions:     false,
			UpdateFollowing:       true,
			UpdateViewedTimestamp: false,
			UpdateParticipants:    false,
		}
		_, e := ss.Thread().MaintainMembership(newPosts[0].UserId, newPosts[0].Id, opts)
		require.NoError(t, e)
		m, err1 := ss.Thread().GetMembershipForUser(newPosts[0].UserId, newPosts[0].Id)
		require.NoError(t, err1)
		m.LastUpdated += 1000
		_, err := ss.Thread().UpdateMembership(m)
		require.NoError(t, err)

		_, err = ss.Channel().UpdateLastViewedAtPost(newPosts[0], newPosts[0].UserId, 0, 0, true, true)
		require.NoError(t, err)

		assert.Eventually(t, func() bool {
			m2, err2 := ss.Thread().GetMembershipForUser(newPosts[0].UserId, newPosts[0].Id)
			require.NoError(t, err2)
			return m2.LastUpdated < m.LastUpdated
		}, time.Second, 10*time.Millisecond)
	})

	t.Run("Updating post does not make thread unread", func(t *testing.T) {
		newPosts := makeSomePosts()
		opts := store.ThreadMembershipOpts{
			Following:             true,
			IncrementMentions:     false,
			UpdateFollowing:       true,
			UpdateViewedTimestamp: false,
			UpdateParticipants:    false,
		}
		m, err := ss.Thread().MaintainMembership(newPosts[0].UserId, newPosts[0].Id, opts)
		require.NoError(t, err)
		th, err := ss.Thread().GetThreadForUser("", m, false)
		require.NoError(t, err)
		require.Equal(t, int64(2), th.UnreadReplies)

		m.LastViewed = newPosts[2].UpdateAt + 1
		_, err = ss.Thread().UpdateMembership(m)
		require.NoError(t, err)
		th, err = ss.Thread().GetThreadForUser("", m, false)
		require.NoError(t, err)
		require.Equal(t, int64(0), th.UnreadReplies)

		editedPost := newPosts[2].Clone()
		editedPost.Message = "This is an edited post"
		_, err = ss.Post().Update(editedPost, newPosts[2])
		require.NoError(t, err)

		th, err = ss.Thread().GetThreadForUser("", m, false)
		require.NoError(t, err)
		require.Equal(t, int64(0), th.UnreadReplies)
	})
}

func testThreadSQLOperations(t *testing.T, ss store.Store, s SqlStore) {
	t.Run("Save", func(t *testing.T) {
		threadToSave := &model.Thread{
			PostId:       model.NewId(),
			ChannelId:    model.NewId(),
			LastReplyAt:  10,
			ReplyCount:   5,
			Participants: model.StringArray{model.NewId(), model.NewId()},
		}
		_, err := ss.Thread().Save(threadToSave)
		require.NoError(t, err)

		th, err := ss.Thread().Get(threadToSave.PostId)
		require.NoError(t, err)
		require.Equal(t, threadToSave, th)
	})
}

func testCRTMigrationFixQuery(t *testing.T, ss store.Store, s SqlStore) {
	t.Run("Get all threads newer than channel LastViewedAt", func(t *testing.T) {
		teamId := model.NewId()
		uId1 := model.NewId()
		uId2 := model.NewId()
		uId3 := model.NewId()

		c1, err := ss.Channel().Save(&model.Channel{
			DisplayName: model.NewId(),
			Type:        model.CHANNEL_OPEN,
			Name:        model.NewId(),
			TeamId:      teamId,
			CreatorId:   uId1,
		}, 999)
		require.NoError(t, err)

		lastPostAt := int64(100)
		// user1 has never seen the channel
		_, err = ss.Channel().SaveMember(&model.ChannelMember{
			ChannelId:    c1.Id,
			UserId:       uId1,
			LastViewedAt: 0,
			NotifyProps:  model.GetDefaultChannelNotifyProps(),
		})
		require.NoError(t, err)

		// user2 has fully read the channel
		_, err = ss.Channel().SaveMember(&model.ChannelMember{
			ChannelId:    c1.Id,
			UserId:       uId2,
			LastViewedAt: lastPostAt,
			NotifyProps:  model.GetDefaultChannelNotifyProps(),
		})
		require.NoError(t, err)

		// user3 has read channel before latest post in a thread
		// in the channel
		cm3, err := ss.Channel().SaveMember(&model.ChannelMember{
			ChannelId:    c1.Id,
			UserId:       uId3,
			LastViewedAt: lastPostAt - 50,
			NotifyProps:  model.GetDefaultChannelNotifyProps(),
		})
		require.NoError(t, err)

		rootPostId := model.NewId()
		_, err = ss.Thread().Save(&model.Thread{
			PostId:      rootPostId,
			ChannelId:   c1.Id,
			LastReplyAt: lastPostAt,
		})
		require.NoError(t, err)

		// Create ThreadMembership
		_, err = ss.Thread().SaveMembership(&model.ThreadMembership{
			PostId: rootPostId,
			UserId: uId1,
		})
		require.NoError(t, err)
		_, err = ss.Thread().SaveMembership(&model.ThreadMembership{
			PostId: rootPostId,
			UserId: uId2,
		})
		require.NoError(t, err)
		_, err = ss.Thread().SaveMembership(&model.ThreadMembership{
			PostId: rootPostId,
			UserId: uId3,
		})
		require.NoError(t, err)

		newerThreadsInfo, err := ss.Thread().GetAllThreadsNewerThanChannelLastViewedAt()
		require.NoError(t, err)
		require.Len(t, newerThreadsInfo, 1)
		require.Equal(t, rootPostId, newerThreadsInfo[0].ThreadId)
		require.Equal(t, uId3, newerThreadsInfo[0].UserId)
		require.Equal(t, teamId, newerThreadsInfo[0].TeamId)
		require.Equal(t, cm3.LastViewedAt, newerThreadsInfo[0].LastViewedAt)
	})
}
