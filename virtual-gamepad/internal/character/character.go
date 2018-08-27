package character

import (
	"bytes"
	"image"

	"github.com/hajimehoshi/ebiten"
	"github.com/kemokemo/ebiten-sketchbook/virtual-gamepad/internal/bullet"
	"github.com/kemokemo/ebiten-sketchbook/virtual-gamepad/internal/images"
	"github.com/kemokemo/ebiten-sketchbook/virtual-gamepad/internal/pad"
)

const bulletsCount = 30

// Player is the player character.
type Player struct {
	baseImg      *ebiten.Image
	normalOp     *ebiten.DrawImageOptions
	size         image.Point
	area         image.Rectangle
	rectangle    image.Rectangle
	bullets      [bulletsCount]*bullet.Bullet
	bulletsIndex int
	bulletSize   image.Point
}

// NewPlayer returns a new Player.
// Please set the area of movement.
func NewPlayer(area image.Rectangle) (*Player, error) {
	m := &Player{area: area}

	img, _, err := image.Decode(bytes.NewReader(images.Fighter_png))
	if err != nil {
		return nil, err
	}
	m.baseImg, err = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		return nil, err
	}
	w, h := m.baseImg.Size()
	m.rectangle = image.Rect(0, 0, w, h)
	m.size = image.Point{w, h}
	m.normalOp = &ebiten.DrawImageOptions{}

	err = m.createBullets(area)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (m *Player) createBullets(area image.Rectangle) error {
	for index := 0; index < bulletsCount; index++ {
		b, err := bullet.NewBullet(image.Point{0, -3}, area)
		if err != nil {
			return err
		}
		m.bullets[index] = b

		if index == 0 {
			m.bulletSize = b.Size()
		}
	}
	return nil
}

// SetLocation sets the location to draw this character.
func (m *Player) SetLocation(point image.Point) {
	sub := point.Sub(m.rectangle.Min)
	m.rectangle.Min = point
	m.rectangle.Max = m.rectangle.Max.Add(sub)

	m.normalOp.GeoM.Reset()
	m.normalOp.GeoM.Translate(float64(point.X), float64(point.Y))
}

// Size returns the size of this character.
func (m *Player) Size() image.Point {
	return m.size
}

// Update updates the internal state.
// Please pass the direction of the pad to move this character.
func (m *Player) Update(direction pad.Direction) {
	// TODO: Make a judgment with enemy bullets
	m.move(direction)
	for index := 0; index < bulletsCount; index++ {
		m.bullets[index].Update()
	}
}

// Draw draws this character.
func (m *Player) Draw(screen *ebiten.Image) error {
	err := screen.DrawImage(m.baseImg, m.normalOp)
	if err != nil {
		return err
	}

	var e error
	for index := 0; index < bulletsCount; index++ {
		e = m.bullets[index].Draw(screen)
		if e != nil {
			return e
		}
	}
	return nil
}

// move moves this character regarding the direction.
// Do not move if the destination is outside the area.
func (m *Player) move(d pad.Direction) {
	switch d {
	case pad.None:
		return
	case pad.UpperLeft:
		m.move4direction(pad.Upper)
		m.move4direction(pad.Left)
	case pad.UpperRight:
		m.move4direction(pad.Upper)
		m.move4direction(pad.Right)
	case pad.LowerLeft:
		m.move4direction(pad.Lower)
		m.move4direction(pad.Left)
	case pad.LowerRight:
		m.move4direction(pad.Lower)
		m.move4direction(pad.Right)
	default:
		m.move4direction(d)
	}
}

func (m *Player) move4direction(d pad.Direction) {
	movement := m.getMove((d))
	moved := m.rectangle.Add(movement)
	if !moved.In(m.area) {
		return
	}

	m.rectangle = moved
	m.normalOp.GeoM.Translate(float64(movement.X), float64(movement.Y))
}

func (m *Player) getMove(d pad.Direction) image.Point {
	switch d {
	case pad.Left:
		return image.Point{-2, 0}
	case pad.Upper:
		return image.Point{0, -2}
	case pad.Right:
		return image.Point{2, 0}
	case pad.Lower:
		return image.Point{0, 2}
	default:
		return image.Point{0, 0}
	}
}

// Fire fires some bullets.
func (m *Player) Fire() {
	if m.bulletsIndex < bulletsCount-1 {
		m.bulletsIndex++
	} else {
		m.bulletsIndex = 0
	}
	m.bullets[m.bulletsIndex].Fire(
		image.Point{
			m.rectangle.Min.X + m.size.X/2 - m.bulletSize.X/2,
			m.rectangle.Min.Y - m.bulletSize.Y/2})
}
