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
	currentMinoXOffset, currentMinoYOffset = 5, 2
	nextMinoXOffset, nextMinoYOffset       = 18, 1
)

var (
	colorMapping = map[rune]termbox.Attribute{
		'k': termbox.ColorBlack,
		'K': termbox.ColorBlack | termbox.AttrBold,
		'r': termbox.ColorRed,
		'R': termbox.ColorRed | termbox.AttrBold,
		'g': termbox.ColorGreen,
		'G': termbox.ColorGreen | termbox.AttrBold,
		'y': termbox.ColorYellow,
		'Y': termbox.ColorYellow | termbox.AttrBold,
		'b': termbox.ColorBlue,
		'B': termbox.ColorBlue | termbox.AttrBold,
		'm': termbox.ColorMagenta,
		'M': termbox.ColorMagenta | termbox.AttrBold,
		'c': termbox.ColorCyan,
		'C': termbox.ColorCyan | termbox.AttrBold,
		'w': termbox.ColorWhite,
		'W': termbox.ColorWhite | termbox.AttrBold,
	}
)

func refreshScreen() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	drawCells(background, 0, 0)
	drawCurrentMino()
	drawNextMino()

	termbox.Flush()
}

func drawCurrentMino() {
	drawMino(currentMino, currentMinoXOffset, currentMinoYOffset)
}

func drawNextMino() {
	drawMino(nextMino, nextMinoXOffset-nextMino.x, nextMinoYOffset-nextMino.y)
}

func drawMino(mino *Mino, xOffset, yOffset int) {
	lines := strings.Split(mino.block, "\n")

	for y, line := range lines {
		for x, char := range line {
			color := colorByChar(char)

			if color != termbox.ColorDefault {
				termbox.SetCell(2*(x+mino.x+xOffset)-1, y+mino.y+yOffset, '▓', color, color^termbox.AttrBold)
				termbox.SetCell(2*(x+mino.x+xOffset), y+mino.y+yOffset, ' ', color, color^termbox.AttrBold)
			}
		}
	}
}

func drawCells(text string, left, top int) {
	lines := strings.Split(text, "\n")

	for y, line := range lines {
		for x, char := range line {
			drawCell(left+x, top+y, colorByChar(char))
		}
	}
}

func drawCell(x, y int, color termbox.Attribute) {
	termbox.SetCell(2*x-1, y, ' ', termbox.ColorDefault, color)
	termbox.SetCell(2*x, y, ' ', termbox.ColorDefault, color)
}

func colorByChar(ch rune) termbox.Attribute {
	return colorMapping[ch]
}
