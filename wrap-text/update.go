package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	mplus "github.com/hajimehoshi/go-mplusbitmap"
	"golang.org/x/image/font"
)

const (
	screenWidth  = 640
	screenHeight = 480
	margin       = 10
	fontSize     = 12
	lineSpacing  = 2
	bgWidth      = 150
	bgHeight     = 150
)

var (
	mplusNormalFont font.Face
	background      *ebiten.Image
)

func update(screen *ebiten.Image) error {
	drawWindow(screen)

	mplusNormalFont = mplus.Gothic12r
	msg := "これはテスト用のテキストです。折り返して表示するためのプラクティスです。"
	splitlen := (bgWidth - margin) / fontSize
	drawMessage(screen, msg, splitlen)
	return nil
}

func drawWindow(screen *ebiten.Image) error {
	background, err := ebiten.NewImage(bgWidth, bgHeight, ebiten.FilterDefault)
	if err != nil {
		return err
	}
	err = background.Fill(color.RGBA{64, 64, 64, 255})
	if err != nil {
		return err
	}
	op := &ebiten.DrawImageOptions{}
	return screen.DrawImage(background, op)
}

func drawMessage(screen *ebiten.Image, msg string, splitlen int) {
	runes := []rune(msg)
	lineNum := 1
	for i := 0; i < len(runes); i += splitlen {
		y := (fontSize+lineSpacing)*lineNum + margin
		if i+splitlen < len(runes) {
			text.Draw(screen, string(runes[i:(i+splitlen)]), mplusNormalFont, margin, y, color.White)
		} else {
			text.Draw(screen, string(runes[i:]), mplusNormalFont, margin, y, color.White)
		}
		lineNum++
	}
}
