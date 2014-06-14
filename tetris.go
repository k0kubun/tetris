package main

import (
	"github.com/nsf/termbox-go"
)

const (
	levelMax = 20
)

var (
	board       *Board
	clock       *Clock
	currentMino *Mino
	nextMino    *Mino
	score       int
	level       int
	deleteLines int
)

func initGame() {
	board = NewBoard()
	initMino()
	score = 0
	level = 1
	deleteLines = 0
	refreshScreen()
}

func initMino() {
	currentMino, nextMino = nil, nil
	pushMino()
	pushMino()
}

func deleteCheck() {
	if !board.hasFullLine() {
		return
	}
	clock.pause()

	lines := board.fullLines()
	board.showDeleteAnimation(lines)
	for _, line := range lines {
		board.deleteLine(line)
	}
	deleteLines += len(lines)
	levelUpdate()

	clock.start()
}

func levelUpdate() {
	if level == levelMax {
		return
	}

	targetLevel := deleteLines / 10
	if level < targetLevel {
		level = targetLevel
		clock.updateInterval()
	}
}

func pushMino() {
	deleteCheck()

	currentMino = nextMino
	if currentMino != nil {
		currentMino.x, currentMino.y = defaultMinoX, defaultMinoY
		if currentMino.conflicts() {
			gameOver()
			return
		}
	}
	nextMino = NewMino()
}

func gameOver() {
	initGame()
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc)
	termbox.Flush()

	initGame()
	clock = NewClock(func() {
		currentMino.applyGravity()
		refreshScreen()
	})
	clock.start()
	waitKeyInput()
}
