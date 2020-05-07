// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package cache2

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNewCache(t *testing.T) {
	p := NewProvider()

	size := 1
	c := p.NewCache(size)

	err := c.Set("key1", "val1")
	require.Nil(t, err)
	err = c.Set("key2", "val2")
	require.Nil(t, err)
	l, err := c.Len()
	require.Nil(t, err)
	require.Equal(t, size, l)
}

func TestNewCacheWithParams(t *testing.T) {
	p := NewProvider()

	size := 1
	expiry := 1
	event := "clusterEvent"
	c := p.NewCacheWithParams(size, "name", expiry, event)

	require.Equal(t, event, c.GetInvalidateClusterEvent())

	err := c.SetWithDefaultExpiry("key1", "val1")
	require.Nil(t, err)
	err = c.SetWithDefaultExpiry("key2", "val2")
	require.Nil(t, err)
	l, err := c.Len()
	require.Nil(t, err)
	require.Equal(t, size, l)

	time.Sleep(time.Duration(expiry+1) * time.Second)

	var v string
	err = c.Get("key1", &v)
	require.Equal(t, ErrKeyNotFound, err)
	err = c.Get("key2", &v)
	require.Equal(t, ErrKeyNotFound, err)

}
