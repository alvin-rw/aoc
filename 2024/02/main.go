package main

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/alvin-rw/aoc/internal/file"
)

const (
	increasing = iota
	decreasing
)

func main() {
	numberOfSafeReportWithoutDampener, err := calculateNumberOfSafeReport("./input.txt", false)
	if err != nil {
		log.Fatalf("error when calculating number of safe report: %v", err)
	}

	numberOfSafeReportWithDampener, err := calculateNumberOfSafeReport("./input.txt", true)
	if err != nil {
		log.Fatalf("error when calculating number of safe report: %v", err)
	}

	fmt.Printf("number of safe report: %d\n", numberOfSafeReportWithoutDampener)
	fmt.Printf("number of safe report with dampener: %d\n", numberOfSafeReportWithDampener)
}

func calculateNumberOfSafeReport(inputFilePath string, isDampened bool) (int, error) {
	numOfSafeReport := 0

	reports, err := getReportsFromFile(inputFilePath)
	if err != nil {
		return -1, fmt.Errorf("error when getting report from file, %v", err)
	}

	for _, report := range reports {
		isSafe := isReportSafeWithDampener(report, isDampened)

		if isSafe {
			numOfSafeReport++
		}
	}

	return numOfSafeReport, nil
}

func isReportSafeWithDampener(report []int, isDampened bool) bool {
	isSafe := isReportSafe(report)

	if !isSafe && isDampened {
		for i := range report {
			tempReport := slices.Clone(report)
			tempReport = slices.Delete(tempReport, i, i+1)

			isSafe = isReportSafe(tempReport)
			if isSafe {
				break
			}
		}
	}

	return isSafe
}

func isReportSafe(report []int) bool {
	var levelDirection int

	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]

		if diff > 3 || diff == 0 || diff < -3 {
			return false
		}

		if diff > 0 {
			if i == 1 {
				levelDirection = increasing
			} else {
				if levelDirection != increasing {
					return false
				}
			}
		} else {
			if i == 1 {
				levelDirection = decreasing
			} else {
				if levelDirection != decreasing {
					return false
				}
			}
		}
	}

	return true
}

func getReportsFromFile(inputFilePath string) ([][]int, error) {
	fileContent := file.ReadFile(inputFilePath)

	reports := [][]int{}
	for _, line := range fileContent {
		stringList := strings.Split(line, " ")

		report := []int{}
		for _, num := range stringList {
			level, err := strconv.Atoi(num)
			if err != nil {
				return nil, fmt.Errorf("error when converting string %s to int, %w", num, err)
			}

			report = append(report, level)
		}
		reports = append(reports, report)
	}

	return reports, nil
}
