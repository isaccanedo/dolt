// Copyright 2023 Dolthub, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sqle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatch(t *testing.T) {
	assert.False(t, matchWildcardPattern("", "abcdefgh"))
	assert.True(t, matchWildcardPattern("*", "abcdefgh"))
	assert.True(t, matchWildcardPattern("**", "abcdefgh"))

	assert.False(t, matchWildcardPattern("*cdefg", "abcdefgh"))
	assert.True(t, matchWildcardPattern("*cdefgh", "abcdefgh"))
	assert.True(t, matchWildcardPattern("*cdef*", "abcdefgh"))
	assert.True(t, matchWildcardPattern("abcd*efgh", "abcdefgh"))
	assert.True(t, matchWildcardPattern("a*cdef*h", "abcdefgh"))
	assert.True(t, matchWildcardPattern("a*", "abcdefgh"))
	assert.True(t, matchWildcardPattern("*h", "abcdefgh"))
	assert.True(t, matchWildcardPattern("*abcdefgh", "abcdefgh"))
	assert.False(t, matchWildcardPattern("*abcdefg", "abcdefgh"))
	assert.True(t, matchWildcardPattern("*abcdefgh*", "abcdefgh"))
}

func TestContainsWildcard(t *testing.T) {
	assert.False(t, containsWildcards(""))
	assert.False(t, containsWildcards("abc"))
	assert.True(t, containsWildcards("*"))
	assert.True(t, containsWildcards("a*c"))
	assert.True(t, containsWildcards("ab*"))
	assert.True(t, containsWildcards("*bc"))
}
