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
func totalDistance(inputPath string) (int, error) {
	f, err := os.Open(inputPath)
	if err != nil {
		return -1, fmt.Errorf("error when opening the input file, %v", err)
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
				return -1, fmt.Errorf("error when reading file, %v", err)
			}
		}

		locationIds := strings.Split(line, "   ")
		for i, locationId := range locationIds {
			locationId = strings.TrimSuffix(locationId, "\n")

			l, err := strconv.Atoi(locationId)
			if err != nil {
				return -1, fmt.Errorf("error when converting string %s to int, %v", locationId, err)
			}

			if i == 0 {
				firstList = append(firstList, l)
			} else {
				secondList = append(secondList, l)
			}
		}
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
