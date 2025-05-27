package main

import (
	"testing"
)

func TestTrailheadScorer(t *testing.T) {
	mapMatrix := getMapMatrix("./test.txt")

	totalTrailheadRatings := new(int)

	trailheadIndexes := getTrailheadIndexes(mapMatrix)
	trailheadPeakMap := getTrailheadsPeakMapAndRatings(mapMatrix, trailheadIndexes, totalTrailheadRatings)
	totalTrailheadScore := calculateTrailheadScore(trailheadPeakMap)

	if totalTrailheadScore != 36 {
		t.Errorf("got %d, want 36", totalTrailheadScore)
	}

	if *totalTrailheadRatings != 81 {
		t.Errorf("got %d trailhead ratings, want 81", totalTrailheadScore)
	}
}
