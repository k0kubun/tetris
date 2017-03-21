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
	board.currentMino = NewMino()
	board.nextMino = NewMino()
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

func (b *Board) HasFullLine() bool {
	for j := 0; j < boardHeight; j++ {
		if b.isFullLine(j) {
			return true
		}
	}
	return false
}

func (b *Board) FullLines() []int {
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

func (b *Board) FullLinesWithMino(mino *Mino) []int {
	fullLines := []int{}
	for j := 0; j < boardHeight; j++ {
		if b.isFullLineWithMino(j, mino) {
			fullLines = append(fullLines, j)
		}
	}
	return fullLines
}

func (b *Board) isFullLineWithMino(y int, mino *Mino) bool {
	for i := 0; i < boardWidth; i++ {
		if b.colors[i][y] == blankColor {
			if !mino.isInLocation(i, y) {
				return false
			}
		}
	}
	return true
}

func (b *Board) HolesChangedWithMino(mino *Mino) int {
	currentHoles := 0
	newHoles := 0
	startY := boardHeight

	for i := 0; i < boardWidth; i++ {
		startY = boardHeight
		for j := 0; j < boardHeight; j++ {
			if b.colors[i][j] != blankColor {
				startY = j + 1
				break
			}
		}
		for j := startY; j < boardHeight; j++ {
			if b.colors[i][j] == blankColor {
				currentHoles++
			}
		}
	}

	for i := 0; i < boardWidth; i++ {
		startY = boardHeight
		for j := 0; j < boardHeight; j++ {
			if b.colors[i][j] != blankColor || mino.isInLocation(i, j) {
				startY = j + 1
				break
			}
		}
		for j := startY; j < boardHeight; j++ {
			if b.colors[i][j] == blankColor && !mino.isInLocation(i, j) {
				newHoles++
			}
		}
	}

	return newHoles - currentHoles
}

func (b *Board) HighestBlock() int {
	for j := 0; j < boardHeight; j++ {
		for i := 0; i < boardWidth; i++ {
			if b.colors[i][j] != blankColor {
				return j
			}
		}
	}
	return boardHeight
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

	b.currentMino = b.nextMino
	b.currentMino.x = defaultMinoX
	b.currentMino.y = defaultMinoY
	if b.currentMino.conflicts() {
		engine.gameOver()
		return
	}
	b.nextMino = NewMino()
}

func (b *Board) ApplyGravity() {
	board.currentMino.moveDown()
}

func isOnBoard(x, y int) bool {
	return (0 <= x && x < boardWidth) && (0 <= y && y < boardHeight)
}
