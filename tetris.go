package main

import (
	"github.com/nsf/termbox-go"
	"gopkg.in/inconshreveable/log15.v2"
	"os"
	"path/filepath"
)

const (
	boardWidth   = 10
	boardHeight  = 20
	boardXOffset = 2
	boardYOffset = 2
	minoWidth    = 4
	minoHeight   = 4
	blankColor   = termbox.ColorBlack
)

var (
	logger log15.Logger
	view   *View
	engine *Engine
	ai     *Ai
	board  *Board
)

func main() {
	baseDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	logger = log15.New()
	if baseDir != "" {
		logger.SetHandler(log15.Must.FileHandler(baseDir+"/tetris.log", log15.LogfmtFormat()))
	}

	view = NewView()
	engine = NewEngine()
	ai = NewAi()

	go ai.Run()
	engine.Run()

	ai.Stop()
	view.Stop()

}
