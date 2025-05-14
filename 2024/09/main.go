package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/alvin-rw/aoc/internal/file"
	"github.com/alvin-rw/aoc/internal/utils"
)

func main() {
	diskMap := getDiskMap("./input.txt")
	compactDiskMap(diskMap)

	checksum := calculateChecksum(diskMap)

	fmt.Printf("checksum: %d\n", checksum)
}

func calculateChecksum(diskMap []string) int {
	checksum := 0

	for i, diskItem := range diskMap {
		if diskItem != "." {
			di, _ := strconv.Atoi(diskItem)

			checksum += di * i
		}
	}

	return checksum
}

func compactDiskMap(diskMap []string) {
	for i, diskItem := range diskMap {
		if diskItem == "." {
			// check if there are any numbers after the current .
			diskMapPart := diskMap[i:]
			if utils.GetNumberOfElementInSlice(diskMapPart, ".") == len(diskMapPart) {
				break
			}

			for k := len(diskMap) - 1; k >= 0; k-- {
				if diskMap[k] != "." {
					diskMap[i] = diskMap[k]
					diskMap[k] = "."
					break
				}
			}
		}
	}
}

func getDiskMap(inputFilePath string) []string {
	fileContent := file.ReadFile(inputFilePath)

	diskMapSlice := strings.Split(fileContent[0], "")

	diskMap := []string{}

	for i, r := range diskMapSlice {
		counter, _ := strconv.Atoi(r)

		for range counter {
			if i%2 == 0 {
				diskMap = append(diskMap, strconv.Itoa(i/2))
			} else {
				diskMap = append(diskMap, ".")
			}
		}
	}

	return diskMap
}
