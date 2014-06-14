package main

import (
	"github.com/nsf/termbox-go"
)

func waitKeyInput() {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Ch == 'q' || ev.Key == termbox.KeyCtrlC || ev.Key == termbox.KeyCtrlD {
				// quit application
				return
			} else if ev.Ch == 'p' {
				if clock.paused {
					// unpause
					clock.start()
				} else {
					// pause
					clock.pause()
				}
			} else if ev.Ch == 'z' {
				// Left Turn
			} else if ev.Ch == 'x' || ev.Key == termbox.KeyArrowUp {
				// Right turn
			} else if ev.Key == termbox.KeySpace {
				// Drop
			} else if ev.Key == termbox.KeyArrowDown {
				// Fast down
			} else if ev.Key == termbox.KeyArrowLeft {
				// Left move
			} else if ev.Key == termbox.KeyArrowRight {
				// Right move
			}
		}
	}
}
