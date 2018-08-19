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
	// Up is the upper direction.
	Up
	// Down is the lower direction.
	Down
)

func getDirectionDegree(d Direction) int {
	switch d {
	case Left:
		return -90
	case Right:
		return 90
	case Down:
		return 180
	default:
		return 0
	}
}
