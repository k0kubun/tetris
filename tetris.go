package main

import (
	"github.com/nsf/termbox-go"
)

var (
	clock       *Clock
	currentMino *Mino
	nextMino    *Mino
)

func initMino() {
	currentMino, nextMino = nil, nil
	pushMino()
	pushMino()
}

func pushMino() {
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
	board = NewBoard()
	initMino()
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc)
	termbox.Flush()

	initMino()
	refreshScreen()
	clock = NewClock(func() {
		currentMino.applyGravity()
		refreshScreen()
	})
	clock.start()
	waitKeyInput()
}
