package main

import (
	"github.com/nsf/termbox-go"
)

type Cell struct {
	x, y  int
	color termbox.Attribute
}

func NewCell(x, y int, ch rune) *Cell {
	return &Cell{x: x, y: y, color: colorMapping[ch]}
}

func (c *Cell) isOnBoard() bool {
	return (0 <= c.x && c.x < boardWidth) && (0 <= c.y && c.y < boardHeight)
}
