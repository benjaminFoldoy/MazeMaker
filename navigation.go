package main

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
//Name:
//Date: 09.08.22
//Time: 14.01
//############################################################
//Dependencies: Position
//Description	: finds the size of a pool
//Usage				:
func (maze Maze) PoolSize(position Position) int {
	listOfCheckedCrds := []Position{}

	listOfCheckedCrds = append(listOfCheckedCrds, position)

	for {

	}
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
func (maze Maze) checkNSEWCrds(positions []Position) []Position {

	allPositions := []Position{
		Position{}.Left(),
		Position{}.Right(),
		Position{}.Up(),
		Position{}.Down(),
	}

	for i := 0; i <= 3; i++ {
		if !maze.OutOfBoundsOrInspected(positions[len(positions)-1].OffsetFromPosition(allPositions[i])) {
			positions = append(positions, positions[len(positions)-1].OffsetFromPosition(allPositions[i]))
		}
	}

	return positions
}
