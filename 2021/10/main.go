package main

import (
	"fmt"
	"slices"

	"github.com/alvin-rw/aoc/internal/file"
)

var (
	illegalCharacterScore = map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}

	pairs = map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}
)

func main() {
	lines := file.ReadFile("input.txt")

	score := calculateSyntaxErrorScore(lines)
	fmt.Printf("syntax error score %d\n", score)
}

func calculateSyntaxErrorScore(lines []string) int {
	score := 0

	for _, line := range lines {
		expectedClosingBrackets := []rune{}

		for _, bracket := range line {
			if checkIfOpeningBracket(bracket) {
				expectedClosingBrackets = append(expectedClosingBrackets, pairs[bracket])
			} else {
				if bracket != expectedClosingBrackets[len(expectedClosingBrackets)-1] {
					score += illegalCharacterScore[bracket]
					break
				} else {
					expectedClosingBrackets = slices.Delete(expectedClosingBrackets, len(expectedClosingBrackets)-1, len(expectedClosingBrackets))
				}
			}
		}
	}

	return score
}

func checkIfOpeningBracket(b rune) bool {
	opening := []rune{'(', '[', '{', '<'}

	return slices.Contains(opening, b)
}
