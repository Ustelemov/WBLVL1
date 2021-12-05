package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
	testTable := []struct {
		name          string
		inputArr      []int
		expectedValue []int
	}{
		{
			name:          "NilOK",
			expectedValue: make([]int, 0),
		},
		{
			name:          "ZeroOK",
			inputArr:      []int{0, 0, 0, 0},
			expectedValue: []int{0},
		},
		{
			name:          "ValuesOK",
			inputArr:      []int{1, 1, 1, 2, 3, 4, 5, 6, 6},
			expectedValue: []int{1, 2, 3, 4, 5, 6},
		},
	}

	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {
			result := removeDuplicates(v.inputArr)
			assert.True(t, reflect.DeepEqual(result, v.expectedValue))
		})
	}
}

func TestSimpleIntersect(t *testing.T) {
	testTable := []struct {
		name          string
		inputSet1     []int
		inputSet2     []int
		expectedValue []int
	}{
		{
			name:          "NilOK",
			expectedValue: make([]int, 0),
		},
		{
			name:          "ZeroOK",
			inputSet1:     []int{0, 0, 0, 0, 0, 0},
			inputSet2:     []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
			expectedValue: []int{0},
		},
		{
			name:          "ValuesOK",
			inputSet1:     []int{1, 2, 3, 4, 5, 5, 5, 5, 6, 7},
			inputSet2:     []int{1, 1, 1, 1, 1, 2, 2, 2, 3, 7, 7},
			expectedValue: []int{1, 2, 3, 7},
		},
	}

	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {
			result := simpleIntersect(v.inputSet1, v.inputSet2)
			assert.True(t, reflect.DeepEqual(result, v.expectedValue))
		})
	}
}

func TestHashMapIntersect(t *testing.T) {
	testTable := []struct {
		name          string
		inputSet1     []int
		inputSet2     []int
		expectedValue []int
	}{
		{
			name:          "NilOK",
			expectedValue: make([]int, 0),
		},
		{
			name:          "ZeroOK",
			inputSet1:     []int{0, 0, 0, 0, 0, 0},
			inputSet2:     []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
			expectedValue: []int{0},
		},
		{
			name:          "ValuesOK",
			inputSet1:     []int{1, 2, 3, 4, 5, 5, 5, 5, 6, 7},
			inputSet2:     []int{1, 1, 1, 1, 1, 2, 2, 2, 3, 7, 7},
			expectedValue: []int{1, 2, 3, 7},
		},
	}

	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {
			result := hashMapIntersect(v.inputSet1, v.inputSet2)
			assert.True(t, reflect.DeepEqual(result, v.expectedValue))
		})
	}
}

func TestSortIntersect(t *testing.T) {
	testTable := []struct {
		name          string
		inputSet1     []int
		inputSet2     []int
		expectedValue []int
	}{
		{
			name:          "NilOK",
			expectedValue: make([]int, 0),
		},
		{
			name:          "ZeroOK",
			inputSet1:     []int{0, 0, 0, 0, 0, 0},
			inputSet2:     []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
			expectedValue: []int{0},
		},
		{
			name:          "ValuesOK",
			inputSet1:     []int{1, 2, 3, 4, 5, 5, 5, 5, 6, 7},
			inputSet2:     []int{1, 1, 1, 1, 1, 2, 2, 2, 3, 7, 7},
			expectedValue: []int{1, 2, 3, 7},
		},
	}

	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {
			result := sortIntersect(v.inputSet1, v.inputSet2)
			assert.True(t, reflect.DeepEqual(result, v.expectedValue))
		})
	}
}
