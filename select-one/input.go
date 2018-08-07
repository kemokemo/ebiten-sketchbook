package main

import (
	"image"
	"sort"

	"github.com/hajimehoshi/ebiten"
)

// Touched returns the state that touched or not.
func Touched(r image.Rectangle) bool {
	IDs := ebiten.TouchIDs()
	if len(IDs) == 0 {
		return false
	}

	// use only the first touched point
	sort.Slice(IDs, func(i, j int) bool {
		return IDs[i] < IDs[j]
	})
	x, y := ebiten.TouchPosition(IDs[0])

	min := r.Min
	max := r.Max
	if min.X < x && x < max.X && min.Y < y && y < max.Y {
		return true
	}
	return false
}
