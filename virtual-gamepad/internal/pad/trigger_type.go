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
	// JustRelease is the button to be triggered when just released only.
	JustRelease TriggerType = iota
	// Pressing is the button to be triggered during pressed every frame
	Pressing
	// JustPressed is the button to be triggered when just pressed only.
	JustPressed
)

func getTriggerButtonImage(tt TriggerType) (*ebiten.Image, error) {
	var b []byte
	switch tt {
	case JustRelease:
		b = images.A_button_png
	case Pressing:
		b = images.B_button_png
	}

	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	return ebiten.NewImageFromImage(img, ebiten.FilterDefault)
}
