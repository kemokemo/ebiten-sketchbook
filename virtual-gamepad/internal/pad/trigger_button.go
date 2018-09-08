package pad

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

type TriggerButton interface {
	SetLocation(x, y int)
	Update()
	IsTriggered() bool
	Draw(*ebiten.Image) error
}

// NewTriggerButton returns a new TriggerButton.
func NewTriggerButton(tt TriggerType) (TriggerButton, error) {
	img, err := getTriggerButtonImage(tt)
	if err != nil {
		return nil, err
	}

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
