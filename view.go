package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"strings"
	"time"
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
	drawDropMarkerDisabled bool
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

	if engine.gameover {
		view.drawGameOver()
	} else {
		view.drawBoard()
		view.drawCurrentMino()
		view.drawDropMarker()
	}

	view.drawNextMino()

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
	yOffset++
	view.drawText(xOffset, yOffset, "u    - level up", termbox.ColorWhite, termbox.ColorBlack)
}

func (view *View) drawText(x, y int, text string, fg, bg termbox.Attribute) {
	for index, ch := range text {
		termbox.SetCell(x+index, y, rune(ch), fg, bg)
	}
}

func (view *View) drawBoard() {
	xOffset := boardXOffset + 2
	yOffset := boardYOffset

	for i := 0; i < boardWidth; i++ {
		for j := 0; j < boardHeight; j++ {
			if board.colors[i][j] != blankColor {
				termbox.SetCell(2*i+xOffset, j+yOffset, '▓', board.colors[i][j], board.colors[i][j]^termbox.AttrBold)
				termbox.SetCell(2*i+1+xOffset, j+yOffset, ' ', board.colors[i][j], board.colors[i][j]^termbox.AttrBold)
			}
		}
	}
}

func (view *View) drawDropMarker() {
	if view.drawDropMarkerDisabled {
		return
	}

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
	view.drawMino(board.nextMino, boardXOffset+boardWidth+4-board.nextMino.x, boardYOffset-board.nextMino.y)
}

func (view *View) drawCurrentMino() {
	view.drawMino(board.currentMino, boardXOffset, boardYOffset)
}

func (view *View) drawMino(mino *Mino, xOffset, yOffset int) {
	lines := strings.Split(mino.block, "\n")

	for y, line := range lines {
		for x, char := range line {
			if isOnBoard(x+mino.x, y+mino.y) {
				color := colorMapping[char]
				view.drawCell(x+mino.x+xOffset, y+mino.y+yOffset, color)
			}
		}
	}
}

func (view *View) drawCells(text string, left, top int) {
	lines := strings.Split(text, "\n")

	for y, line := range lines {
		for x, char := range line {
			view.drawCell(left+x, top+y, colorMapping[char])
		}
	}
}

func (view *View) drawCell(x, y int, color termbox.Attribute) {
	if color != termbox.ColorDefault && color != blankColor {
		if color == colorMapping['K'] {
			termbox.SetCell(2*x-1, y, '▓', color, termbox.ColorWhite)
			termbox.SetCell(2*x, y, ' ', color, termbox.ColorWhite)
		} else {
			termbox.SetCell(2*x-1, y, '▓', color, color^termbox.AttrBold)
			termbox.SetCell(2*x, y, ' ', color, color^termbox.AttrBold)
		}
	}
}

func (view *View) drawGameOver() {
	xOffset := boardXOffset + 4
	yOffset := boardYOffset + 2

	view.drawText(xOffset, yOffset, "   GAME OVER", termbox.ColorWhite, termbox.ColorBlack)
	yOffset += 2
	view.drawText(xOffset, yOffset, "sbar for new game", termbox.ColorWhite, termbox.ColorBlack)
	yOffset += 2
	xOffset += 2
	ranking := NewRanking()
	for index, line := range ranking.scores {
		view.drawText(xOffset, yOffset+index, fmt.Sprintf("%2d: %6d", index+1, line), termbox.ColorWhite, termbox.ColorBlack)
	}
}

func (view *View) ShowDeleteAnimation(lines []int) {
	view.drawDropMarkerDisabled = true

	view.RefreshScreen()

	for times := 0; times < 3; times++ {
		for _, y := range lines {
			view.colorizeLine(y, termbox.ColorCyan)
		}
		termbox.Flush()
		time.Sleep(160 * time.Millisecond)

		view.RefreshScreen()
		time.Sleep(160 * time.Millisecond)
	}

	view.drawDropMarkerDisabled = false
}

func (view *View) ShowGameOverAnimation() {
	view.drawDropMarkerDisabled = true

	view.RefreshScreen()

	for y := boardHeight - 1; y >= 0; y-- {
		view.colorizeLine(y, termbox.ColorBlack)
		termbox.Flush()
		time.Sleep(50 * time.Millisecond)
	}

	view.drawDropMarkerDisabled = false
}

func (view *View) colorizeLine(y int, color termbox.Attribute) {
	for x := 0; x < boardWidth; x++ {
		termbox.SetCell((x+boardXOffset)*2-1, y+boardYOffset, ' ', termbox.ColorDefault, color)
		termbox.SetCell((x+boardXOffset)*2, y+boardYOffset, ' ', termbox.ColorDefault, color)
	}
}
