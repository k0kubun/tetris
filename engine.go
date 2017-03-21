package main

import (
	"github.com/nsf/termbox-go"
	"time"
)

const (
	levelMax         = 20
	scoreMax         = 999999
	gameoverDuration = 50
)

type Engine struct {
	score       int
	level       int
	initLevel   int
	deleteLines int
}

func NewEngine() *Engine {
	return &Engine{}
}

func initGame() {
	board = NewBoard()
	engine.score = 0
	engine.level = engine.initLevel
	engine.deleteLines = 0
	refreshScreen()
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
	engine.deleteLines += len(lines)
	switch len(lines) {
	case 1:
		addScore(40 * (engine.level + 1))
	case 2:
		addScore(100 * (engine.level + 1))
	case 3:
		addScore(300 * (engine.level + 1))
	case 4:
		addScore(1200 * (engine.level + 1))
	}
	levelUpdate()

	clock.start()
}

func levelUpdate() {
	if engine.level == levelMax {
		return
	}

	targetLevel := engine.deleteLines / 10
	if engine.level < targetLevel {
		engine.level = targetLevel
		clock.updateInterval()
	}
}

func addScore(add int) {
	engine.score += add
	if engine.score > scoreMax {
		engine.score = scoreMax
	}
}

func gameOver() {
	clock.over()

	clock.lock = true
	for j := 0; j < boardHeight; j++ {
		rewriteScreen(func() {
			for y := boardHeight - 1; y > boardHeight-1-j; y -= 1 {
				board.colorizeLine(y, termbox.ColorBlack)
			}
		})
		timer := time.NewTimer(gameoverDuration * time.Millisecond)
		<-timer.C
	}
	clock.lock = false
}
