package aoc2024

import "testing"

func TestCalculateSumOfValidPagesMiddleValue(t *testing.T) {
	sumOfValidPagesMidValue, err := CalculateSumOfValidPagesMiddleValue("./day5_input.txt")
	if err != nil {
		t.Errorf("an error occured, %+v", err)
	}

	t.Logf("sum of valid pages mid value: %d", sumOfValidPagesMidValue)
}
