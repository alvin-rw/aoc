package aoc2024

import "testing"

func TestCalculateMultiplicationResult(t *testing.T) {
	res, err := CalculateMultiplicationResult("./day3_input.txt", false)
	if err != nil {
		t.Errorf("an error occured, %+v", err)
	}

	t.Logf("multiplication result: %d", res)

	resWithDoDonts, err := CalculateMultiplicationResult("./day3_input.txt", true)
	if err != nil {
		t.Errorf("an error occured, %+v", err)
	}

	t.Logf("multiplication result with do and dont commands: %d", resWithDoDonts)
}
