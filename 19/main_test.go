package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseString(t *testing.T) {
	assert.Equal(t, reverseString(""), "")
	assert.Equal(t, reverseString("X"), "X")
	assert.Equal(t, reverseString("The quick bròwn 狐 jumped over the lazy 犬"), "犬 yzal eht revo depmuj 狐 nwòrb kciuq ehT")
	assert.Equal(t, reverseString("😎⚽"), "⚽😎")
}
