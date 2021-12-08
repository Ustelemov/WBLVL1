package main

import (
	"testing"
	"time"
)

func TestSleep(t *testing.T) {
	const secs = 10
	start := time.Now()
	Sleep(10)
	duration := time.Since(start).Seconds()

	if duration < secs || duration > duration*1.05 {
		t.Error("Sleep duration is not correct")
	}
}
