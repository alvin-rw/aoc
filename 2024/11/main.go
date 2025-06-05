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
	blinkCounts := []int{25, 75}

	for _, bb := range blinkCounts {
		stones := getInputStones("./input.txt")
		count := getStoneCountAfterBlinks(bb, stones)

		fmt.Printf("stone count after blinking %d times: %d\n", bb, count)
	}
}

func getStoneCountAfterBlinks(blinkCount int, stones []int) int {
	finalStoneCount := 0

	stoneCountList := getStoneCountList(blinkCount)

	for i := range blinkCount {
		// find all 0 stones
		for {
			z := slices.Index(stones, 0)
			if z == -1 {
				break
			}
			stones = slices.Delete(stones, z, z+1)

			finalStoneCount += stoneCountList[blinkCount-i-1]
		}

		// find all 1 stones
		for {
			z := slices.Index(stones, 1)
			if z == -1 {
				break
			}
			stones = slices.Delete(stones, z, z+1)

			finalStoneCount += stoneCountList[blinkCount-i]
		}

		stones = blink(stones)
	}

	finalStoneCount += len(stones)

	return finalStoneCount
}

func getStoneCountList(blinkCount int) []int {
	stones := []int{1}
	stoneCountList := []int{1}

	zeroAndOneStoneCounter := map[int][]int{
		0: {},
		1: {},
	}

	for blinkCount := range blinkCount {
		fmt.Println(blinkCount)
		currentBlinkStoneCount := 0
		stones = blink(stones)

		for n, stoneBlinkStartingPointList := range zeroAndOneStoneCounter {
			for _, stoneBlinkStartingPoint := range stoneBlinkStartingPointList {
				switch n {
				case 0:
					if blinkCount-stoneBlinkStartingPoint < 0 {
						currentBlinkStoneCount += 1
					} else {
						currentBlinkStoneCount += stoneCountList[blinkCount-stoneBlinkStartingPoint-1]
					}
				case 1:
					currentBlinkStoneCount += stoneCountList[blinkCount-stoneBlinkStartingPoint]
				}
			}
		}

		currentBlinkStoneCount += len(stones)
		stoneCountList = append(stoneCountList, currentBlinkStoneCount)

		for n := range zeroAndOneStoneCounter {
			ii := slices.Index(stones, n)
			for ii != -1 {
				zeroAndOneStoneCounter[n] = append(zeroAndOneStoneCounter[n], blinkCount)
				stones = slices.Delete(stones, ii, ii+1)

				ii = slices.Index(stones, n)
			}
		}
	}

	return stoneCountList
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
