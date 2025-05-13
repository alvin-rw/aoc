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
	mapMatrix := createMapMatrix("./input.txt")

	guardSymbol := "^"
	guardFirstPosition, err := getGuardFirstPosition(mapMatrix, guardSymbol)
	if err != nil {
		log.Fatalf("error when getting guard first position %v", err)
	}

	numberOfDistinctPosition, guardVisitedPoints := calculateAndListDistinctPosition(mapMatrix, guardFirstPosition)

	fmt.Printf("guard's number of distinct position: %d\n", numberOfDistinctPosition)

	numOfPossibleObstruction := calculateNumberOfPossibleObstructions(mapMatrix, guardFirstPosition, guardVisitedPoints)

	fmt.Printf("number of possible obstruction to cause a loop: %d\n", numOfPossibleObstruction)
}

func calculateNumberOfPossibleObstructions(mapMatrix [][]string, guardFirstPosition []int, guardVisitedPoints [][]int) int {
	numberOfPossibleObstructionLocation := 0

	maxRow := len(mapMatrix)
	maxColumn := func(row int) int {
		return len(mapMatrix[row])
	}

	for _, pointsVisited := range guardVisitedPoints {
		direction := matrix.Up

		obstructionRow := pointsVisited[0]
		obstructionColumn := pointsVisited[1]

		// make a deep copy of the mapMatrix and put the obstruction in the map
		tempMapMatrix := make([][]string, len(mapMatrix))
		for i := range mapMatrix {
			tempMapMatrix[i] = slices.Clone(mapMatrix[i])
		}

		tempMapMatrix[obstructionRow][obstructionColumn] = "#"

		// record all the coordinates that the guard used to turn
		// in row,column,direction string format
		// this is used for determining if this route is a loop
		turningPoints := []string{}

		row := guardFirstPosition[0]
		column := guardFirstPosition[1]

		for {
			nextRow := matrix.GetNextRow(row, direction, 1)
			nextColumn := matrix.GetNextColumn(column, direction, 1)

			if nextColumn >= 0 && nextRow >= 0 && nextRow < maxRow && nextColumn < maxColumn(nextRow) {
				if tempMapMatrix[nextRow][nextColumn] == "#" {
					turningPoint := fmt.Sprintf("%d,%d,%d", row, column, direction)

					direction = matrix.ChangeDirection90Degree(direction)

					if slices.Contains(turningPoints, turningPoint) {
						numberOfPossibleObstructionLocation++

						break
					} else {
						turningPoints = append(turningPoints, turningPoint)
					}
				} else {
					row = nextRow
					column = nextColumn
				}
			} else {
				break
			}
		}
	}

	return numberOfPossibleObstructionLocation
}

func calculateAndListDistinctPosition(mapMatrix [][]string, guardFirstPosition []int) (int, [][]int) {
	numberOfDistinctPosition := 1
	direction := matrix.Up

	maxRow := len(mapMatrix)
	maxColumn := func(row int) int {
		return len(mapMatrix[row])
	}

	row := guardFirstPosition[0]
	column := guardFirstPosition[1]

	// list the points visited by guard, excluding the first position
	pointsVisited := [][]int{}

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

					pointsVisited = append(pointsVisited, []int{row, column})

					mapMatrix[nextRow][nextColumn] = "X"
				}
			}
		} else {
			break
		}
	}

	return numberOfDistinctPosition, pointsVisited
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

func createMapMatrix(inputFilePath string) [][]string {
	fileContent := file.ReadFile(inputFilePath)

	mapMatrix := [][]string{}
	for _, line := range fileContent {
		mapLevel := strings.Split(line, "")

		mapMatrix = append(mapMatrix, mapLevel)
	}

	return mapMatrix
}
