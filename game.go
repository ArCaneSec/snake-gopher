package main

import (
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
)

type game struct {
	screen tcell.Screen
	style  *tcell.Style
	board  *board
	isOver bool
	score  int
	snake *snake
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
		board:  &board{width: 100, height: 50, nodes: make(map[position]string)},
		isOver: false,
		snake: newSnake(),
	}
}

func (g *game) quit() {
	g.screen.Show()
	os.Exit(0)
}

func (g *game) drawText(x, y int, style tcell.Style, text string) {
	for _, r := range text {
		g.screen.SetContent(x, y, r, nil, style)
		x++
	}
}

func (g *game) run() {
	for {
		g.screen.Show()
		g.board.draw(g)
		g.board.addApple(g)
		err := g.snake.draw(g)
		if err != nil {
			g.drawText(50, 25, g.style.Foreground(tcell.ColorDarkRed), "You Died!")
			g.quit()
		}
		if g.snake.canEatApple(g) {
			g.score += 1
			g.board.apple = nil
			g.snake.addBody()
		}

		time.Sleep(time.Duration(100 - g.score * 1) * time.Millisecond)
	}
}

func (g *game) start() {
	go func() {
		g.run()
	}()
	for {
		ev := g.screen.PollEvent()
		if ev != nil {
			switch ev := ev.(type) {
			case *tcell.EventResize:
				g.screen.Sync()

			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyCtrlC:
					g.quit()
				case tcell.KeyUp:
					g.snake.changeDirection(TOP)
				case tcell.KeyDown:
					g.snake.changeDirection(BUTTOM)
				case tcell.KeyRight:
					g.snake.changeDirection(RIGHT)
				case tcell.KeyLeft:
					g.snake.changeDirection(LEFT)
				}
			}
		}
	}
}
