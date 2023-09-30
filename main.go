package main

//to use the maze struct within the maze package, we need to import it.
//This is done by importing the package, and then using the package name as a prefix.
import (
	"MazeModule/maze"
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
// letting us pass arguments when calling main from console, with default values:
// to pass values for width, height, minNumberOfTurns, maxLength and margin, we need to pass them as arguments when calling main from console.
// Example: go run main.go 20 60 70 3 30
func main() {
	//Create the maze struct
	width, height, minNumberOfTurns, maxLength, margin := 20, 60, 2, 3, 30
	minNumberOfTurns = width*height/100 + 1

	//convert the arguments from string to int
	//if the conversion fails, the default values will be used.
	if len(os.Args) < 4 {
		panic("Not enough arguments, using default values\n")
	}
	w, error1 := strconv.Atoi(os.Args[1])
	h, error2 := strconv.Atoi(os.Args[2])
	n := os.Args[3]

	if w < 10 || h < 10 {
		panic("Width and height must be at least 10")
	}

	if error1 == nil {
		width = w
	}
	if error2 == nil {
		height = h
	}
	maze := maze.Maze{
		Size:             maze.Size{Width: width, Height: height},
		MinNumberOfTurns: minNumberOfTurns,
		MaxLength:        maxLength,
		Margin:           margin,
	}
	maze.Create("output_images/"+n, time.Now().UnixNano())
	print("\033[32mSuccess!\033[35m\nCreated image at path: output_images/" + n + "\n\033[0m")
}
