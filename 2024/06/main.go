package main

import (
	"fmt"
	"log"
	"slices"
	"strings"

	"github.com/alvin-rw/aoc/internal/file"
	"github.com/alvin-rw/aoc/internal/matrix"
)

func main() {
	guardsNumberOfDistinctPosition, err := getGuardsNumberOfDistinctPosition("./input.txt")
	if err != nil {
		log.Fatalf("error when calculating guard's number of distinct position: %v", err)
	}

	fmt.Printf("guard's number of distinct position: %d\n", guardsNumberOfDistinctPosition)
}

func getGuardsNumberOfDistinctPosition(inputFilePath string) (int, error) {
	mapMatrix, err := createMapMatrix(inputFilePath)
	if err != nil {
		return -1, err
	}

	guardSymbol := "^"
	guardFirstPosition, err := getGuardFirstPosition(mapMatrix, guardSymbol)
	if err != nil {
		return -1, err
	}

	numberOfDistinctPosition := calculateDistinctPosition(mapMatrix, guardFirstPosition)

	return numberOfDistinctPosition, nil
}

func calculateDistinctPosition(mapMatrix [][]string, guardFirstPosition []int) int {
	numberOfDistinctPosition := 1
	direction := matrix.Up

	maxRow := len(mapMatrix)
	maxColumn := func(row int) int {
		return len(mapMatrix[row])
	}

	row := guardFirstPosition[0]
	column := guardFirstPosition[1]

	// change guard first position to X
	mapMatrix[row][column] = "X"

	for {
		nextRow := matrix.GetNextRow(row, direction, 1)
		nextColumn := matrix.GetNextColumn(column, direction, 1)

		if nextColumn >= 0 && nextRow >= 0 && nextRow < maxRow && nextColumn < maxColumn(nextRow) {
			if mapMatrix[nextRow][nextColumn] == "#" {
				direction = matrix.ChangeDirection90Degree(direction)
			} else {
				row = nextRow
				column = nextColumn

				if mapMatrix[nextRow][nextColumn] == "." {
					numberOfDistinctPosition++

					mapMatrix[nextRow][nextColumn] = "X"
				}
			}
		} else {
			break
		}
	}

	return numberOfDistinctPosition
}

func getGuardFirstPosition(mapMatrix [][]string, guardSymbol string) ([]int, error) {
	for row, level := range mapMatrix {
		if slices.Contains(level, guardSymbol) {
			column := slices.Index(level, guardSymbol)
			return []int{row, column}, nil
		}
	}

	return nil, fmt.Errorf("cannot find guard in the map matrix")
}

func createMapMatrix(inputFilePath string) ([][]string, error) {
	fileContent, err := file.ReadFile(inputFilePath)
	if err != nil {
		return nil, err
	}

	mapMatrix := [][]string{}
	for _, line := range fileContent {
		mapLevel := strings.Split(line, "")

		mapMatrix = append(mapMatrix, mapLevel)
	}

	return mapMatrix, nil
}
