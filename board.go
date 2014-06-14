package main

import (
	"github.com/nsf/termbox-go"
)

const (
	BOARD_WIDTH  = 10
	BOARD_HEIGHT = 18
)

var (
	board = [BOARD_HEIGHT][BOARD_WIDTH]Cell{}
)

type Cell struct {
	color termbox.Attribute
}
