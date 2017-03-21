package main

const (
	levelMax         = 20
	scoreMax         = 999999
	gameoverDuration = 50
)

var (
	score       int
	level       int
	initLevel   int
	deleteLines int
)

func initGame() {
	board = NewBoard()
	score = 0
	level = initLevel
	deleteLines = 0
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
	deleteLines += len(lines)
	switch len(lines) {
	case 1:
		addScore(40 * (level + 1))
	case 2:
		addScore(100 * (level + 1))
	case 3:
		addScore(300 * (level + 1))
	case 4:
		addScore(1200 * (level + 1))
	}
	levelUpdate()

	clock.start()
}

func levelUpdate() {
	if level == levelMax {
		return
	}

	targetLevel := deleteLines / 10
	if level < targetLevel {
		level = targetLevel
		clock.updateInterval()
	}
}

func addScore(add int) {
	score += add
	if score > scoreMax {
		score = scoreMax
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
