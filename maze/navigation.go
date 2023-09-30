package maze

//############################################################
//Name: Directions
//Date: 31.07.22
//Time: 16.02
//############################################################
//Description: Used to see what directions are avalable
//############################################################
type Directions struct {
	Up, Down, Left, Right bool
}

//############################################################
//Name: randomDirection
//Date: 31.07.22
//Time: 16.08
//############################################################
//Description: Used to choose one of the possible directions.
//############################################################
func (directions Directions) RandomDirection() Position {
	allDirections := []Position{}
	if directions.Up {
		allDirections = append(allDirections, Position{0, -1})
	}
	if directions.Down {
		allDirections = append(allDirections, Position{0, 1})
	}
	if directions.Left {
		allDirections = append(allDirections, Position{-1, 0})
	}
	if directions.Right {
		allDirections = append(allDirections, Position{1, 0})
	}

	numberOfDirections := len(allDirections)

	n := RandomLength(0, numberOfDirections)

	return allDirections[n]
}

//############################################################
//Name: position
//Date: 31.07.22
//Time: 15.51
//############################################################
//Description: Used to describe position as coordinate
//############################################################
type Position struct {
	X, Y int
}

//############################################################
//Name: OffsetFromPosition
//Date: 05.08.22
//Time: 21.59
//############################################################
//Description:
//############################################################
func (position Position) OffsetFromPosition(offset Position) Position {
	position.X += offset.X
	position.Y += offset.Y
	return position
}

//############################################################
//Name: Multiplcation
//Date: 06.08.22
//Time: 14.46
//############################################################
//Dependencies: Position
//Description	: Multiplies a the x and y by a scalar
//Usage				: position.multiplication(3) - Position{1, 2} -> Position{3, 6}
//############################################################
func (position *Position) Multiplication(mutiplicand int) {
	position.X *= mutiplicand
	position.Y *= mutiplicand
}

//############################################################
//Name: Match
//Date: 06.08.22
//Time: 14.50
//############################################################
//Dependencies: Position
//Description	: Checks if two
//Usage				: position.multiplication(3) - Position{1, 2} -> Position{3, 6}
//############################################################
func (position *Position) Match(Pos2 Position) bool {
	return position.X == Pos2.X && position.Y == Pos2.Y
}

//############################################################
//Name: Left, Right, Up, Down
//Date: 07.08.22
//Time: 14.05
//############################################################
//Dependencies: Position
//Description	: returns a directions in a way Position{0,1} is down, Position{1,0} is right
//Usage				: position.Left() - Position{-1, 0}
//############################################################
func (position Position) Left() Position {
	position.X = -1
	position.Y = 0
	return position
}

func (position Position) Right() Position {
	position.X = 1
	position.Y = 0
	return position
}

func (position Position) Up() Position {
	position.X = 0
	position.Y = -1
	return position
}

func (position Position) Down() Position {
	position.X = 0
	position.Y = 1
	return position
}

//############################################################
//Name: IsAtBoundry
//Date: 09.08.22
//Time: 07.12
//############################################################
//Dependencies: Position
//Description	: returns true if position is at
//x||y = 0 or if x||y = their respective widths
//Usage				: position.IsAtBoundry() - true
//############################################################
func (position Position) IsAtBoundry(size Size) bool {
	return position.X == 0 || position.Y == 0 || position.X == size.Width-1 || position.Y == size.Height-1
}

//############################################################
//Name: IsHorizontal
//Date: 09.08.22
//Time: 06.56
//############################################################
//Dependencies: Direction
//Description	: Checks if direction is either Left or Right
//Usage				: direction.IsHorizontal(Direction.Left()) - true
func (direction Directions) IsHorizontal() bool {
	if direction.Down || direction.Up {
		return false
	}
	return true
}

//############################################################
//Name: IsHorizontal
//Date: 09.08.22
//Time: 06.56
//############################################################
//Dependencies: Direction
//Description	: Checks if direction is either Left or Right
//Usage				: direction.IsHorizontal(Direction.Left()) - true
func (directions Directions) directionsToPositions() []Position {
	positions := []Position{}
	if directions.Up {
		positions = append(positions, Position{0, -1})
	}
	if directions.Down {
		positions = append(positions, Position{0, 1})
	}
	if directions.Left {
		positions = append(positions, Position{-1, 0})
	}
	if directions.Right {
		positions = append(positions, Position{1, 0})
	}
	return positions
}

//############################################################
//Name:Pool
//Date: 11.08.22
//Time: 07.39
//############################################################
//Dependencies: Position
//Description	: returns a list of all Tile positions in the same pool as
//the input CRD. A pool is all Tiles with the property IsInspected = false
//that is next to each other without being blocked off by a isInspected = True tile.
//Usage				:
func (maze Maze) Pool(position Position) []Position {
	listOfCheckedCrds := []Position{}
	listOfCheckedCrds = append(listOfCheckedCrds, position)
	n := 0
	for {
		dummyList := maze.getUncheckedNeighbors(listOfCheckedCrds[n], listOfCheckedCrds)
		for i := 0; i < len(dummyList); i++ {
			listOfCheckedCrds = append(listOfCheckedCrds, dummyList[i])
		}
		n++
		if len(listOfCheckedCrds) == n {
			break
		}
	}
	return listOfCheckedCrds
}

//############################################################
//Name:
//Date: 10.08.22
//Time: 07.11
//############################################################
//Dependencies: Maze, Position
//Description	: checks neighbouring coordinates, and returns the coordinates
//that is still not checked, and is not in the Positions list
//Usage				:
func (maze Maze) getUncheckedNeighbors(origoPosition Position, listOfCheckedPositions []Position) []Position {
	var positions []Position

	allPositions := []Position{
		Position{}.Left(),
		Position{}.Right(),
		Position{}.Up(),
		Position{}.Down(),
	}
	for i := 0; i <= 3; i++ {
		newPosition := origoPosition.OffsetFromPosition(allPositions[i])

		if !maze.OutOfBoundsOrInspected(newPosition) && !newPosition.PositionListContainsPosition(listOfCheckedPositions) {
			positions = append(positions, newPosition)
		}
	}
	return positions
}

//############################################################
//Name: PositionListContainsPosition
//Date: 11.08.22
//Time: 07.04
//############################################################
//Dependencies: Maze, Position
//Description	: returns true if input position mach any of the positions in the list
//Usage				: Maze.PositionLostContainsPosition
func (position Position) PositionListContainsPosition(positionList []Position) bool {
	for i := 0; i < len(positionList); i++ {
		if positionList[i].X == position.X && positionList[i].Y == position.Y {
			return true
		}
	}
	return false
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
func (maze Maze) IsEnclosed(position []Position) bool {
	for i := 0; i < len(position); i++ {
		if position[i].IsAtBoundry(maze.Size) {
			return false
		}
	}
	return true
}

//############################################################
//Name:RemoveLine
//Date:12.08.22
//Time:06.55
//############################################################
//Dependencies:Maze Position
//Description : Substitutes every Tile from Start pos to End pos
//with a blank tile
//Usage				: Maze.RemoveLine(POSITION, POSITION)
//############################################################
func (maze *Maze) RemoveLine(start Position, end Position) {
	if start.Y == end.Y { //If it is horizontal
		i := end.X
		for {
			//fmt.Println("The tiles ", i, ", ", start.Y, " is reverted")
			if i != start.X {
				maze.TileMatrix[i][start.Y] = Tile{false, false, false, false, false}
			}
			if i < start.X {
				i++
			} else if i > start.X {
				i--
			} else {
				break
			}
		}
	} else if start.X == end.X { //If it is vertical
		i := end.Y
		for {
			if i != start.Y {
				maze.TileMatrix[start.X][i] = Tile{false, false, false, false, false}
			}
			if i < start.Y {
				i++
			} else if i > start.Y {
				i--
			} else {
				break
			}
		}

	} else {
		print("ERROR")
	}
	//time.Sleep(time.Millisecond * 500)
}

//############################################################
//Name:GetDirectionFromTwoPoints
//Date:12.08.22
//Time:19.00
//############################################################
//Dependencies:Position
//Description :Returns the direction that has to be taken from the
//START(Position) parameter to the END(Position) parameter.
//Usage				: Position.GetDirectionFromTwoPoints(Position, Position)
//############################################################
func (position Position) GetDirectionFromTwoPoints(start Position, end Position) Position {
	x := end.X - start.X
	y := end.Y - start.Y

	if x < 0 {
		return Position{-1, 0}
	}
	if x > 0 {
		return Position{1, 0}
	}
	if y < 0 {
		return Position{0, -1}
	}
	if y > 0 {
		return Position{0, 1}
	}
	return Position{}
}
