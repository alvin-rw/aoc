package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

// movement direction inside a matrix
const (
	Up Direction = iota
	Down
	Right
	Left
	UpRight
	UpLeft
	DownRight
	DownLeft
)

type Direction uint8

func GetNextColumn(column int, direction Direction, modifier int) int {
	nextColumn := column

	switch direction {
	case Up:
		nextColumn = column
	case Down:
		nextColumn = column
	case Right:
		nextColumn = column + modifier
	case Left:
		nextColumn = column - modifier
	case UpRight:
		nextColumn = column + modifier
	case UpLeft:
		nextColumn = column - modifier
	case DownRight:
		nextColumn = column + modifier
	case DownLeft:
		nextColumn = column - modifier
	}

	return nextColumn
}

func GetNextRow(row int, direction Direction, modifier int) int {
	nextRow := row

	switch direction {
	case Up:
		nextRow = row - modifier
	case Down:
		nextRow = row + modifier
	case Right:
		nextRow = row
	case Left:
		nextRow = row
	case UpRight:
		nextRow = row - modifier
	case UpLeft:
		nextRow = row - modifier
	case DownRight:
		nextRow = row + modifier
	case DownLeft:
		nextRow = row + modifier
	}

	return nextRow
}

func ChangeDirection90Degree(dir Direction) Direction {
	switch dir {
	case Up:
		return Right
	case Right:
		return Down
	case Down:
		return Left
	case Left:
		return Up
	default:
		panic("error when changing direction")
	}
}

// CheckCoordinateInsideMatrix checks if coordinate (in the format []int{Row, Column})
// is located inside the matrix
func CheckCoordinateInsideMatrix(coordinate []int, numOfRows int, numOfColumns int) bool {
	if coordinate[0] >= 0 && coordinate[1] >= 0 && coordinate[0] < numOfRows && coordinate[1] < numOfColumns {
		return true
	} else {
		return false
	}
}

// Change coordinate to string format (row,col)
func CoordToString(row, col int) string {
	return fmt.Sprintf("%d,%d", row, col)
}

// Extract row and column from string formatted (row,col)
// Return values are row, column
func StringToCoord(s string) (int, int) {
	coord := strings.Split(s, ",")
	row, _ := strconv.Atoi(coord[0])
	col, _ := strconv.Atoi(coord[1])

	return row, col
}
