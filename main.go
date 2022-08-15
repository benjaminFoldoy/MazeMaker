package main

//############################################################
//Name: Main
//Date: Who knows
//Time:
//############################################################
//Dependencies: Everything
//Description : Everything
//############################################################
func main() {
	//Create the maze struct
	maze := Maze{
		Size:             Size{Width: 16, Height: 16},
		MinNumberOfTurns: 10,
		MaxLength:        4,
	}
	//maze.Create("Test", time.Now().UnixNano())
	maze.Create("test", 1660019299025486300)
}
