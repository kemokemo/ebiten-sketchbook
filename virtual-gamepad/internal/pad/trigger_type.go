package pad

import (
	"bytes"
	"image"

	"github.com/hajimehoshi/ebiten"
	"github.com/kemokemo/ebiten-sketchbook/virtual-gamepad/internal/images"
)

// TriggerType is the type for the TriggerButton.
type TriggerType int

const (
	// AButton is the button for A. (ex: jump, bomb, etc..)
	AButton TriggerType = iota
	// BButton is the button for B. (ex: dash, shot, etc..)
	BButton
)

func getTriggerButtonImage(tt TriggerType) (*ebiten.Image, error) {
	var b []byte
	switch tt {
	case AButton:
		b = images.A_button_png
	case BButton:
		b = images.B_button_png
	}

	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	return ebiten.NewImageFromImage(img, ebiten.FilterDefault)
}
