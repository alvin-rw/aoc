package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/alvin-rw/aoc/internal/file"
)

func main() {
	calibrationElementsList, err := getCalibrationElementsList("./input.txt")
	if err != nil {
		log.Fatalf("error when getting calibration elements: %v", err)
	}

	partOperator := [][]string{
		{"+", "*"},
		{"+", "*", "||"},
	}

	for i, operators := range partOperator {
		result := 0

		for _, calibrationElements := range calibrationElementsList {
			desiredResult := calibrationElements[0]
			input := calibrationElements[1:]

			if calibrate(input, desiredResult, operators) {
				result += desiredResult
			}
		}

		fmt.Printf("total calibration result part %d: %d\n", i+1, result)
	}
}

func calibrate(input []int, desiredResult int, operators []string) bool {
	if len(input) == 1 {
		return input[0] == desiredResult
	}

	for _, o := range operators {
		calculated := calculateResult(input, o)
		if calibrate(calculated, desiredResult, operators) {
			return true
		}
	}

	return false
}

func calculateResult(input []int, operator string) []int {
	if len(input) >= 2 {
		r := 0
		switch operator {
		case "+":
			r = input[0] + input[1]
		case "*":
			r = input[0] * input[1]
		case "||":
			result := fmt.Sprintf("%d%d", input[0], input[1])
			r, _ = strconv.Atoi(result)
		}

		calculated := []int{r}
		if len(input) > 2 {
			calculated = append(calculated, input[2:]...)
		}

		return calculated
	} else {
		return input
	}
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
