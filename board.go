package main

import (
	"github.com/nsf/termbox-go"
)

const (
	boardWidth  = 10
	boardHeight = 18
	blankColor  = 'k'
)

var (
	board = NewBoard()
)

type Board struct {
	cells [boardWidth][boardHeight]termbox.Attribute
}

func NewBoard() *Board {
	board := &Board{}
	for i := 0; i < boardWidth; i++ {
		for j := 0; j < boardHeight; j++ {
			board.cells[i][j] = blankColor
		}
	}
	return board
}
