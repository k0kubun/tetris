package main

import (
	"github.com/nsf/termbox-go"
	"os"
	"strconv"
)

const (
	boardWidth  = 10
	boardHeight = 18
	minoWidth   = 4
	minoHeight  = 4
)

var (
	engine *Engine
	board  *Board
	clock  *Clock
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc)
	termbox.Flush()

	engine = NewEngine()

	clock = NewClock(func() {
		board.ApplyGravity()
		refreshScreen()
	})

	engine.initLevel = 1
	if len(os.Args) > 1 {
		num, err := strconv.Atoi(os.Args[1])
		if err != nil {
			panic(err)
		}
		if 0 < num && num < 10 {
			engine.initLevel = num
		}
	}
	initGame()
	clock.start()
	waitKeyInput()
}
