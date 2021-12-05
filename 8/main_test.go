package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetBit(t *testing.T) {
	testTable := []struct {
		name          string
		inputValue    int64
		position      uint
		expectedValue int64
	}{
		{
			name:          "zeroOK",
			inputValue:    0,
			position:      7,
			expectedValue: 128,
		},
		{
			name:          "negativeOK",
			inputValue:    -32,
			position:      0,
			expectedValue: -31,
		},
		{
			name:          "positiveOK",
			inputValue:    1,
			position:      5,
			expectedValue: 33,
		},
	}

	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {
			res := setBit(v.inputValue, v.position)
			assert.Equal(t, res, v.expectedValue)
		})
	}
}

func TestClearBit(t *testing.T) {
	testTable := []struct {
		name          string
		inputValue    int64
		position      uint
		expectedValue int64
	}{
		{
			name:          "zeroOK",
			inputValue:    0,
			position:      7,
			expectedValue: 0,
		},
		{
			name:          "negativeOK",
			inputValue:    -32,
			position:      6,
			expectedValue: -96,
		},
		{
			name:          "positiveOK",
			inputValue:    129,
			position:      7,
			expectedValue: 1,
		},
	}

	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {
			res := clearBit(v.inputValue, v.position)
			assert.Equal(t, res, v.expectedValue)
		})
	}
}
