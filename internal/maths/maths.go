package maths

import (
	"strconv"
	"strings"
)

func GetNumberOfDigits(n int) int {
	s := strconv.Itoa(n)

	return len(s)
}

func SplitNumberIntoTwo(n int) (int, int) {
	number := strconv.Itoa(n)
	separateIndex := GetNumberOfDigits(n) / 2

	a := []string{}
	b := []string{}
	for i, r := range number {
		if i < separateIndex {
			a = append(a, string(r))
		} else {
			b = append(b, string(r))
		}
	}

	aa, _ := strconv.Atoi(strings.Join(a, ""))
	bb, _ := strconv.Atoi(strings.Join(b, ""))

	return aa, bb
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}
