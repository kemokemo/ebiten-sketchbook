package bullet

import (
	"bytes"
	"image"

	"github.com/hajimehoshi/ebiten"
	"github.com/kemokemo/ebiten-sketchbook/virtual-gamepad/internal/images"
)

// Bullet is the bullet triggered by characters.
type Bullet struct {
	baseImg  *ebiten.Image
	op       *ebiten.DrawImageOptions
	visible  bool
	point    image.Point
	velocity image.Point
	area     image.Rectangle
	size     image.Point
}

// NewBullet returns a new Bullet.
// Please set the velocity of this bullet and the area of movement.
func NewBullet(velocity image.Point, area image.Rectangle) (*Bullet, error) {
	b := &Bullet{velocity: velocity, area: area}
	img, _, err := image.Decode(bytes.NewReader(images.Bullet_png))
	if err != nil {
		return nil, err
	}
	b.baseImg, err = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		return nil, err
	}
	w, h := b.baseImg.Size()
	b.size = image.Point{w, h}

	b.op = &ebiten.DrawImageOptions{}
	return b, nil
}

// Update update the internal state of this bullet.
func (b *Bullet) Update() {
	if !b.visible {
		return
	}
	b.move()
	b.checkArea()
}

// move moves this bullet.
func (b *Bullet) move() {
	b.point = b.point.Add(b.velocity)
	b.op.GeoM.Translate(float64(b.velocity.X), float64(b.velocity.Y))
}

// checkAreac checks whether this bullet is out of the
func (b *Bullet) checkArea() {
	if !b.point.In(b.area) {
		b.visible = false
	}
}

// Draw draws the image of this bullet.
func (b *Bullet) Draw(screen *ebiten.Image) error {
	if !b.visible {
		return nil
	}
	return screen.DrawImage(b.baseImg, b.op)
}

// Fire sets the initial position and make this bullet
func (b *Bullet) Fire(point image.Point) {
	b.point = point
	b.visible = true

	b.op.GeoM.Reset()
	b.op.GeoM.Translate(float64(b.point.X), float64(b.point.Y))
}

// GetRectangle returns the ractangle of this bullet.
func (b *Bullet) GetRectangle() image.Rectangle {
	return b.op.SourceRect.Bounds()
}

// Size returns the size of this bullet.
func (b *Bullet) Size() image.Point {
	return b.size
}
