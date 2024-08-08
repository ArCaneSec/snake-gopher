package main

import (
	"os"

	"github.com/gdamore/tcell/v2"
)

type game struct {
	screen tcell.Screen
	style  *tcell.Style
	board  *board
	isOver bool
}

func newGame() *game {
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	s, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	s.SetStyle(defStyle)
	s.Clear()

	return &game{
		screen: s,
		style:  &defStyle,
		board:  &board{width: 100, height: 50},
		isOver: false,
	}
}

func (g *game) quit() {
	g.screen.Show()
	os.Exit(0)
}

func (g *game)drawText(x, y int, style tcell.Style, text string) {
	for _, r := range text {
		g.screen.SetContent(x, y, r, nil, style)
		x++
	}
}

