package main

import "testing"

func TestGetChecksumPart1(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "test case",
			input: "./test.txt",
			want:  1928,
		},
		{
			name:  "input case",
			input: "./input.txt",
			want:  6307275788409,
		},
	}

	for _, tt := range cases {
		diskMap := getDiskMap(tt.input)
		compactDiskMap(diskMap)
		if got := calculateChecksum(diskMap); got != tt.want {
			t.Errorf("%s, got %d, want %d", tt.name, got, tt.want)
		}
	}
}

func TestGetChecksumPart2(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "test case",
			input: "./test.txt",
			want:  2858,
		},
		{
			name:  "input case",
			input: "./input.txt",
			want:  6327174563252,
		},
	}

	for _, tt := range cases {
		diskMapv2 := getDiskMapv2(tt.input)
		compactDiskMapv2(diskMapv2)
		if got := calculateChecksumv2(diskMapv2); got != tt.want {
			t.Errorf("%s, got %d, want %d", tt.name, got, tt.want)
		}
	}
}
