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

// MainCharacter is the main character for users to control.
type MainCharacter struct {
	baseImg      *ebiten.Image
	normalOp     *ebiten.DrawImageOptions
	size         image.Point
	area         image.Rectangle
	point        image.Point
	bullets      [bulletsCount]*bullet.Bullet
	bulletsIndex int
}

// NewMainCharacter returns a character.
// Please set the area of movement.
func NewMainCharacter(area image.Rectangle) (*MainCharacter, error) {
	m := &MainCharacter{area: area}

	img, _, err := image.Decode(bytes.NewReader(images.Fighter_png))
	if err != nil {
		return nil, err
	}
	m.baseImg, err = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		return nil, err
	}
	w, h := m.baseImg.Size()
	m.size = image.Point{X: w, Y: h}
	m.normalOp = &ebiten.DrawImageOptions{}

	err = m.createBullets(area)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (m *MainCharacter) createBullets(area image.Rectangle) error {
	for index := 0; index < bulletsCount; index++ {
		b, err := bullet.NewBullet(image.Point{0, -3}, area)
		if err != nil {
			return err
		}
		m.bullets[index] = b
	}
	return nil
}

// SetLocation sets the location to draw this character.
func (m *MainCharacter) SetLocation(point image.Point) {
	m.point = point

	m.normalOp.GeoM.Reset()
	m.normalOp.GeoM.Translate(float64(m.point.X), float64(m.point.Y))
}

// Size returns the size of this character.
func (m *MainCharacter) Size() image.Point {
	return m.size
}

// Update updates the internal state.
// Please pass the direction of the pad to move this character.
func (m *MainCharacter) Update(direction pad.Direction) {
	// TODO: Make a judgment with enemy bullets
	m.move(direction)
	for index := 0; index < bulletsCount; index++ {
		m.bullets[index].Update()
	}
}

// Draw draws this character.
func (m *MainCharacter) Draw(screen *ebiten.Image) error {
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
func (m *MainCharacter) move(d pad.Direction) {
	switch d {
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

func (m *MainCharacter) move4direction(d pad.Direction) {
	movement := m.getMove((d))
	moved := m.point.Add(movement)
	if !moved.In(m.area) {
		return
	}
	m.point = moved
	m.normalOp.GeoM.Translate(float64(movement.X), float64(movement.Y))
}

func (m *MainCharacter) getMove(d pad.Direction) image.Point {
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
func (m *MainCharacter) Fire() {
	if m.bulletsIndex < bulletsCount-1 {
		m.bulletsIndex++
	} else {
		m.bulletsIndex = 0
	}
	m.bullets[m.bulletsIndex].Fire(image.Point{
		m.point.X + m.size.X/2, m.point.Y})
}
