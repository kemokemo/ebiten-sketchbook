package scenes

import (
	"bytes"
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/kemokemo/ebiten-sketchbook/virtual-gamepad/internal/character"
	"github.com/kemokemo/ebiten-sketchbook/virtual-gamepad/internal/gun"
	"github.com/kemokemo/ebiten-sketchbook/virtual-gamepad/internal/images"
	"github.com/kemokemo/ebiten-sketchbook/virtual-gamepad/internal/pad"
	"github.com/kemokemo/ebiten-sketchbook/virtual-gamepad/internal/ui"
)

// GameScene is the scene of the main game screen.
type GameScene struct {
	dpad      *pad.DirectionalPad
	aButton   pad.TriggerButton
	bButton   pad.TriggerButton
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
	img, _, err := image.Decode(bytes.NewReader(images.Directional_pad_png))
	if err != nil {
		return err
	}
	padImg, err := ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		return err
	}
	img, _, err = image.Decode(bytes.NewReader(images.Directional_button_png))
	if err != nil {
		return err
	}
	buttonImg, err := ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		return err
	}
	g.dpad, err = pad.NewDirectionalPad(padImg, buttonImg)
	if err != nil {
		return err
	}
	g.dpad.SetLocation(10, height-130)

	img, _, err = image.Decode(bytes.NewReader(images.A_button_png))
	if err != nil {
		return err
	}
	aButton, err := ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		return err
	}
	g.aButton, err = pad.NewTriggerButton(aButton, pad.JustRelease)
	if err != nil {
		return err
	}
	g.aButton.SetLocation(width-120, height-220)

	img, _, err = image.Decode(bytes.NewReader(images.B_button_png))
	if err != nil {
		return err
	}
	bButton, err := ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		return err
	}
	g.bButton, err = pad.NewTriggerButton(bButton, pad.Pressing)
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
	img, _, err := image.Decode(bytes.NewReader(images.Bullet_png))
	if err != nil {
		return err
	}
	bImg, err := ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		return err
	}
	gun, err := gun.NewGun(bImg, area)
	if err != nil {
		return err
	}

	img, _, err = image.Decode(bytes.NewReader(images.Fighter_png))
	if err != nil {
		return err
	}
	pImg, err := ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		return err
	}

	g.character, err = character.NewPlayer(pImg, area, gun)
	if err != nil {
		return err
	}
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
	g.aButton.Update()
	if g.aButton.IsTriggered() {
		g.character.ChangeMode()
	}
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
