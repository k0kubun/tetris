package main

import (
	"github.com/nsf/termbox-go"
)

var clock *Clock

func initMino() {
	pushMino()
	pushMino()
}

func pushMino() {
	currentMino = nextMino
	if currentMino != nil {
		currentMino.x = defaultMinoX
		currentMino.y = defaultMinoY
	}
	nextMino = NewMino()
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
