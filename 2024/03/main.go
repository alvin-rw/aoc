package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
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
	d
	o
	n
	singleQuote
	t
)

// command types
const (
	invalidCommand = iota
	mulCommand
	doCommand
	dontCommand
)

func main() {
	multiplicationResult, err := calculateMultiplicationResult("./input.txt", false)
	if err != nil {
		log.Fatalf("error when calculating multiplication result: %v", err)
	}

	multiplicationResultWithDoDonts, err := calculateMultiplicationResult("./input.txt", true)
	if err != nil {
		log.Fatalf("error when calculating multiplication result: %v", err)
	}

	fmt.Printf("multiplication result: %d\n", multiplicationResult)
	fmt.Printf("multiplication result with do and don'ts: %d\n", multiplicationResultWithDoDonts)
}

func calculateMultiplicationResult(inputFilePath string, enableDoDonts bool) (int, error) {
	numbersToMultiply, err := readNumbersToMultiplyFromFile(inputFilePath, enableDoDonts)
	if err != nil {
		return -1, err
	}

	result := multiply(numbersToMultiply)
	return result, nil
}

func readNumbersToMultiplyFromFile(inputFilePath string, enableDoDonts bool) ([][]int, error) {
	f, err := os.Open(inputFilePath)
	if err != nil {
		return nil, fmt.Errorf("error when opening the file, %v", err)
	}
	defer f.Close()

	// numbersToMultiply stores the numbers that needs to be multiplied from all the mul() instructions
	numbersToMultiply := [][]int{}

	// tempNumbers stores the first and second numbers obtained from reading the files in a string format
	// these will be converted when we reached ')' character and appened to numbersToMultiply slice
	tempFirstNumber := ""
	tempSecondNumber := ""

	enabled := true
	commandCounter := invalid
	commandType := invalidCommand

	cleanParameters := func() {
		commandCounter = invalid
		tempFirstNumber = ""
		tempSecondNumber = ""
		commandType = invalidCommand
	}

	updateCommandVerification := func(valid bool, nextValue int) {
		if valid {
			if commandCounter == l {
				commandType = mulCommand
			} else if commandCounter == o {
				commandType = doCommand
			} else if commandCounter == t {
				commandType = dontCommand
			}

			commandCounter = nextValue

			if commandCounter == rPharenthesis {
				switch commandType {
				case mulCommand:
					if enabled {
						currentFirstNumber, _ := strconv.Atoi(tempFirstNumber)
						currentSecondNumber, _ := strconv.Atoi(tempSecondNumber)

						numbersToMultiply = append(numbersToMultiply, []int{currentFirstNumber, currentSecondNumber})

						cleanParameters()
					}

				case doCommand:
					enabled = true
					cleanParameters()

				case dontCommand:
					if enableDoDonts {
						enabled = false
					}
					cleanParameters()
				}
			}
		} else {
			cleanParameters()
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
			updateCommandVerification(true, m)
		case char == 'u':
			updateCommandVerification(commandCounter == m, u)
		case char == 'l':
			updateCommandVerification(commandCounter == u, l)

		case char == '(':
			updateCommandVerification(commandCounter == l || commandCounter == o || commandCounter == t, lPharenthesis)
		case char == ',':
			updateCommandVerification(commandCounter >= lNumFirstChar && commandCounter < comma && commandType == mulCommand, comma)
		case char == ')':
			if commandType == mulCommand {
				updateCommandVerification(commandCounter >= rNumFirstChar && commandCounter < rPharenthesis, rPharenthesis)
			} else if commandType == doCommand || commandType == dontCommand {
				updateCommandVerification(commandCounter == lPharenthesis, rPharenthesis)
			}

		case char == 'd':
			updateCommandVerification(true, d)
		case char == 'o':
			updateCommandVerification(commandCounter == d, o)
		case char == 'n':
			updateCommandVerification(commandCounter == o, n)
		case char == '\'':
			updateCommandVerification(commandCounter == n, singleQuote)
		case char == 't':
			updateCommandVerification(commandCounter == singleQuote, t)

		case slices.Contains(numbers, char):
			if commandType == mulCommand {
				if commandCounter >= lPharenthesis && commandCounter < lNumThirdChar {
					commandCounter++

					tempFirstNumber = tempFirstNumber + string(char)
				} else if commandCounter >= comma && commandCounter < rNumThirdChar {
					commandCounter++

					tempSecondNumber = tempSecondNumber + string(char)
				} else {
					updateCommandVerification(false, invalid)
				}
			} else {
				updateCommandVerification(false, invalid)
			}

		default:
			updateCommandVerification(false, invalid)
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
