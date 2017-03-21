package main

import (
	"github.com/nsf/termbox-go"
	"time"
)

const (
	levelMax         = 20
	scoreMax         = 999999
	gameoverDuration = 1
)

type Engine struct {
	stopped     bool
	chanStop    chan struct{}
	keyInput    *KeyInput
	ticker      *time.Ticker
	paused      bool
	gameover    bool
	score       int
	level       int
	deleteLines int
}

func NewEngine() *Engine {
	return &Engine{
		chanStop: make(chan struct{}, 1),
		gameover: true,
	}
}

func (engine *Engine) Run() {
	logger.Info("Engine Run start")

	var event *termbox.Event

	engine.ticker = time.NewTicker(time.Second)
	engine.ticker.Stop()

	board = NewBoard()
	view.RefreshScreen()

	engine.keyInput = NewKeyInput()
	go engine.keyInput.Run()

	for {
		select {
		case <-engine.chanStop:
			return
		default:
			select {
			case event = <-engine.keyInput.chanKeyInput:
				engine.keyInput.ProcessEvent(event)
				view.RefreshScreen()
			case <-engine.ticker.C:
				engine.tick()
			case <-engine.chanStop:
				return
			}
		}
	}

	logger.Info("Engine Run end")
}

func (engine *Engine) Stop() {
	if !engine.stopped {
		engine.stopped = true
		close(engine.chanStop)
	}
	engine.ticker.Stop()
}

func (engine *Engine) setTickDuration() {
	tickDuration := time.Duration(10*(50-2*engine.level)) * time.Millisecond
	engine.ticker.Stop()
	engine.ticker = time.NewTicker(tickDuration)
}

func (engine *Engine) tick() {
	board.ApplyGravity()
	view.RefreshScreen()
}

func (engine *Engine) gameOver() {
	engine.Pause()
	engine.gameover = true
	ranking := NewRanking()
	ranking.insertScore(engine.score)
	ranking.save()
	view.RefreshScreen()
	select {
	case <-engine.keyInput.chanKeyInput:
	default:
	}
}

func (engine *Engine) Pause() {
	engine.ticker.Stop()
	engine.paused = true
}

func (engine *Engine) UnPause() {
	engine.setTickDuration()
	engine.paused = false
}

func (engine *Engine) NewGame() {
	board = NewBoard()
	engine.score = 0
	engine.level = 1
	engine.deleteLines = 0
	engine.gameover = false
	engine.UnPause()
	view.RefreshScreen()
}

func (engine *Engine) DeleteCheck() {
	if !board.HasFullLine() {
		return
	}

	engine.Pause()

	lines := board.FullLines()
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
