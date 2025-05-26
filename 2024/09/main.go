package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/alvin-rw/aoc/internal/file"
)

type diskItem struct {
	fileId        int
	startingIndex int
	size          int
}

func main() {
	diskMap := getDiskMap("./input.txt")
	compactDiskMap(diskMap)
	checksum := calculateChecksum(diskMap)

	fmt.Printf("checksum: %d\n", checksum)

	diskMapv2 := getDiskMapv2("./input.txt")
	compactDiskMapv2(diskMapv2)
	checksumv2 := calculateChecksumv2(diskMapv2)

	fmt.Printf("checksum part 2: %d\n", checksumv2)
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
			for k := len(diskMap) - 1; k >= i; k-- {
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

func calculateChecksumv2(diskMap []diskItem) int {
	checksum := 0

	counter := 0

	for _, di := range diskMap {
		if di.fileId != -1 {
			for range di.size {
				checksum = checksum + (counter * di.fileId)
				counter++
			}
		} else {
			counter += di.size
		}
	}

	return checksum
}

func compactDiskMapv2(diskMap []diskItem) {
	for i := len(diskMap) - 1; i >= 0; i-- {
		if diskMap[i].fileId != -1 {
			file := diskMap[i]

			for j := range i {
				if diskMap[j].fileId == -1 && diskMap[j].size >= file.size {
					file.startingIndex = diskMap[j].startingIndex

					diskMap[i].fileId = -1
					if diskMap[j].size == file.size {
						diskMap[j] = file
					} else if diskMap[j].size > file.size {
						diskMap[j].startingIndex += file.size
						diskMap[j].size -= file.size
						diskMap = slices.Insert(diskMap, j, file)
					}
					break
				}
			}
		}
	}
}

func getDiskMapv2(inputFilePath string) []diskItem {
	diskMap := []diskItem{}

	fileContent := file.ReadFile(inputFilePath)
	diskMapSlice := strings.Split(fileContent[0], "")

	readDiskSize := 0

	for i, n := range diskMapSlice {
		itemSize, _ := strconv.Atoi(n)

		fileId := -1  // fileId for empty disk space
		if i%2 == 0 { // file
			fileId = i / 2
		}

		di := diskItem{
			fileId:        fileId,
			startingIndex: readDiskSize,
			size:          itemSize,
		}

		diskMap = append(diskMap, di)

		readDiskSize += itemSize
	}

	return diskMap
}

/**
LOGIC

example disk
123456

0..111....22222......

diskMap := []diskItem{
		{
			fileId:        0,
			startingIndex: 0,
			size:    			 1,
		},
		{
			fileId:        -1,
			startingIndex: 1,
			size:  			   2,
		},
		{
			fileId:        1,
			startingIndex: 3,
			size:  			   3,
		},
		{
			fileId:        -1,
			startingIndex: 6,
			size:          4,
		},
		{
			fileId:        2,
			startingIndex: 10,
			size:          5,
		},
		{
			fileId:        -1,
			startingIndex: 15,
			size:          6,
		},
	}
**/
