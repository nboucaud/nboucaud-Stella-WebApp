// Copyright (c) 2016-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package model

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserSearchJson(t *testing.T) {
	userSearch := UserSearch{Term: NewId(), TeamId: NewId()}
	json := userSearch.ToJson()
	ruserSearch := UserSearchFromJson(bytes.NewReader(json))

	assert.Equal(t, userSearch.Term, ruserSearch.Term, "Terms do not match")
}
