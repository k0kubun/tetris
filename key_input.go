package main

import (
	"github.com/nsf/termbox-go"
)

func waitKeyInput() {
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
	} else if clock.gameover {
		if event.Key == termbox.KeySpace {
			initGame()
			clock.start()
		}
		return
	} else if clock.paused {
		if event.Ch == 'p' {
			clock.start()
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
		}
	} else {
		switch event.Ch {
		case 'p':
			clock.pause()
		case 'z':
			board.currentMino.rotateLeft()
		case 'x':
			board.currentMino.rotateRight()
		}
	}

	refreshScreen()
}
