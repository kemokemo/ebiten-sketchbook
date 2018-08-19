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
	baseImg  *ebiten.Image
	normalOp *ebiten.DrawImageOptions
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

	m.normalOp = &ebiten.DrawImageOptions{}
	return m, nil
}

// SetLocation sets the location to draw this character.
func (m *MainCharacter) SetLocation(x, y int) {
	m.normalOp.GeoM.Reset()
	m.normalOp.GeoM.Translate(float64(x), float64(y))
}

// Update updates the internal state.
func (m *MainCharacter) Update() {
	// TODO: Make a judgment with enemy bullets
}

// Move moves this character regarding the direction.
func (m *MainCharacter) Move(direc pad.Direction) {
	switch direc {
	case pad.Left:
		m.normalOp.GeoM.Translate(float64(-2), 0.0)
	case pad.Up:
		m.normalOp.GeoM.Translate(0.0, float64(-2))
	case pad.Right:
		m.normalOp.GeoM.Translate(float64(2), 0.0)
	case pad.Down:
		m.normalOp.GeoM.Translate(0.0, float64(2))
	default:

	}
}

// Draw draws this character.
func (m *MainCharacter) Draw(screen *ebiten.Image) error {
	return screen.DrawImage(m.baseImg, m.normalOp)
}
