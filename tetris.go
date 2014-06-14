package main

import (
	"github.com/nsf/termbox-go"
)

const (
	levelMax = 20
	scoreMax = 999999
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
	switch len(lines) {
	case 1:
		addScore(40 * (level + 1))
	case 2:
		addScore(100 * (level + 1))
	case 3:
		addScore(300 * (level + 1))
	case 4:
		addScore(1200 * (level + 1))
	}
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

func addScore(add int) {
	score += add
	if score > scoreMax {
		score = scoreMax
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
