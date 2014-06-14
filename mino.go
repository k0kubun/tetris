package main

import (
	"math/rand"
	"strings"
	"time"
)

const (
	defaultMinoX, defaultMinoY = 3, -1
)

var (
	blocks = []string{
		`
		  ....
			.gg.
			gg..
			....
		`, `
		  ....
			.rr.
			..rr
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

var (
	currentMino *Mino
	nextMino    *Mino
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
	return rune(m.block[x+5*y])
}

func (m *Mino) setCell(x, y int, cell rune) {
	buf := []rune(m.block)
	buf[x+5*y] = cell
	m.block = string(buf)
}

func (m *Mino) moveDown() {
	m.y++
}

func (m *Mino) moveLeft() {
	m.x--
}

func (m *Mino) moveRight() {
	m.x++
}

func (m *Mino) applyGravity() {
	m.moveDown()
}

func (m *Mino) rotateRight() {
	oldMino := *m

	for j := 0; j < 4; j++ {
		for i := 0; i < 4; i++ {
			m.setCell(3-j, i, oldMino.cell(i, j))
		}
	}
}

func (m *Mino) rotateLeft() {
	oldMino := *m

	for j := 0; j < 4; j++ {
		for i := 0; i < 4; i++ {
			m.setCell(j, 3-i, oldMino.cell(i, j))
		}
	}
}
