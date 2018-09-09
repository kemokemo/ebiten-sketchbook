package character

import (
	"image"

	"github.com/hajimehoshi/ebiten"
	"github.com/kemokemo/ebiten-sketchbook/virtual-gamepad/internal/gun"
	"github.com/kemokemo/ebiten-sketchbook/virtual-gamepad/internal/pad"
)

// Player is the player character.
type Player struct {
	img        *ebiten.Image
	baseImg    *ebiten.Image
	anotherImg *ebiten.Image
	op         *ebiten.DrawImageOptions
	size       image.Point
	area       image.Rectangle
	rectangle  image.Rectangle
	gun        *gun.Gun
	baseGun    *gun.Gun
	anotherGun *gun.Gun
}

// NewPlayer returns a new Player.
// Please set the area of movement.
func NewPlayer(img, anotherImg *ebiten.Image, area image.Rectangle, bGun, aGun *gun.Gun) (*Player, error) {
	w, h := img.Size()
	return &Player{
		img:        img,
		baseImg:    img,
		anotherImg: anotherImg,
		op:         &ebiten.DrawImageOptions{},
		size:       image.Point{w, h},
		area:       area,
		rectangle:  image.Rect(0, 0, w, h),
		gun:        bGun,
		baseGun:    bGun,
		anotherGun: aGun,
	}, nil
}

// SetLocation sets the location to draw this character.
func (p *Player) SetLocation(point image.Point) {
	sub := point.Sub(p.rectangle.Min)
	p.rectangle.Min = point
	p.rectangle.Max = p.rectangle.Max.Add(sub)

	p.op.GeoM.Reset()
	p.op.GeoM.Translate(float64(point.X), float64(point.Y))
}

// Size returns the size of this character.
func (p *Player) Size() image.Point {
	return p.size
}

// Update updates the internal state.
// Please pass the direction of the pad to move this character.
func (p *Player) Update(direction pad.Direction) {
	// TODO: Make a judgment with enemy bullets
	p.move(direction)
	p.baseGun.Update()
	p.anotherGun.Update()
}

// Draw draws this character.
func (p *Player) Draw(screen *ebiten.Image) error {
	err := screen.DrawImage(p.img, p.op)
	if err != nil {
		return err
	}

	err = p.baseGun.Draw(screen)
	if err != nil {
		return err
	}
	err = p.anotherGun.Draw(screen)
	if err != nil {
		return err
	}

	return nil
}

// move moves this character regarding the direction.
// Do not move if the destination is outside the area.
func (p *Player) move(d pad.Direction) {
	switch d {
	case pad.None:
		return
	case pad.UpperLeft:
		p.move4direction(pad.Upper)
		p.move4direction(pad.Left)
	case pad.UpperRight:
		p.move4direction(pad.Upper)
		p.move4direction(pad.Right)
	case pad.LowerLeft:
		p.move4direction(pad.Lower)
		p.move4direction(pad.Left)
	case pad.LowerRight:
		p.move4direction(pad.Lower)
		p.move4direction(pad.Right)
	default:
		p.move4direction(d)
	}
}

func (p *Player) move4direction(d pad.Direction) {
	movement := p.getMove((d))
	moved := p.rectangle.Add(movement)
	if !moved.In(p.area) {
		return
	}

	p.rectangle = moved
	p.op.GeoM.Translate(float64(movement.X), float64(movement.Y))
}

func (p *Player) getMove(d pad.Direction) image.Point {
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

// Fire fires onley selected gun.
func (p *Player) Fire() {
	p.gun.Fire(image.Point{
		p.rectangle.Min.X + p.size.X/2,
		p.rectangle.Min.Y})
}

// ChangeMode changes the image and the gun of this character.
func (p *Player) ChangeMode() {
	if p.img == p.baseImg {
		p.img = p.anotherImg
	} else {
		p.img = p.baseImg
	}

	if p.gun == p.baseGun {
		p.gun = p.anotherGun
	} else {
		p.gun = p.baseGun
	}
}
