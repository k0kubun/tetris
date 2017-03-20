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

func (c *Cell) conflicts() bool {
	return c.isOnWall() || c.isOverlapped()
}

func (c *Cell) isOverlapped() bool {
	if !isOnBoard(c.x, c.y) {
		return false
	}
	return board.colors[c.x][c.y] != blankColor
}

func (c *Cell) isOnWall() bool {
	return c.x < 0 || boardWidth <= c.x || boardHeight <= c.y
}

func (c *Cell) isInLocation(x int, y int) bool {
	return c.x == x && c.y == y
}
