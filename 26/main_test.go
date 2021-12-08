package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckUnique(t *testing.T) {
	assert.False(t, checkUnique("AaaaBcc"))
	assert.True(t, checkUnique("qweasdrt"))
	assert.True(t, checkUnique(""))
}
