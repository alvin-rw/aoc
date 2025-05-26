package matrix

// movement direction inside a matrix
const (
	Up = iota
	Down
	Right
	Left
	UpRight
	UpLeft
	DownRight
	DownLeft
)

func GetNextColumn(column int, direction int, modifier int) int {
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

func GetNextRow(row int, direction int, modifier int) int {
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

func ChangeDirection90Degree(dir int) int {
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

func CheckCoordinateInsideMatrix(coordinate []int, maxRow int, maxColumn int) bool {
	if coordinate[0] >= 0 && coordinate[1] >= 0 && coordinate[0] < maxRow && coordinate[1] < maxColumn {
		return true
	} else {
		return false
	}
}
