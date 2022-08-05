package main

/*
//############################################################
//Name: Maze
//Date: 31.07.22
//Time: 16.01
//############################################################
//Dependencies: Tile
//Description : Maze is the structure that contains the laberynth.
//############################################################
type Maze struct {
	size             tools.Size
	tileMatrix       [][]Tile
	minNumberOfTurns int
}

//############################################################
//Name:
//Date:
//Time:
//############################################################
//Dependencies:
//Description :
//Usage				:
//############################################################
func (maze Maze) create(fileName string) {
	maze.tileMatrix = maze.createBlankMatrix()
}

//############################################################
//Name: Initialize
//Date: 03.08.22
//Time: 08.23
//############################################################
//Dependencies: Maze
//Description : Used to create an instance of Maze
//Usage				: maze.initialize(Size, minimum number of turns)
//############################################################
func (maze Maze) initialize(size Size, minNumberOfTurns int) Maze {
	maze.size = size
	maze.minNumberOfTurns = minNumberOfTurns
	return maze
}

//############################################################
//Name: width and height getters
//Date: 03.08.22
//Time: 09.19
//############################################################
//Dependencies: Maze
//Description : Getters for Width and Height.
//This is just to make it so that you dont have to write maze.size.height every time...
//Usage				: maze.width() - maze.height()
//############################################################
func (maze Maze) width() int {
	return maze.size.width
}
func (maze Maze) height() int {
	return maze.size.height
}

//############################################################
//Name: CreateMatrix
//Date: 03.08.22
//Time: 08.38
//############################################################
//Dependencies: Maze
//Description : Used to create a 2d matrix with empty Tile that have the properties
//Left, Right, Up, Down = false
//Usage				: maze.createBlankMatrix()
//############################################################
func (maze Maze) createBlankMatrix() [][]Tile {
	output := [][]Tile{}

	for x := 0; x < maze.width(); x++ {
		output = append(output, []Tile{})
		for y := 0; y < maze.height(); y++ {
			output[x] = append(output[x], Tile{})
		}
	}
	return output
}
*/