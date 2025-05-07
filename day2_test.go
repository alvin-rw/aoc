package aoc2024

import "testing"

func TestCalculateNumberOfSafeReport(t *testing.T) {
	numOfSafeReportWithoutTolerations, err := CalculateNumberOfSafeReport("./day2_input.txt", false)
	if err != nil {
		t.Errorf("an error occurred, %+v", err)
	}

	t.Logf("number of safe report: %d", numOfSafeReportWithoutTolerations)

	numOfSafeReportWithOneToleration, err := CalculateNumberOfSafeReport("./day2_input.txt", true)
	if err != nil {
		t.Errorf("an error occurred, %+v", err)
	}

	t.Logf("number of safe report with 1 tolerations: %d", numOfSafeReportWithOneToleration)
}
