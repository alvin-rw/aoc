package main

import (
	"strconv"
	"strings"

	"github.com/alvin-rw/aoc/internal/file"
	"github.com/alvin-rw/aoc/internal/matrix"
)

func main() {

}

func checkTrail(mapMatrix [][]int, startingPoint []int, startingSlope int, directions []int) []int {
	maxRow := len(mapMatrix)
	maxColumn := len(mapMatrix[0])

	row := startingPoint[0]
	column := startingPoint[1]

	if mapMatrix[row][column] == 9 {
		return []int{row, column}
	}

	for _, direction := range directions {
		nextRow := matrix.GetNextRow(row, direction, 1)
		nextColumn := matrix.GetNextColumn(column, direction, 1)

		if matrix.CheckCoordinateInsideMatrix([]int{nextRow, nextColumn}, maxRow, maxColumn) {
			if mapMatrix[nextRow][nextColumn] == startingSlope+1 {
				if coord := checkTrail(mapMatrix, []int{nextRow, nextColumn}, mapMatrix[nextRow][nextColumn], directions); coord != nil {
					return coord
				}
			}
		}
	}

	return nil
}

func getTrailheadIndexes(mapMatrix [][]int) [][]int {
	trailheadIndexes := [][]int{}

	for row, line := range mapMatrix {
		for column, num := range line {
			if num == 0 {
				trailheadIndexes = append(trailheadIndexes, []int{row, column})
			}
		}
	}

	return trailheadIndexes
}

func getMapMatrix(inputFilePath string) [][]int {
	mapMatrix := [][]int{}

	fileContent := file.ReadFile(inputFilePath)

	for _, line := range fileContent {
		lineSlice := strings.Split(line, "")

		intSlice := []int{}
		for _, e := range lineSlice {
			n, _ := strconv.Atoi(e)

			intSlice = append(intSlice, n)
		}

		mapMatrix = append(mapMatrix, intSlice)
	}

	return mapMatrix
}

/**
Search trailhead recursion

Base case: (returns true if peak, returns false if terminated early)
	- if it's peak (9)
	- if there's no number exactly 1 higher than the current number (after using all the available directions)

Recursion case:
	- if we're not in peak and if we're in the number 1 higher than the previous number

Input:
	-
**/
