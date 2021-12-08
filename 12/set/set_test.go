package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetToSlice(t *testing.T) {
	set := NewSet()
	expectedResult := []string{"test"}

	set.Add("test")
	result := set.ToSlice()

	assert.Equal(t, expectedResult, result)

}

func TestSetAdd(t *testing.T) {
	set := NewSet()

	set.Add("test")

	assert.True(t, set.Has("test"))

}

func TestSetAddMul(t *testing.T) {
	set := NewSet()

	set.AddMul("test2", "test1")

	assert.True(t, set.Has("test1"))
	assert.True(t, set.Has("test2"))
}

func TestSetHas(t *testing.T) {
	set := NewSet()

	set.Add("test")

	assert.True(t, set.Has("test"))
	assert.False(t, set.Has("test1"))
}

func TestSize(t *testing.T) {
	set := NewSet()

	set.Add("test")
	set.Add("test1")

	assert.Equal(t, set.Size(), 2)
}
