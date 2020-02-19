package rubik

import (
	"log"
	"sync"
)

var wg = sync.WaitGroup{}
var solved bool

var visited map[string]int
var mutex sync.Mutex

// Solve the cube.
// experiment2 - 6 moves: 1:05 mins, 7 moves: no response in more thant 1 hour
func Solve(cube Cube, max int) string {
	visited = make(map[string]int)
	for i := 1; i <= max; i++ {
		log.Printf("Visited nodes @l%d: %d", i, len(visited))
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

			if isVisited(newCube, max) {
				continue
			}

			wg.Add(1)
			go solve(newCube, moves+move, m, i, max-1)
		}
	}
}

func isVisited(cube Cube, moves int) bool {
	defer mutex.Unlock()
	mutex.Lock()

	k := cube.Key()
	if visited[k] >= moves {
		return true
	}

	visited[k] = moves
	return false
}
