package main

import (
	"testing"
)

func TestCalculateTotalDistance(t *testing.T) {
	want := 11
	got, err := calculateTotalDistance("./test.txt")
	if err != nil {
		t.Errorf("an error occured, %+v", err)
	}

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCalculateSimilarityScore(t *testing.T) {
	want := 31
	got, err := calculateSimilarityScore("./test.txt")
	if err != nil {
		t.Errorf("an error occured, %+v", err)
	}

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
