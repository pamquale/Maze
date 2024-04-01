package main

import (
	"fmt"
)

func main() {
	var width, height int
	fmt.Print("Enter maze width: ")
	fmt.Scanln(&width)
	fmt.Print("Enter maze height: ")
	fmt.Scanln(&height)

	maze := NewMaze(width, height)
	maze.generate()
	maze.display()
}
