package pad

import (
	"image"

	"github.com/hajimehoshi/ebiten"
)

// PressingButton is implementation of the TriggerButton to be
// triggered during being pressed.
type PressingButton struct {
	baseImg     *ebiten.Image
	tType       TriggerType
	normalOp    *ebiten.DrawImageOptions
	selectedOp  *ebiten.DrawImageOptions
	rectangle   image.Rectangle
	isTriggered bool
}

// SetLocation sets the location to draw this button.
func (t *PressingButton) SetLocation(x, y int) {
	w, h := t.baseImg.Size()
	t.rectangle = image.Rect(x, y, x+w, y+h)

	t.normalOp.GeoM.Translate(float64(x), float64(y))
	t.selectedOp.GeoM.Concat(t.normalOp.GeoM)
}

// Update updates the internal state of this button.
// Please call this before using IsTriggered method.
func (t *PressingButton) Update() {
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
func (t *PressingButton) IsTriggered() bool {
	return t.isTriggered
}

// Draw draws this button.
func (t *PressingButton) Draw(screen *ebiten.Image) error {
	if t.isTriggered {
		return screen.DrawImage(t.baseImg, t.selectedOp)
	}
	return screen.DrawImage(t.baseImg, t.normalOp)
}
