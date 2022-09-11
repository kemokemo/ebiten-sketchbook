package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := NewGame()
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Alpha Blending Image Demo App for Ebiten")
	if err := ebiten.RunGame(g); err != nil {
		log.Println("failed to run game", err)
	}
}
