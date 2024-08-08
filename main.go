package main

import (
	"time"

	"github.com/gdamore/tcell/v2"
)

type position struct {
	x int
	y int
}

func main() {
	game := newGame()
	snake := newSnake()

	// Event loop
	go func() {
		for {
			game.screen.Show()
			game.board.draw(game)
			err := snake.draw(game, 100, 50)
			if err != nil {
				game.drawText(50, 25, game.style.Foreground(tcell.ColorDarkRed), "You Died!")
				game.quit()
			}

			time.Sleep(100 * time.Millisecond)
		}
	}()
	for {
		ev := game.screen.PollEvent()
		if ev != nil {

			// Process event
			switch ev := ev.(type) {
			case *tcell.EventResize:
				game.screen.Sync()

			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyCtrlC:
					game.quit()
				case tcell.KeyUp:
					snake.direction = 't'
				case tcell.KeyDown:
					snake.direction = 'b'
				case tcell.KeyRight:
					snake.direction = 'r'
				case tcell.KeyLeft:
					snake.direction = 'l'
				}
			}
		}
	}
}
