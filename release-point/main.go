package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"

	"github.com/hajimehoshi/ebiten"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

const (
	exitCodeOK int = iota
	exitCodeFailed
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
	err := ebiten.Run(update, screenWidth, screenHeight, 2, "Release point")
	if err != nil {
		return err
	}
	return nil
}

var (
	message = ""
	touches = make(map[*touch]struct{})
)

func update(screen *ebiten.Image) error {
	IDs := inpututil.JustPressedTouchIDs()
	if len(IDs) != 0 {
		for _, id := range IDs {
			touches[&touch{id: id}] = struct{}{}
			fmt.Printf("Pressed IDs: %v.\n", id)
		}
	}

	for t := range touches {
		t.Update()
		if t.IsReleased() {
			fmt.Printf("ID %v was released.\n", t.id)
			message = fmt.Sprintf("Released point: (%v, %v)\n", t.x, t.y)
			delete(touches, t)
		}
	}

	ebitenutil.DebugPrint(screen, message)

	return nil
}
