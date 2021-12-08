package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveIndex(t *testing.T) {
	assert.True(t, reflect.DeepEqual([]int{2, 3, 4, 5}, RemoveIndex([]int{1, 2, 3, 4, 5}, 0)))
	assert.True(t, reflect.DeepEqual([]int{1, 2, 3, 4}, RemoveIndex([]int{1, 2, 3, 4, 5}, 4)))
	assert.True(t, reflect.DeepEqual([]int{}, RemoveIndex([]int{5}, 0)))
}
