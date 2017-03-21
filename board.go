package main

import (
	"github.com/nsf/termbox-go"
)

const (
	defaultMinoX = 3
	defaultMinoY = -1
)

type Board struct {
	colors      [boardWidth][boardHeight]termbox.Attribute
	currentMino *Mino
	nextMino    *Mino
}

func NewBoard() *Board {
	board := &Board{}
	for i := 0; i < boardWidth; i++ {
		for j := 0; j < boardHeight; j++ {
			board.colors[i][j] = blankColor
		}
	}
	board.nextMino = NewMino()
	board.currentMino = NewMino()
	board.currentMino.x = defaultMinoX
	board.currentMino.y = defaultMinoY
	return board
}

func (b *Board) deleteLine(y int) {
	for i := 0; i < boardWidth; i++ {
		b.colors[i][y] = blankColor
	}
	for j := y; j > 0; j-- {
		for i := 0; i < boardWidth; i++ {
			b.colors[i][j] = b.colors[i][j-1]
		}
	}
	for i := 0; i < boardWidth; i++ {
		b.colors[i][0] = blankColor
	}
}

func (b *Board) hasFullLine() bool {
	for j := 0; j < boardHeight; j++ {
		if b.isFullLine(j) {
			return true
		}
	}
	return false
}

func (b *Board) fullLines() []int {
	fullLines := []int{}
	for j := 0; j < boardHeight; j++ {
		if b.isFullLine(j) {
			fullLines = append(fullLines, j)
		}
	}
	return fullLines
}

func (b *Board) isFullLine(y int) bool {
	for i := 0; i < boardWidth; i++ {
		if b.colors[i][y] == blankColor {
			return false
		}
	}
	return true
}

func (b *Board) setCells(cells []*Cell) {
	for _, cell := range cells {
		b.setCell(cell)
	}
}

func (b *Board) setCell(cell *Cell) {
	b.colors[cell.x][cell.y] = cell.color
}

func (b *Board) addMino() {
	engine.DeleteCheck()

	b.nextMino.x = defaultMinoX
	b.nextMino.y = defaultMinoY
	if b.nextMino.conflicts() {
		b.nextMino.x = 0
		b.nextMino.y = 0
		engine.gameOver()
		return
	}
	b.currentMino = b.nextMino
	b.nextMino = NewMino()
}

func (b *Board) ApplyGravity() {
	board.currentMino.moveDown()
}

func isOnBoard(x, y int) bool {
	return (0 <= x && x < boardWidth) && (0 <= y && y < boardHeight)
}
