package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/mimshin/rubik"
)

func main() {

	var c rubik.Cube
	c.Init()

	max, _ := strconv.Atoi(os.Args[1])

	for i := 0; i < max; i++ {
		fmt.Println(c.RandomMove())
	}
	c.Print()

	start := time.Now()
	solution := rubik.Solve(c, max)
	fmt.Println(solution)
	fmt.Printf("Elapsed time: %s\n", time.Since(start))
}
