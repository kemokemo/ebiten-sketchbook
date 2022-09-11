package main

import (
	"bytes"
	"image"
	"image/color"
	"log"

	_ "image/png" // to load png images

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 480
	screenHeight = 320
)

func NewGame() *Game {
	bkImg, _, err := image.Decode(bytes.NewReader(bk_png))
	if err != nil {
		log.Println("failed to load dot png,", err)
	}

	bk := ebiten.NewImageFromImage(bkImg)
	colorS := color.RGBA{R: 50, G: 50, B: 50, A: 200}
	img := ebiten.NewImage(20, 20)
	img.Fill(colorS)
	squares := []square{}

	x := 15.0
	y := 70.0
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	squares = append(squares, square{img: img, op: op})

	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x+40.0, y)
	op.ColorM.Scale(1.0, 1.0, 1.0, 0.5)
	squares = append(squares, square{img: img, op: op})

	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x+85.0, y)
	op.ColorM.Translate(0.1, 0.1, 0.2, 1.0)
	squares = append(squares, square{img: img, op: op})

	alpha := 0.0
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y+45.0)
	op.ColorM.Scale(1.0, 1.0, 1.0, alpha)
	squares = append(squares, square{img: img, op: op})

	return &Game{
		bk:         bk,
		squares:    squares,
		alpha:      alpha,
		counter:    0,
		increasing: true,
	}
}

type Game struct {
	bk         *ebiten.Image
	squares    []square
	alpha      float64
	counter    int
	increasing bool
}

type square struct {
	img *ebiten.Image
	op  *ebiten.DrawImageOptions
}

func (g *Game) Update() error {
	if g.increasing && g.counter < 10 {
		g.counter++
	} else if g.increasing && g.counter == 10 {
		g.increasing = false
		g.counter--
	} else if !g.increasing && 0 < g.counter {
		g.counter--
	} else {
		g.increasing = true
		g.counter++
	}
	g.alpha = 0.1 * float64(g.counter)

	// 4th square blinks.
	g.squares[3].op.ColorM.Reset()
	g.squares[3].op.ColorM.Scale(1.0, 1.0, 1.0, g.alpha)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.bk, &ebiten.DrawImageOptions{})
	for i := range g.squares {
		screen.DrawImage(g.squares[i].img, g.squares[i].op)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
