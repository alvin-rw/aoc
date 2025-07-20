package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/alvin-rw/aoc/internal/file"
	"github.com/alvin-rw/aoc/internal/maths"
)

func main() {
	crabLocations, lowest, highest := getCrabLocationsAndBoundaries("input.txt")
	fuelConsumed := getLowestFuel(crabLocations, lowest, highest)

	fmt.Printf("lowest possible fuel consumed %d\n", fuelConsumed)
}

func getLowestFuel(crabLocations map[int]int, lowest, highest int) int {
	fuelConsumed := 0

	for i := lowest; i <= highest; i++ {
		fuelConsumedForCurrentIteration := 0
		for position, numOfCrabs := range crabLocations {
			fuelConsumedForCurrentIteration += maths.Abs(position-i) * numOfCrabs
		}

		if i == lowest {
			fuelConsumed = fuelConsumedForCurrentIteration
		} else if fuelConsumedForCurrentIteration < fuelConsumed {
			fuelConsumed = fuelConsumedForCurrentIteration
		}
	}

	return fuelConsumed
}

// returns map[int]int of (map[crab location]number of crabs in that location)
// and lowest and highest location
func getCrabLocationsAndBoundaries(fileInputPath string) (map[int]int, int, int) {
	lines := file.ReadFile(fileInputPath)

	positions := strings.Split(lines[0], ",")

	// initialize lowest and highest value by using the first crab position
	lowest, _ := strconv.Atoi(positions[0])
	highest, _ := strconv.Atoi(positions[0])

	crabPositions := make(map[int]int)
	for _, pos := range positions {
		crabPosition, _ := strconv.Atoi(pos)

		if crabPosition < lowest {
			lowest = crabPosition
		}
		if crabPosition > highest {
			highest = crabPosition
		}

		crabPositions[crabPosition]++
	}

	return crabPositions, lowest, highest
}
