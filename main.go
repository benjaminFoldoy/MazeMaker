package main

//to use the maze struct within the maze package, we need to import it.
//This is done by importing the package, and then using the package name as a prefix.
import (
	"MazeModule/maze"
	"math"
	"os"
	"strconv"
	"time"
)

// ############################################################
// Name: Main
// Date: Who knows
// Time:
// ############################################################
// Dependencies: Everything
// Description : Everything
// ############################################################
func main() {
	//Create the maze struct
	if len(os.Args) < 4 {
		panic("Not enough arguments, using default values\n")
	}
	w, error1 := strconv.Atoi(os.Args[1])
	h, error2 := strconv.Atoi(os.Args[2])
	n := os.Args[3]

	//convert the arguments from string to int
	//if the conversion fails, the default values will be used.

	if w < 10 || h < 10 {
		panic("Width and height must be at least 10")
	}
	width, height, margin := 20, 60, 30

	if error1 == nil {
		width = w
	}
	if error2 == nil {
		height = h
	}
	minNumberOfTurns := int(math.Sqrt(float64(width*height)) * 2)
	maxLength := int(min(min(width, height)/3+1, 12))

	maze := maze.Maze{
		Size:             maze.Size{Width: width, Height: height},
		MinNumberOfTurns: minNumberOfTurns,
		MaxLength:        maxLength,
		Margin:           margin,
	}
	//maze.Create("output_images/"+n, )
	running := true
	for running {
		running = !maze.Create("output_images/"+n, time.Now().UnixNano())
	}
	print("\033[32mSuccess!\033[35m\nCreated image at path: output_images/" + n + "\n\033[0m")
}
