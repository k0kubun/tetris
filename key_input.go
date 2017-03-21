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
			currentMino.drop()
		case termbox.KeyArrowUp:
			currentMino.rotateRight()
		case termbox.KeyArrowDown:
			currentMino.moveDown()
		case termbox.KeyArrowLeft:
			currentMino.moveLeft()
		case termbox.KeyArrowRight:
			currentMino.moveRight()
		}
	} else {
		switch event.Ch {
		case 'p':
			clock.pause()
		case 'z':
			currentMino.rotateLeft()
		case 'x':
			currentMino.rotateRight()
		}
	}

	refreshScreen()
}
