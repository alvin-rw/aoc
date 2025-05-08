package aoc2024

import "testing"

func TestCalculateMultiplicationResult(t *testing.T) {
	res, err := CalculateMultiplicationResult("./day3_input.txt")
	if err != nil {
		t.Errorf("an error occured, %+v", err)
	}

	t.Logf("multiplication result: %d", res)
}
