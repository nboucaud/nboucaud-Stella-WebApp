// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package config

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestWatcherInvalidDirectory(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping watcher test in short mode")
	}

	callback := func() {}
	_, err := newWatcher("/does/not/exist", callback)
	require.Error(t, err, "should have failed to watch a non-existent directory")
}

func TestWatcher(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping watcher test in short mode")
	}

	tempDir, err := ioutil.TempDir("", "TestWatcher")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	f, err := ioutil.TempFile(tempDir, "TestWatcher")
	require.NoError(t, err)
	defer f.Close()
	defer os.Remove(f.Name())

	called := make(chan bool)
	callback := func() {
		called <- true
	}
	watcher, err := newWatcher(f.Name(), callback)
	require.NoError(t, err)
	defer watcher.Close()

	// Write to a different file
	ioutil.WriteFile(filepath.Join(tempDir, "unrelated"), []byte("data"), 0644)
	require.False(t, wasCalled(called, 1*time.Second), "callback should not have been called for unrelated file")

	// Write to the watched file
	ioutil.WriteFile(f.Name(), []byte("data"), 0644)
	require.True(t, wasCalled(called, 5*time.Second), "callback should have been called when file written")
}

// wasCalled reports whether a given callback channel was called
// within the specified time duration or not.
func wasCalled(c chan bool, duration time.Duration) bool {
	wasCalled := false
	select {
	case <-c:
		wasCalled = true
	case <-time.After(duration):
	}
	return wasCalled
}
