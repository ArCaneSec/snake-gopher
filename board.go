package main

import (
	"math/rand"
)

type position struct {
	x int
	y int
}

type board struct {
	width  int
	height int
	apple  *position
	nodes  map[position]string
}

func (b *board) draw(g *game) {
	for x := 0; x <= b.width; x++ {
		g.drawText(x, 0, *g.style, "-")
		g.drawText(x, b.height, *g.style, "-")

	}

	for y := 0; y <= b.height; y++ {
		g.drawText(0, y, *g.style, "|")
		g.drawText(b.width, y, *g.style, "|")
	}

	for x := 1; x < b.width; x++ {
		for y := 1; y < b.height; y++ {
			g.drawText(x, y, *g.style, " ")
			g.board.nodes[position{x, y}] = " "
		}
	}

	if b.apple != nil {
		g.drawText(b.apple.x, b.apple.y, *g.style, "@")
		g.board.nodes[position{b.apple.x, b.apple.y}] = "@"
	}
}

func (b *board) addApple(g *game) {
	if b.apple != nil {
		return
	}

	b.apple = &position{x: rand.Intn(b.width-1) + 1, y: rand.Intn(b.height-1) + 1}
}
