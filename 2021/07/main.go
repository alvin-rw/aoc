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
	fuelConsumedPartOne := getLowestFuel(crabLocations, lowest, highest, 1)
	fuelConsumedPartTwo := getLowestFuel(crabLocations, lowest, highest, 2)

	fmt.Printf("lowest possible fuel consumed part one: %d\n", fuelConsumedPartOne)
	fmt.Printf("lowest possible fuel consumed part two: %d\n", fuelConsumedPartTwo)
}

func getLowestFuel(crabLocations map[int]int, lowest, highest int, part int) int {
	fuelConsumed := 0

	for i := lowest; i <= highest; i++ {
		fuelConsumedForCurrentIteration := 0
		for position, numOfCrabs := range crabLocations {
			switch part {
			case 1:
				fuelConsumedForCurrentIteration += maths.Abs(position-i) * numOfCrabs
			case 2:
				fuelConsumedForCurrentIteration += getArithmeticFuelCost(maths.Abs(position-i)) * numOfCrabs
			}
		}

		if i == lowest {
			fuelConsumed = fuelConsumedForCurrentIteration
		} else if fuelConsumedForCurrentIteration < fuelConsumed {
			fuelConsumed = fuelConsumedForCurrentIteration
		}
	}

	return fuelConsumed
}

func getArithmeticFuelCost(distance int) int {
	switch distance {
	case 1:
		return 1
	case 0:
		return 0
	}

	return distance + getArithmeticFuelCost(distance-1)
	// or we can use aritmethic function
	// sum = n(n-1)/2
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
