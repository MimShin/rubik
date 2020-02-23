package rubik

import (
	"fmt"
	"math/rand"
)

const RevMask = 0b1_000_00
const MovMask = 0b0_111_00
const RowMask = 0b0_000_11
const NoMove = 0b1_111_11
const RevMovMask = RevMask | MovMask

var Moves = [...]int{
	0b0_001_00, 0b0_001_01, 0b0_001_10, 
	0b1_001_00, 0b1_001_01, 0b1_001_10, 

	0b0_010_00, 0b0_010_01, 0b0_010_10, 
	0b1_010_00, 0b1_010_01, 0b1_010_10, 

	0b0_100_00, 0b0_100_01, 0b0_100_10, 
	0b1_100_00, 0b1_100_01, 0b1_100_10, 
}

func (cube *Cube) RowTurnCW(r int) string {
	notation := fmt.Sprintf("%d%s  ", r+1, left)
	for i := 0; i < size; i++ {
		cube.faces[2][r][i], cube.faces[4][r][i], cube.faces[3][r][i], cube.faces[5][r][i] =
			cube.faces[4][r][i], cube.faces[3][r][i], cube.faces[5][r][i], cube.faces[2][r][i]
	}

	if r == 0 {
		cube.faces[0].turnCW()
	}

	if r == size-1 {
		cube.faces[1].turnCCW()
	}

	return notation
}

func (cube *Cube) RowTurnCCW(r int) string {
	notation := fmt.Sprintf("%d%s  ", r+1, right)
	for i := 0; i < size; i++ {
		cube.faces[3][r][i], cube.faces[4][r][i], cube.faces[2][r][i], cube.faces[5][r][i] =
			cube.faces[4][r][i], cube.faces[2][r][i], cube.faces[5][r][i], cube.faces[3][r][i]
	}

	if r == 0 {
		cube.faces[0].turnCCW()
	}

	if r == size-1 {
		cube.faces[1].turnCW()
	}

	return notation
}

func (cube *Cube) ColTurnUp(c int) string {
	notation := fmt.Sprintf("%d%s  ", c+1, up)
	for i := 0; i < size; i++ {
		cube.faces[0][i][c], cube.faces[4][i][c], cube.faces[1][i][c], cube.faces[5][size-1-i][size-1-c] =
			cube.faces[4][i][c], cube.faces[1][i][c], cube.faces[5][size-i-1][size-c-1], cube.faces[0][i][c]
	}

	if c == 0 {
		cube.faces[2].turnCCW()
	}

	if c == size-1 {
		cube.faces[3].turnCW()
	}

	return notation
}

func (cube *Cube) ColTurnDn(c int) string {
	notation := fmt.Sprintf("%d%s  ", c+1, down)
	for i := 0; i < size; i++ {
		cube.faces[4][i][c], cube.faces[1][i][c], cube.faces[5][size-i-1][size-c-1], cube.faces[0][i][c] =
			cube.faces[0][i][c], cube.faces[4][i][c], cube.faces[1][i][c], cube.faces[5][size-1-i][size-1-c]
	}

	if c == 0 {
		cube.faces[2].turnCW()
	}

	if c == size-1 {
		cube.faces[3].turnCCW()
	}

	return notation
}

func (cube *Cube) FaceTurnCW(f int) string {
	notation := fmt.Sprintf("%d%s  ", f+1, cw)
	for i := 0; i < size; i++ {
		cube.faces[2][i][size-1-f], cube.faces[0][size-1-f][size-1-i], cube.faces[3][size-1-i][f], cube.faces[1][f][i] =
			cube.faces[1][f][i], cube.faces[2][i][size-1-f], cube.faces[0][size-1-f][size-1-i], cube.faces[3][size-1-i][f]
	}

	if f == 0 {
		cube.faces[4].turnCW()
	}

	if f == size-1 {
		cube.faces[5].turnCCW()
	}

	return notation
}

func (cube *Cube) FaceTurnCCW(f int) string {
	notation := fmt.Sprintf("%d%s  ", f+1, ccw)
	for i := 0; i < size; i++ {
		cube.faces[1][f][i], cube.faces[2][i][size-1-f], cube.faces[0][size-1-f][size-1-i], cube.faces[3][size-1-i][f] =
			cube.faces[2][i][size-1-f], cube.faces[0][size-1-f][size-1-i], cube.faces[3][size-1-i][f], cube.faces[1][f][i]
	}

	if f == 0 {
		cube.faces[4].turnCCW()
	}

	if f == size-1 {
		cube.faces[5].turnCW()
	}

	return notation
}

func (cube *Cube) Move(m, i int) string {
	switch m {
	case 0:
		return cube.RowTurnCW(i)
	case 1:
		return cube.ColTurnUp(i)
	case 2:
		return cube.FaceTurnCW(i)
	case 3:
		return cube.FaceTurnCCW(i)
	case 4:
		return cube.ColTurnDn(i)
	case 5:
		return cube.RowTurnCCW(i)
	}
	return "undefined move!"
}


func (cube *Cube) MoveX(bitmap int) string {
	i := bitmap & RowMask
	m := bitmap >> 2

	switch m {
	case 0b0001:
		return cube.RowTurnCW(i)
	case 0b1001:
		return cube.RowTurnCCW(i)
	case 0b0010:
		return cube.ColTurnUp(i)
	case 0b1010:
		return cube.ColTurnDn(i)
	case 0b0100:
		return cube.FaceTurnCW(i)
	case 0b1100:
		return cube.FaceTurnCCW(i)
	}
	return "undefined move!"
}

func (cube *Cube) RandomMove() string {
	//rand.Seed(time.Now().UnixNano())
	m := rand.Intn(6)
	i := rand.Intn(size)

	return cube.Move(m, i)
}

func (face *Face) turnCW() {
	for r := 0; r < (size+1)/2; r++ {
		for c := 0; c < size/2; c++ {
			face[r][c], face[size-1-c][r], face[size-1-r][size-1-c], face[c][size-1-r] =
				face[size-1-c][r], face[size-1-r][size-1-c], face[c][size-1-r], face[r][c]
		}
	}
}

func (face *Face) turnCCW() {
	for r := 0; r < (size+1)/2; r++ {
		for c := 0; c < size/2; c++ {
			face[size-1-c][r], face[size-1-r][size-1-c], face[c][size-1-r], face[r][c] =
				face[r][c], face[size-1-c][r], face[size-1-r][size-1-c], face[c][size-1-r]
		}
	}
}
