package main

type board struct {
	width  int
	height int
	nodes  []*position
}

func (b *board) draw(g *game) {
	for x := 0; x <= b.width; x++ {
		g.screen.SetContent(x, 0, '-', nil, *g.style)
		g.screen.SetContent(x, b.height, '-', nil, *g.style)

	}

	for y := 0; y <= b.height; y++ {
		g.screen.SetContent(0, y, '|', nil, *g.style)
		g.screen.SetContent(b.width, y, '|', nil, *g.style)
	}

	for x := 1; x < b.width; x++ {
		for y := 1; y < b.height; y++ {
			g.screen.SetContent(x, y, ' ', nil, *g.style)
			b.nodes = append(b.nodes, &position{x, y})
		}
	}

}
