package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"time"
)

const (
	defaultMinoX      = 3
	defaultMinoY      = -1
	blankColor        = termbox.ColorBlack
	animationDuration = 160
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
	deleteCheck()

	b.currentMino = b.nextMino
	if b.currentMino != nil {
		b.currentMino.x = defaultMinoX
		b.currentMino.y = defaultMinoY
		if b.currentMino.conflicts() {
			ranking := NewRanking()
			ranking.insertScore(engine.score)
			ranking.save()
			gameOver()
			return
		}
	}
	b.nextMino = NewMino()
}

func (b *Board) ApplyGravity() {
	board.currentMino.moveDown()
}

func isOnBoard(x, y int) bool {
	return (0 <= x && x < boardWidth) && (0 <= y && y < boardHeight)
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

func (b *Board) showDeleteAnimation(lines []int) {
	for times := 0; times < 3; times++ {
		rewriteScreen(func() {
			for _, line := range lines {
				b.colorizeLine(line, termbox.ColorCyan)
			}
		})
		timer := time.NewTimer(animationDuration * time.Millisecond)
		<-timer.C

		rewriteScreen(func() {})
		timer = time.NewTimer(animationDuration * time.Millisecond)
		<-timer.C
	}
}

func (b *Board) colorizeLine(line int, color termbox.Attribute) {
	for i := 0; i < boardWidth; i++ {
		drawBack(i+boardXOffset, line+boardYOffset, color)
	}
}
