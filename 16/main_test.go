package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	testTable := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name: "NilOK",
		},
		{
			name:     "ZeroOK",
			input:    make([]int, 0),
			expected: make([]int, 0),
		},
		{
			name:     "OneOK",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "PositiveOK",
			input:    []int{10, 5, 3, 11, 1, 12},
			expected: []int{1, 3, 5, 10, 11, 12},
		},
		{
			name:     "NegativeOK",
			input:    []int{-10, 5, 3, 11, -1, 12},
			expected: []int{-10, -1, 3, 5, 11, 12},
		},
	}

	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {
			result := quickSort(v.input)
			assert.True(t, reflect.DeepEqual(result, v.expected))
		})
	}
}
