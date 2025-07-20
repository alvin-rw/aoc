package main

import (
	"testing"
)

func Test(t *testing.T) {
	measurements := getMeasurements("test.txt")

	want := 7

	if got := getNumOfLargerMeasurements(measurements); got != want {
		t.Errorf("part 1: got %d, want %d", got, want)
	}

	measurementsSlidingWindow := getMeasurementSlidingWindow(measurements)
	wantSlidingWindow := 5

	if got := getNumOfLargerMeasurements(measurementsSlidingWindow); got != wantSlidingWindow {
		t.Errorf("part 2: got %d, want %d", got, wantSlidingWindow)
	}
}
