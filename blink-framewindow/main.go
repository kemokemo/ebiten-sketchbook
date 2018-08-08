package main

import (
	"image/color"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten"
	"github.com/kemokemo/ebiten-sketchbook/blink-framewindow/internal/ui"
)

const (
	screenWidth      = 320
	screenHeight     = 240
	exitCodeOK   int = iota
	exitCodeFailed
)

var (
	frameWindow *ui.FrameWindow
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
	err := prepare()
	if err != nil {
		return err
	}
	err = ebiten.Run(update, screenWidth, screenHeight, 2, "Blink frame-window")
	if err != nil {
		return err
	}
	return nil
}

func prepare() error {
	var err error
	frameWindow, err = ui.NewFrameWindow(20, 10, 100, 120, 5)
	if err != nil {
		return err
	}

	frameWindow.SetColors(
		color.RGBA{64, 64, 64, 255},
		color.RGBA{192, 192, 192, 255},
		color.RGBA{0, 148, 255, 255})
	frameWindow.SetBlink(true)
	return nil
}

func update(screen *ebiten.Image) error {
	if ebiten.IsRunningSlowly() {
		return nil
	}

	frameWindow.DrawWindow(screen)
	return nil
}
