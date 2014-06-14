package main

type Position struct {
	x, y int
}

func NewPosition(x, y int) *Position {
	return &Position{x: x, y: y}
}

func (p *Position) isOnBoard() bool {
	return (0 <= p.x && p.x < boardWidth) && (0 <= p.y && p.y < boardHeight)
}
