package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/hajimehoshi/ebiten/ebitenutil"

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
	err := ebiten.Run(update, screenWidth, screenHeight, 2, "Touch point")
	if err != nil {
		return err
	}
	return nil
}

func update(screen *ebiten.Image) error {
	message := "touching points:\n"
	IDs := ebiten.TouchIDs()
	sort.Slice(IDs, func(i, j int) bool {
		return IDs[i] < IDs[j]
	})
	for _, id := range IDs {
		x, y := ebiten.TouchPosition(id)
		message = fmt.Sprintf("%v x: %v, y: %v\n", message, x, y)
	}
	ebitenutil.DebugPrint(screen, message)
	return nil
}
