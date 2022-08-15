package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/rand"
	"os"
)

//############################################################
//Name: Maze
//Date: 31.07.22
//Time: 16.01
//############################################################
//Dependencies: Tile
//Description : Maze is the structure that contains the laberynth.
//############################################################
type Maze struct {
	Size             Size
	TileMatrix       [][]Tile
	MinNumberOfTurns int
	MaxLength        int
}

//############################################################
//Name: CrdIsInspected
//Date: 05.08.22
//Time: 21.47
//############################################################
//Dependencies: Tile, Position, Maze
//Description : Checks if a given coordinate is inspected
//############################################################
func (maze *Maze) CrdIsInspected(Crd Position) bool {
	if !maze.CrdIsOutOfBounds(Crd) {
		return maze.TileMatrix[Crd.X][Crd.Y].IsInspected
	} else {
		return false
	}
}

//############################################################
//Name: CrdIsOutOfBounds
//Date: 05.08.22
//Time: 21.47
//############################################################
//Dependencies: Tile, Position, Maze
//Description : Checks if a given coordinate is bigger than or
//smaller than the height or width of the matrix
//############################################################
func (maze *Maze) CrdIsOutOfBounds(Crd Position) bool {
	return Crd.X < 0 || Crd.Y < 0 || Crd.X >= maze.Width() || Crd.Y >= maze.Height()
}

//############################################################
//Name: OutOfBoundsOrInspected
//Date: 05.08.22
//Time: 21.53
//############################################################
//Dependencies: Tile, Position, Maze
//Description : Checks if a given coordinate is bigger than or
//smaller than the height or width of the matrix, as well as
//if the coordinate is already inspected.
//############################################################
func (maze *Maze) OutOfBoundsOrInspected(Crd Position) bool {
	return maze.CrdIsInspected(Crd) || maze.CrdIsOutOfBounds(Crd)
}

//############################################################
//Name: Create
//Date: 06.08.22
//Time: 11.19
//############################################################
//Dependencies: Maze, math.rand
//Description : goes through the process of generating the maze
//and creating the .png file with the visualized maze.
//Usage				:
//############################################################
func (maze *Maze) Create(fileName string, seed int64) {

	rand.Seed(seed)
	println("seed : ", seed)

	maze.TileMatrix = maze.CreateBlankMatrix()
	startPos := maze.SetStartPosition()
	maze.TileMatrix[startPos.X][startPos.Y].IsInspected = true

	println("start : ", startPos.X, startPos.Y)
	lines := []Position{}
	lines = append(lines, maze.NewLine(startPos))
	for i := 0; i <= maze.MinNumberOfTurns; i++ {
		//println(lines[i].X, lines[i].Y)
		lines = append(lines, maze.NewLine(lines[len(lines)-1]))
	}

	maze.CreatePng(fileName)
	maze.PrintInspectedTiles()
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
func (maze *Maze) Initialize(size Size, minNumberOfTurns int) {
	maze.Size = size
	maze.MinNumberOfTurns = minNumberOfTurns
}

//############################################################
//Name: Width and Height getters
//Date: 03.08.22
//Time: 09.19
//############################################################
//Dependencies: Maze
//Description : Getters for Width and Height.
//This is just to make it so that you dont have to write maze.Size.height every time...
//Usage				: maze.width() - maze.height(T
//#########################################M##################
func (maze Maze) Width() int {
	return maze.Size.Width
}
func (maze Maze) Height() int {
	return maze.Size.Height
}

//############################################################
//Name: CreateMatrix
//Date: 03.08.22
//############################################################
//Dependencies: Maze
//Description : Used to create a 2d matrix with empty Tile that have the properties
//Left, Right, Up, Down = false
//Usage				: maze.createBlankMatrix()
//############################################################
func (maze Maze) CreateBlankMatrix() [][]Tile {
	output := [][]Tile{}

	for x := 0; x < maze.Width(); x++ {
		output = append(output, []Tile{})
		for y := 0; y < maze.Height(); y++ {
			output[x] = append(output[x], Tile{})
		}
	}
	return output
}

//############################################################
//Name: RandomLenght
//Date: 05.08.22
//Time: 10.13
//############################################################
//Dependencies	: math.rand
//Description 	: inputs a min and max value, and returns a value between the numbers
//Usage			: randomLength(3, 7) - can return values: 3, 4, 5, 6 and 7
//############################################################
func RandomLength(min int, max int) int {
	return min + int(math.Floor(rand.Float64()*float64((max-min))))
	//return min + rand.Intn(max-min+1)
}

//############################################################
//Name: SetStartPosition
//Date: 05.08.22
//Time: 10.38
//############################################################
//Dependencies: math.rand, Maze.tileMatrix
//Description : Sets the starting position of the maze
//Usage				: maze.setStartPosition()
//############################################################
func (maze *Maze) SetStartPosition() Position {
	directions := Directions{true, true, true, true}
	direction := directions.RandomDirection()

	startPosition := Position{0, 0}
	if direction.Match(Position{0, -1}) {
		startPosition.X = RandomLength(1, maze.Width()-2)
	} else if direction.Match(Position{0, 1}) {
		startPosition.Y = maze.Height() - 1
		startPosition.X = RandomLength(1, maze.Width()-2)
	} else if direction.Match(Position{-1, 0}) {
		startPosition.Y = RandomLength(0, maze.Height()-1)
	} else if direction.Match(Position{1, 0}) {
		startPosition.Y = RandomLength(0, maze.Height()-1)
		startPosition.X = maze.Height() - 1
	}

	return startPosition
}

//############################################################
//Name: CheckPossibleDirecions
//Date: 05.08.22
//Time: 21.42
//############################################################
//Dependencies: Maze.tileMatrix
//Description : Takes a coordinate as input and checks what
//neighbouring tiles are taken. Returns Directions.
//Usage				: maze.chackPossibleDirections
//############################################################
func (maze Maze) CheckPossibleDirections(startPosition Position) Directions {
	possibleDirections := Directions{true, true, true, true}

	if maze.OutOfBoundsOrInspected(startPosition.OffsetFromPosition(Position{}.Right())) {
		possibleDirections.Right = false
	}

	if maze.OutOfBoundsOrInspected(startPosition.OffsetFromPosition(Position{}.Left())) {
		possibleDirections.Left = false
	}

	if maze.OutOfBoundsOrInspected(startPosition.OffsetFromPosition(Position{}.Down())) {
		possibleDirections.Down = false
	}

	if maze.OutOfBoundsOrInspected(startPosition.OffsetFromPosition(Position{}.Up())) {
		possibleDirections.Up = false
	}

	return possibleDirections
}

//############################################################
//Name: printInspectedTiles
//Date: 06.08.22
//Time: 00.21
//############################################################
//Dependencies: Maze.tileMatrix
//Description : Prints the isInspected property of each tile in maze.TileMatrix
//Usage				: maze.printInspectedTiles()
//############################################################
func (maze Maze) PrintInspectedTiles() {
	for y := 0; y < maze.Height(); y++ {
		for x := 0; x < maze.Width(); x++ {
			if maze.TileMatrix[x][y].IsInspected {
				print(" \033[31mX\033[0m ")
			} else {
				print(" \033[36mO\033[0m ")
			}
		}
		println()
	}
}

//############################################################
//Name: NewLine
//Date: 06.08.22
//Time: 11.23
//############################################################
//Dependencies: Maze.tileMatrix
//Description :
//Usage				: maze.newLine(position)
//############################################################
func (maze *Maze) NewLine(startPos Position) Position {
	lineDirection := maze.CheckPossibleDirections(startPos).RandomDirection()
	lineLength := RandomLength(1, maze.MaxLength)
	lastPosition := startPos

	Crds := []Position{}
	for i := 1; i <= lineLength; i++ {
		Crds = append(Crds, Position{startPos.X + lineDirection.X*i, startPos.Y + lineDirection.Y*i})
		if !maze.OutOfBoundsOrInspected(Crds[len(Crds)-1]) {
			tile := &maze.TileMatrix[Crds[len(Crds)-1].X][Crds[len(Crds)-1].Y]

			tile.IsInspected = true
			lastPosition = Position{startPos.X + lineDirection.X*i, startPos.Y + lineDirection.Y*i}

			//SIDELINES
			tile.WallSides(lineDirection)
			//ENDLINE
			if i == lineLength {
				tile.WallEnd(lineDirection)
			}
		} else {
			tile := &maze.TileMatrix[Crds[len(Crds)-2].X][Crds[len(Crds)-2].Y]
			//ENDLINE
			tile.WallEnd(lineDirection)
			break
		}
	}

	maze.TileMatrix[startPos.X][startPos.Y].ClearWallStart(lineDirection)

	return lastPosition
}

//############################################################
//Name: WallSides
//Date: 06.08.22
//Time: 18.31
//############################################################
//Dependencies: Maze, Tile
//Description : Adds a walls perpendicular to the linedirection input
//Usage				: tile.wallEnd(direction)
//############################################################
func (tile *Tile) WallSides(direction Position) {
	if direction.Y == 0 { //if the direction is horizontal
		tile.Bottom = true
		tile.Top = true
	} else {
		tile.Left = true
		tile.Right = true
	}
}

//############################################################
//Name: WallEnd
//Date: 06.08.22
//Time: 18.26
//############################################################
//Dependencies: Maze, Tile
//Description : Adds a wall depending on the linedirection input
//Usage				: tile.wallEnd(direction)
//############################################################
func (tile *Tile) WallEnd(direction Position) {
	switch direction.X {
	case 1:
		tile.Right = true
	case -1:
		tile.Left = true
	}

	switch direction.Y {
	case 1:
		tile.Bottom = true
	case -1:
		tile.Top = true
	}
}

//############################################################
//Name: ClearWallStart
//Date: 07.08.22
//Time: 15.10
//############################################################
//Dependencies: Maze, Tile
//Description : Adds a wall depending on the linedirection input
//Usage				: tile.wallEnd(direction)
//############################################################
func (tile *Tile) ClearWallStart(direction Position) {
	switch direction.X {
	case 1:
		tile.Right = false
	case -1:
		tile.Left = false
	}

	switch direction.Y {
	case 1:
		tile.Bottom = false
	case -1:
		tile.Top = false
	}
}

//############################################################
//Name: CreatePng
//Date: 07.08.22
//Time: 09.41
//############################################################
//Dependencies: Maze, Tile
//Description : The whole pipeline for creating a png of the maze.
//Usage				: maze.createPng("FILE NAME")
//############################################################
func (maze *Maze) CreatePng(fileName string) {
	width := maze.Width()*9 + 3
	height := maze.Height()*9 + 3

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	for x := 0; x < maze.Size.Width*9+3; x++ {
		for y := 0; y < maze.Size.Width*9+3; y++ {
			img.Set(x, y, color.White)
		}
	}

	for x1 := 0; x1 < maze.Width(); x1++ { //FOR EVERY CRD X
		for y1 := 0; y1 < maze.Height(); y1++ { //AND EVERY CRD X
			shape := Shape{}.Walls(maze.TileMatrix[x1][y1]) //CREATE THE REPRESENTATIVE SHAPE

			for i := 0; i < shape.Length(); i++ { //AND FOR EVERY WALL IN THE SHAPE
				for x2 := shape.Start[i].X; x2 < shape.End[i].X; x2++ { //AND EVERY STARTING POINT...
					for y2 := shape.Start[i].Y; y2 < shape.End[i].Y; y2++ { /// AND ENDPOINT OF THE WALLS
						img.Set(x2+x1*9, y2+y1*9, color.Black)
					}
				}
			}
		}
	}

	// Encode as PNG.
	f, _ := os.Create(fileName + ".png")
	png.Encode(f, img)
}

//############################################################
//Name: IsEnclosed
//Date: 09.08.22
//Time: 06.49
//############################################################
//Dependencies:
//Description : Checks if a "pool" can reach one of the walls
//Usage				: maze.IsEnclosed($COORDINATE)
//############################################################
func (maze Maze) IsEnclosed(position Position, direction Directions) bool {
	//if position.IsAtBoundry() {
	//	return true
	//}

	//CREATE AND FILL CHECKED CRDS MARTIX
	checkedCrds := [][]bool{}
	for x := 0; x < maze.Width(); x++ {
		for y := 0; y < maze.Height(); y++ {
			checkedCrds[x][y] = false
		}
	}

	for {
		if !direction.IsHorizontal() {

		}
	}
}
