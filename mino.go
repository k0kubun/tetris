package main

import (
	"math/rand"
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
	blockType      int
	rightTurnCount int
	x              int
	y              int
}

func NewMino() *Mino {
	rand.Seed(time.Now().UnixNano())
	return &Mino{blockType: rand.Intn(len(blocks))}
}

func initMino() {
	pushMino()
	pushMino()
}

func pushMino() {
	currentMino = nextMino
	if currentMino != nil {
		currentMino.x = defaultMinoX
		currentMino.y = defaultMinoY
	}
	nextMino = NewMino()
}

func (m *Mino) block() string {
	return blocks[m.blockType]
}
