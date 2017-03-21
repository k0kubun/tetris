package main

import (
	"github.com/nsf/termbox-go"
	"os"
	"strconv"
	"time"
)

const (
	boardWidth  = 10
	boardHeight = 18
	minoWidth   = 4
	minoHeight  = 4
)

var (
	board *Board
	clock *Clock
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc)
	termbox.Flush()

	clock = NewClock(func() {
		board.ApplyGravity()
		refreshScreen()
	})

	initLevel = 1
	if len(os.Args) > 1 {
		num, err := strconv.Atoi(os.Args[1])
		if err != nil {
			panic(err)
		}
		if 0 < num && num < 10 {
			initLevel = num
		}
	}
	initGame()
	clock.start()
	waitKeyInput()
}
