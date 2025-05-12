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
	numberOfXMAS, err := calculateNumberOfXMAS("./input.txt")
	if err != nil {
		log.Fatalf("error when calculating number of XMAS: %v", err)
	}

	numberOfX_MAS, err := calculateNumberOfX_MAS("./input.txt")
	if err != nil {
		log.Fatalf("error when calculating number of X-MAS: %v", err)
	}

	fmt.Printf("number of XMAS: %d\n", numberOfXMAS)
	fmt.Printf("number of X-MAS: %d\n", numberOfX_MAS)
}

func calculateNumberOfXMAS(inputFilePath string) (int, error) {
	inputMatrix, err := getInputMatrix(inputFilePath)
	if err != nil {
		return -1, err
	}

	xmas := []string{"X", "M", "A", "S"}

	xIndexes := findAllStartingCharOccurenceInMatrix(inputMatrix, xmas[0])

	numberOfXMAS := findWordFromIndex(xIndexes, inputMatrix, xmas)

	return numberOfXMAS, nil
}

func calculateNumberOfX_MAS(inputFilePath string) (int, error) {
	inputMatrix, err := getInputMatrix(inputFilePath)
	if err != nil {
		return -1, err
	}

	mas := []string{"M", "A", "S"}

	aIndexes := findAllStartingCharOccurenceInMatrix(inputMatrix, "A")

	numberOfX_MAS := findX_MASFromCenterIndex(aIndexes, inputMatrix, mas)

	return numberOfX_MAS, nil
}

func findAllStartingCharOccurenceInMatrix(matrix [][]string, char string) [][]int {
	indexes := [][]int{}

	for _, list := range matrix {
		currentLevelIndexes := []int{}

		for i, e := range list {
			if e == char {
				currentLevelIndexes = append(currentLevelIndexes, i)
			}
		}

		indexes = append(indexes, currentLevelIndexes)
	}

	return indexes
}

func findWordFromIndex(startingCharOccurenceIndexes [][]int, inputMatrix [][]string, word []string) int {
	numberOfWords := 0

	maxRow := len(inputMatrix)
	maxColumn := func(row int) int {
		return len(inputMatrix[row])
	}

	wordLength := len(word)

	searchDirections := []int{matrix.Up, matrix.Down, matrix.Right, matrix.Left, matrix.UpRight, matrix.UpLeft, matrix.DownRight, matrix.DownLeft}

	for startingCharIndexRow, startingCharIndex := range startingCharOccurenceIndexes {
		row := startingCharIndexRow
		for _, startingCharIndexColumn := range startingCharIndex {
			column := startingCharIndexColumn

			for _, direction := range searchDirections {

			wordSearch:
				for i := 1; i < wordLength; i++ {
					nextColumn := matrix.GetNextColumn(column, direction, i)
					nextRow := matrix.GetNextRow(row, direction, i)

					if nextColumn >= 0 && nextRow >= 0 && nextRow < maxRow && nextColumn < maxColumn(nextRow) {
						if inputMatrix[nextRow][nextColumn] == word[i] {
							if i == wordLength-1 {
								numberOfWords++
							} else {
								continue
							}
						} else {
							break wordSearch
						}
					} else {
						break wordSearch
					}
				}
			}
		}
	}

	return numberOfWords
}

func findX_MASFromCenterIndex(centerCharOccurenceIndexes [][]int, inputMatrix [][]string, word []string) int {
	numberOfWords := 0

	maxRow := len(inputMatrix)
	maxColumn := func(row int) int {
		return len(inputMatrix[row])
	}

	halfWordLength := len(word) / 2

	searchDirectionGroup := [][]int{
		{matrix.UpRight, matrix.DownLeft},
		{matrix.UpLeft, matrix.DownRight},
	}

	for startingCharIndexRow, startingCharIndex := range centerCharOccurenceIndexes {
		row := startingCharIndexRow
		for _, startingCharIndexColumn := range startingCharIndex {
			column := startingCharIndexColumn

			wordFound := 0

			for _, searchDirections := range searchDirectionGroup {
				// get the word without its center character
				wordWithoutCenterChar := slices.Delete(slices.Clone(word), halfWordLength, halfWordLength+1)

				for _, direction := range searchDirections {
				wordSearch:
					for i := 1; i <= halfWordLength; i++ {
						nextColumn := matrix.GetNextColumn(column, direction, i)
						nextRow := matrix.GetNextRow(row, direction, i)

						if nextColumn >= 0 && nextRow >= 0 && nextRow < maxRow && nextColumn < maxColumn(nextRow) {
							if slices.Contains(wordWithoutCenterChar, inputMatrix[nextRow][nextColumn]) {
								k := slices.Index(wordWithoutCenterChar, inputMatrix[nextRow][nextColumn])
								wordWithoutCenterChar = slices.Delete(wordWithoutCenterChar, k, k+1)

								if len(wordWithoutCenterChar) == 0 {
									wordFound++
								}
							} else {
								break wordSearch
							}
						} else {
							break wordSearch
						}
					}
				}
			}

			if wordFound == 2 {
				numberOfWords++
			}
		}
	}

	return numberOfWords
}

func getInputMatrix(inputFilePath string) ([][]string, error) {
	fileContent, err := file.ReadFile(inputFilePath)
	if err != nil {
		return nil, err
	}

	matrix := [][]string{}

	for _, line := range fileContent {
		charactersList := strings.Split(line, "")

		matrix = append(matrix, charactersList)
	}

	return matrix, nil
}
