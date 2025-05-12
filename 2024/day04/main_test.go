package main

import "testing"

func TestCalculateNumberOfXMAS(t *testing.T) {
	want := 18
	got, err := calculateNumberOfXMAS("./test.txt")
	if err != nil {
		t.Errorf("an error occured, %+v", err)
	}

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
func TestCalculateNumberOfX_MAS(t *testing.T) {
	want := 9
	got, err := calculateNumberOfX_MAS("./test.txt")
	if err != nil {
		t.Errorf("an error occured, %+v", err)
	}

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
