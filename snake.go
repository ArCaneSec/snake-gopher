package main

type snake struct {
	head      *position
	positions []position
	direction int
}

const (
	TOP = iota
	BUTTOM
	LEFT
	RIGHT
)

type errBadPosition struct{}

func (e errBadPosition) Error() string {
	return ""
}

func (s *snake) move() {
	switch s.direction {
	case RIGHT:
		s.head.x += 1
	case LEFT:
		s.head.x -= 1
	case TOP:
		s.head.y -= 1
	case BUTTOM:
		s.head.y += 1
	}
	s.positions = append([]position{*s.head}, s.positions[:len(s.positions)-1]...)
}

func (s *snake) draw(g *game) error {
	for _, position := range s.positions {
		g.drawText(position.x, position.y, *g.style, "0")
		g.board.nodes[position] = "0"
	}
	s.move()
	if !s.wasMoveLegal(g) {
		return errBadPosition{}
	}

	return nil
}

func (s *snake) wasMoveLegal(g *game) bool {
	hittedWall := (s.head.x >= g.board.width || s.head.x <= 0) || (s.head.y >= g.board.height || s.head.y <= 0)
	hittedBody := g.board.nodes[*s.head] == "0"

	return !hittedWall && !hittedBody
}

func (s *snake) canEatApple(g *game) bool {
	return g.board.nodes[*s.head] == "@"
}

func (s *snake) addBody() {
	lastCell := s.positions[len(s.positions)-1]
	switch s.direction {
	case TOP:
		s.positions = append(s.positions, position{lastCell.x, lastCell.y + 1})
	case BUTTOM:
		s.positions = append(s.positions, position{lastCell.x, lastCell.y - 1})
	case RIGHT:
		s.positions = append(s.positions, position{lastCell.x - 1, lastCell.y})
	case LEFT:
		s.positions = append(s.positions, position{lastCell.x + 1, lastCell.y})
	}

}

func (s *snake) changeDirection(newDirection int) {
	if newDirection == s.direction {
		return
	}
	if s.direction == LEFT && newDirection == RIGHT || s.direction == RIGHT && newDirection == LEFT {
		return
	}
	if s.direction == TOP && newDirection == BUTTOM || s.direction == BUTTOM && newDirection == TOP {
		return
	}

	s.direction = newDirection
}

func newSnake() *snake {
	return &snake{
		head:      &position{10, 10},
		positions: []position{{10, 10}, {9, 10}, {8, 10}},
		direction: RIGHT,
	}
}
