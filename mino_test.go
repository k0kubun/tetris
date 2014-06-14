package main

import (
	"testing"
)

func TestNewMino(t *testing.T) {
	mino := NewMino()
	if len(mino.block) != 19 {
		t.Errorf("mino unexpected length: %d", len(mino.block))
	}
	for i := 0; i < 19; i++ {
		if i%5 == 4 && mino.block[i] != '\n' {
			t.Errorf("unexpected mino: %s", mino.block)
		} else if i%5 != 4 && mino.block[i] == '\n' {
			t.Errorf("unexpected mino: %s", mino.block)
		}
	}
}

func TestMinoCell(t *testing.T) {
	mino := NewMino()
	mino.block = "1234\n5678\n9abc\ndefg"
	mino.assertCellResult(t, 0, 0, '1')
	mino.assertCellResult(t, 1, 0, '2')
	mino.assertCellResult(t, 2, 0, '3')
	mino.assertCellResult(t, 3, 0, '4')
	mino.assertCellResult(t, 0, 1, '5')
	mino.assertCellResult(t, 1, 1, '6')
	mino.assertCellResult(t, 2, 1, '7')
	mino.assertCellResult(t, 3, 1, '8')
	mino.assertCellResult(t, 0, 2, '9')
	mino.assertCellResult(t, 1, 2, 'a')
	mino.assertCellResult(t, 2, 2, 'b')
	mino.assertCellResult(t, 3, 2, 'c')
	mino.assertCellResult(t, 0, 3, 'd')
	mino.assertCellResult(t, 1, 3, 'e')
	mino.assertCellResult(t, 2, 3, 'f')
	mino.assertCellResult(t, 3, 3, 'g')
}

func TestMinoSetCell(t *testing.T) {
	mino := NewMino()
	mino.block = "1234\n5678\n9abc\ndefg"
	mino.assertSetCellResult(t, 0, 0, 'A')
	mino.assertSetCellResult(t, 1, 0, 'B')
	mino.assertSetCellResult(t, 2, 0, 'C')
	mino.assertSetCellResult(t, 3, 0, 'D')
	mino.assertSetCellResult(t, 0, 1, 'E')
	mino.assertSetCellResult(t, 1, 1, 'F')
	mino.assertSetCellResult(t, 2, 1, 'G')
	mino.assertSetCellResult(t, 3, 1, 'H')
	mino.assertSetCellResult(t, 0, 2, 'I')
	mino.assertSetCellResult(t, 1, 2, 'J')
	mino.assertSetCellResult(t, 2, 2, 'K')
	mino.assertSetCellResult(t, 3, 2, 'L')
	mino.assertSetCellResult(t, 0, 3, 'M')
	mino.assertSetCellResult(t, 1, 3, 'N')
	mino.assertSetCellResult(t, 2, 3, 'O')
	mino.assertSetCellResult(t, 3, 3, 'P')
}

func TestMinoRotateRight(t *testing.T) {
	mino := NewMino()
	mino.block = "1234\n5678\n9abc\ndefg"
	mino.rotateRight()
	if mino.block != "d951\nea62\nfb73\ngc84" {
		t.Errorf("unexpected mino:\n%s", mino.block)
	}
}

func TestMinoRotateLeft(t *testing.T) {
	mino := NewMino()
	mino.block = "1234\n5678\n9abc\ndefg"
	mino.rotateLeft()
	if mino.block != "48cg\n37bf\n26ae\n159d" {
		t.Errorf("unexpected mino:\n%s", mino.block)
	}
}

func (m *Mino) assertCellResult(t *testing.T, x, y int, ch rune) {
	if m.cell(x, y) != ch {
		t.Errorf("expected: %c, but got: %c", ch, m.cell(x, y))
	}
}

func (m *Mino) assertSetCellResult(t *testing.T, x, y int, ch rune) {
	m.setCell(x, y, ch)
	if m.cell(x, y) != ch {
		t.Errorf("expected: %c, but got: %c", ch, m.cell(x, y))
	}
}
