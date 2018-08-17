package pad

import (
	"bytes"
	"image"
	"math"

	"image/color"

	"github.com/hajimehoshi/ebiten"

	"github.com/kemokemo/ebiten-sketchbook/virtual-gamepad/internal/images"
)

// debugirectionalButton is the button of the directional pad.
type directionalButton struct {
	baseImg    *ebiten.Image
	selected   bool
	rectangle  image.Rectangle
	normalOp   *ebiten.DrawImageOptions
	selectedOp *ebiten.DrawImageOptions
}

// newDirectionalButton returns a new DirectionalButton.
func newDirectionalButton(x, y, degree int) (*directionalButton, error) {
	d := &directionalButton{}
	img, _, err := image.Decode(bytes.NewReader(images.Directional_button_png))
	if err != nil {
		return nil, err
	}
	d.baseImg, err = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		return nil, err
	}

	w, h := d.baseImg.Size()
	d.rectangle = image.Rect(x, y, x+w, y+h)
	halfW := float64(w) / 2
	halfH := float64(h) / 2

	d.normalOp = &ebiten.DrawImageOptions{}
	d.normalOp.GeoM.Translate(-halfW, -halfH)
	d.normalOp.GeoM.Rotate(float64(degree) * 2 * math.Pi / 360)
	d.normalOp.GeoM.Translate(getRePosition(halfW, halfH, degree))
	d.normalOp.GeoM.Translate(float64(x), float64(y))

	d.selectedOp = &ebiten.DrawImageOptions{}
	d.selectedOp.GeoM.Add(d.normalOp.GeoM)
	d.selectedOp.ColorM.Scale(colorScale(color.RGBA{0, 148, 255, 255}))

	return d, nil
}

func getRePosition(halfW, halfH float64, degree int) (float64, float64) {
	if (degree/90)%2 != 0 {
		return halfH, halfW
	}
	return halfW, halfH
}

// SelectButton sets the argument for selected flag of this button.
func (d *directionalButton) SelectButton(selected bool) {
	d.selected = selected
}

// Draw draws this button.
func (d *directionalButton) Draw(screen *ebiten.Image) error {
	if d.selected {
		return screen.DrawImage(d.baseImg, d.selectedOp)
	}
	return screen.DrawImage(d.baseImg, d.normalOp)
}

func (d *directionalButton) GetRectangle() image.Rectangle {
	return d.rectangle
}

func colorScale(clr color.Color) (rf, gf, bf, af float64) {
	r, g, b, a := clr.RGBA()
	if a == 0 {
		return 0, 0, 0, 0
	}

	rf = float64(r) / float64(a)
	gf = float64(g) / float64(a)
	bf = float64(b) / float64(a)
	af = float64(a) / 0xffff
	return
}
