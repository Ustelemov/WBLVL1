package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPowsSum(t *testing.T) {

	testTable := []struct {
		name           string
		inputSlice     []int
		expectedResult int
	}{
		{
			name:           "Not nulls",
			inputSlice:     []int{10, 10, 20, 6, 6},
			expectedResult: 672,
		},
		{
			name:           "Empty",
			expectedResult: 0,
		},
		{
			name:           "Nulls",
			inputSlice:     []int{0, 0, 0, 0, 0},
			expectedResult: 0,
		},
	}

	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {
			result := getPowsSum(v.inputSlice...)
			assert.Equal(t, v.expectedResult, result)
		})
	}
}
