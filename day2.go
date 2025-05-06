package aoc2024

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	increasing = iota
	decreasing
)

func calculateSafeReport(inputFilePath string) (int, error) {
	numOfSafeReport := 0

	reports, err := getReportsFromFile(inputFilePath)
	if err != nil {
		return -1, fmt.Errorf("error when getting report from file, %v", err)
	}

	for _, report := range reports {
		isSafe := checkReportSafety(report)

		if isSafe {
			numOfSafeReport++
		}
	}

	return numOfSafeReport, nil
}

func checkReportSafety(report []int) bool {
	var levelDirection int

	for i, level := range report {
		if i == 0 {
			continue
		}

		diff := level - report[i-1]
		switch {
		case diff > 0:
			if diff > 3 {
				return false
			}

			if i == 1 {
				levelDirection = increasing
			} else {
				if levelDirection != increasing {
					return false
				}
			}
		case diff < 0:
			if diff < -3 {
				return false
			}

			if i == 1 {
				levelDirection = decreasing
			} else {
				if levelDirection != decreasing {
					return false
				}
			}
		case diff == 0:
			return false
		}
	}

	return true
}

func getReportsFromFile(inputFilePath string) ([][]int, error) {
	f, err := os.Open(inputFilePath)
	if err != nil {
		return nil, fmt.Errorf("error when opening the file, %v", err)
	}
	defer f.Close()

	reports := [][]int{}

	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			} else {
				return nil, fmt.Errorf("error when reading file, %v", err)
			}
		}

		line = strings.TrimSuffix(line, "\n")
		stringList := strings.Split(line, " ")

		report := []int{}
		for _, num := range stringList {
			level, err := strconv.Atoi(num)
			if err != nil {
				return nil, fmt.Errorf("error when converting string %s to int, %v", num, err)
			}

			report = append(report, level)
		}

		reports = append(reports, report)
	}

	return reports, nil
}
