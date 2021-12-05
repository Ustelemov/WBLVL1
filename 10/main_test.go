package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTemperatureToGroup(t *testing.T) {
	testTable := []struct {
		name          string
		inputArr      []float64
		expectedValue map[int][]float64
	}{
		{
			name:          "NilOK",
			expectedValue: make(map[int][]float64),
		},
		{
			name:          "ZeroOK",
			inputArr:      make([]float64, 0),
			expectedValue: make(map[int][]float64),
		},
		{
			name:     "ValuesOK",
			inputArr: []float64{-25.3, 0, 5, 12, -30, -80, 4, 4},
			expectedValue: map[int][]float64{
				-80: {-80},
				-30: {-30},
				-20: {-25.3},
				0:   {0, 5, 4, 4},
				10:  {12},
			},
		},
	}

	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {
			result := getTemperatureToGroups(v.inputArr)
			assert.True(t, reflect.DeepEqual(result, v.expectedValue))
		})
	}
}
