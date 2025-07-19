package main

import (
	"fmt"
	"slices"

	"github.com/alvin-rw/aoc/internal/file"
)

func main() {
	polymer := getPolymer("input.txt")

	polymer = react(polymer, 0)

	fmt.Printf("units remaining after reacting %d\n", len(polymer))

	shortestPolymerLength := getShortestPolymerLength("input.txt", len(polymer))

	fmt.Printf("shortest polymer we can produce by removing 1 unit type %d\n", shortestPolymerLength)
}

func react(polymer []rune, index int) []rune {
	if index >= len(polymer)-1 {
		return polymer
	} else if index < 0 {
		index = 0
	}

	diff := 'a' - 'A'

	if polymer[index]-polymer[index+1] == diff || polymer[index]-polymer[index+1] == -diff {
		polymer = slices.Delete(polymer, index, index+2)

		return react(polymer, index-1)
	} else {
		return react(polymer, index+1)
	}
}

func getShortestPolymerLength(inputFilePath string, originalPolymerLength int) int {
	shortestPolymerLength := originalPolymerLength

	for i := 'a'; i <= 'z'; i++ {
		reducedPolymer := getPolymer(inputFilePath)

		unit := i
		reducedPolymer = removeUnit(reducedPolymer, unit)
		reducedPolymer = react(reducedPolymer, 0)

		if len(reducedPolymer) < shortestPolymerLength {
			shortestPolymerLength = len(reducedPolymer)
		}
	}

	return shortestPolymerLength
}

func removeUnit(polymer []rune, unit rune) []rune {
	diff := 'a' - 'A'
	result := []rune{}

	for _, r := range polymer {
		if r != unit && r != unit-diff && r != unit+diff {
			result = append(result, r)
		}
	}

	return result
}

func getPolymer(inputFilePath string) []rune {
	line := file.ReadFile(inputFilePath)

	polymer := []rune{}
	for _, r := range line[0] {
		polymer = append(polymer, r)
	}

	return polymer
}
