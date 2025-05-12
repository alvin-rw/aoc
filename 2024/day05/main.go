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
	sumOfValidPagesMidValue, err := calculateSumOfValidPagesMiddleValue("./input.txt")
	if err != nil {
		log.Fatalf("error when calculating valid pages value: %v", err)
	}

	sumOfInvalidPagesMidValue, err := calculateSumOfInvalidPagesMiddleValue("./input.txt")
	if err != nil {
		log.Fatalf("error when calculating valid pages value: %v", err)
	}

	fmt.Printf("sum of valid pages middle values: %d\n", sumOfValidPagesMidValue)
	fmt.Printf("sum of invalid pages middle values: %d\n", sumOfInvalidPagesMidValue)
}

func calculateSumOfValidPagesMiddleValue(inputFilePath string) (int, error) {
	sum := 0

	pageOrderingRules, printedPageMatrix, err := readPageOrderingRulesAndPrintedPagesFromFile(inputFilePath)
	if err != nil {
		return -1, err
	}

	for _, printedPagesList := range printedPageMatrix {
		currentPrintedPageMap := createPrintedPagesMap(printedPagesList)

		if isCorrectPrintedPage(currentPrintedPageMap, pageOrderingRules) {
			sum = sum + getMiddlePageNumber(printedPagesList)
		}
	}

	return sum, nil
}

func calculateSumOfInvalidPagesMiddleValue(inputFilePath string) (int, error) {
	sum := 0

	pageOrderingRules, printedPageMatrix, err := readPageOrderingRulesAndPrintedPagesFromFile(inputFilePath)
	if err != nil {
		return -1, err
	}

	for _, printedPagesList := range printedPageMatrix {
		currentPrintedPageMap := createPrintedPagesMap(printedPagesList)

		if !isCorrectPrintedPage(currentPrintedPageMap, pageOrderingRules) {
			fixPrintedPagesOrder(printedPagesList, pageOrderingRules)

			sum = sum + getMiddlePageNumber(printedPagesList)
		}
	}

	return sum, nil
}

func readPageOrderingRulesAndPrintedPagesFromFile(inputFilePath string) (map[int][]int, [][]int, error) {
	fileContent, err := file.ReadFile(inputFilePath)
	if err != nil {
		return nil, nil, err
	}

	// pageOrderingRules stores the list of numbers that need to be behind a certain number
	// for example
	// 1|2
	// 1|3
	// 2|3
	// will produce 1: {2,3}, 2: {3}
	pageOrderingRules := make(map[int][]int)
	printedPageMatrix := [][]int{}

	readingPageOrderingRules := true

	for _, line := range fileContent {
		// change to reading printed pages after a blank line
		if line == "" {
			readingPageOrderingRules = false
			continue
		}

		if readingPageOrderingRules {
			orderingRules := strings.Split(line, "|")

			rule, err := strconv.Atoi(orderingRules[0])
			if err != nil {
				return nil, nil, err
			}

			value, err := strconv.Atoi(orderingRules[1])
			if err != nil {
				return nil, nil, err
			}

			if _, ok := pageOrderingRules[rule]; !ok {
				pageOrderingRules[rule] = []int{value}
			} else {
				pageOrderingRules[rule] = append(pageOrderingRules[rule], value)
			}
		} else {
			lineSplitted := strings.Split(line, ",")

			printedPages := []int{}

			for _, p := range lineSplitted {
				pageNumber, err := strconv.Atoi(p)
				if err != nil {
					return nil, nil, err
				}

				printedPages = append(printedPages, pageNumber)
			}

			printedPageMatrix = append(printedPageMatrix, printedPages)
		}
	}

	return pageOrderingRules, printedPageMatrix, nil
}

func createPrintedPagesMap(printedPages []int) map[int][]int {
	printedPageMap := make(map[int][]int)

	for _, checkedPage := range printedPages {
		// add current page to all the entries in the map
		for page := range printedPageMap {
			printedPageMap[page] = append(printedPageMap[page], checkedPage)
		}

		if _, ok := printedPageMap[checkedPage]; !ok {
			printedPageMap[checkedPage] = []int{}
		}
	}

	return printedPageMap
}

func isCorrectPrintedPage(printedPageMap map[int][]int, orderingRules map[int][]int) bool {
	for pageNumber, pagesAfter := range printedPageMap {
		for _, page := range pagesAfter {
			if slices.Contains(orderingRules[page], pageNumber) {
				return false
			}
		}
	}

	return true
}

func getMiddlePageNumber(pages []int) int {
	midIndex := len(pages) / 2

	return pages[midIndex]
}

func fixPrintedPagesOrder(printedPageList []int, orderingRules map[int][]int) {
	printedPageMap := createPrintedPagesMap(printedPageList)

	if !isCorrectPrintedPage(printedPageMap, orderingRules) {
	checker:
		for pageNumber, pagesAfter := range printedPageMap {
			for _, checkedPage := range pagesAfter {
				if slices.Contains(orderingRules[checkedPage], pageNumber) {
					pageNumberIndex := slices.Index(printedPageList, pageNumber)
					checkedPageIndex := slices.Index(printedPageList, checkedPage)

					printedPageList = slices.Delete(printedPageList, checkedPageIndex, checkedPageIndex+1)
					printedPageList = slices.Insert(printedPageList, pageNumberIndex, checkedPage)

					break checker
				}
			}
		}

		fixPrintedPagesOrder(printedPageList, orderingRules)
	} else {
		return
	}
}

/**
NOTES: logic
1|2
1|3
1|4
2|3
2|4
3|4

3,1,2
1,4,3
1,5,2

rules:
[1]{2,3,4}
[2]{3,4}
[3]{4}


first:
[3]{1,2} -> wrong, there's 3 in rules[1]. wrong, there's 3 in rules[2]
[1]{2} -> correct, no 1 in rule[2]

second:
[1]{4,3} -> correct correct
[4]{3} -> wrong, there's 4 in rules 3

third:
[1]{5,2}
[5]{2}
**/
