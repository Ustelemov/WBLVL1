package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInputChan(t *testing.T) {
	inputArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	resultArr := make([]int, len(inputArr))

	ind := 0
	for val := range inputChan(inputArr...) {
		resultArr[ind] = val
		ind += 1
	}

	assert.Equal(t, inputArr, resultArr)
}

func TestMultChan(t *testing.T) {
	inputArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expectedArr := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}

	resultArr := make([]int, len(inputArr))

	ind := 0
	for val := range multChan(inputChan(inputArr...)) {
		resultArr[ind] = val
		ind += 1
	}

	assert.Equal(t, expectedArr, resultArr)
}
