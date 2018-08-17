package main

import (
	"log"
	"os"

	_ "image/png"

	"github.com/hajimehoshi/ebiten"
	"github.com/kemokemo/ebiten-sketchbook/virtual-gamepad/internal/scenes"
)

const (
	screenWidth      = 640
	screenHeight     = 360
	exitCodeOK   int = iota
	exitCodeFailed
)

var gameScene *scenes.GameScene

func main() {
	os.Exit(run(os.Args))
}

func run(args []string) int {
	err := prepare()
	if err != nil {
		log.Println("Failed to preparing:", err)
		return exitCodeFailed
	}

	err = work()
	if err != nil {
		log.Println("Failed to execute the work:", err)
		return exitCodeFailed
	}
	return exitCodeOK
}

func work() error {
	err := ebiten.Run(update, screenWidth, screenHeight, 2, "Virtual gamepad")
	if err != nil {
		return err
	}
	return nil
}

func update(screen *ebiten.Image) error {
	if ebiten.IsRunningSlowly() {
		return nil
	}
	err := gameScene.Draw(screen)
	if err != nil {
		return err
	}
	return nil
}

func prepare() error {
	var err error
	gameScene, err = scenes.NewGameScene(screenWidth, screenHeight)
	if err != nil {
		return err
	}
	return nil
}
