package pad

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

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

func isTouched(touchedID int, bounds image.Rectangle) bool {
	x, y := ebiten.TouchPosition(touchedID)
	min := bounds.Min
	max := bounds.Max
	if min.X < x && x < max.X && min.Y < y && y < max.Y {
		return true
	}
	return false
}
