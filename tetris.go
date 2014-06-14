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

func drawByText(text string) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	lines := strings.Split(text, "\n")

	for y, line := range lines {
		for x, char := range line {
			switch char {
			case 'w':
				termbox.SetCell(2*x-1, y, ' ', termbox.ColorDefault, termbox.ColorWhite|termbox.AttrBold)
				termbox.SetCell(2*x, y, ' ', termbox.ColorDefault, termbox.ColorWhite|termbox.AttrBold)
			case '.':
				termbox.SetCell(2*x-1, y, ' ', termbox.ColorDefault, termbox.ColorBlack)
				termbox.SetCell(2*x, y, ' ', termbox.ColorDefault, termbox.ColorBlack)
			}
		}
	}
	termbox.Flush()
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc)
	termbox.Flush()

	drawByText(background)
	waitKeyInput()
}
