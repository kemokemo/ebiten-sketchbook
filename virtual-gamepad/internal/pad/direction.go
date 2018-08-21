package pad

// Direction is the direction for the pad.
type Direction int

const (
	// None describes that no direction buttons are pressed.
	None Direction = iota
	// Left is left direction
	Left
	// Right is right direction.
	Right
	// Upper is the upper direction.
	Upper
	// UpperLeft is the upper left direction
	UpperLeft
	// UpperRight is the upper right direction
	UpperRight
	// Lower is the lower direction.
	Lower
	// LowerLeft is the lower left direction
	LowerLeft
	// LowerRight is the lower right direction
	LowerRight
)

func getDirectionDegree(d Direction) int {
	switch d {
	case Left:
		return -90
	case Right:
		return 90
	case Lower:
		return 180
	default:
		return 0
	}
}

func getMergedDirection(previous, current Direction) Direction {
	switch previous {
	case Left:
		if current == Upper {
			return UpperLeft
		} else if current == Lower {
			return LowerLeft
		} else {
			return previous
		}
	case Upper:
		if current == Left {
			return UpperLeft
		} else if current == Right {
			return UpperRight
		} else {
			return previous
		}
	case Right:
		if current == Upper {
			return UpperRight
		} else if current == Lower {
			return LowerRight
		} else {
			return previous
		}
	case Lower:
		if current == Right {
			return LowerRight
		} else if current == Left {
			return LowerLeft
		} else {
			return previous
		}
	default:
		return current
	}
}
