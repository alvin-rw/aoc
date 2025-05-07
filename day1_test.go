package aoc2024

import (
	"testing"
)

func TestCalculateTotalDistance(t *testing.T) {
	totalDist, err := CalculateTotalDistance("./day1_input.txt")
	if err != nil {
		t.Errorf("an error occured, %+v", err)
	}

	t.Logf("total distance: %d", totalDist)
}

func TestCalculateSimilarityScore(t *testing.T) {
	similarityScore, err := CalculateSimilarityScore("./day1_input.txt")
	if err != nil {
		t.Errorf("an error occured, %+v", err)
	}

	t.Logf("similarity score: %d", similarityScore)
}
