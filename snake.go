package main

type snake struct {
	head      *position
	positions []position
	direction rune
}

type errBadPosition struct{}

func (e errBadPosition) Error() string {
	return ""
}

func (s *snake) move() {
	switch s.direction {
	case 'r':
		s.head.x += 1
	case 'l':
		s.head.x -= 1
	case 't':
		s.head.y -= 1
	case 'b':
		s.head.y += 1
	}
	s.positions = append([]position{*s.head}, s.positions[:len(s.positions)-1]...)
}

func (s *snake) draw(g *game, width, height int) error {
	s.move()
	for _, position := range s.positions {
		g.drawText(position.x, position.y, *g.style, "0")
	}
	if !s.wasMoveLegal(width, height) {
		return errBadPosition{}
	}
	return nil
}

func (s *snake) wasMoveLegal(width, height int) bool {
	return (s.head.x < width && s.head.x > 0) && (s.head.y < height && s.head.y > 0)
}

func newSnake() *snake {
	return &snake{
		head:      &position{10, 10},
		positions: []position{{10, 10}, {9, 10}, {8, 10}},
		direction: 'r',
	}
}
