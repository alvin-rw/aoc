package main

import "testing"

func TestCalculateSumOfValidPagesMiddleValue(t *testing.T) {
	want := 143
	got, err := calculateSumOfValidPagesMiddleValue("./test.txt")
	if err != nil {
		t.Errorf("an error occured, %+v", err)
	}

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCalculateSumOfInvalidPagesMiddleValue(t *testing.T) {
	want := 123
	got, err := calculateSumOfInvalidPagesMiddleValue("./test.txt")
	if err != nil {
		t.Errorf("an error occured, %+v", err)
	}

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
