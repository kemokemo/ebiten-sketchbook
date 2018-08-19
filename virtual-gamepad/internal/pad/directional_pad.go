package pad

import (
	"bytes"
	"image"
	"sort"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/kemokemo/ebiten-sketchbook/virtual-gamepad/internal/images"
)

// DirectionalPad is the directional pad for a game.
type DirectionalPad struct {
	baseImg           *ebiten.Image
	buttons           map[Direction]*directionalButton
	selectedDirection Direction
	op                *ebiten.DrawImageOptions
}

// NewDirectionalPad returns a new DirectionalPad.
func NewDirectionalPad() (*DirectionalPad, error) {
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

	err = dp.createButtons()
	if err != nil {
		return nil, err
	}
	return dp, nil
}

func (dp *DirectionalPad) createButtons() error {
	if dp.buttons == nil {
		dp.buttons = make(map[Direction]*directionalButton, 4)
	}

	b, err := newDirectionalButton(Left)
	if err != nil {
		return err
	}
	dp.buttons[Left] = b

	b, err = newDirectionalButton(Up)
	if err != nil {
		return err
	}
	dp.buttons[Up] = b

	return nil
}

// SetLocation sets the location to draw this directional pad.
func (dp *DirectionalPad) SetLocation(x, y int) {
	dp.op.GeoM.Translate(float64(x), float64(y))

	wp, _ := dp.baseImg.Size()
	halfWp := int(wp / 2)
	wb, _ := dp.buttons[Left].Size()
	halfWb := int(wb / 2)
	outerMargin := int(0.25 * float64(halfWp))
	//innerMargin := int(0.08 * float64(halfWp))

	dp.buttons[Left].SetLocation(x+outerMargin, y+halfWp-halfWb)
	dp.buttons[Up].SetLocation(x+halfWp-halfWb, y+outerMargin)
}

// Update updates the internal status of this struct.
func (dp *DirectionalPad) Update() error {
	dp.updateDirection()
	dp.updateButtons()

	return nil
}

func (dp *DirectionalPad) updateDirection() {
	IDs := ebiten.TouchIDs()

	// There are no touches
	if len(IDs) == 0 {
		dp.selectedDirection = None
		return
	}

	// Just touched!
	jIDs := inpututil.JustPressedTouchIDs()
	sort.Slice(jIDs, func(i, j int) bool {
		return jIDs[i] < jIDs[j]
	})
	for index := range jIDs {
		for key := range dp.buttons {
			if isTouched(jIDs[index], dp.buttons[key].GetRectangle()) {
				dp.selectedDirection = key
				return
			}
		}
	}
}

func (dp *DirectionalPad) updateButtons() {
	for key := range dp.buttons {
		if key == dp.selectedDirection {
			dp.buttons[key].SelectButton(true)
		} else {
			dp.buttons[key].SelectButton(false)
		}
	}
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

// GetDirection returns the currently selected direction.
func (dp *DirectionalPad) GetDirection() Direction {
	return dp.selectedDirection
}

func isTouched(touchedID int, bounds image.Rectangle) bool {
	x, y := ebiten.TouchPosition(touchedID)
	min := bounds.Min
	max := bounds.Max
	if min.X < x && x < max.X && min.Y < y && y < max.Y {
		return true
	}
	return false
}
