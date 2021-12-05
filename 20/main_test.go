package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseString(t *testing.T) {
	assert.Equal(t, reverseWords(""), "")
	assert.Equal(t, reverseWords("X"), "X")
	assert.Equal(t, reverseWords("XX"), "XX")
	assert.Equal(t, reverseWords("The quick bròwn 狐 jumped over the lazy 犬"), "犬 lazy the over jumped 狐 bròwn quick The")
	assert.Equal(t, reverseWords("😎 ⚽"), "⚽ 😎")
}
