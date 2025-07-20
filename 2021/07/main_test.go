package main

import (
	"testing"
)

func Test(t *testing.T) {
	crabLocations, lowest, highest := getCrabLocationsAndBoundaries("test.txt")

	want := 37

	if fuelConsumed := getLowestFuel(crabLocations, lowest, highest); fuelConsumed != want {
		t.Errorf("part 1: got %d, want %d", fuelConsumed, want)
	}
}
