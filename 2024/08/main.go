package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/alvin-rw/aoc/internal/file"
)

func main() {
	mapMatrix := getMapMatrix("./input.txt")

	maxRow := len(mapMatrix)
	maxColumn := len(mapMatrix[0])

	antennaCoordinateMap := getAntennaCoordinateMap(mapMatrix)

	numberOfAntinode := calculateNumberOfAntinode(antennaCoordinateMap, maxRow, maxColumn)

	fmt.Printf("number of antinode inside map: %d\n", numberOfAntinode)
}

func calculateNumberOfAntinode(antennaCoordinateMap map[string][][]int, maxRow, maxColumn int) int {
	numberOfAntinode := 0

	antinodesCoordinates := []string{}

	addUniqueNumberOfAntinodes := func(coordinate []int) {
		c := fmt.Sprintf("%v", coordinate)
		if !slices.Contains(antinodesCoordinates, c) {
			antinodesCoordinates = append(antinodesCoordinates, c)
			numberOfAntinode++
		}
	}

	for _, coordinates := range antennaCoordinateMap {
		for i, firstCoord := range coordinates {
			for j := i + 1; j < len(coordinates); j++ {
				secondCoord := coordinates[j]

				rowDist := firstCoord[0] - secondCoord[0]
				columnDist := firstCoord[1] - secondCoord[1]

				firstAntinodeCoord := []int{firstCoord[0] + rowDist, firstCoord[1] + columnDist}
				if isCoordinateInsideMap(firstAntinodeCoord, maxRow, maxColumn) {
					addUniqueNumberOfAntinodes(firstAntinodeCoord)
				}

				secondAntinodeCoord := []int{secondCoord[0] - rowDist, secondCoord[1] - columnDist}
				if isCoordinateInsideMap(secondAntinodeCoord, maxRow, maxColumn) {
					addUniqueNumberOfAntinodes(secondAntinodeCoord)
				}
			}
		}
	}

	return numberOfAntinode
}

func isCoordinateInsideMap(coordinate []int, maxRow int, maxColumn int) bool {
	if coordinate[0] >= 0 && coordinate[1] >= 0 && coordinate[0] < maxRow && coordinate[1] < maxColumn {
		return true
	} else {
		return false
	}
}

func getAntennaCoordinateMap(mapMatrix [][]string) map[string][][]int {
	antennaCoordinateMap := make(map[string][][]int)

	for row, line := range mapMatrix {
		for column, char := range line {
			if char != "." {
				antennaCoordinateMap[char] = append(antennaCoordinateMap[char], []int{row, column})
			}
		}
	}

	return antennaCoordinateMap
}

func getMapMatrix(inputFilePath string) [][]string {
	mapMatrix := [][]string{}

	fileContent := file.ReadFile(inputFilePath)

	for _, line := range fileContent {
		lineSlice := strings.Split(line, "")

		mapMatrix = append(mapMatrix, lineSlice)
	}

	return mapMatrix
}
