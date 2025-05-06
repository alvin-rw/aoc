package aoc2024

import "testing"

func TestCalculateSafeReport(t *testing.T) {
	numOfSafeReport, err := calculateSafeReport("./day2_input.txt")
	if err != nil {
		t.Errorf("an error occurred, %+v", err)
	}

	t.Logf("number of safe report: %d", numOfSafeReport)
}
