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

func (m machine) checkPrizeReached() bool {
	return m.a.x*m.buttonAPushedTimes+m.b.x*m.buttonBPushedTimes == m.prize.x && m.a.y*m.buttonAPushedTimes+m.b.y*m.buttonBPushedTimes == m.prize.y
}

func (m machine) calculateCost() int {
	return m.buttonAPushedTimes*3 + m.buttonBPushedTimes
}

type button struct {
	x    int
	y    int
	cost int
}

// returns the most times we can push a button without exceeding the destination
// retuns true if destination reached
func (b button) checkButtonCanReachDestination(xDest, yDest int) (int, bool) {
	// number of button pressed to reach destination X & Y coordinate
	toX := xDest / b.x
	toY := yDest / b.y

	buttonPress := 0
	buttonPress = min(toX, toY)

	destinationReached := toX == toY && // same number of button press required to reach X and Y
		xDest%b.x == 0 && // button can cover X distance
		yDest%b.y == 0 // button can cover Y distance

	return buttonPress, destinationReached
}

type location struct {
	x int
	y int
}

func main() {
	machines := getMachinesDetails("input.txt")

	totalCost := 0
	for _, m := range machines {
		totalCost += calculateMachineCost(&m)
	}

	fmt.Printf("fewest token used to win all possible prizes: %d\n", totalCost)
}

func calculateMachineCost(m *machine) int {
	// check if pressing 1 button can take us to the prize

	aTimes, possibleAOnly := m.a.checkButtonCanReachDestination(m.prize.x, m.prize.y)
	bTimes, possibleBOnly := m.b.checkButtonCanReachDestination(m.prize.x, m.prize.y)

	switch {
	case possibleAOnly && possibleBOnly:
		if bTimes < aTimes*3 {
			return bTimes
		} else {
			return aTimes
		}
	case possibleAOnly:
		return aTimes * 3
	case possibleBOnly:
		return bTimes
	// we can return early if pressing b or a button once will take us beyond destination
	// because this means not a b&a combination and they're checked already
	case bTimes == 0 || aTimes == 0:
		return 0
	}

	// check button combination
	// press B button as much as we can without exceeding the coordinate
	m.buttonBPushedTimes = bTimes

	for range m.buttonBPushedTimes {
		if m.checkPrizeReached() {
			break
		}

		xDistanceToCover := m.prize.x - (m.b.x * m.buttonBPushedTimes)
		yDistanceToCover := m.prize.y - (m.b.y * m.buttonBPushedTimes)

		if push, destReached := m.a.checkButtonCanReachDestination(xDistanceToCover, yDistanceToCover); destReached {
			m.buttonAPushedTimes = push
			break
		}
		m.buttonBPushedTimes--
	}

	return m.calculateCost()
}

func getMachinesDetails(inputFilePath string) []machine {
	machines := []machine{}

	fileContent := file.ReadFile(inputFilePath)

	m := &machine{}
	for _, line := range fileContent {
		if strings.Contains(line, "Button A") {
			m.a = parseButtonString(line)
		} else if strings.Contains(line, "Button B") {
			m.b = parseButtonString(line)
		} else if strings.Contains(line, "Prize") {
			m.prize = parsePrizeString(line)

			machines = append(machines, *m)
			m = &machine{}
		}
	}

	return machines
}

func parseButtonString(s string) button {
	cost := 0
	if strings.Contains(s, "Button A") {
		cost = 3
	} else if strings.Contains(s, "Button B") {
		cost = 1
	}

	elem := strings.Split(s, "")

	xString := ""
	for i := (slices.Index(elem, "+") + 1); i < slices.Index(elem, ","); i++ {
		xString += elem[i]
	}
	x, _ := strconv.Atoi(xString)

	elem = slices.Delete(elem, 0, slices.Index(elem, ",")+1)

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

	elem = slices.Delete(elem, 0, slices.Index(elem, ",")+1)

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
