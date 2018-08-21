package character

import (
	"bytes"
	"image"

	"github.com/hajimehoshi/ebiten"
	"github.com/kemokemo/ebiten-sketchbook/virtual-gamepad/internal/images"
	"github.com/kemokemo/ebiten-sketchbook/virtual-gamepad/internal/pad"
)

// MainCharacter is the main character for users to control.
type MainCharacter struct {
	baseImg    *ebiten.Image
	normalOp   *ebiten.DrawImageOptions
	size       image.Point
	area       image.Rectangle
	currentPos image.Point
}

// NewMainCharacter returns a main character.
func NewMainCharacter() (*MainCharacter, error) {
	m := &MainCharacter{}
	img, _, err := image.Decode(bytes.NewReader(images.Fighter_png))
	if err != nil {
		return nil, err
	}
	m.baseImg, err = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		return nil, err
	}
	w, h := m.baseImg.Size()
	m.size = image.Point{w, h}

	m.normalOp = &ebiten.DrawImageOptions{}
	return m, nil
}

// SetLocation sets the location to draw this character.
func (m *MainCharacter) SetLocation(x, y int) {
	m.normalOp.GeoM.Reset()

	m.currentPos.X = x
	m.currentPos.Y = y
	m.normalOp.GeoM.Translate(float64(x), float64(y))
}

// SetArea sets the area for this character to move.
func (m *MainCharacter) SetArea(rect image.Rectangle) {
	m.area = rect
}

// GetSize returns the size this character.
func (m *MainCharacter) GetSize() image.Point {
	return m.size
}

// Update updates the internal state.
func (m *MainCharacter) Update() {
	// TODO: Make a judgment with enemy bullets
}

// Move moves this character regarding the direction.
// Do not move if the destination is outside the area.
func (m *MainCharacter) Move(direc pad.Direction) {
	switch direc {
	case pad.UpperLeft:
		m.move4direction(pad.Upper)
		m.move4direction(pad.Left)
	case pad.UpperRight:
		m.move4direction(pad.Upper)
		m.move4direction(pad.Right)
	case pad.LowerLeft:
		m.move4direction(pad.Lower)
		m.move4direction(pad.Left)
	case pad.LowerRight:
		m.move4direction(pad.Lower)
		m.move4direction(pad.Right)
	default:
		m.move4direction(direc)
	}
}

func (m *MainCharacter) move4direction(direc pad.Direction) {
	switch direc {
	case pad.Left:
		if m.area.Min.X > m.currentPos.X-2 {
			return
		}
		m.currentPos.X -= 2
		m.normalOp.GeoM.Translate(float64(-2), 0.0)
	case pad.Upper:
		if m.area.Min.Y > m.currentPos.Y-2 {
			return
		}
		m.currentPos.Y -= 2
		m.normalOp.GeoM.Translate(0.0, float64(-2))
	case pad.Right:
		if m.area.Max.X < m.currentPos.X+m.size.X+2 {
			return
		}
		m.currentPos.X += 2
		m.normalOp.GeoM.Translate(float64(2), 0.0)
	case pad.Lower:
		if m.area.Max.Y < m.currentPos.Y+m.size.Y+2 {
			return
		}
		m.currentPos.Y += 2
		m.normalOp.GeoM.Translate(0.0, float64(2))
	default:
		return
	}

}

// Draw draws this character.
func (m *MainCharacter) Draw(screen *ebiten.Image) error {
	return screen.DrawImage(m.baseImg, m.normalOp)
}
