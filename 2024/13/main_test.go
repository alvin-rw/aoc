package main

import (
	"testing"
)

func Test(t *testing.T) {
	machines := getMachinesDetails("test.txt")

	totalCost := 0
	for _, m := range machines {
		totalCost += calculateMachineCost(&m)
	}

	want := 480

	if totalCost != want {
		t.Errorf("got %d, want %d\n", totalCost, want)
	}
}
