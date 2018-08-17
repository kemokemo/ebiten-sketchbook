package pad

import (
	"bytes"
	"image"
	"sort"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/kemokemo/ebiten-sketchbook/virtual-gamepad/internal/images"
)

const (
	shortMargin = 3
	longMargin  = 5
)

// DirectionalPad is the directional pad for a game.
type DirectionalPad struct {
	baseImg           *ebiten.Image
	buttons           map[Direction]*directionalButton
	selectedDirection Direction
	op                *ebiten.DrawImageOptions
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
	if dp.buttons == nil {
		dp.buttons = make(map[Direction]*directionalButton, 4)
	}
	b, err := newDirectionalButton(x+longMargin, y+35+shortMargin, -90)
	if err != nil {
		return err
	}
	dp.buttons[Left] = b
	return nil
}

// Update updates the internal status of this struct.
func (dp *DirectionalPad) Update() error {
	dp.updateDirection()
	dp.updateButtons()

	return nil
}

func (dp *DirectionalPad) updateDirection() {
	IDs := ebiten.TouchIDs()
	if len(IDs) == 0 {
		dp.selectedDirection = None
		return
	}

	jIDs := inpututil.JustPressedTouchIDs()
	if len(jIDs) == 0 && dp.selectedDirection == None {
		return
	}

	sort.Slice(jIDs, func(i, j int) bool {
		return jIDs[i] < jIDs[j]
	})
	for index := range jIDs {
		if isTouched(jIDs[index], dp.buttons[Left].GetRectangle()) {
			dp.selectedDirection = Left
			return
		}
	}
}

func (dp *DirectionalPad) updateButtons() {
	for key := range dp.buttons {
		dp.buttons[key].SelectButton(false)
	}
	if dp.selectedDirection != None {
		dp.buttons[dp.selectedDirection].SelectButton(true)
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
