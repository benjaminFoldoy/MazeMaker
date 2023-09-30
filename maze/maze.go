package maze

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/rand"
	"os"
)

// ############################################################
// Name: Maze
// Date: 31.07.22
// Time: 16.01
// ############################################################
// Dependencies: Tile
// Description : Maze is the structure that contains the laberynth.
// ############################################################
type Maze struct {
	Size             Size
	TileMatrix       [][]Tile
	MinNumberOfTurns int
	MaxLength        int
	Start            Position
	End              Position
	Margin           int
}

// ############################################################
// Name: CrdIsInspected
// Date: 05.08.22
// Time: 21.47
// ############################################################
// Dependencies: Tile, Position, Maze
// Description : Checks if a given coordinate is inspected
// ############################################################
func (maze *Maze) CrdIsInspected(Crd Position) bool {
	if !maze.CrdIsOutOfBounds(Crd) {
		return maze.TileMatrix[Crd.X][Crd.Y].IsInspected
	} else {
		return false
	}
}

// ############################################################
// Name: CrdIsOutOfBounds
// Date: 05.08.22
// Time: 21.47
// ############################################################
// Dependencies: Tile, Position, Maze
// Description : Checks if a given coordinate is bigger than or
// smaller than the height or width of the matrix
// ############################################################
func (maze *Maze) CrdIsOutOfBounds(Crd Position) bool {
	return Crd.X < 0 || Crd.Y < 0 || Crd.X >= maze.Width() || Crd.Y >= maze.Height()
}

// ############################################################
// Name: OutOfBoundsOrInspected
// Date: 05.08.22
// Time: 21.53
// ############################################################
// Dependencies: Tile, Position, Maze
// Description : Checks if a given coordinate is bigger than or
// smaller than the height or width of the matrix, as well as
// if the coordinate is already inspected.
// ############################################################
func (maze *Maze) OutOfBoundsOrInspected(Crd Position) bool {
	return maze.CrdIsInspected(Crd) || maze.CrdIsOutOfBounds(Crd)
}

// ############################################################
// Name: Create
// Date: 06.08.22
// Time: 11.19
// ############################################################
// Dependencies: Maze, math.rand
// Description : goes through the process of generating the maze
// and creating the .png file with the visualized maze.
// Usage				:
// ############################################################
func (maze *Maze) Create(fileName string, seed int64) bool {
	rand.Seed(seed)
	//println("seed : ", seed)

	maze.TileMatrix = maze.CreateBlankMatrix()
	startPos := maze.SetStartPosition()
	maze.Start = startPos
	if maze.OutOfBoundsOrInspected(startPos) {
		return false
	}
	maze.TileMatrix[startPos.X][startPos.Y].IsInspected = true

	//println("start : ", startPos.X, startPos.Y)
	lines := []Position{}
	lastPos, lastDirection := maze.NewLine(startPos, Position{0, 0})
	lines = append(lines, lastPos)

	i := 0
	for {
		//for i := 0; i <= maze.MinNumberOfTurns; {
		//println(lines[i].X, lines[i].Y)
		lastPos, lastDirection = maze.NewLine(lastPos, lastDirection)
		lines = append(lines, lastPos)
		possibleDirections := maze.CheckPossibleDirections(lastPos).directionsToPositions()

		shouldRevert := false

		//if there is no possible directions, it should revert
		if len(possibleDirections) == 0 {
			shouldRevert = true
			return false
		}

		//if all possible directions are enclosed, it should revert
		var onlyEnclosedDirections = true
		for j := 0; j < len(possibleDirections); j++ {
			if !maze.IsEnclosed(maze.Pool(lastPos.OffsetFromPosition(possibleDirections[j]))) {
				onlyEnclosedDirections = false
			}
		}
		if onlyEnclosedDirections {
			shouldRevert = true
			return false
		}

		//the actual revertion
		if shouldRevert {
			//println("last line1: ", lines[len(lines)-1].X, " , ", lines[len(lines)-1].Y)
			startPoint := lines[len(lines)-2]
			endPoint := lines[len(lines)-1]
			fixForStartTile := Position{}.GetDirectionFromTwoPoints(startPoint, endPoint)
			maze.RemoveLine(startPoint, endPoint)
			lines = lines[:len(lines)-1]
			lastPos = lines[len(lines)-1]
			maze.TileMatrix[startPoint.X][startPoint.Y].WallEnd(fixForStartTile)
			//ln("last line2: ", lines[len(lines)-1].X, " , ", lines[len(lines)-1].Y)
		} else {
			i++
		}

		if lastPos.IsAtBoundry(maze.Size) && i >= maze.MinNumberOfTurns {
			maze.End = lastPos
			//println("Start: ", maze.Start.X, ", ", maze.Start.Y, ", End: ", maze.End.X, ", ", maze.End.Y)
			break
		}
	}

	//ADD NEW "WRONG" LINES

	//pool := maze.Pool(Position{12, 0})
	//maze.NewLine(test[RandomLength(0, len(test)-1)], Position{0, 1})

	for {
		test := maze.GetAllTilesWithUnInspectedNeighbors()
		if len(test) == 0 {
			break
		}
		lastPos = test[RandomLength(0, len(test))]
		for i := 0; i < RandomLength(0, maze.MaxLength); i++ {
			lastPos, lastDirection = maze.NewLine(lastPos, lastDirection)
			if lastPos.Match(Position{}) {
				i = maze.MaxLength
			} else {
				lines = append(lines, lastPos)
			}
		}
	}

	maze.PrettyUpStartAndEnd()
	maze.CreatePng(fileName)
	//maze.PrintInspectedTiles()
	return true
}

// ############################################################
// Name: Initialize
// Date: 03.08.22
// Time: 08.23
// ############################################################
// Dependencies: Maze
// Description : Used to create an instance of Maze
// Usage				: maze.initialize(Size, minimum number of turns)
// ############################################################
func (maze *Maze) Initialize(size Size, minNumberOfTurns int) {
	maze.Size = size
	maze.MinNumberOfTurns = minNumberOfTurns
}

// ############################################################
// Name: Width and Height getters
// Date: 03.08.22
// Time: 09.19
// ############################################################
// Dependencies: Maze
// Description : Getters for Width and Height.
// This is just to make it so that you dont have to write maze.Size.height every time...
// Usage				: maze.width() - maze.height(T
// #########################################M##################
func (maze Maze) Width() int {
	return maze.Size.Width
}
func (maze Maze) Height() int {
	return maze.Size.Height
}

// ############################################################
// Name: CreateMatrix
// Date: 03.08.22
// ############################################################
// Dependencies: Maze
// Description : Used to create a 2d matrix with empty Tile that have the properties
// Left, Right, Up, Down = false
// Usage				: maze.createBlankMatrix()
// ############################################################
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

// ############################################################
// Name: RandomLenght
// Date: 05.08.22
// Time: 10.13
// ############################################################
// Dependencies	: math.rand
// Description 	: inputs a min and max value, and returns a value between the numbers
// Usage			: randomLength(3, 7) - can return values: 3, 4, 5, 6 and 7
// ############################################################
func RandomLength(min int, max int) int {
	return min + int(math.Floor(rand.Float64()*float64((max-min))))
	//return min + rand.Intn(max-min+1)
}

// ############################################################
// Name: SetStartPosition
// Date: 05.08.22
// Time: 10.38
// ############################################################
// Dependencies: math.rand, Maze.tileMatrix
// Description : Sets the starting position of the maze
// Usage				: maze.setStartPosition()
// ############################################################
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

// ############################################################
// Name: CheckPossibleDirecions
// Date: 05.08.22
// Time: 21.42
// ############################################################
// Dependencies: Maze.tileMatrix
// Description : Takes a coordinate as input and checks what
// neighbouring tiles are taken. Returns Directions.
// Usage				: maze.chackPossibleDirections
// ############################################################
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

// ############################################################
// Name: printInspectedTiles
// Date: 06.08.22
// Time: 00.21
// ############################################################
// Dependencies: Maze.tileMatrix
// Description : Prints the isInspected property of each tile in maze.TileMatrix
// Usage				: maze.printInspectedTiles()
// ############################################################
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

// ############################################################
// Name: NewLine
// Date: 06.08.22
// Time: 11.23
// ############################################################
// Dependencies: Maze.tileMatrix
// Description :
// Usage				: maze.newLine(position)
// ############################################################
func (maze *Maze) NewLine(startPos Position, lastDirection Position) (Position, Position) {

	lineDirection := Position{} //The Direction of the new line

	//ADD CHECK, Does the end of this line end up only with pools without access to boundries?

	//If cutoff, choose the direction with access to bounds, and biggest pool
	if maze.OutOfBoundsOrInspected(startPos.OffsetFromPosition(lastDirection)) {
		possibleDirections := maze.CheckPossibleDirections(startPos).directionsToPositions()
		pools := [][]Position{}
		var indexes []int
		bestIndex := 0
		//the directions that can lead to one of the sides ==> pools
		for i := 0; i < len(possibleDirections); i++ {
			pool := maze.Pool(startPos.OffsetFromPosition(possibleDirections[i]))
			pools = append(pools, pool)
			//print(startPos.X, startPos.Y, maze.IsEnclosed(pool))
			if !maze.IsEnclosed(pool) {
				indexes = append(indexes, i)
			}
			//println(len(pools))
		}

		for i := 0; i < len(indexes); i++ {
			if len(pools[indexes[i]]) >= len(pools[bestIndex]) {
				bestIndex = indexes[i]
			}
		}
		//maze.PrintInspectedTiles()
		//println()
		if len(possibleDirections) == 0 {
			return Position{}, Position{}
		}
		lineDirection = possibleDirections[bestIndex]
	} else {
		lineDirection = maze.CheckPossibleDirections(startPos).RandomDirection()
	}
	lineLength := RandomLength(1, maze.MaxLength)
	lastPosition := startPos
	//fmt.Println(lineLength, lineDirection.X, lineDirection.Y)
	Crds := []Position{}
	for i := 1; i <= lineLength; i++ {
		Crds = append(Crds, Position{startPos.X + lineDirection.X*i, startPos.Y + lineDirection.Y*i})
		if !maze.OutOfBoundsOrInspected(Crds[len(Crds)-1]) {
			tile := &maze.TileMatrix[Crds[len(Crds)-1].X][Crds[len(Crds)-1].Y]

			//hasHitWall = maze.OutOfBoundsOrInspected(Position{startPos.X + lineDirection.X*(i+1), startPos.Y + lineDirection.Y*(i+1)})

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

	return lastPosition, lineDirection
}

// ############################################################
// Name: WallSides
// Date: 06.08.22
// Time: 18.31
// ############################################################
// Dependencies: Maze, Tile
// Description : Adds a walls perpendicular to the linedirection input
// Usage				: tile.wallEnd(direction)
// ############################################################
func (tile *Tile) WallSides(direction Position) {
	if direction.Y == 0 { //if the direction is horizontal
		tile.Bottom = true
		tile.Top = true
	} else {
		tile.Left = true
		tile.Right = true
	}
}

// ############################################################
// Name: WallEnd
// Date: 06.08.22
// Time: 18.26
// ############################################################
// Dependencies: Maze, Tile
// Description : Adds a wall depending on the linedirection input
// Usage				: tile.wallEnd(direction)
// ############################################################
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

// ############################################################
// Name: ClearWallStart
// Date: 07.08.22
// Time: 15.10
// ############################################################
// Dependencies: Maze, Tile
// Description : Adds a wall depending on the linedirection input
// Usage				: tile.wallEnd(direction)
// ############################################################
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

func (maze *Maze) PrettyUpStartAndEnd() {
	startTile := &maze.TileMatrix[maze.Start.X][maze.Start.Y]
	endTile := &maze.TileMatrix[maze.End.X][maze.End.Y]

	//START TILE
	switch maze.Start.X {
	case 0:
		startTile.Left = false
	case maze.Width() - 1:
		startTile.Right = false
	}

	switch maze.Start.Y {
	case 0:
		startTile.Top = false
	case maze.Height() - 1:
		startTile.Bottom = false
	}
	if (maze.Start.X == 0 || maze.Start.X == maze.Width()-1) &&
		maze.Start.Y == 0 {
		startTile.Top = true
	} else if (maze.Start.X == 0 || maze.Start.X == maze.Width()-1) &&
		maze.Start.Y == maze.Height()-1 {
		startTile.Bottom = true
	}

	//END TILE
	switch maze.End.X {
	case 0:
		endTile.Left = false
	case maze.Width() - 1:
		endTile.Right = false
	}

	switch maze.End.Y {
	case 0:
		endTile.Top = false
	case maze.Height() - 1:
		endTile.Bottom = false
	}

	if (maze.End.X == 0 || maze.End.X == maze.Width()-1) &&
		maze.End.Y == 0 {
		endTile.Top = true
	} else if (maze.End.X == 0 || maze.End.X == maze.Width()-1) &&
		maze.End.Y == maze.Height()-1 {
		endTile.Bottom = true
	}
}

// ############################################################
// Name: CreatePng
// Date: 07.08.22
// Time: 09.41
// ############################################################
// Dependencies: Maze, Tile
// Description : The whole pipeline for creating a png of the maze.
// Usage				: maze.createPng("FILE NAME")
// ############################################################
func (maze *Maze) CreatePng(fileName string) {
	width := maze.Width()*9 + 3 + maze.Margin*2
	height := maze.Height()*9 + 3 + maze.Margin*2

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, color.White)
		}
	}

	//SETTING THE START ARROWS
	for x := maze.Start.X*9 + maze.Margin; x < (maze.Start.X+1)*9+3+maze.Margin; x++ {
		for y := maze.Start.Y*9 + maze.Margin; y < (maze.Start.Y+1)*9+3+maze.Margin; y++ {
			img.Set(x, y, color.RGBA{0, 255, 0, 255})
		}
	}
	//SETTING THE END ARROWS
	for x := maze.End.X*9 + maze.Margin; x < (maze.End.X+1)*9+3+maze.Margin; x++ {
		for y := maze.End.Y*9 + maze.Margin; y < (maze.End.Y+1)*9+3+maze.Margin; y++ {
			img.Set(x, y, color.RGBA{255, 0, 0, 255})
		}
	}

	for x1 := 0; x1 < maze.Width(); x1++ { //FOR EVERY CRD X
		for y1 := 0; y1 < maze.Height(); y1++ { //AND EVERY CRD X
			shape := Shape{}.Walls(maze.TileMatrix[x1][y1]) //CREATE THE REPRESENTATIVE SHAPE

			for i := 0; i < shape.Length(); i++ { //AND FOR EVERY WALL IN THE SHAPE
				for x2 := shape.Start[i].X; x2 < shape.End[i].X; x2++ { //AND EVERY STARTING POINT...
					for y2 := shape.Start[i].Y; y2 < shape.End[i].Y; y2++ { /// AND ENDPOINT OF THE WALLS
						img.Set(x2+x1*9+maze.Margin, y2+y1*9+maze.Margin, color.Black)
					}
				}
			}
		}
	}

	// Encode as PNG.
	f, _ := os.Create(fileName + ".png")
	png.Encode(f, img)
}

// ############################################################
// Name:GetAllTilesWithUnInspected
// Date:12.08.22
// Time:20.00
// ############################################################
// Dependencies: Maze
// Description : returns a list of Positions only with
// Tiles in the Maze.TileMatrix that has at least one blank neighbor
// Usage				: Position.GetDirectionFromTwoPoints(Position, Position)
// ############################################################
func (maze Maze) GetAllTilesWithUnInspectedNeighbors() []Position {
	var checkedPositions []Position
	for x := 0; x < maze.Width(); x++ {
		for y := 0; y < maze.Height(); y++ {
			newCheckedPositions := maze.getUncheckedNeighbors(Position{x, y}, checkedPositions)
			if len(newCheckedPositions) != 0 && maze.TileMatrix[x][y].IsInspected {
				for i := 0; i < len(newCheckedPositions); i++ {
					checkedPositions = append(checkedPositions, Position{x, y})
				}
			}
		}
	}
	return checkedPositions
}
