package scenes

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/kemokemo/ebiten-sketchbook/virtual-gamepad/internal/pad"
)

// GameScene is the scene of the main game screen.
type GameScene struct {
	dpad *pad.DirectionalPad
}

// NewGameScene returns a new GemeScene instance.
func NewGameScene(width, height int) (*GameScene, error) {
	g := &GameScene{}
	d, err := pad.NewDirectionalPad(10, height-110)
	if err != nil {
		return nil, err
	}
	g.dpad = d
	return g, nil
}

// Draw draws the objects contained in this scene.
func (g *GameScene) Draw(screen *ebiten.Image) error {
	return g.dpad.Draw(screen)
}
