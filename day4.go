package aoc2024

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
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
					nextColumn := column
					nextRow := row

					switch direction {
					case up:
						nextColumn = column
						nextRow = row - i

					case down:
						nextColumn = column
						nextRow = row + i

					case right:
						nextColumn = column + i
						nextRow = row

					case left:
						nextColumn = column - i
						nextRow = row

					case upRight:
						nextColumn = column + i
						nextRow = row - i

					case upLeft:
						nextColumn = column - i
						nextRow = row - i

					case downRight:
						nextColumn = column + i
						nextRow = row + i

					case downLeft:
						nextColumn = column - i
						nextRow = row + i
					}

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
