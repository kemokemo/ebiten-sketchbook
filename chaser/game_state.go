package main

// GameState is the state of this game.
type GameState int

const (
	prepared GameState = iota
	running
	goal
	gameover
)
