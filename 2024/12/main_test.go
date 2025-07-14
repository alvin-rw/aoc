package main

import "testing"

func TestCalculateFencePrice(t *testing.T) {
	mapMatrix := getMapMatrix("test.txt")

	price := calculateFencePrice(mapMatrix)
	want := 1930

	if price != want {
		t.Errorf("got %d, want %d", price, want)
	}
}
