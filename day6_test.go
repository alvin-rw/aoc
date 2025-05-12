package aoc2024

import "testing"

func TestGetGuardNumberOfDistinctPosition(t *testing.T) {
	numberOfDistinctPosition, err := GetGuardNumberOfDistinctPosition("./day6_input.txt")
	if err != nil {
		t.Errorf("an error occured, %+v", err)
	}

	t.Logf("number of distinct position visited by guard: %d", numberOfDistinctPosition)
}
