package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/alvin-rw/aoc/internal/file"
	"github.com/alvin-rw/aoc/internal/matrix"
)

type field struct {
	area      int
	perimeter int
}

func main() {
	mapMatrix := getMapMatrix("input.txt")

	fencePrice := calculateFencePrice(mapMatrix)

	fmt.Printf("total fence price: %d\n", fencePrice)
}

func calculateFencePrice(mapMatrix [][]string) int {
	totalPrice := 0

	visited := make(map[string]struct{})

	numOfRows := len(mapMatrix)
	numOfColumns := len(mapMatrix[0])

	directionsToCheck := []int{matrix.Up, matrix.Right, matrix.Down, matrix.Left}
	for row, line := range mapMatrix {
		for column, plant := range line {
			if _, ok := visited[matrix.CoordToString(row, column)]; !ok {
				visited[matrix.CoordToString(row, column)] = struct{}{}

				toVisit := []string{}

				currentField := &field{
					area:      1,
					perimeter: 0,
				}

				for _, direction := range directionsToCheck {
					checkCell(visited, mapMatrix, row, column, direction, &toVisit, currentField, numOfRows, numOfColumns, plant)
				}

				for {
					if len(toVisit) == 0 {
						break
					}

					rowToVisit, columnToVisit := matrix.StringToCoord(toVisit[0])
					currentField.area++
					visited[matrix.CoordToString(rowToVisit, columnToVisit)] = struct{}{}
					toVisit = slices.Delete(toVisit, 0, 1)

					for _, direction := range directionsToCheck {
						checkCell(visited, mapMatrix, rowToVisit, columnToVisit, direction, &toVisit, currentField, numOfRows, numOfColumns, plant)
					}
				}

				price := currentField.area * currentField.perimeter
				totalPrice += price
			}
		}
	}

	return totalPrice
}

func checkCell(visited map[string]struct{}, mapMatrix [][]string, row, col, direction int, toVisit *[]string, currentField *field, maxRow int, maxCol int, plant string) {
	nextRow := matrix.GetNextRow(row, direction, 1)
	nextColumn := matrix.GetNextColumn(col, direction, 1)

	if matrix.CheckCoordinateInsideMatrix([]int{nextRow, nextColumn}, maxRow, maxCol) {
		if mapMatrix[nextRow][nextColumn] == plant {
			if _, ok := visited[matrix.CoordToString(nextRow, nextColumn)]; !ok {
				if !slices.Contains(*toVisit, matrix.CoordToString(nextRow, nextColumn)) {
					*toVisit = append(*toVisit, matrix.CoordToString(nextRow, nextColumn))
				}
			}
		} else {
			currentField.perimeter++
		}
	} else {
		currentField.perimeter++
	}
}

func getMapMatrix(inputFilePath string) [][]string {
	mapMatrix := [][]string{}

	fileContent := file.ReadFile(inputFilePath)

	for _, line := range fileContent {
		row := strings.Split(line, "")

		mapMatrix = append(mapMatrix, row)
	}

	return mapMatrix
}

/**
recursion

base case
no cell in the list to check

recursive case
there's still cell in the list to check
**/
