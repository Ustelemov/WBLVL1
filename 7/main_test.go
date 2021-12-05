package main

import (
	"math/rand"
	"testing"
)

func TestConcurrentMap(t *testing.T) {
	t.Run("RaceOK", func(t *testing.T) {
		key := "key"
		m := NewConcurrentMap()

		for i := 0; i < 10; i++ {
			go func() {
				m.Store(key, rand.Intn(100))
			}()
		}
	})
}
