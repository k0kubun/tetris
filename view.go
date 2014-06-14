package main

import (
	"github.com/nsf/termbox-go"
	"strings"
)

const (
	background = `
    WWWWWWWWWWWW WWWWWW
    WkkkkkkkkkkW WkkkkW
    WkkkkkkkkkkW WkkkkW
    WkkkkkkkkkkW WkkkkW
    WkkkkkkkkkkW WkkkkW
    WkkkkkkkkkkW WWWWWW
    WkkkkkkkkkkW
    WkkkkkkkkkkW
    WkkkkkkkkkkW
    WkkkkkkkkkkW
    WkkkkkkkkkkW
    WkkkkkkkkkkW
    WkkkkkkkkkkW
    WkkkkkkkkkkW
    WkkkkkkkkkkW
    WkkkkkkkkkkW
    WkkkkkkkkkkW
    WkkkkkkkkkkW
    WkkkkkkkkkkW
    WWWWWWWWWWWW
	`
)

func refreshScreen() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	drawBlocks(background, 0, 0)

	termbox.Flush()
}

func drawBlocks(text string, left, top int) {
	lines := strings.Split(text, "\n")

	for y, line := range lines {
		for x, char := range line {
			termbox.SetCell(2*left+2*x-1, top+y, ' ', termbox.ColorDefault, colorByChar(char))
			termbox.SetCell(2*left+2*x, top+y, ' ', termbox.ColorDefault, colorByChar(char))
		}
	}
}

func colorByChar(ch rune) termbox.Attribute {
	switch ch {
	case 'W':
		return termbox.ColorWhite | termbox.AttrBold
	case 'k':
		return termbox.ColorBlack
	default:
		return termbox.ColorDefault
	}
}
