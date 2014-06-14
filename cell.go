package main

type Cell struct {
	x, y int
}

func NewCell(x, y int) *Cell {
	return &Cell{x: x, y: y}
}

func (c *Cell) isOnBoard() bool {
	return (0 <= c.x && c.x < boardWidth) && (0 <= c.y && c.y < boardHeight)
}
