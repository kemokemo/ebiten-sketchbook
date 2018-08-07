package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/kemokemo/ebiten-sketchbook/select-one/internal/ui"
)

const (
	frameWidth    = 2
	margin        = 20
	scale         = 2
	windowSpacing = 15
	windowMargin  = 20
	fontSize      = 12
	lineSpacing   = 2
)

// SelectScene is the select scene.
type SelectScene struct {
	windowMap map[int]*ui.FrameWindow
	selector  int
}

// NewSelectScene returns a new SelectScene instance.
func NewSelectScene() *SelectScene {
	num := 3
	s := SelectScene{}
	s.windowMap = make(map[int]*ui.FrameWindow, num)

	windowWidth := (screenWidth - windowSpacing*2 - windowMargin*2) / num
	windowHeight := screenHeight - windowMargin*2 - 100
	for index := 0; index < num; index++ {
		win, err := ui.NewFrameWindow(
			windowMargin+(windowWidth+windowSpacing)*index,
			windowMargin*2, windowWidth, windowHeight, frameWidth)
		if err != nil {
			log.Println("failed to create a new frame window", err)
		}
		win.SetColors(
			color.RGBA{64, 64, 64, 255},
			color.RGBA{192, 192, 192, 255},
			color.RGBA{0, 148, 255, 255})
		if index == 0 {
			s.selector = 0
			win.SetBlink(true)
		}
		s.windowMap[index] = win
	}

	return &s
}

// Update updates the selection state.
func (s *SelectScene) Update() error {
	if ebiten.IsRunningSlowly() {
		return nil
	}
	s.checkSelectorChanged()
	return nil
}

// Draw draws the select windows.
func (s *SelectScene) Draw(r *ebiten.Image) {
	for index := range s.windowMap {
		if index == s.selector {
			s.windowMap[index].SetBlink(true)
		} else {
			s.windowMap[index].SetBlink(false)
		}
		s.windowMap[index].DrawWindow(r)
	}
}

func (s *SelectScene) checkSelectorChanged() {
	for index := range s.windowMap {
		if Touched(s.windowMap[index].GetWindowRect()) {
			s.selector = index
			return
		}
	}
}
