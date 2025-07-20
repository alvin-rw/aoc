package main

import (
	"fmt"
	"strconv"

	"github.com/alvin-rw/aoc/internal/file"
)

func main() {
	measurements := getMeasurements("input.txt")

	numOfLargerMeasurements := getNumOfLargerMeasurements(measurements)

	fmt.Printf("number of larger measurements %d\n", numOfLargerMeasurements)
}

func getNumOfLargerMeasurements(measurements []int) int {
	numOfLargerMeasurements := 0

	for i := 1; i < len(measurements); i++ {
		if measurements[i] > measurements[i-1] {
			numOfLargerMeasurements++
		}
	}

	return numOfLargerMeasurements
}

func getMeasurements(inputFilePath string) []int {
	measurements := []int{}
	lines := file.ReadFile(inputFilePath)

	for _, line := range lines {
		n, _ := strconv.Atoi(line)
		measurements = append(measurements, n)
	}

	return measurements
}
