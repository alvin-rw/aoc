package main

import "testing"

func TestGetChecksum(t *testing.T) {
	diskMap := getDiskMap("./test.txt")
	compactDiskMap(diskMap)
	if got := calculateChecksum(diskMap); got != 1928 {
		t.Errorf("diskmap v2, got %d, want 1928", got)
	}

	diskMapv2 := getDiskMapv2("./test.txt")
	compactDiskMapv2(diskMapv2)
	if got := calculateChecksumv2(diskMapv2); got != 2858 {
		t.Errorf("diskmap v2, got %d, want 2858", got)
	}
}
