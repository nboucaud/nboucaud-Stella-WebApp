// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package app

import (
	"testing"

	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/stretchr/testify/require"
)

func TestGetPostAccessibleBounds(t *testing.T) {
	var p = func(at int64) *model.Post {
		return &model.Post{CreateAt: at}
	}

	t.Run("nil returns all accessible posts", func(t *testing.T) {
		bounds := getPostAccessibleBounds(nil, 0)
		require.True(t, bounds.allAccessible())
	})

	t.Run("nil posts returns all accessible posts", func(t *testing.T) {
		bounds := getPostAccessibleBounds(&model.PostList{
			Posts: nil,
		}, 0)
		require.True(t, bounds.allAccessible())
	})

	t.Run("empty posts returns all accessible posts", func(t *testing.T) {
		bounds := getPostAccessibleBounds(&model.PostList{
			Posts: map[string]*model.Post{},
		}, 0)
		require.True(t, bounds.allAccessible())
	})

	t.Run("one accessible post returns all accessible posts", func(t *testing.T) {
		bounds := getPostAccessibleBounds(&model.PostList{
			Posts: map[string]*model.Post{
				"post_a": p(1),
			},
			Order: []string{"post_a"},
		}, 0)
		require.True(t, bounds.allAccessible())
	})

	t.Run("one inaccessible post returns no accessible posts", func(t *testing.T) {
		bounds := getPostAccessibleBounds(&model.PostList{
			Posts: map[string]*model.Post{
				"post_a": p(0),
			},
			Order: []string{"post_a"},
		}, 1)
		require.True(t, bounds.allInaccessible())
	})

	t.Run("all accessible posts returns all accessible posts", func(t *testing.T) {
		bounds := getPostAccessibleBounds(&model.PostList{
			Posts: map[string]*model.Post{
				"post_a": p(1),
				"post_b": p(2),
				"post_c": p(3),
				"post_d": p(4),
				"post_e": p(5),
				"post_f": p(6),
			},
			Order: []string{"post_a", "post_b", "post_c", "post_d", "post_e", "post_f"},
		}, 0)
		require.True(t, bounds.allAccessible())
	})

	t.Run("all inaccessible posts returns all inaccessible posts", func(t *testing.T) {
		bounds := getPostAccessibleBounds(&model.PostList{
			Posts: map[string]*model.Post{
				"post_a": p(1),
				"post_b": p(2),
				"post_c": p(3),
				"post_d": p(4),
				"post_e": p(5),
				"post_f": p(6),
			},
			Order: []string{"post_a", "post_b", "post_c", "post_d", "post_e", "post_f"},
		}, 7)
		require.True(t, bounds.allInaccessible())
	})

	t.Run("two posts, first accessible", func(t *testing.T) {
		bounds := getPostAccessibleBounds(&model.PostList{
			Posts: map[string]*model.Post{
				"post_a": p(1),
				"post_b": p(0),
			},
			Order: []string{"post_a", "post_b"},
		}, 1)
		require.Equal(t, postAccessibleBounds{accessible: 0, inaccessible: 1}, bounds)
	})

	t.Run("two posts, second accessible", func(t *testing.T) {
		bounds := getPostAccessibleBounds(&model.PostList{
			Posts: map[string]*model.Post{
				"post_a": p(0),
				"post_b": p(1),
			},
			Order: []string{"post_a", "post_b"},
		}, 1)
		require.Equal(t, postAccessibleBounds{accessible: 1, inaccessible: 0}, bounds)
	})

	t.Run("picks the right post for boundaries when there are time ties", func(t *testing.T) {
		bounds := getPostAccessibleBounds(&model.PostList{
			Posts: map[string]*model.Post{
				"post_a": p(0),
				"post_b": p(1),
				"post_c": p(1),
				"post_d": p(2),
			},
			Order: []string{"post_a", "post_b", "post_c", "post_d"},
		}, 2)
		require.Equal(t, postAccessibleBounds{accessible: 3, inaccessible: 2}, bounds)
	})

	t.Run("picks the right post for boundaries when there are time ties, reverse order", func(t *testing.T) {
		bounds := getPostAccessibleBounds(&model.PostList{
			Posts: map[string]*model.Post{
				"post_a": p(0),
				"post_b": p(1),
				"post_c": p(1),
				"post_d": p(2),
			},
			Order: []string{"post_d", "post_c", "post_b", "post_a"},
		}, 2)
		require.Equal(t, postAccessibleBounds{accessible: 0, inaccessible: 1}, bounds)
	})

	t.Run("odd number of posts and reverse time selects right boundaries", func(t *testing.T) {
		bounds := getPostAccessibleBounds(&model.PostList{
			Posts: map[string]*model.Post{
				"post_a": p(0),
				"post_b": p(1),
				"post_c": p(2),
				"post_d": p(3),
				"post_e": p(4),
			},
			Order: []string{"post_e", "post_d", "post_c", "post_b", "post_a"},
		}, 2)
		require.Equal(t, postAccessibleBounds{accessible: 2, inaccessible: 3}, bounds)
	})
}
