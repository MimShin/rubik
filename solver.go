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
func Solve(cube Cube, max int) string {
	for i := 1; i <= max; i++ {
		log.Printf("Looking for a solution with %d moves\n", i)
		wg.Add(1)
		go solve(cube, "", -1, -1, -1, i)
		wg.Wait()
		if solved {
			return "solved"
		}
	}
	return "not solved"
}

func solve(cube Cube, movesStr string, lastLastM, lastM, lastI, max int) {
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

	for m := 0; m < 6; m++ {

		// 3 moves of the same type can always be replaced with 2 or less
		if lastLastM == lastM && (lastM == m || lastM+m == 5) {
			continue
		}
		if lastLastM+lastM == 5 && (lastM == m || lastLastM == m) {
			continue
		}

		for i := 0; i < size; i++ {

			// Don't undo the last move
			if lastI == i && lastM+m == 5 {
				continue
			}

			newCube := cube
			move := newCube.Move(m, i)

			wg.Add(1)
			go solve(newCube, movesStr+move, lastM, m, i, max-1)
		}
	}
}
