package main

import (
	"image"

	"github.com/hajimehoshi/ebiten"
	"github.com/kemokemo/ebiten-sketchbook/chaser/assets/images"
)

// NewChaser returns a new chaser instance.
func NewChaser(x, y int) (*Chaser, error) {
	c := &Chaser{
		point:    image.Point{X: x, Y: y},
		velocity: image.Point{X: 2, Y: 0},
		op:       &ebiten.DrawImageOptions{},
	}

	img, err := images.LoadImage(images.EnemyImage)
	if err != nil {
		return nil, err
	}
	c.image = img

	c.op.GeoM.Translate(float64(c.point.X), float64(c.point.Y))

	return c, nil
}

// Chaser is an enemy charactor to chase the player.
type Chaser struct {
	point    image.Point
	velocity image.Point
	image    *ebiten.Image
	op       *ebiten.DrawImageOptions
}

// Update updates this chaser's internal state.
func (c *Chaser) Update() {
	c.point.Add(c.velocity)
	c.op.GeoM.Translate(float64(c.velocity.X), float64(c.velocity.Y))
}

// Draw draws this chaser.
func (c *Chaser) Draw(screen *ebiten.Image) error {
	screen.DrawImage(c.image, c.op)
	return nil
}
