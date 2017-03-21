package main

import (
	"time"
)

const (
	levelMax = 20
	scoreMax = 999999
)

type Engine struct {
	ticker      *time.Ticker
	paused      bool
	gameover    bool
	score       int
	level       int
	initLevel   int
	deleteLines int
}

func NewEngine() *Engine {
	engine := &Engine{
		gameover: true,
	}
	engine.ticker = time.NewTicker(time.Hour)
	engine.ticker.Stop()
	board = NewBoard()
	return engine
}

func (engine *Engine) Pause() {
	engine.ticker.Stop()
	engine.paused = true
}

func (engine *Engine) UnPause() {
	engine.setTickDuration()
	engine.paused = false
	clock.start()
}

func (engine *Engine) setTickDuration() {
	engine.ticker.Stop()
	engine.ticker = time.NewTicker(time.Duration(10*(50-2*engine.level)) * time.Millisecond)
}

func (engine *Engine) NewGame() {
	board = NewBoard()
	engine.score = 0
	engine.level = engine.initLevel
	engine.deleteLines = 0
	engine.gameover = false
	view.RefreshScreen()
	engine.UnPause()
}

func (engine *Engine) tick() {
	board.ApplyGravity()
	view.RefreshScreen()
}

func (engine *Engine) DeleteCheck() {
	if !board.hasFullLine() {
		return
	}

	engine.Pause()

	lines := board.fullLines()
	view.ShowDeleteAnimation(lines)
	for _, line := range lines {
		board.deleteLine(line)
	}
	engine.deleteLines += len(lines)

	switch len(lines) {
	case 1:
		engine.AddScore(40 * (engine.level + 1))
	case 2:
		engine.AddScore(100 * (engine.level + 1))
	case 3:
		engine.AddScore(300 * (engine.level + 1))
	case 4:
		engine.AddScore(1200 * (engine.level + 1))
	}

	if engine.level < engine.deleteLines/10 {
		engine.LevelUp()
	}

	view.RefreshScreen()

	engine.UnPause()
}

func (engine *Engine) AddScore(add int) {
	engine.score += add
	if engine.score > scoreMax {
		engine.score = scoreMax
	}
}

func (engine *Engine) LevelUp() {
	if engine.level >= levelMax {
		return
	}
	engine.level++
	engine.setTickDuration()
}

func (engine *Engine) GameOver() {
	engine.Pause()
	clock.lock = true

	view.ShowGameOverAnimation()

	engine.gameover = true

	ranking := NewRanking()
	ranking.insertScore(engine.score)
	ranking.save()

	clock.lock = false
}
