package main

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/alvin-rw/aoc/internal/file"
)

func main() {
	totalDistance, err := calculateTotalDistance("./input.txt")
	if err != nil {
		log.Fatalf("error when calculating total distance: %v", err)
	}

	similarityScore, err := calculateSimilarityScore("./input.txt")
	if err != nil {
		log.Fatalf("error when calculating similarity score: %v", err)
	}

	fmt.Printf("total distance: %d\n", totalDistance)
	fmt.Printf("similarity score: %d\n", similarityScore)
}

func calculateTotalDistance(inputFilePath string) (int, error) {
	firstList, secondList, err := getSlicesFromInputFile(inputFilePath)
	if err != nil {
		return -1, err
	}

	slices.Sort(firstList)
	slices.Sort(secondList)

	var totalDistance int

	for i, firstListLocationId := range firstList {
		secondListLocationId := secondList[i]

		distance := firstListLocationId - secondListLocationId
		if distance < 0 {
			distance = distance * -1
		}

		totalDistance = totalDistance + distance
	}

	return totalDistance, nil
}

func calculateSimilarityScore(inputFilePath string) (int, error) {
	firstList, secondList, err := getSlicesFromInputFile(inputFilePath)
	if err != nil {
		return -1, fmt.Errorf("error when getting slice from input file: %v", err)
	}

	secondListOccurenceMap := createOccurenceMap(secondList)

	var similarityScore int
	for _, num := range firstList {
		if occurence, ok := secondListOccurenceMap[num]; ok {
			score := num * occurence

			similarityScore = similarityScore + score
		}
	}

	return similarityScore, nil
}

func getSlicesFromInputFile(inputFilePath string) ([]int, []int, error) {
	fileContent := file.ReadFile(inputFilePath)

	firstList := []int{}
	secondList := []int{}

	for _, line := range fileContent {
		locationIds := strings.Split(line, "   ")

		for i, locationId := range locationIds {
			id, err := strconv.Atoi(locationId)
			if err != nil {
				return nil, nil, fmt.Errorf("error when converting string %s to int: %w", locationId, err)
			}

			if i == 0 {
				firstList = append(firstList, id)
			} else {
				secondList = append(secondList, id)
			}
		}
	}

	return firstList, secondList, nil
}

func createOccurenceMap(list []int) map[int]int {
	occurenceMap := make(map[int]int)

	for _, n := range list {
		occurenceMap[n]++
	}

	return occurenceMap
}
