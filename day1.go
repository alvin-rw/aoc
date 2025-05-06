package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

// part 1
func totalDistance(inputFilePath string) (int, error) {
	firstList, secondList, err := getSlicesFromInputFile(inputFilePath)
	if err != nil {
		return -1, fmt.Errorf("error when getting slice from input file: %v", err)
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

// part 2
func getSimilarityScore(inputFilePath string) (int, error) {
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
	f, err := os.Open(inputFilePath)
	if err != nil {
		return nil, nil, fmt.Errorf("error when opening the input file, %v", err)
	}
	defer f.Close()

	firstList := []int{}
	secondList := []int{}

	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			} else {
				return nil, nil, fmt.Errorf("error when reading file, %v", err)
			}
		}

		locationIds := strings.Split(line, "   ")
		for i, locationId := range locationIds {
			locationId = strings.TrimSuffix(locationId, "\n")

			l, err := strconv.Atoi(locationId)
			if err != nil {
				return nil, nil, fmt.Errorf("error when converting string %s to int, %v", locationId, err)
			}

			if i == 0 {
				firstList = append(firstList, l)
			} else {
				secondList = append(secondList, l)
			}
		}
	}

	return firstList, secondList, nil
}

func createOccurenceMap(list []int) map[int]int {
	occurenceMap := make(map[int]int)

	for _, n := range list {
		if _, ok := occurenceMap[n]; ok {
			occurenceMap[n] = occurenceMap[n] + 1
		} else if !ok {
			occurenceMap[n] = 1
		}
	}

	return occurenceMap
}
