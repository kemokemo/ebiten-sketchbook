package gun

import (
	"image"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/kemokemo/ebiten-sketchbook/virtual-gamepad/internal/bullet"
)

const (
	bulletsCount       = 30
	noramalGunInterval = 300 * time.Millisecond
)

// Gun is a gun for fighter.
type Gun struct {
	interval     time.Duration
	fired        time.Time
	bullets      [bulletsCount]*bullet.Bullet
	bulletsIndex int
	bulletSize   image.Point
}

// NewGun returns a new gun.
// Please sets the area of movement for bullets.
func NewGun(area image.Rectangle) (*Gun, error) {
	g := &Gun{interval: time.Duration(noramalGunInterval)}
	err := g.createBullets(area)
	if err != nil {
		return nil, err
	}
	return g, nil
}

func (g *Gun) createBullets(area image.Rectangle) error {
	for index := 0; index < bulletsCount; index++ {
		b, err := bullet.NewBullet(image.Point{0, -3}, area)
		if err != nil {
			return err
		}
		g.bullets[index] = b

		if index == 0 {
			g.bulletSize = b.Size()
		}
	}
	return nil
}

// Update updates the bullets status.
func (g *Gun) Update() {
	for index := 0; index < bulletsCount; index++ {
		g.bullets[index].Update()
	}
}

// Draw draws bullets.
func (g *Gun) Draw(screen *ebiten.Image) error {
	var e error
	for index := 0; index < bulletsCount; index++ {
		e = g.bullets[index].Draw(screen)
		if e != nil {
			return e
		}
	}
	return nil
}

// Fire fires a bullet.
// If the duration time is less than the interval from previous fire,
// this function do nothing.
func (g *Gun) Fire(point image.Point) {
	t := time.Now()
	if t.Sub(g.fired) < g.interval {
		return
	}
	g.fired = t

	if g.bulletsIndex < bulletsCount-1 {
		g.bulletsIndex++
	} else {
		g.bulletsIndex = 0
	}
	g.bullets[g.bulletsIndex].Fire(
		image.Point{
			point.X - g.bulletSize.X/2,
			point.Y - g.bulletSize.Y/2})
}
