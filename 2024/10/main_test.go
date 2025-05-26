package main

import (
	"testing"

	"github.com/alvin-rw/aoc/internal/matrix"
)

func TestTrailheadScorer(t *testing.T) {
	mapMatrix := getMapMatrix("./test.txt")
	trailheadIndexes := getTrailheadIndexes(mapMatrix)

	totalTrailHeadScore := 0
	for _, trailheadIndex := range trailheadIndexes {
		directions := []int{matrix.Up, matrix.Right, matrix.Down, matrix.Left}

		for _, direction := range directions {

		}
	}

	if totalTrailHeadScore != 36 {
		t.Errorf("got %d, want 36", totalTrailHeadScore)
	}
}
