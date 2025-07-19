package main

import (
	"testing"
)

func Test(t *testing.T) {
	polymer := getPolymer("test.txt")

	polymer = react(polymer, 0)
	want := 10

	if len(polymer) != want {
		t.Errorf("part 1: got %d, want %d\n", len(polymer), want)
	}

	shortestPolymerLength := getShortestPolymerLength("test.txt", len(polymer))
	wantPart2 := 4
	if shortestPolymerLength != wantPart2 {
		t.Errorf("part 2: got %d, want %d\n", shortestPolymerLength, wantPart2)
	}
}
