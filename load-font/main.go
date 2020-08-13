// Copyright 2016 Hajime Hoshi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build example jsgo

package main

import (
	"fmt"
	"image/color"
	"log"

	"bitbucket.org/KemoKemo/ebiten-font/assets/fonts"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
)

const (
	screenWidth  = 800
	screenHeight = 600
)

var (
	smallFont  font.Face
	normalFont font.Face
	bigFont    font.Face
)

func init() {
	tt, err := truetype.Parse(fonts.TheStrongGamer_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	smallFont = truetype.NewFace(tt, &truetype.Options{
		Size:    16,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	normalFont = truetype.NewFace(tt, &truetype.Options{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	bigFont = truetype.NewFace(tt, &truetype.Options{
		Size:    36,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
}

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	const x = 10

	// Draw info
	msg := fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS())
	text.Draw(screen, msg, smallFont, x, 20, color.White)

	// Draw the sample text
	text.Draw(screen, "オン アルファ ...", bigFont, x, 80, color.White)
	text.Draw(screen, "オン アルファ エト オメガ ...", bigFont, x, 140, color.White)
	text.Draw(screen, "オン アルファ エト オメガ イオト テトラグマン!", bigFont, x, 200, color.White)

	text.Draw(screen, "この すてきフォントは かんじゃちょうひっく さんの TheStrongGamer フォントです。", normalFont, x, 280, color.White)
	text.Draw(screen, "じゅもん は めがみてんせい ラストバイブル GGばん の ランカイン です。", normalFont, x, 310, color.White)

	return nil
}

func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "ebiten-font (The Strong Gamer)"); err != nil {
		log.Fatal(err)
	}
}
