package aoc2024

import "testing"

func TestCalculateNumberOfXMAS(t *testing.T) {
	numberOfXMAS, err := CalculateNumberOfXMAS("./day4_input.txt")
	if err != nil {
		t.Errorf("an error occured, %+v", err)
	}

	t.Logf("number of XMAS: %d", numberOfXMAS)

	numberOfX_MAS, err := CalculateNumberOfX_MAS("./day4_input.txt")
	if err != nil {
		t.Errorf("an error occured, %+v", err)
	}

	t.Logf("number of X_MAS: %d", numberOfX_MAS)
}
