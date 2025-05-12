package main

import "testing"

func TestGetGuardsNumberOfDistinctPosition(t *testing.T) {
	want := 41
	got, err := getGuardsNumberOfDistinctPosition("./test.txt")
	if err != nil {
		t.Errorf("an error occured, %+v", err)
	}

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
