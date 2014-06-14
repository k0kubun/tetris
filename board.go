package main

import (
	"github.com/nsf/termbox-go"
)

const (
	boardWidth  = 10
	boardHeight = 18
)

var (
	board = [boardWidth][boardHeight]Cell{}
)

type Cell struct {
	color termbox.Attribute
}
