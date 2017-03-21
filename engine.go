package main

const (
	levelMax = 20
	scoreMax = 999999
)

type Engine struct {
	gameover    bool
	score       int
	level       int
	initLevel   int
	deleteLines int
}

func NewEngine() *Engine {
	return &Engine{}
}

func (engine *Engine) tick() {
	board.ApplyGravity()
	view.RefreshScreen()
}

func (engine *Engine) NewGame() {
	board = NewBoard()
	engine.score = 0
	engine.level = engine.initLevel
	engine.deleteLines = 0
	engine.gameover = false
	view.RefreshScreen()
}

func (engine *Engine) DeleteCheck() {
	if !board.hasFullLine() {
		return
	}

	clock.pause()

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
	engine.levelUpdate()

	clock.start()
}

func (engine *Engine) levelUpdate() {
	if engine.level == levelMax {
		return
	}

	targetLevel := engine.deleteLines / 10
	if engine.level < targetLevel {
		engine.level = targetLevel
		clock.updateInterval()
	}
}

func (engine *Engine) AddScore(add int) {
	engine.score += add
	if engine.score > scoreMax {
		engine.score = scoreMax
	}
}

func (engine *Engine) gameOver() {
	clock.over()
	clock.lock = true

	view.ShowGameOverAnimation()

	engine.gameover = true

	ranking := NewRanking()
	ranking.insertScore(engine.score)
	ranking.save()

	clock.lock = false
}
