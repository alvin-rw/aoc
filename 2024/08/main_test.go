package main

import "testing"

func TestCalculateNumberOfAntinode(t *testing.T) {
	mapMatrix := getMapMatrix("./test.txt")

	maxRow := len(mapMatrix)
	maxColumn := len(mapMatrix[0])

	antennaCoordinateMap := getAntennaCoordinateMap(mapMatrix)

	want := 14
	got := calculateNumberOfAntinode(antennaCoordinateMap, maxRow, maxColumn)

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCalculateNumberOfAntinodeWithResonanceHarmonics(t *testing.T) {
	mapMatrix := getMapMatrix("./test.txt")

	maxRow := len(mapMatrix)
	maxColumn := len(mapMatrix[0])

	antennaCoordinateMap := getAntennaCoordinateMap(mapMatrix)

	want := 34
	got := calculateNumberOfAntinodeWithResonanceHarmonics(antennaCoordinateMap, maxRow, maxColumn)

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
