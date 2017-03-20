package main

import (
	"math/rand"
	"strings"
	"time"
)

var (
	blocks = []string{
		`....
.GG.
GG..
....`,
		`....
.RR.
..RR
....`,
		`....
.YY.
.YY.
....`,
		`....
....
CCCC
....`,
		`....
.M..
MMM.
....`,
		`....
.b..
.bbb
....`,
		`....
..m.
mmm.
....`,
	}
	bag      []int
	bagIndex int
)

func init() {
	rand.Seed(time.Now().UnixNano())
	bag = rand.Perm(7)
}

type Mino struct {
	block string
	x     int
	y     int
}

func NewMino() *Mino {
	block := blocks[bag[bagIndex]]
	block = strings.Trim(block, "\n")
	bagIndex++
	if bagIndex > 6 {
		bagIndex = 0
		bag = rand.Perm(7)
	}
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

func (m *Mino) drop() {
	distance := 0
	dstMino := *m
	dstMino.y++
	for !dstMino.conflicts() {
		distance++
		dstMino.y++
	}

	if distance > 0 {
		dstMino := *m
		dstMino.y += distance
		*m = dstMino
	}

	if !m.isOnBoard() {
		engine.gameOver()
		return
	}

	engine.AddScore(distance)
	board.setCells(m.cells())
	board.addMino()
}

func (m *Mino) moveDown() {
	dstMino := *m
	dstMino.y++
	if dstMino.conflicts() {
		if !m.isOnBoard() {
			engine.gameOver()
			return
		}
		board.setCells(m.cells())
		board.addMino()
		return
	}

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
		return
	}
	dstMino.x--
	if !dstMino.conflicts() {
		m.x--
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
		return
	}
	dstMino.x++
	if !dstMino.conflicts() {
		m.x++
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

func (m *Mino) isOnBoard() bool {
	for _, cell := range m.cells() {
		if !isOnBoard(cell.x, cell.y) {
			return false
		}
	}
	return true
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

func (m *Mino) isInLocation(x int, y int) bool {
	for _, cell := range m.cells() {
		if cell.isInLocation(x, y) {
			return true
		}
	}
	return false
}
