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

	if event.Ch == 'p' {
		clock.pause()
	} else if event.Ch == 'z' {
		currentMino.rotateLeft()
	} else if event.Ch == 'x' || event.Key == termbox.KeyArrowUp {
		currentMino.rotateRight()
	} else if event.Key == termbox.KeySpace {
		currentMino.drop()
	} else if event.Key == termbox.KeyArrowDown {
		currentMino.moveDown()
	} else if event.Key == termbox.KeyArrowLeft {
		currentMino.moveLeft()
	} else if event.Key == termbox.KeyArrowRight {
		currentMino.moveRight()
	}

	refreshScreen()
}
