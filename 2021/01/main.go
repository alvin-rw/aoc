package main

import (
	"fmt"
	"strconv"

	"github.com/alvin-rw/aoc/internal/file"
)

func main() {
	measurements := getMeasurements("input.txt")
	measurementsSlidingWindow := getMeasurementSlidingWindow(measurements)

	numOfLargerMeasurements := getNumOfLargerMeasurements(measurements)
	numOfLargerMeasurementsSlidingWindow := getNumOfLargerMeasurements(measurementsSlidingWindow)

	fmt.Printf("number of larger measurements %d\n", numOfLargerMeasurements)
	fmt.Printf("number of larger measurements in the sliding window %d\n", numOfLargerMeasurementsSlidingWindow)
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

func getMeasurementSlidingWindow(measurements []int) []int {
	measurementsSlidingWindow := make([]int, len(measurements)-2)

	for i := range measurementsSlidingWindow {
		measurementsSlidingWindow[i] = measurements[i] + measurements[i+1] + measurements[i+2]
	}

	return measurementsSlidingWindow
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
