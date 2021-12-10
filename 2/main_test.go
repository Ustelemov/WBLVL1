package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPowChan(t *testing.T) {

	testTable := []struct {
		name           string
		inputArr       []int
		expectedResult []int
	}{
		{
			name:           "Not nulls",
			inputArr:       []int{1, 2, 3, 4, 5},
			expectedResult: []int{1, 4, 9, 16, 25},
		},
		{
			name:           "Nulls",
			expectedResult: []int{},
		},
	}

	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {

			channel := make(chan int)
			result := make([]int, len(v.inputArr))

			go getPowChan(channel, v.inputArr)
			index := 0
			for pow := range channel {
				result[index] = pow
				index = index + 1
			}

			assert.True(t, reflect.DeepEqual(v.expectedResult, result))

		})
	}
}

func TestGetPoWWG(t *testing.T) {

	testTable := []struct {
		name           string
		inputArr       []int
		expectedResult []int
	}{
		{
			name:           "Not nulls",
			inputArr:       []int{1, 2, 3, 4, 5},
			expectedResult: []int{1, 4, 9, 16, 25},
		},
		{
			name:           "Nulls",
			expectedResult: []int{},
		},
	}

	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {
			result := getPowWG(v.inputArr)
			assert.True(t, reflect.DeepEqual(v.expectedResult, result))
		})
	}
}
