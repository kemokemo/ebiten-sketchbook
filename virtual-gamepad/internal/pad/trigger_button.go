package pad

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

// TriggerButton is the button to fire guns.
type TriggerButton struct {
	baseImg     *ebiten.Image
	normalOp    *ebiten.DrawImageOptions
	selectedOp  *ebiten.DrawImageOptions
	rectangle   image.Rectangle
	isTriggered bool
}

// NewTriggerButton returns a new TriggerButton.
func NewTriggerButton(tt TriggerType) (*TriggerButton, error) {
	img, err := getTriggerButtonImage(tt)
	if err != nil {
		return nil, err
	}

	tb := &TriggerButton{}
	tb.baseImg = img
	tb.normalOp = &ebiten.DrawImageOptions{}
	tb.selectedOp = &ebiten.DrawImageOptions{}
	tb.selectedOp.ColorM.Scale(colorScale(color.RGBA{0, 148, 255, 255}))

	return tb, nil
}

// SetLocation sets the location to draw this button.
func (t *TriggerButton) SetLocation(x, y int) {
	w, h := t.baseImg.Size()
	t.rectangle = image.Rect(x, y, x+w, y+h)

	t.normalOp.GeoM.Translate(float64(x), float64(y))
	t.selectedOp.GeoM.Add(t.normalOp.GeoM)
}

// Update updates the internal state of this button.
// Please call this before using IsTriggered method.
func (t *TriggerButton) Update() {
	t.isTriggered = false

	IDs := ebiten.TouchIDs()
	if len(IDs) == 0 {
		return
	}

	for i := range IDs {
		if isTouched(IDs[i], t.rectangle) {
			t.isTriggered = true
			return
		}
	}
}

// IsTriggered returns the state of this trigger is pressed.
// If result is 'true', this is pressed now.
func (t *TriggerButton) IsTriggered() bool {
	return t.isTriggered
}

// Draw draws this button.
func (t *TriggerButton) Draw(screen *ebiten.Image) error {
	if t.isTriggered {
		return screen.DrawImage(t.baseImg, t.selectedOp)
	}
	return screen.DrawImage(t.baseImg, t.normalOp)
}
