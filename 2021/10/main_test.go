package main

import (
	"testing"

	"github.com/alvin-rw/aoc/internal/file"
)

func Test(t *testing.T) {
	lines := file.ReadFile("test.txt")

	wantCorrupted := 26397
	wantAutoComplete := 288957
	corruptedScore, autoCompleteScore := calculateSyntaxErrorScore(lines)

	if corruptedScore != wantCorrupted {
		t.Errorf("corruptedScore %d, want %d\n", corruptedScore, wantCorrupted)
	}

	if autoCompleteScore != wantAutoComplete {
		t.Errorf("autoCompleteScore %d, want %d\n", autoCompleteScore, wantCorrupted)
	}
}
