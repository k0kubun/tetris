package main

import (
	"github.com/nsf/termbox-go"
	"strings"
)

const (
	background = `
    wwwwwwwwwwww wwwwww
    w..........w w....w
    w..........w w....w
    w..........w w....w
    w..........w w....w
    w..........w wwwwww
    w..........w
    w..........w
    w..........w
    w..........w
    w..........w
    w..........w
    w..........w
    w..........w
    w..........w
    w..........w
    w..........w
    w..........w
    w..........w
    wwwwwwwwwwww
	`
)

func refreshScreen() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	drawByText(background, 0, 0)

	termbox.Flush()
}

func drawByText(text string, left, top int) {
	lines := strings.Split(text, "\n")

	for y, line := range lines {
		for x, char := range line {
			switch char {
			case 'w':
				termbox.SetCell(2*left+2*x-1, top+y, ' ', termbox.ColorDefault, termbox.ColorWhite|termbox.AttrBold)
				termbox.SetCell(2*left+2*x, top+y, ' ', termbox.ColorDefault, termbox.ColorWhite|termbox.AttrBold)
			case '.':
				termbox.SetCell(2*left+2*x-1, top+y, ' ', termbox.ColorDefault, termbox.ColorBlack)
				termbox.SetCell(2*left+2*x, top+y, ' ', termbox.ColorDefault, termbox.ColorBlack)
			}
		}
	}
}
