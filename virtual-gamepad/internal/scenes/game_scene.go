package scenes

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/kemokemo/ebiten-sketchbook/virtual-gamepad/internal/character"
	"github.com/kemokemo/ebiten-sketchbook/virtual-gamepad/internal/pad"
	"github.com/kemokemo/ebiten-sketchbook/virtual-gamepad/internal/ui"
)

// GameScene is the scene of the main game screen.
type GameScene struct {
	dpad      *pad.DirectionalPad
	aButton   *pad.TriggerButton
	bButton   *pad.TriggerButton
	baseImg   *ebiten.Image
	op        *ebiten.DrawImageOptions
	window    *ui.FrameWindow
	character *character.Player
}

// NewGameScene returns a new GemeScene instance.
func NewGameScene(width, height int) (*GameScene, error) {
	g := &GameScene{}
	err := g.createImage(width, height)
	if err != nil {
		return nil, err
	}

	err = g.createButtons(width, height)
	if err != nil {
		return nil, err
	}

	err = g.createWindow(width, height)
	if err != nil {
		return nil, err
	}

	err = g.createCharacter(g.window.GetWindowRect())
	if err != nil {
		return nil, err
	}

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

func (g *GameScene) createButtons(width, height int) error {
	var err error
	g.dpad, err = pad.NewDirectionalPad()
	if err != nil {
		return err
	}
	g.dpad.SetLocation(10, height-130)

	g.aButton, err = pad.NewTriggerButton(pad.AButton)
	if err != nil {
		return err
	}
	g.aButton.SetLocation(width-120, height-220)

	g.bButton, err = pad.NewTriggerButton(pad.BButton)
	if err != nil {
		return err
	}
	g.bButton.SetLocation(width-120, height-110)

	return nil
}

func (g *GameScene) createWindow(width, height int) error {
	var err error
	margin := 12
	g.window, err = ui.NewFrameWindow(140, margin, width-140*2, height-margin*2, 2)
	if err != nil {
		return err
	}
	g.window.SetColors(
		color.RGBA{27, 24, 44, 255},
		color.RGBA{255, 255, 255, 255},
		color.RGBA{0, 148, 255, 255},
	)
	return nil
}

func (g *GameScene) createCharacter(area image.Rectangle) error {
	c, err := character.NewPlayer(area)
	if err != nil {
		return err
	}
	g.character = c

	cSize := g.character.Size()
	wSize := area.Size()
	g.character.SetLocation(
		image.Point{
			X: area.Min.X + wSize.X/2 - cSize.X/2,
			Y: area.Max.Y - cSize.Y - 2,
		})
	return nil
}

// Update updates the inner state of this scene.
func (g *GameScene) Update() error {
	err := g.dpad.Update()
	if err != nil {
		return err
	}
	g.character.Update(g.dpad.GetDirection())

	// TODO: test code
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.character.Update(pad.Left)
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.character.Update(pad.Right)
	}

	g.aButton.Update()
	g.bButton.Update()
	if g.bButton.IsTriggered() {
		g.character.Fire()
	}
	return nil
}

// Draw draws the objects contained in this scene.
func (g *GameScene) Draw(screen *ebiten.Image) error {
	err := screen.DrawImage(g.baseImg, g.op)
	if err != nil {
		return err
	}

	err = g.aButton.Draw(screen)
	if err != nil {
		return err
	}
	err = g.bButton.Draw(screen)
	if err != nil {
		return err
	}

	err = g.dpad.Draw(screen)
	if err != nil {
		return err
	}

	g.window.DrawWindow(screen)
	err = g.character.Draw(screen)
	if err != nil {
		return err
	}

	return nil
}
