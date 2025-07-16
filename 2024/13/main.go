package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/alvin-rw/aoc/internal/file"
)

type machine struct {
	a                  button
	b                  button
	prize              location
	buttonAPushedTimes int
	buttonBPushedTimes int
}

func (m machine) checkXReached() bool {
	return m.a.x*m.buttonAPushedTimes+m.b.x*m.buttonBPushedTimes == m.prize.x
}

func (m machine) checkYReached() bool {
	return m.a.y*m.buttonAPushedTimes+m.b.y*m.buttonBPushedTimes == m.prize.y
}

func (m machine) calculateCost() int {
	return m.buttonAPushedTimes*3 + m.buttonBPushedTimes
}

type button struct {
	x    int
	y    int
	cost int
}

type location struct {
	x int
	y int
}

func (l *location) pushButton(b button) {
	l.x += b.x
	l.y += b.y
}

func main() {
	machines := getButtonAndPrizeList("input.txt")

	totalCost := 0
	for _, m := range machines {
		totalCost += calculateMachineCost(&m)
	}

	fmt.Printf("fewest token used to win all possible prizes: %d\n", totalCost)
}

func calculateMachineCost(m *machine) int {
	// check if pressing 1 button can take us to the prize

	// number of button B pressed to reach prize's X & Y coordinate
	bToX := m.prize.x / m.b.x
	bToY := m.prize.y / m.b.y
	// bOnly shows if it's possible to reach the prize using only B
	bOnly := bToX == bToY && m.prize.x%m.b.x == 0 && m.prize.y%m.b.y == 0

	// number of button A pressed to reach prize's X & Y coordinate
	aToX := m.prize.x / m.a.x
	aToY := m.prize.y / m.a.y
	// aOnly shows if it's possible to reach the prize using only A
	aOnly := aToX == aToY && m.prize.x%m.a.x == 0 && m.prize.y%m.a.y == 0

	switch {
	case aOnly && bOnly:
		if bToX < aToX*3 {
			return bToX
		} else {
			return aToX
		}
	case aOnly:
		return aToX * 3
	case bOnly:
		return bToX
	}

	// check button combination
	// press B button as much as we can without exceeding the coordinate
	if bToX < bToY {
		m.buttonAPushedTimes = bToX
	} else {
		m.buttonBPushedTimes = bToY
	}

	return m.calculateCost()
}

// buttonMash is the recursive function to find the button combination
// that will take us to the prize
func buttonMash(m *machine) {
	if m.checkXReached() && m.checkYReached() {
		return
	}
}

func getButtonAndPrizeList(inputFilePath string) []machine {
	machines := []machine{}

	fileContent := file.ReadFile(inputFilePath)

	m := &machine{}
	for _, line := range fileContent {
		if strings.Contains("Button A", line) {
			m.a = parseButtonString(line)
		} else if strings.Contains("Button B", line) {
			m.b = parseButtonString(line)
		} else if strings.Contains("Prize", line) {
			m.prize = parsePrizeString(line)

			machines = append(machines, *m)
			m = &machine{}
		}
	}

	return machines
}

func parseButtonString(s string) button {
	cost := 0
	if strings.Contains("Button A", s) {
		cost = 3
	} else if strings.Contains("Button B", s) {
		cost = 1
	}

	elem := strings.Split(s, "")

	xString := ""
	for i := (slices.Index(elem, "+") + 1); i < slices.Index(elem, ","); i++ {
		xString += elem[i]
	}
	x, _ := strconv.Atoi(xString)

	elem = slices.Delete(elem, slices.Index(elem, ","), slices.Index(elem, ",")+1)

	yString := ""
	for i := (slices.Index(elem, "+") + 1); i < len(elem); i++ {
		yString += elem[i]
	}
	y, _ := strconv.Atoi(yString)

	return button{
		x:    x,
		y:    y,
		cost: cost,
	}
}

func parsePrizeString(s string) location {
	elem := strings.Split(s, "")

	xString := ""
	for i := (slices.Index(elem, "=") + 1); i < slices.Index(elem, ","); i++ {
		xString += elem[i]
	}
	x, _ := strconv.Atoi(xString)

	elem = slices.Delete(elem, slices.Index(elem, ","), slices.Index(elem, ",")+1)

	yString := ""
	for i := (slices.Index(elem, "=") + 1); i < len(elem); i++ {
		yString += elem[i]
	}
	y, _ := strconv.Atoi(yString)

	return location{
		x: x,
		y: y,
	}
}

/**
recursion

base case
x = x
y = y

or
button A = 0
button B = all



**/
