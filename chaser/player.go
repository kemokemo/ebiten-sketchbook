package main

import (
	"image"

	"github.com/hajimehoshi/ebiten"
	"github.com/kemokemo/ebiten-sketchbook/chaser/assets/images"
)

// NewPlayer returns a new player instance.
func NewPlayer(x, y int) (*Player, error) {
	p := &Player{
		point:    image.Point{X: x, Y: y},
		velocity: image.Point{X: 1, Y: 0},
		op:       &ebiten.DrawImageOptions{},
	}

	img, err := images.LoadImage(images.PlayerImage)
	if err != nil {
		return nil, err
	}
	p.image = img

	p.op.GeoM.Translate(float64(p.point.X), float64(p.point.Y))

	return p, nil
}

// Player is the player charactor.
type Player struct {
	point    image.Point
	velocity image.Point
	image    *ebiten.Image
	op       *ebiten.DrawImageOptions
}

// Update updates this player's internal state.
func (p *Player) Update() {
	p.point.Add(p.velocity)
	p.op.GeoM.Translate(float64(p.velocity.X), float64(p.velocity.Y))
}

// Draw draws this player.
func (p *Player) Draw(screen *ebiten.Image) error {
	screen.DrawImage(p.image, p.op)
	return nil
}
