package pad

import (
	"bytes"
	"image"
	"math"

	"image/color"

	"github.com/hajimehoshi/ebiten"

	"github.com/kemokemo/ebiten-sketchbook/virtual-gamepad/internal/images"
)

const (
	longMargin  = 40
	shortMargin = 20
)

// debugirectionalButton is the button of the directional pad.
type directionalButton struct {
	baseImg    *ebiten.Image
	selected   bool
	direction  Direction
	rectangle  image.Rectangle
	normalOp   *ebiten.DrawImageOptions
	selectedOp *ebiten.DrawImageOptions
}

// newDirectionalButton returns a new DirectionalButton.
func newDirectionalButton(direc Direction) (*directionalButton, error) {
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
	halfW := float64(w) / 2
	halfH := float64(h) / 2

	d.direction = direc
	degree := getDirectionDegree(direc)

	d.normalOp = &ebiten.DrawImageOptions{}
	d.normalOp.GeoM.Translate(-halfW, -halfH)
	d.normalOp.GeoM.Rotate(float64(degree) * 2 * math.Pi / 360)
	d.normalOp.GeoM.Translate(getRePosition(halfW, halfH, degree))

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

func (d *directionalButton) SetLocation(x, y int) {
	d.calcRectangle(x, y)

	d.normalOp.GeoM.Translate(float64(x), float64(y))
	d.selectedOp.GeoM.Translate(float64(x), float64(y))
}

func (d *directionalButton) calcRectangle(x, y int) {
	w, h := d.baseImg.Size()

	switch d.direction {
	case Left:
		d.rectangle = image.Rect(x-h, y-w, x+h, y+w*2)
	case Right:
		d.rectangle = image.Rect(x, y-w, x+h*2, y+w*2)
	case Up:
		d.rectangle = image.Rect(x-w, y-h, x+w*2, y+h)
	case Down:
		d.rectangle = image.Rect(x-w, y, x+w*2, y+h*2)
	}
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

func (d *directionalButton) Size() (width int, height int) {
	return d.baseImg.Size()
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
