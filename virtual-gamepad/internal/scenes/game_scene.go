package scenes

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/kemokemo/ebiten-sketchbook/virtual-gamepad/internal/pad"
)

// GameScene is the scene of the main game screen.
type GameScene struct {
	dpad    *pad.DirectionalPad
	baseImg *ebiten.Image
	op      *ebiten.DrawImageOptions
}

// NewGameScene returns a new GemeScene instance.
func NewGameScene(width, height int) (*GameScene, error) {
	g := &GameScene{}

	err := g.createImage(width, height)
	if err != nil {
		return nil, err
	}

	d, err := pad.NewDirectionalPad()
	if err != nil {
		return nil, err
	}
	d.SetLocation(30, height-150)
	g.dpad = d
	return g, nil
}

func (g *GameScene) createImage(width, height int) error {
	var err error
	g.baseImg, err = ebiten.NewImage(width, height, ebiten.FilterDefault)
	if err != nil {
		return err
	}
	err = g.baseImg.Fill(color.RGBA{27, 24, 44, 255})
	if err != nil {
		return err
	}
	g.op = &ebiten.DrawImageOptions{}

	return nil
}

// Update updates the inner state of this scene.
func (g *GameScene) Update() error {
	return g.dpad.Update()
}

// Draw draws the objects contained in this scene.
func (g *GameScene) Draw(screen *ebiten.Image) error {
	err := screen.DrawImage(g.baseImg, g.op)
	if err != nil {
		return err
	}
	return g.dpad.Draw(screen)
}
