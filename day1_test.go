package main

import (
	"testing"
)

func TestTotalDistance(t *testing.T) {
	int, err := totalDistance("./day1_input.txt")
	if err != nil {
		t.Errorf("an error occured, %+v", err)
	}

	t.Logf("total distance: %d", int)
}
