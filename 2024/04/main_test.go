package main

import "testing"

func TestCalculateNumberOfXMAS(t *testing.T) {
	want := 18
	got := calculateNumberOfXMAS("./test.txt")

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCalculateNumberOfX_MAS(t *testing.T) {
	want := 9
	got := calculateNumberOfX_MAS("./test.txt")

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
