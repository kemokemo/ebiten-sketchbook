package pad

import (
	"image"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

// JustReleaseButton is implementation of the TriggerButton to be
// triggered when just released.
type JustReleaseButton struct {
	baseImg     *ebiten.Image
	normalOp    *ebiten.DrawImageOptions
	selectedOp  *ebiten.DrawImageOptions
	rectangle   image.Rectangle
	isTriggered bool
	touches     map[*touch]struct{}
}

// SetLocation sets the location to draw this button.
func (b *JustReleaseButton) SetLocation(x, y int) {
	w, h := b.baseImg.Size()
	b.rectangle = image.Rect(x, y, x+w, y+h)

	b.normalOp.GeoM.Translate(float64(x), float64(y))
	b.selectedOp.GeoM.Concat(b.normalOp.GeoM)
}

// Update updates the internal state of this button.
// Please call this before using IsTriggered method.
func (b *JustReleaseButton) Update() {
	b.isTriggered = false

	IDs := inpututil.JustPressedTouchIDs()
	if len(IDs) != 0 {
		for _, id := range IDs {
			b.touches[&touch{id: id}] = struct{}{}
		}
	}

	for t := range b.touches {
		t.Update()
		if t.IsReleased() {
			delete(b.touches, t)
			in := image.Point{t.x, t.y}.In(b.rectangle)
			if in {
				b.isTriggered = true
				return
			}
		}
	}
}

// IsTriggered returns the state of this trigger is pressed.
// If result is 'true', this is pressed now.
func (b *JustReleaseButton) IsTriggered() bool {
	return b.isTriggered
}

// Draw draws this button.
func (b *JustReleaseButton) Draw(screen *ebiten.Image) error {
	if b.isTriggered {
		return screen.DrawImage(b.baseImg, b.selectedOp)
	}
	return screen.DrawImage(b.baseImg, b.normalOp)
}
