package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPoint(t *testing.T) {
	p := NewPoint(-1, 100)
	assert.Equal(t, float64(-1), p.getX())
	assert.Equal(t, float64(100), p.getY())
}

func TestFindDist(t *testing.T) {
	p1 := NewPoint(7, 6)
	p2 := NewPoint(15, 12)

	assert.Equal(t, findDist(*p1, *p2), float64(10))

	p1 = NewPoint(0, 0)
	p2 = NewPoint(0, 0)

	assert.Equal(t, findDist(*p1, *p2), float64(0))

}
