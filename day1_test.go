package main

import (
	"testing"
)

func TestTotalDistance(t *testing.T) {
	totalDist, err := totalDistance("./day1_input.txt")
	if err != nil {
		t.Errorf("an error occured, %+v", err)
	}

	t.Logf("total distance: %d", totalDist)
}

func TestGetSimilarityScore(t *testing.T) {
	similarityScore, err := getSimilarityScore("./day1_input.txt")
	if err != nil {
		t.Errorf("an error occured, %+v", err)
	}

	t.Logf("similarity score: %d", similarityScore)
}
