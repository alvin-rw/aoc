package main

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/alvin-rw/aoc/internal/file"
)

const (
	addition       = "+"
	multiplication = "*"
)

func main() {
	calibrationElementsList, err := getCalibrationElementsList("./input.txt")
	if err != nil {
		log.Fatalf("error when getting calibration elements: %v", err)
	}

	totalCalibrationResult := 0

	for _, calibrationElements := range calibrationElementsList {
		if res, ok := checkCalibrationResult(calibrationElements); ok {
			totalCalibrationResult += res
		}
	}

	fmt.Printf("total calibration result: %d\n", totalCalibrationResult)
}

func checkCalibrationResult(calibrationElements []int) (int, bool) {
	desiredResult := calibrationElements[0]

	// remove the calibration result from the slice
	calibrationInputs := slices.Delete(calibrationElements, 0, 1)

	if getHighestPossibleResult(calibrationInputs) < desiredResult {
		return 0, false
	}

	if tryCalibration(desiredResult, calibrationInputs) {
		return desiredResult, true
	}

	return 0, false
}

func getHighestPossibleResult(calibrationInputs []int) int {
	r := calibrationInputs[0]

	for i := 1; i < len(calibrationInputs); i++ {
		if calibrationInputs[i] > 1 {
			r = r * calibrationInputs[i]
		} else {
			r = r + calibrationInputs[i]
		}
	}

	return r
}

// tryCalibration will try to get the calibration result using all the available
// operators. Returns true if the calibration result can be obtained from the input
func tryCalibration(desiredResult int, inputs []int) bool {
	numberOfOperators := len(inputs) - 1

	for numberOfAdditions := range inputs {
		additionStartingIndexes := initializeAdditionStartingIndexesSlice(numberOfAdditions)

		for {
			// make an operators slice and fill it with multiplication
			operators := make([]string, numberOfOperators)
			for i := range operators {
				operators[i] = multiplication
			}

			addAdditionsToOperatorsSlice(operators, numberOfAdditions, additionStartingIndexes)

			if validateCalibrationOperators(desiredResult, inputs, operators) {
				return true
			}

			if ok := incrementAdditionStartingIndex(numberOfOperators-1, numberOfAdditions-1, additionStartingIndexes); !ok {
				break
			}
		}
	}

	return false
}

// Initializes the indexes of additions depending on the number of additions.
//
// First element will represent the index of the first addition, second element will represent the index
// of the second addition, etc.
//
// Example resulting slice:
//   - 1 addition: [0] -> +****
//   - 2 additions: [0 1] -> ++***
//   - 3 additions: [0 1 2] -> +++**
func initializeAdditionStartingIndexesSlice(numberOfAdditions int) []int {
	additionStartingIndexes := make([]int, numberOfAdditions)
	if numberOfAdditions > 0 {
		for i := range additionStartingIndexes {
			additionStartingIndexes[i] = i
		}
	}

	return additionStartingIndexes
}

func addAdditionsToOperatorsSlice(operatorsSlice []string, desiredNumberOfAdditions int, additionStartingIndexes []int) {
	for i := range desiredNumberOfAdditions {
		operatorsSlice[additionStartingIndexes[i]] = addition
	}
}

// validateCalibrationOperators will calculate the result from the inputs and the operators and
// compare it with desiredResult. Returns true if the calibration result can be obtained from the provided inputs and operators
func validateCalibrationOperators(desiredResult int, inputs []int, operators []string) bool {
	result := inputs[0]

	for i := 1; i < len(inputs); i++ {
		if operators[i-1] == addition {
			result = result + inputs[i]
		} else {
			result = result * inputs[i]
		}
	}

	return result == desiredResult
}

func incrementAdditionStartingIndex(maxIndex int, indexToIncrement int, additionStartingIndexes []int) bool {
	if indexToIncrement < 0 {
		return false
	}

	additionStartingIndexes[indexToIncrement]++

	if additionStartingIndexes[indexToIncrement] > maxIndex {
		if indexToIncrement == 0 {
			return false
		}

		if !incrementAdditionStartingIndex(maxIndex-1, indexToIncrement-1, additionStartingIndexes) {
			return false
		}

		additionStartingIndexes[indexToIncrement] = additionStartingIndexes[indexToIncrement-1] + 1
	}

	return true
}

// return the calibration elements (result and inputs) as slice of []int
func getCalibrationElementsList(inputFilePath string) ([][]int, error) {
	fileContent := file.ReadFile(inputFilePath)

	calibrationElementsList := [][]int{}

	for _, line := range fileContent {
		calibrationElements := []int{}

		resultAndInput := strings.Split(line, ": ")

		calibrationResult, err := strconv.Atoi(resultAndInput[0])
		if err != nil {
			return nil, err
		}

		calibrationElements = append(calibrationElements, calibrationResult)

		calibrationInputs := strings.Split(resultAndInput[1], " ")

		for _, input := range calibrationInputs {
			in, err := strconv.Atoi(input)
			if err != nil {
				return nil, err
			}

			calibrationElements = append(calibrationElements, in)
		}

		calibrationElementsList = append(calibrationElementsList, calibrationElements)
	}

	return calibrationElementsList, nil
}
