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

func TestStoneCounter(t *testing.T) {
	cases := []struct {
		name       string
		input      string
		blinkCount int
		want       int
	}{
		// {
		// 	name:       "test",
		// 	input:      "./test.txt",
		// 	blinkCount: 6,
		// 	want:       22,
		// },
		// {
		// 	name:       "input part 1",
		// 	input:      "./input.txt",
		// 	blinkCount: 25,
		// 	want:       218079,
		// },
		{
			name:       "input part 1",
			input:      "./input.txt",
			blinkCount: 75,
			want:       218079,
		},
	}

	for _, tt := range cases {
		stones := getInputStones(tt.input)

		if got := getStoneCountAfterBlinks(tt.blinkCount, stones); got != tt.want {
			t.Errorf("%s, got %d, want %d", tt.name, got, tt.want)
		}
	}
}
