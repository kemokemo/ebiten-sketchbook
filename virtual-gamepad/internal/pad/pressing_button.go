package pad

import (
	"image"

	"github.com/hajimehoshi/ebiten"
)

// PressingButton is implementation of the TriggerButton to be
// triggered during being pressed.
type PressingButton struct {
	baseImg     *ebiten.Image
	normalOp    *ebiten.DrawImageOptions
	selectedOp  *ebiten.DrawImageOptions
	rectangle   image.Rectangle
	isSelected  bool
	isTriggered bool
}

// SetLocation sets the location to draw this button.
func (b *PressingButton) SetLocation(x, y int) {
	w, h := b.baseImg.Size()
	b.rectangle = image.Rect(x, y, x+w, y+h)

	b.normalOp.GeoM.Translate(float64(x), float64(y))
	b.selectedOp.GeoM.Concat(b.normalOp.GeoM)
}

// Update updates the internal state of this button.
// Please call this before using IsTriggered method.
func (b *PressingButton) Update() {
	b.updateSelect()
	b.updateTrigger()
}

func (b *PressingButton) updateSelect() {
	b.isSelected = false

	IDs := ebiten.TouchIDs()
	if len(IDs) == 0 {
		return
	}

	for i := range IDs {
		if isTouched(IDs[i], b.rectangle) {
			b.isSelected = true
			return
		}
	}
}

func (b *PressingButton) updateTrigger() {
	b.isTriggered = false
	IDs := ebiten.TouchIDs()
	if len(IDs) == 0 {
		return
	}

	for i := range IDs {
		if isTouched(IDs[i], b.rectangle) {
			b.isTriggered = true
			return
		}
	}
}

// IsTriggered returns the state of this trigger is pressed.
// If result is 'true', this is pressed now.
func (b *PressingButton) IsTriggered() bool {
	return b.isTriggered
}

// Draw draws this button.
func (b *PressingButton) Draw(screen *ebiten.Image) error {
	if b.isSelected {
		return screen.DrawImage(b.baseImg, b.selectedOp)
	}
	return screen.DrawImage(b.baseImg, b.normalOp)
}
