package pad

import (
	"bytes"
	"image"

	"github.com/hajimehoshi/ebiten"
	"github.com/kemokemo/ebiten-sketchbook/virtual-gamepad/internal/images"
)

const (
	shortMargin = 3
	longMargin  = 5
)

// DirectionalPad is the directional pad for a game.
type DirectionalPad struct {
	baseImg *ebiten.Image
	buttons []*directionalButton
	op      *ebiten.DrawImageOptions
}

// NewDirectionalPad returns a new DirectionalPad.
func NewDirectionalPad(x, y int) (*DirectionalPad, error) {
	dp := &DirectionalPad{}
	img, _, err := image.Decode(bytes.NewReader(images.Directional_pad_png))
	if err != nil {
		return nil, err
	}
	dp.baseImg, err = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		return nil, err
	}

	dp.op = &ebiten.DrawImageOptions{}
	dp.op.GeoM.Translate(float64(x), float64(y))

	err = dp.createButtons(x, y)
	if err != nil {
		return nil, err
	}
	return dp, nil
}

func (dp *DirectionalPad) createButtons(x, y int) error {
	b, err := newDirectionalButton(x+longMargin, y+35+shortMargin, -90)
	if err != nil {
		return err
	}
	dp.buttons = append(dp.buttons, b)

	return nil
}

// Update updates the internal status of this struct.
func (dp *DirectionalPad) Update() error {
	return nil
}

// Draw draws the directional buttons belong this struct.
func (dp *DirectionalPad) Draw(screen *ebiten.Image) error {
	err := screen.DrawImage(dp.baseImg, dp.op)
	if err != nil {
		return err
	}

	for index := range dp.buttons {
		e := dp.buttons[index].Draw(screen)
		if err != nil {
			return e
		}
	}
	return nil
}
