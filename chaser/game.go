package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
)

// Game is a game object.
type Game struct {
	state  GameState
	enemy  *Chaser
	player *Player
}

// NewGame returns a new game instance.
func NewGame() (*Game, error) {
	g := &Game{}
	var err error

	g.enemy, err = NewChaser(10, 10)
	if err != nil {
		return nil, fmt.Errorf("failed to create chaser,%v", err)
	}
	g.player, err = NewPlayer(80, 10)
	if err != nil {
		return nil, fmt.Errorf("failed to create player,%v", err)
	}

	return g, nil
}

// Update updates the game screen.
func (g *Game) Update(screen *ebiten.Image) error {
	g.updateState()
	// TODO: Check game over or goal in.

	if ebiten.IsRunningSlowly() {
		return nil
	}
	return game.draw(screen)
}

func (g *Game) updateState() {
	g.enemy.Update()
	g.player.Update()
}

func (g *Game) draw(screen *ebiten.Image) error {
	err := g.enemy.Draw(screen)
	if err != nil {
		return err
	}
	err = g.player.Draw(screen)
	if err != nil {
		return err
	}
	return nil
}
