package main

import (
	"math/rand"
	"strings"
	"time"
)

const (
	defaultMinoX, defaultMinoY = 3, -1
	minoWidth, minoHeight      = 4, 4
)

var (
	blocks = []string{
		`
			....
			.GG.
			GG..
			....
		`, `
			....
			.RR.
			..RR
			....
		`, `
			....
			.YY.
			.YY.
			....
		`, `
			....
			....
			CCCC
			....
		`, `
			....
			.M..
			MMM.
			....
		`, `
			....
			.b..
			.bbb
			....
		`, `
			....
			..m.
			mmm.
			....
		`,
	}
)

type Mino struct {
	block string
	x     int
	y     int
}

func NewMino() *Mino {
	rand.Seed(time.Now().UnixNano())
	block := blocks[rand.Intn(len(blocks))]
	block = strings.Replace(block, "\t", "", -1)
	block = strings.Replace(block, " ", "", -1)
	block = strings.Trim(block, "\n")
	return &Mino{block: block}
}

func (m *Mino) cell(x, y int) rune {
	return rune(m.block[x+(minoWidth+1)*y])
}

func (m *Mino) setCell(x, y int, cell rune) {
	buf := []rune(m.block)
	buf[x+(minoWidth+1)*y] = cell
	m.block = string(buf)
}

func (m *Mino) applyGravity() {
	m.moveDown()
}

func (m *Mino) drop() {
	addScore(m.putBottom())
	board.setCells(m.cells())
	pushMino()
}

func (m *Mino) putBottom() int {
	distance := -1
	dstMino := *m
	for !dstMino.conflicts() {
		*m = dstMino
		dstMino.forceMoveDown()
		distance++
	}
	if distance < 0 {
		distance = 0
	}
	return distance
}

func (m *Mino) moveDown() {
	dstMino := *m
	dstMino.forceMoveDown()

	if dstMino.conflicts() {
		board.setCells(m.cells())
		pushMino()
	} else {
		m.forceMoveDown()
	}
}

func (m *Mino) forceMoveDown() {
	m.y++
}

func (m *Mino) moveLeft() {
	dstMino := *m
	dstMino.x--
	if !dstMino.conflicts() {
		m.x--
	}
}

func (m *Mino) moveRight() {
	dstMino := *m
	dstMino.x++
	if !dstMino.conflicts() {
		m.x++
	}
}

func (m *Mino) rotateRight() {
	dstMino := *m
	dstMino.forceRotateRight()
	if !dstMino.conflicts() {
		m.forceRotateRight()
	}
}

func (m *Mino) forceRotateRight() {
	oldMino := *m
	for j := 0; j < minoHeight; j++ {
		for i := 0; i < minoWidth; i++ {
			m.setCell(minoHeight-j-1, i, oldMino.cell(i, j))
		}
	}
}

func (m *Mino) rotateLeft() {
	dstMino := *m
	dstMino.forceRotateLeft()
	if !dstMino.conflicts() {
		m.forceRotateLeft()
	}
}

func (m *Mino) forceRotateLeft() {
	oldMino := *m
	for j := 0; j < minoHeight; j++ {
		for i := 0; i < minoWidth; i++ {
			m.setCell(j, minoWidth-i-1, oldMino.cell(i, j))
		}
	}
}

func (m *Mino) conflicts() bool {
	for _, cell := range m.cells() {
		if cell.conflicts() {
			return true
		}
	}
	return false
}

func (m *Mino) cells() []*Cell {
	cells := []*Cell{}
	for i := 0; i < minoWidth; i++ {
		for j := 0; j < minoHeight; j++ {
			if m.cell(i, j) != '.' {
				cells = append(cells, NewCell(m.x+i, m.y+j, m.cell(i, j)))
			}
		}
	}
	return cells
}
