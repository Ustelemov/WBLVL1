package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	testTable := []struct {
		name      string
		input     []int
		searching int
		expected  int
	}{
		{
			name:     "NilOK",
			expected: -1,
		},
		{
			name:     "ZeroOK",
			input:    make([]int, 0),
			expected: -1,
		},
		{
			name:      "ExistsOK",
			input:     []int{-10, 1, 3, 5, 10, 20},
			searching: 3,
			expected:  2,
		},
		{
			name:      "NotExistsOK",
			input:     []int{-2, 1, 2, 10, 15},
			searching: 200,
			expected:  -1,
		},
	}

	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {
			result := binarySearch(v.input, v.searching)
			assert.Equal(t, result, v.expected)
		})
	}
}
