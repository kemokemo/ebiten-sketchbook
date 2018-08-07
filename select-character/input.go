package main

import (
	"image"
	"sort"

	"github.com/hajimehoshi/ebiten"
)

type input struct {
}

func (i *input) Touched(r image.Rectangle) bool {
	min := r.Min
	max := r.Max

	IDs := ebiten.TouchIDs()
	sort.Slice(IDs, func(i, j int) bool {
		return IDs[i] < IDs[j]
	})
	for _, id := range IDs {
		x, y := ebiten.TouchPosition(id)
		if min.X < x && x < max.X && min.Y < y && y < max.Y {
			return true
		}
	}
	return false
}
