package aoc2024

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
)

var (
	numbers = []rune{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
)

// correct mul() instruction format
const (
	invalid = iota
	m
	u
	l
	lPharenthesis
	lNumFirstChar
	lNumSecondChar
	lNumThirdChar
	comma
	rNumFirstChar
	rNumSecondChar
	rNumThirdChar
	rPharenthesis
)

func CalculateMultiplicationResult(inputFilePath string) (int, error) {
	numbersToMultiply, err := readNumbersToMultiplyFromFile(inputFilePath)
	if err != nil {
		return -1, err
	}

	result := multiply(numbersToMultiply)
	return result, nil
}

func readNumbersToMultiplyFromFile(inputFilePath string) ([][]int, error) {
	f, err := os.Open(inputFilePath)
	if err != nil {
		return nil, fmt.Errorf("error when opening the file, %v", err)
	}
	defer f.Close()

	// numbersToMultiply stores the numbers that needs to be multiplied from all the mul() instructions
	numbersToMultiply := [][]int{}

	validMulCounter := invalid

	// tempNumbers stores the first and second numbers obtained from reading the files in a string format
	// these will be converted when we reached ')' character and appened to numbersToMultiply slice
	tempFirstNumber := ""
	tempSecondNumber := ""

	updateCurrentMul := func(valid bool, nextValue int) {
		if valid {
			validMulCounter = nextValue

			if validMulCounter == rPharenthesis {
				currentFirstNumber, _ := strconv.Atoi(tempFirstNumber)
				currentSecondNumber, _ := strconv.Atoi(tempSecondNumber)

				numbersToMultiply = append(numbersToMultiply, []int{currentFirstNumber, currentSecondNumber})

				// clear all the counter and temp numbers after adding it to the numbersToMultiply slice
				validMulCounter = invalid
				tempFirstNumber = ""
				tempSecondNumber = ""
			}
		} else {
			validMulCounter = invalid
			tempFirstNumber = ""
			tempSecondNumber = ""
		}
	}

	reader := bufio.NewReader(f)
	for {
		char, _, err := reader.ReadRune()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			} else {
				return nil, fmt.Errorf("error when reading file, %v", err)
			}
		}

		switch {
		case char == 'm':
			updateCurrentMul(true, m)

		case char == 'u':
			updateCurrentMul(validMulCounter == m, u)

		case char == 'l':
			updateCurrentMul(validMulCounter == u, l)

		case char == '(':
			updateCurrentMul(validMulCounter == l, lPharenthesis)

		case char == ',':
			updateCurrentMul(validMulCounter >= lNumFirstChar && validMulCounter < comma, comma)

		case char == ')':
			updateCurrentMul(validMulCounter >= rNumFirstChar && validMulCounter < rPharenthesis, rPharenthesis)

		case slices.Contains(numbers, char):
			if validMulCounter >= lPharenthesis && validMulCounter < lNumThirdChar {
				validMulCounter++

				tempFirstNumber = tempFirstNumber + string(char)
			} else if validMulCounter >= comma && validMulCounter < rNumThirdChar {
				validMulCounter++

				tempSecondNumber = tempSecondNumber + string(char)
			} else {
				updateCurrentMul(false, invalid)
			}

		default:
			updateCurrentMul(false, invalid)
		}
	}

	return numbersToMultiply, nil
}

func multiply(numbers [][]int) int {
	result := 0

	for _, nums := range numbers {
		result = result + (nums[0] * nums[1])
	}

	return result
}
