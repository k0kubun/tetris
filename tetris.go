package main

import (
	"github.com/nsf/termbox-go"
	"gopkg.in/inconshreveable/log15.v2"
	"os"
	"path/filepath"
	"strconv"
)

const (
	boardWidth  = 10
	boardHeight = 18
	minoWidth   = 4
	minoHeight  = 4
)

var (
	logger log15.Logger
	engine *Engine
	board  *Board
	clock  *Clock
)

func main() {
	baseDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	logger = log15.New()
	if baseDir != "" {
		logger.SetHandler(log15.Must.FileHandler(baseDir+"/tetris.log", log15.LogfmtFormat()))
	}

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc)
	termbox.Flush()

	engine = NewEngine()
	clock = NewClock()

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
