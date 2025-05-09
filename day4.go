package aoc2024

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

const (
	up = iota
	down
	right
	left
	upRight
	upLeft
	downRight
	downLeft
)

func CalculateNumberOfXMAS(inputFilePath string) (int, error) {
	inputMatrix, err := getInputMatrix(inputFilePath)
	if err != nil {
		return -1, err
	}

	xmas := []string{"X", "M", "A", "S"}

	xIndexes := findAllStartingCharOccurenceInMatrix(inputMatrix, xmas[0])

	numberOfXMAS := findWordFromIndex(xIndexes, inputMatrix, xmas)

	return numberOfXMAS, nil
}

func CalculateNumberOfX_MAS(inputFilePath string) (int, error) {
	inputMatrix, err := getInputMatrix(inputFilePath)
	if err != nil {
		return -1, err
	}

	mas := []string{"M", "A", "S"}

	aIndexes := findAllStartingCharOccurenceInMatrix(inputMatrix, "A")

	numberOfX_MAS := findX_MASFromCenterIndex(aIndexes, inputMatrix, mas)

	return numberOfX_MAS, nil
}

func getInputMatrix(inputFilePath string) ([][]string, error) {
	f, err := os.Open(inputFilePath)
	if err != nil {
		return nil, fmt.Errorf("error when opening the file, %v", err)
	}
	defer f.Close()

	matrix := [][]string{}

	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			} else {
				return nil, fmt.Errorf("error when reading file, %v", err)
			}
		}

		line = strings.TrimSuffix(line, "\n")
		lineSlice := strings.Split(line, "")

		matrix = append(matrix, lineSlice)
	}

	return matrix, nil
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

	searchDirections := []int{up, down, right, left, upRight, upLeft, downRight, downLeft}

	for startingCharIndexRow, startingCharIndex := range startingCharOccurenceIndexes {
		row := startingCharIndexRow
		for _, startingCharIndexColumn := range startingCharIndex {
			column := startingCharIndexColumn

			for _, direction := range searchDirections {

			wordSearch:
				for i := 1; i < wordLength; i++ {
					nextColumn := getNextColumn(column, direction, i)
					nextRow := getNextRow(row, direction, i)

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
		{upRight, downLeft},
		{upLeft, downRight},
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
						nextColumn := getNextColumn(column, direction, i)
						nextRow := getNextRow(row, direction, i)

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

func getNextColumn(column int, direction int, modifier int) int {
	nextColumn := column

	switch direction {
	case up:
		nextColumn = column
	case down:
		nextColumn = column
	case right:
		nextColumn = column + modifier
	case left:
		nextColumn = column - modifier
	case upRight:
		nextColumn = column + modifier
	case upLeft:
		nextColumn = column - modifier
	case downRight:
		nextColumn = column + modifier
	case downLeft:
		nextColumn = column - modifier
	}

	return nextColumn
}

func getNextRow(row int, direction int, modifier int) int {
	nextRow := row

	switch direction {
	case up:
		nextRow = row - modifier
	case down:
		nextRow = row + modifier
	case right:
		nextRow = row
	case left:
		nextRow = row
	case upRight:
		nextRow = row - modifier
	case upLeft:
		nextRow = row - modifier
	case downRight:
		nextRow = row + modifier
	case downLeft:
		nextRow = row + modifier
	}

	return nextRow
}
