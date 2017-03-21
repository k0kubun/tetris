package main

import (
	"github.com/nsf/termbox-go"
	"runtime"
)

func waitKeyInput() {
	view.RefreshScreen()

	for {
		event := termbox.PollEvent()
		if event.Type == termbox.EventKey {
			if event.Ch == 'q' || event.Key == termbox.KeyCtrlC || event.Key == termbox.KeyCtrlD {
				return
			}
			ProcessEvent(&event)
		}
	}
}

func ProcessEvent(event *termbox.Event) {

	if clock.lock {
		return
	} else if engine.gameover {
		if event.Key == termbox.KeySpace {
			engine.NewGame()
		}
		return
	} else if engine.paused {
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
		}
	}

	view.RefreshScreen()
}
