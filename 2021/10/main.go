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

	autoCompleteCharacterScore = map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
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

	corruptedScore, autoCompleteScore := calculateSyntaxErrorScore(lines)
	fmt.Printf("corrupted error score %d\n", corruptedScore)
	fmt.Printf("autocomplete error score %d\n", autoCompleteScore)
}

func calculateSyntaxErrorScore(lines []string) (int, int) {
	corruptedScore := 0
	autoCompleteScores := []int{}

	for _, line := range lines {
		expectedClosingBrackets := []rune{}
		corrupted := false

		for _, bracket := range line {
			if checkIfOpeningBracket(bracket) {
				expectedClosingBrackets = append(expectedClosingBrackets, pairs[bracket])
			} else {
				if bracket != expectedClosingBrackets[len(expectedClosingBrackets)-1] {
					corruptedScore += illegalCharacterScore[bracket]
					corrupted = true
					break
				} else {
					expectedClosingBrackets = slices.Delete(expectedClosingBrackets, len(expectedClosingBrackets)-1, len(expectedClosingBrackets))
				}
			}
		}

		if !corrupted {
			score := 0
			for _, bracket := range slices.Backward(expectedClosingBrackets) {
				score = score * 5
				score += autoCompleteCharacterScore[bracket]
			}
			autoCompleteScores = append(autoCompleteScores, score)
		}
	}

	slices.Sort(autoCompleteScores)
	autoCompleteScore := autoCompleteScores[len(autoCompleteScores)/2]

	return corruptedScore, autoCompleteScore
}

func checkIfOpeningBracket(b rune) bool {
	opening := []rune{'(', '[', '{', '<'}

	return slices.Contains(opening, b)
}
