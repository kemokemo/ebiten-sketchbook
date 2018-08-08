package main

import (
	"log"
	"os"

	"github.com/hajimehoshi/ebiten"
)

const (
	screenWidth      = 320
	screenHeight     = 240
	exitCodeOK   int = iota
	exitCodeFailed
)

var (
	selectScene *SelectScene
)

func main() {
	os.Exit(run(os.Args))
}

func run(args []string) int {
	err := work()
	if err != nil {
		log.Println("Failed to execute the work:", err)
		return exitCodeFailed
	}
	return exitCodeOK
}

func work() error {
	err := ebiten.Run(update, screenWidth, screenHeight, 2, "Select one")
	if err != nil {
		return err
	}
	return nil
}

func update(screen *ebiten.Image) error {
	if selectScene == nil {
		selectScene = NewSelectScene()
	}

	err := selectScene.Update()
	if err != nil {
		return err
	}
	selectScene.Draw(screen)

	return nil
}
