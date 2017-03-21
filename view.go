package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"strings"
	"time"
)

const (
	nextMinoXOffset, nextMinoYOffset = 17, 2
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

type View struct {
	deleteAnimation bool
}

func NewView() *View {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	termbox.SetInputMode(termbox.InputEsc)
	termbox.Flush()
	return &View{}
}

func (view *View) Stop() {
	termbox.Close()
}

func (view *View) RefreshScreen() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	view.drawBackground()
	view.drawTexts()

	view.drawCells(board.text(), boardXOffset, boardYOffset)
	view.drawNextMino()

	if !view.deleteAnimation {
		view.drawDropMarker()
	}
	view.drawCurrentMino()
	if clock.gameover {
		for j := 0; j < boardHeight; j++ {
			view.colorizeLine(j, termbox.ColorBlack)
		}
		view.drawText(10, 4, "GAME OVER", termbox.ColorWhite, termbox.ColorBlack)
		view.drawText(7, 6, "<SPC> to continue", termbox.ColorWhite, termbox.ColorBlack)

		ranking := NewRanking()
		for idx, sc := range ranking.scores {
			view.drawText(9, 8+idx, fmt.Sprintf("%2d: %6d", idx+1, sc), termbox.ColorWhite, termbox.ColorBlack)
		}
	}

	termbox.Flush()
}

func (view *View) drawBackground() {
	colorBackgroundBorder := termbox.ColorWhite | termbox.AttrBold
	colorBackgroundFill := termbox.ColorBlack

	// playing board
	xOffset := boardXOffset
	yOffset := boardYOffset - 1
	xEnd := boardXOffset + boardWidth*2 + 4
	yEnd := boardYOffset + boardHeight + 1
	for x := xOffset; x < xEnd; x++ {
		for y := yOffset; y < yEnd; y++ {
			if x == xOffset || x == xOffset+1 || x == xEnd-1 || x == xEnd-2 ||
				y == yOffset || y == yEnd-1 {
				termbox.SetCell(x, y, ' ', termbox.ColorDefault, colorBackgroundBorder)
			} else {
				termbox.SetCell(x, y, ' ', termbox.ColorDefault, colorBackgroundFill)
			}
		}
	}

	// piece preview
	xOffset = boardXOffset + boardWidth*2 + 7
	yOffset = boardYOffset - 1
	xEnd = xOffset + minoWidth*2 + 6
	yEnd = yOffset + minoHeight + 2
	for x := xOffset; x < xEnd; x++ {
		for y := yOffset; y < yEnd; y++ {
			if x == xOffset || x == xOffset+1 || x == xEnd-1 || x == xEnd-2 ||
				y == yOffset || y == yEnd-1 {
				termbox.SetCell(x, y, ' ', termbox.ColorDefault, colorBackgroundBorder)
			} else {
				termbox.SetCell(x, y, ' ', termbox.ColorDefault, colorBackgroundFill)
			}
		}
	}

}

func (view *View) drawTexts() {
	xOffset := boardXOffset + boardWidth*2 + 7
	yOffset := boardYOffset + 6

	view.drawText(xOffset, yOffset, "SCORE:", termbox.ColorWhite, termbox.ColorBlue)
	view.drawText(xOffset+7, yOffset, fmt.Sprintf("%7d", engine.score), termbox.ColorBlack, termbox.ColorWhite)

	yOffset += 2

	view.drawText(xOffset, yOffset, "LEVEL:", termbox.ColorWhite, termbox.ColorBlue)
	view.drawText(xOffset+7, yOffset, fmt.Sprintf("%5d", engine.level), termbox.ColorBlack, termbox.ColorWhite)

	yOffset += 2

	view.drawText(xOffset, yOffset, "LINES:", termbox.ColorWhite, termbox.ColorBlue)
	view.drawText(xOffset+7, yOffset, fmt.Sprintf("%5d", engine.deleteLines), termbox.ColorBlack, termbox.ColorWhite)

	yOffset += 2

	// ascii arrow characters add extra two spaces
	view.drawText(xOffset, yOffset, "←  - left", termbox.ColorWhite, termbox.ColorBlack)
	yOffset++
	view.drawText(xOffset, yOffset, "z    - rotate left", termbox.ColorWhite, termbox.ColorBlack)
	yOffset++
	view.drawText(xOffset, yOffset, "x,↑- rotate right", termbox.ColorWhite, termbox.ColorBlack)
	yOffset++
	view.drawText(xOffset, yOffset, "→  - right", termbox.ColorWhite, termbox.ColorBlack)
	yOffset++
	view.drawText(xOffset, yOffset, "↓  - slow drop", termbox.ColorWhite, termbox.ColorBlack)
	yOffset++
	view.drawText(xOffset, yOffset, "sbar - full drop", termbox.ColorWhite, termbox.ColorBlack)
	yOffset++
	view.drawText(xOffset, yOffset, "p    - pause", termbox.ColorWhite, termbox.ColorBlack)
	yOffset++
	view.drawText(xOffset, yOffset, "q    - quit", termbox.ColorWhite, termbox.ColorBlack)
}

func (view *View) drawText(x, y int, text string, fg, bg termbox.Attribute) {
	for index, ch := range text {
		termbox.SetCell(x+index, y, rune(ch), fg, bg)
	}
}

func (view *View) drawCurrentMino() {
	view.drawMino(board.currentMino, boardXOffset, boardYOffset)
}

func (view *View) drawDropMarker() {
	marker := *board.currentMino
	for !marker.conflicts() {
		marker.y++
	}
	if marker.conflicts() {
		marker.y--
	}

	lines := strings.Split(marker.block, "\n")
	for y, line := range lines {
		for x, char := range line {
			if isOnBoard(x+marker.x, y+marker.y) && colorMapping[char] != blankColor &&
				colorMapping[char] != termbox.ColorDefault {
				view.drawCell(x+marker.x+boardXOffset, y+marker.y+boardYOffset, colorMapping['K'])
			}
		}
	}
}

func (view *View) drawNextMino() {
	view.drawMino(board.nextMino, nextMinoXOffset-board.nextMino.x, nextMinoYOffset-board.nextMino.y)
}

func (view *View) drawMino(mino *Mino, xOffset, yOffset int) {
	lines := strings.Split(mino.block, "\n")

	for y, line := range lines {
		for x, char := range line {
			if isOnBoard(x+mino.x, y+mino.y) {
				color := view.colorByChar(char)
				view.drawCell(x+mino.x+xOffset, y+mino.y+yOffset, color)
			}
		}
	}
}

func (view *View) drawCells(text string, left, top int) {
	lines := strings.Split(text, "\n")

	for y, line := range lines {
		for x, char := range line {
			view.drawCell(left+x, top+y, view.colorByChar(char))
		}
	}
}

func (view *View) drawCell(x, y int, color termbox.Attribute) {
	if color != termbox.ColorDefault && color != blankColor {
		if color == view.colorByChar('K') {
			termbox.SetCell(2*x-1, y, '▓', color, termbox.ColorWhite)
			termbox.SetCell(2*x, y, ' ', color, termbox.ColorWhite)
		} else {
			termbox.SetCell(2*x-1, y, '▓', color, color^termbox.AttrBold)
			termbox.SetCell(2*x, y, ' ', color, color^termbox.AttrBold)
		}
	}
}

func (view *View) drawBacks(text string, left, top int) {
	lines := strings.Split(text, "\n")

	for y, line := range lines {
		for x, char := range line {
			view.drawBack(left+x, top+y, view.colorByChar(char))
		}
	}
}

func (view *View) drawBack(x, y int, color termbox.Attribute) {
	termbox.SetCell(2*x-1, y, ' ', termbox.ColorDefault, color)
	termbox.SetCell(2*x, y, ' ', termbox.ColorDefault, color)
}

func (view *View) colorByChar(ch rune) termbox.Attribute {
	return colorMapping[ch]
}

func (view *View) charByColor(color termbox.Attribute) rune {
	for ch, currentColor := range colorMapping {
		if currentColor == color {
			return ch
		}
	}
	return '.'
}

func (view *View) ShowDeleteAnimation(lines []int) {
	view.deleteAnimation = true

	view.RefreshScreen()

	for times := 0; times < 3; times++ {
		for _, line := range lines {
			view.colorizeLine(line, termbox.ColorCyan)
		}
		termbox.Flush()
		time.Sleep(160 * time.Millisecond)

		view.RefreshScreen()
		time.Sleep(160 * time.Millisecond)
	}

	view.deleteAnimation = false
}

func (view *View) ShowGameOverAnimation() {
	for y := boardHeight - 1; y >= 0; y-- {
		view.colorizeLine(y, termbox.ColorBlack)
		termbox.Flush()
		time.Sleep(50 * time.Millisecond)
	}
}

func (view *View) colorizeLine(line int, color termbox.Attribute) {
	for i := 0; i < boardWidth; i++ {
		view.drawBack(i+boardXOffset, line+boardYOffset, color)
	}
}
