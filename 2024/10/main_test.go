package main

import (
	"fmt"
	"slices"
	"testing"

	"github.com/alvin-rw/aoc/internal/matrix"
)

func TestTrailheadScorer(t *testing.T) {
	mapMatrix := getMapMatrix("./test.txt")
	trailheadIndexes := getTrailheadIndexes(mapMatrix)

	trailPeaksList := make(map[string][]string)

	for _, trailheadIndex := range trailheadIndexes {
		directions := []int{matrix.Up, matrix.Right, matrix.Down, matrix.Left}
		startingSlope := mapMatrix[trailheadIndex[0]][trailheadIndex[1]]

		if coord := checkTrail(mapMatrix, trailheadIndex, startingSlope, directions); coord != nil {
			trailheadCode := fmt.Sprintf("%v", trailheadIndex)
			trailheadIndexCode := fmt.Sprintf("%v", trailheadIndex)
			if peaks, ok := trailPeaksList[trailheadCode]; ok {
				if !slices.Contains(peaks, trailheadIndexCode) {
					trailPeaksList[trailheadCode] = append(trailPeaksList[trailheadCode], trailheadIndexCode)
				} else {
					continue
				}
			} else {
				trailPeaksList[trailheadCode] = []string{trailheadIndexCode}
			}
		}
	}

	totalTrailheadScore := 0
	for _, achievablePeaks := range trailPeaksList {
		totalTrailheadScore += len(achievablePeaks)
	}

	if totalTrailheadScore != 36 {
		t.Errorf("got %d, want 36", totalTrailheadScore)
	}
}
