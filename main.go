package main

type Maze struct {
	Grid   []string
	Width  int
	Height int
}

func createMaze() Maze {
	layout := Maze{
		Grid: []string{
			"##########",
			"#        #",
			"# ####### #",
			"# #     # #",
			"# # ### # #",
			"# # # # # #",
			"#   # #   #",
			"####### ###",
			"#C       #",
			"##########",
		},
		Width:  10,
		Height: 10,
	}
	return layout
}

func main() {
	maze := createMaze()

	// TODO: line-by-line maze printing
	printMaze()
}
