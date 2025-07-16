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

	totalTrailheadRatings := new(int)

	trailheadIndexes := getTrailheadIndexes(mapMatrix)
	trailheadPeakMap := getTrailheadsPeakMapAndRatings(mapMatrix, trailheadIndexes, totalTrailheadRatings)
	totalTrailheadScore := calculateTrailheadScore(trailheadPeakMap)

	fmt.Printf("total trailhead score: %d\n", totalTrailheadScore)
	fmt.Printf("total trailhead ratings: %d\n", *totalTrailheadRatings)
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

func getTrailheadsPeakMapAndRatings(mapMatrix [][]int, trailheadIndexes [][]int, totalTrailheadRatings *int) map[string][]string {
	trailheadPeakMap := make(map[string][]string)

	directions := []matrix.Direction{matrix.Up, matrix.Right, matrix.Down, matrix.Left}
	for _, trailheadIndex := range trailheadIndexes {
		trailheadStartingSlope := 0

		checkTrail(mapMatrix, trailheadIndex, trailheadStartingSlope, directions, trailheadPeakMap, trailheadIndex, totalTrailheadRatings)
	}

	return trailheadPeakMap
}

func checkTrail(mapMatrix [][]int, startingPoint []int, startingSlope int, directions []matrix.Direction, trailheadPeakMap map[string][]string, trailheadIndex []int, totalTrailheadRatings *int) {
	row := startingPoint[0]
	column := startingPoint[1]

	if mapMatrix[row][column] == 9 {
		trailheadIndexCode := fmt.Sprintf("%v", trailheadIndex)
		peakIndexCode := fmt.Sprintf("%v", []int{row, column})

		*totalTrailheadRatings = *totalTrailheadRatings + 1

		if peaks, ok := trailheadPeakMap[trailheadIndexCode]; ok {
			if !slices.Contains(peaks, peakIndexCode) {
				trailheadPeakMap[trailheadIndexCode] = append(trailheadPeakMap[trailheadIndexCode], peakIndexCode)
			}
		} else {
			trailheadPeakMap[trailheadIndexCode] = []string{peakIndexCode}
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
				checkTrail(mapMatrix, []int{nextRow, nextColumn}, mapMatrix[nextRow][nextColumn], directions, trailheadPeakMap, trailheadIndex, totalTrailheadRatings)
			}
		}
	}
}

func calculateTrailheadScore(trailheadPeakMap map[string][]string) int {
	totalTrailheadScore := 0
	for _, achievablePeaks := range trailheadPeakMap {
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
