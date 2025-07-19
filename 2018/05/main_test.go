package main

import (
	"testing"
)

func Test(t *testing.T) {
	polymer := getPolymer("test.txt")

	polymer = react(polymer, 0)
	want := 10

	if len(polymer) != want {
		t.Errorf("got %d, want %d\n", len(polymer), want)
	}
}
