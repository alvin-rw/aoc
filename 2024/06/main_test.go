package main

import (
	"testing"
)

func TestGetGuardsNumberOfDistinctPosition(t *testing.T) {
	mapMatrix := createMapMatrix("./test.txt")

	guardSymbol := "^"
	guardFirstPosition, err := getGuardFirstPosition(mapMatrix, guardSymbol)
	if err != nil {
		t.Errorf("error when getting guard first position: %v", err)
	}

	gotNumberOfDistinctPosition, guardVisitedPoints := calculateAndListDistinctPosition(mapMatrix, guardFirstPosition)

	gotNumOfPossibleObstruction := calculateNumberOfPossibleObstructions(mapMatrix, guardFirstPosition, guardVisitedPoints)

	wantNumberOfDistinctPosition := 41
	if gotNumberOfDistinctPosition != wantNumberOfDistinctPosition {
		t.Errorf("error when getting guard number of distinct position: got %d, want %d", gotNumberOfDistinctPosition, wantNumberOfDistinctPosition)
	}

	wantNumberOfPossibleObstruction := 6
	if gotNumOfPossibleObstruction != wantNumberOfPossibleObstruction {
		t.Errorf("error when getting number of possible obstructions: got %d, want %d", gotNumOfPossibleObstruction, wantNumberOfPossibleObstruction)
	}
}
