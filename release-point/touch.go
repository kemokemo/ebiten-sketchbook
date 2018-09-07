package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type touch struct {
	id int
	x  int
	y  int
}

func (t *touch) Update() {
	if inpututil.IsTouchJustReleased(t.id) {
		return
	}
	t.x, t.y = ebiten.TouchPosition(t.id)
}

func (t *touch) IsReleased() bool {
	return inpututil.IsTouchJustReleased(t.id)
}
