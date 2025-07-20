package main

import (
	"testing"
)

func Test(t *testing.T) {
	cases := []struct {
		part          int
		inputFilePath string
		want          int
	}{
		{
			part:          1,
			inputFilePath: "test.txt",
			want:          37,
		},
		{
			part:          2,
			inputFilePath: "test.txt",
			want:          168,
		},
	}

	for _, tt := range cases {
		crabLocations, lowest, highest := getCrabLocationsAndBoundaries(tt.inputFilePath)

		if fuelConsumed := getLowestFuel(crabLocations, lowest, highest, tt.part); fuelConsumed != tt.want {
			t.Errorf("part %d: got %d, want %d", tt.part, fuelConsumed, tt.want)
		}
	}
}
