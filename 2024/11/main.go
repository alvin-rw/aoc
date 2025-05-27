package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/alvin-rw/aoc/internal/file"
	"github.com/alvin-rw/aoc/internal/maths"
)

func main() {
	stones := getInputStones("./input.txt")
	for range 25 {
		stones = blink(stones)
	}

	fmt.Printf("number of stones after blinking 25 times: %d\n", len(stones))
}

func blink(stones []int) []int {
	for i := 0; i < len(stones); i++ {
		if stones[i] == 0 {
			stones[i] = 1
		} else if maths.GetNumberOfDigits(stones[i])%2 == 0 {
			first, second := maths.SplitNumberIntoTwo(stones[i])

			stones[i] = first
			stones = slices.Insert(stones, i+1, second)
			i++
		} else {
			stones[i] = stones[i] * 2024
		}
	}
	return stones
}

func getInputStones(inputFilePath string) []int {
	stones := []int{}

	fileContent := file.ReadFile(inputFilePath)

	for _, line := range fileContent {
		lineSlice := strings.Split(line, " ")

		for _, s := range lineSlice {
			stone, _ := strconv.Atoi(s)

			stones = append(stones, stone)
		}
	}

	return stones
}
