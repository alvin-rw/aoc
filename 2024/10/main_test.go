package main

import (
	"testing"
)

func TestTrailheadScorer(t *testing.T) {
	mapMatrix := getMapMatrix("./test.txt")

	trailheadIndexes := getTrailheadIndexes(mapMatrix)
	trailheadPeaksMap := getTrailheadsPeaksMap(mapMatrix, trailheadIndexes)
	totalTrailheadScore := calculateTrailheadScore(trailheadPeaksMap)

	if totalTrailheadScore != 36 {
		t.Errorf("got %d, want 36", totalTrailheadScore)
	}
}
