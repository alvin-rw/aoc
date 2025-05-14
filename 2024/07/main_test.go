package main

import "testing"

func TestCalibrationResult(t *testing.T) {
	cases := []struct {
		name          string
		inputFilePath string
		want          int
	}{
		{
			name:          "test file",
			inputFilePath: "./test.txt",
			want:          3749,
		},
		{
			name:          "input file",
			inputFilePath: "./input.txt",
			want:          1620690235709,
		},
	}

	for _, tt := range cases {
		calibrationElementsList, err := getCalibrationElementsList(tt.inputFilePath)
		if err != nil {
			t.Errorf("an error occured: %v", err)
		}

		got := 0

		for _, calibrationElements := range calibrationElementsList {
			if res, ok := checkCalibrationResult(calibrationElements); ok {
				got += res
			}
		}

		if got != tt.want {
			t.Errorf("test name: %s, got %d, want %d", tt.name, got, tt.want)
		}
	}

}
