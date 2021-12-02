package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPowChan(t *testing.T) {

	testTable := []struct {
		name           string
		inputArr       [5]int
		expectedResult [5]int
	}{
		{
			name:           "Not nulls",
			inputArr:       [5]int{1, 2, 3, 4, 5},
			expectedResult: [5]int{1, 4, 9, 16, 25},
		},
		{
			name:           "Nulls",
			expectedResult: [5]int{0, 0, 0, 0, 0},
		},
	}

	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {

			channel := make(chan int)
			var result [5]int

			go getPowChan(channel, v.inputArr)
			index := 0
			for pow := range channel {
				result[index] = pow
				index = index + 1
			}

			assert.Equal(t, v.expectedResult, result)

		})
	}
}

func TestGetPoWWG(t *testing.T) {

	testTable := []struct {
		name           string
		inputArr       [5]int
		expectedResult []int
	}{
		{
			name:           "Not nulls",
			inputArr:       [5]int{1, 2, 3, 4, 5},
			expectedResult: []int{1, 4, 9, 16, 25},
		},
		{
			name:           "Nulls",
			expectedResult: []int{0, 0, 0, 0, 0},
		},
	}

	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {
			result := getPowWG(v.inputArr)
			assert.Equal(t, v.expectedResult, result)
		})
	}
}
