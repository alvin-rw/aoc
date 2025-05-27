package main

import (
	"slices"
	"testing"
)

func TestBlink(t *testing.T) {
	stones := getInputStones("./test.txt")

	for range 6 {
		stones = blink(stones)
	}

	wantPartOne := []int{
		2097446912,
		14168,
		4048,
		2,
		0,
		2,
		4,
		40,
		48,
		2024,
		40,
		48,
		80,
		96,
		2,
		8,
		6,
		7,
		6,
		0,
		3,
		2,
	}

	if slices.Compare(stones, wantPartOne) != 0 {
		t.Errorf("got %v", stones)
	}
}
