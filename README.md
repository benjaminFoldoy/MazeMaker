# MazeMaker
 A random maze-generator written in golang.

## Usage

`go run main.go *WIDTH* *HEIGHT* *OUTPUT FILE NAME*`<br>

## Examples
### 10 x 10 Maze
Command: <br>
`go run main.go 10 10 small_maze` <br>
Output image: <br>
![small_maze](output_images/10x10.png)

### 20 x 20 Maze
Command: <br>
`go run main.go 10 20 small_rectangle_maze` <br>
Output image: <br>
![rectangle_maze](output_images/10x20.png)

### 30 x 20 Maze
Command: <br>
`go run main.go 20 30 medium_rectangle_maze` <br>
Output image: <br>
![rectangle_maze](output_images/20x30.png)

### 70 x 70 Maze
Command: <br>
`go run main.go 70 70 large_maze` <br>
Output image: <br>
![large_maze](output_images/70x70.png)

### 100 x 100 Maze
Command: <br>
`go run main.go 100 100 huge_maze` <br>
Output image: <br>
![huge_maze](output_images/100x100.png)

## The Algorithm
The algorithm works by choosing a direction to draw a line in, and then drawing a line in that direction. It has a minimum amount of lines it has to draw for the "Solution". When the solution is drawn, the algorithm chooses one of the "cells" that has already been drawn in. It then draws a line from that cell, and continues this process until all cells have been drawn in.
