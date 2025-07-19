package main

import (
	"testing"

	"github.com/alvin-rw/aoc/internal/file"
)

func Test(t *testing.T) {
	lines := file.ReadFile("test.txt")

	want := 26397
	if score := calculateSyntaxErrorScore(lines); score != want {
		t.Errorf("score %d, want %d\n", score, want)
	}
}
