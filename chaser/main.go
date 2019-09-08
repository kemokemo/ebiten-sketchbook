package main

import (
	"log"
	"os"

	"github.com/hajimehoshi/ebiten"
)

const (
	width  = 320
	height = 240
	scale  = 2.0
	title  = "chaser!"
)

const (
	exitCodeOK int = iota
	exitCodeFailed
)

var (
	game *Game
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
	var err error
	game, err = NewGame()
	if err != nil {
		return err
	}
	return ebiten.Run(game.Update, width, height, scale, title)
}
