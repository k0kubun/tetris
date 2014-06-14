package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

const (
	boardWidth  = 10
	boardHeight = 18
	blankColor  = termbox.ColorBlack
)

var (
	board = NewBoard()
)

type Board struct {
	colors [boardWidth][boardHeight]termbox.Attribute
}

func NewBoard() *Board {
	board := &Board{}
	for i := 0; i < boardWidth; i++ {
		for j := 0; j < boardHeight; j++ {
			board.colors[i][j] = blankColor
		}
	}
	return board
}

func (b *Board) text() string {
	text := ""
	for j := 0; j < boardHeight; j++ {
		for i := 0; i < boardWidth; i++ {
			text = fmt.Sprintf("%s%c", text, charByColor(b.colors[i][j]))
		}
		text = fmt.Sprintf("%s\n", text)
	}
	return text
}

func (b *Board) setCell(cell *Cell) {
	b.colors[cell.x][cell.y] = cell.color
}

func (b *Board) setCells(cells []*Cell) {
	for _, cell := range cells {
		b.setCell(cell)
	}
}
