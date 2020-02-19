package rubik

import (
	"log"
	"sync"
)

var wg = sync.WaitGroup{}
var solved bool

// Solve the cube.
// experiment1 - 6 moves: 12 secs, 7 moves 5 mins
func Solve(cube Cube, max int) string {
	for i := 1; i <= max; i++ {
		log.Printf("Looking for a solution with %d moves\n", i)
		wg.Add(1)
		go solve(cube, "", -1, -1, i)
		wg.Wait()
		if solved {
			return "solved"
		}
	}
	return "not solved"
}

func solve(cube Cube, moves string, lastM, lastI, max int) {
	defer wg.Done()
	if solved {
		return
	}

	if cube.Solved() {
		solved = true
		log.Println(moves)
		return
	}

	if max <= 0 {
		return
	}

	for i := 0; i < size; i++ {
		for m := 0; m < 6; m++ {

			if lastI == i && lastM+m == 5 { // is reverse move
				continue
			}

			newCube := cube
			move := newCube.Move(m, i)

			wg.Add(1)
			go solve(newCube, moves+move, m, i, max-1)
		}
	}
}
