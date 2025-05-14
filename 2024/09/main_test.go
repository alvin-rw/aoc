package main

import "testing"

func TestGetChecksum(t *testing.T) {
	diskMap := getDiskMap("./test.txt")

	t.Logf("%v\n", diskMap)

	compactDiskMap(diskMap)

	t.Logf("%v\n", diskMap)

	wantChecksum := 1928
	checksum := calculateChecksum(diskMap)

	if checksum != wantChecksum {
		t.Errorf("got %d, want %d", checksum, wantChecksum)
	}
}
