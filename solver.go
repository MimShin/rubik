package rubik

import (
	"log"
	"sync"
)

var wg = sync.WaitGroup{}
var solved bool

// Solve the cube.
// experiment1 - pruning undo move -> 6 moves: 12 secs, 7 moves 5 mins
// experiment2 - pruning 3 moves of the same type -> 6 moves: 10 secs, 7 moves 3:30 mins
// experiment3 - using bitmap for moves -> 6  moves: 9 secs, 7 moves: 3-7 mins
func Solve(cube Cube, max int) string {
	for i := 1; i <= max; i++ {
		log.Printf("Looking for a solution with %d moves\n", i)
		wg.Add(1)
		go solve(cube, "", NoMove, NoMove, i)
		wg.Wait()
		if solved {
			return "solved"
		}
	}
	return "not solved"
}

func solve(cube Cube, movesStr string, lastLastM, lastM, max int) {
	defer wg.Done()
	if solved {
		return
	}

	if cube.Solved() {
		solved = true
		log.Println(movesStr)
		return
	}

	if max <= 0 {
		return
	}

	for i := 0; i < len(Moves); {
		m := Moves[i]

		// Don't undo the last move
		if m^lastM == RevMask {
			i++
			continue
		}

		// 3 moves of the same type can always be replaced with 2 or less
		if lastLastM&lastM&MovMask == m&MovMask {
			i += 6
			continue
		}

		newCube := cube
		move := newCube.MoveX(m)
		i++

		wg.Add(1)
		go solve(newCube, movesStr+move, lastM, m, max-1)
	}
}
