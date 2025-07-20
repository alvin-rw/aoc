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
}
