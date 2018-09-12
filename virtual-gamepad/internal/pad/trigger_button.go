package pad

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten"
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

type TriggerButton interface {
	SetLocation(x, y int)
	Update()
	IsTriggered() bool
	Draw(*ebiten.Image) error
}

// NewTriggerButton returns a new TriggerButton.
func NewTriggerButton(img *ebiten.Image, tt TriggerType) (TriggerButton, error) {
	sop := &ebiten.DrawImageOptions{}
	sop.ColorM.Scale(colorScale(color.RGBA{0, 148, 255, 255}))
	switch tt {
	case JustRelease:
		return &JustReleaseButton{
			baseImg:    img,
			normalOp:   &ebiten.DrawImageOptions{},
			selectedOp: sop,
			touches:    make(map[*touch]struct{}),
		}, nil
	case Pressing:
		return &PressingButton{
			baseImg:    img,
			normalOp:   &ebiten.DrawImageOptions{},
			selectedOp: sop,
		}, nil
	default:
		return nil, fmt.Errorf("unknown trigger type: %v", tt)
	}
}
