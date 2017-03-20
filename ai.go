package main

import (
	"time"
)

const (
	aiQueueSize  = boardWidth + 6
	aiMoveLength = (boardWidth + 1) / 2
)

type Ai struct {
	stopped    bool
	chanStop   chan struct{}
	chanTicker *time.Ticker
	enabled    bool
	queue      []rune
	index      int
}

func NewAi() *Ai {
	ai := Ai{}
	ai.chanTicker = time.NewTicker(200 * time.Millisecond)
	ai.chanStop = make(chan struct{}, 1)
	ai.queue = make([]rune, aiQueueSize, aiQueueSize)
	for i := 0; i < aiQueueSize; i++ {
		ai.queue[i] = 's'
	}
	return &ai
}

func (ai *Ai) toggle() {
	if ai.enabled {
		ai.enabled = false
	} else {
		ai.enabled = true
		ai.addMovesToQueue(ai.getBestQueue())
	}
}

func (ai *Ai) Run() {
	logger.Info("Ai Run start")

	for {
		select {
		case <-ai.chanTicker.C:
		case <-ai.chanStop:
			return
		}
		if !ai.enabled {
			continue
		}
		if engine.gameover {
			continue
		}
		if engine.paused {
			continue
		}
		switch ai.queue[ai.index] {
		case 's':
			board.currentMino.drop()
			view.RefreshScreen()
			ai.addMovesToQueue(ai.getBestQueue())
			continue
		case 'a':
			board.currentMino.moveLeft()
		case 'd':
			board.currentMino.moveRight()
		case 'q':
			board.currentMino.rotateLeft()
		case 'e':
			board.currentMino.rotateRight()
		}
		ai.index++
		if ai.index == aiQueueSize {
			ai.index = 0
		}
		view.RefreshScreen()
	}

	logger.Info("Ai Run end")
}

func (ai *Ai) Stop() {
	if !ai.stopped {
		ai.stopped = true
		close(ai.chanStop)
	}
}

func (ai *Ai) getBestQueue() []rune {
	var rotateScore int
	bestQueue := make([]rune, 0, 0)
	bestHoles := 999
	bestLines := 0
	bestY := 0
	bestRotate := 9
	highestBlock := board.HighestBlock()

	for rotate := -2; rotate < 3; rotate++ {
		if rotate < 0 {
			rotateScore = -rotate
		} else {
			rotateScore = rotate
		}

		for move := -aiMoveLength; move < aiMoveLength; move++ {
			queue := make([]rune, 0, 4)
			mino := *board.currentMino

			if rotate < 0 {
				for i := 0; i < 0-rotate; i++ {
					mino.forceRotateLeft()
					if mino.conflicts() {
						break
					}
					queue = append(queue, 'q')
				}
			} else {
				for i := 0; i < rotate; i++ {
					mino.forceRotateRight()
					if mino.conflicts() {
						break
					}
					queue = append(queue, 'e')
				}
			}
			if mino.conflicts() {
				break
			}

			if move < 0 {
				for i := 0; i < 0-move; i++ {
					mino.x--
					if mino.conflicts() {
						mino.x++
						break
					}
					queue = append(queue, 'a')
				}
			} else {
				for i := 0; i < move; i++ {
					mino.x++
					if mino.conflicts() {
						mino.x--
						break
					}
					queue = append(queue, 'd')
				}
			}

			for !mino.conflicts() {
				mino.y++
			}
			if mino.conflicts() {
				mino.y--
			}
			if mino.conflicts() {
				continue
			}

			lines := board.FullLinesWithMino(&mino)
			numLines := len(lines)
			holes := board.HolesChangedWithMino(&mino)

			if highestBlock < boardHeight/2-2 {

				if numLines < bestLines {
					continue
				}

				if numLines > bestLines || holes < bestHoles ||
					(holes == bestHoles && bestY < mino.y) ||
					(holes == bestHoles && bestY == mino.y && rotateScore < bestRotate) {
					bestQueue = queue
					bestHoles = holes
					bestLines = numLines
					bestY = mino.y
					bestRotate = rotateScore
				}

			} else {

				if holes > bestHoles {
					continue
				}

				if holes < bestHoles || numLines > bestLines ||
					(numLines == bestLines && bestY < mino.y) ||
					(numLines == bestLines && bestY == mino.y && rotateScore < bestRotate) {
					bestQueue = queue
					bestHoles = holes
					bestLines = numLines
					bestY = mino.y
					bestRotate = rotateScore
				}

			}

		}

	}

	return bestQueue
}

func (ai *Ai) addMovesToQueue(queue []rune) {
	insertIndex := ai.index

	for _, char := range queue {
		ai.queue[insertIndex] = char
		insertIndex++
		if insertIndex == aiQueueSize {
			insertIndex = 0
		}
	}
	ai.queue[insertIndex] = 's'
}
