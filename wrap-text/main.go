package main

import (
	"fmt"
	"os"

	"github.com/hajimehoshi/ebiten"
)

const (
	exitCodeOK int = iota
	exitCodeFailed
)

func main() {
	os.Exit(run(os.Args))
}

func run(args []string) int {
	err := ebiten.Run(update, screenWidth, screenHeight, 1, "tryrune")
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to run ebiten game: ", err)
		return exitCodeFailed
	}
	return exitCodeOK
}
