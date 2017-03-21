package main

import (
	"github.com/nsf/termbox-go"
	"gopkg.in/inconshreveable/log15.v2"
	"os"
	"path/filepath"
)

const (
	boardWidth   = 10
	boardHeight  = 18
	boardXOffset = 3
	boardYOffset = 2
	minoWidth    = 4
	minoHeight   = 4
	blankColor   = termbox.ColorBlack
)

var (
	logger log15.Logger
	view   *View
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

	view = NewView()
	engine = NewEngine()

	clock = NewClock()
	waitKeyInput()

	view.Stop()

}
