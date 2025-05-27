package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/alvin-rw/aoc/internal/file"
	"github.com/alvin-rw/aoc/internal/matrix"
)

func main() {
	mapMatrix := getMapMatrix("./input.txt")

	trailheadIndexes := getTrailheadIndexes(mapMatrix)
	trailheadPeaksMap := getTrailheadsPeaksMap(mapMatrix, trailheadIndexes)
	totalTrailheadScore := calculateTrailheadScore(trailheadPeaksMap)

	fmt.Printf("total trailhead score: %d\n", totalTrailheadScore)
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

func getTrailheadsPeaksMap(mapMatrix [][]int, trailheadIndexes [][]int) map[string][]string {
	trailheadPeaksMap := make(map[string][]string)

	directions := []int{matrix.Up, matrix.Right, matrix.Down, matrix.Left}
	for _, trailheadIndex := range trailheadIndexes {
		trailheadStartingSlope := 0

		checkTrail(mapMatrix, trailheadIndex, trailheadStartingSlope, directions, trailheadPeaksMap, trailheadIndex)
	}

	return trailheadPeaksMap
}

func checkTrail(mapMatrix [][]int, startingPoint []int, startingSlope int, directions []int, trailPeaksList map[string][]string, trailheadIndex []int) {
	row := startingPoint[0]
	column := startingPoint[1]

	if mapMatrix[row][column] == 9 {
		trailheadIndexCode := fmt.Sprintf("%v", trailheadIndex)
		peakIndexCode := fmt.Sprintf("%v", []int{row, column})

		if peaks, ok := trailPeaksList[trailheadIndexCode]; ok {
			if !slices.Contains(peaks, peakIndexCode) {
				trailPeaksList[trailheadIndexCode] = append(trailPeaksList[trailheadIndexCode], peakIndexCode)
			}
		} else {
			trailPeaksList[trailheadIndexCode] = []string{peakIndexCode}
		}
		return
	}

	maxRow := len(mapMatrix)
	maxColumn := len(mapMatrix[0])

	for _, direction := range directions {
		nextRow := matrix.GetNextRow(row, direction, 1)
		nextColumn := matrix.GetNextColumn(column, direction, 1)

		if matrix.CheckCoordinateInsideMatrix([]int{nextRow, nextColumn}, maxRow, maxColumn) {
			if mapMatrix[nextRow][nextColumn] == startingSlope+1 {
				checkTrail(mapMatrix, []int{nextRow, nextColumn}, mapMatrix[nextRow][nextColumn], directions, trailPeaksList, trailheadIndex)
			}
		}
	}
}

func calculateTrailheadScore(trailheadPeaksMap map[string][]string) int {
	totalTrailheadScore := 0
	for _, achievablePeaks := range trailheadPeaksMap {
		totalTrailheadScore += len(achievablePeaks)
	}

	return totalTrailheadScore
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

**/
