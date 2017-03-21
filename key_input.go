package main

import (
	"github.com/nsf/termbox-go"
	"runtime"
)

type KeyInput struct {
	stopped      bool
	chanStop     chan struct{}
	chanKeyInput chan *termbox.Event
}

func NewKeyInput() *KeyInput {
	return &KeyInput{
		chanStop:     make(chan struct{}, 1),
		chanKeyInput: make(chan *termbox.Event, 1),
	}
}

func (keyInput *KeyInput) Run() {
	logger.Info("KeyInput Run start")

	for {
		select {
		case <-keyInput.chanStop:
			return
		default:
		}
		event := termbox.PollEvent()
		if event.Type == termbox.EventKey {
			select {
			case <-keyInput.chanStop:
				return
			default:
				select {
				case keyInput.chanKeyInput <- &event:
				case <-keyInput.chanStop:
					return
				}
			}
		}
	}

	logger.Info("KeyInput Run end")
}

func (keyInput *KeyInput) ProcessEvent(event *termbox.Event) {
	if event.Ch == 'q' || event.Key == termbox.KeyCtrlC {
		if !keyInput.stopped {
			keyInput.stopped = true
			close(keyInput.chanStop)
		}
		engine.Stop()
		return
	}

	if engine.gameover {
		if event.Key == termbox.KeySpace {
			engine.NewGame()
		}
		return
	}
	if engine.paused {
		if event.Ch == 'p' {
			engine.UnPause()
		}
		return
	}

	if event.Ch == 0 {
		switch event.Key {
		case termbox.KeySpace:
			board.currentMino.drop()
		case termbox.KeyArrowUp:
			board.currentMino.rotateRight()
		case termbox.KeyArrowDown:
			board.currentMino.moveDown()
		case termbox.KeyArrowLeft:
			board.currentMino.moveLeft()
		case termbox.KeyArrowRight:
			board.currentMino.moveRight()
		case termbox.KeyCtrlBackslash:
			// ctrl \ to log stack trace
			buffer := make([]byte, 1<<16)
			length := runtime.Stack(buffer, true)
			logger.Debug("Stack trace", "buffer", string(buffer[:length]))
		}
	} else {
		switch event.Ch {
		case 'p':
			engine.Pause()
		case 'z':
			board.currentMino.rotateLeft()
		case 'x':
			board.currentMino.rotateRight()
		case 'u':
			engine.LevelUp()
		}
	}

}
