package main

import (
	"slices"
	"testing"
)

func TestCalibrationResult(t *testing.T) {
	cases := []struct {
		name          string
		inputFilePath string
		operators     []string
		want          int
	}{
		{
			name:          "test part 1",
			inputFilePath: "./test.txt",
			operators:     []string{"+", "*"},
			want:          3749,
		},
		{
			name:          "input part 1",
			inputFilePath: "./input.txt",
			operators:     []string{"+", "*"},
			want:          1620690235709,
		},
		{
			name:          "test part 2",
			inputFilePath: "./test.txt",
			operators:     []string{"+", "*", "||"},
			want:          11387,
		},
	}

	for _, tt := range cases {
		calibrationElementsList, err := getCalibrationElementsList(tt.inputFilePath)
		if err != nil {
			t.Errorf("an error occured: %v", err)
		}

		result := 0

		for _, calibrationElements := range calibrationElementsList {
			desiredResult := calibrationElements[0]
			input := slices.Delete(calibrationElements, 0, 1)

			if calibrate(input, desiredResult, tt.operators) {
				result += desiredResult
			}
		}

		if result != tt.want {
			t.Errorf("test name: %s, got %d, want %d", tt.name, result, tt.want)
		}
	}

}
