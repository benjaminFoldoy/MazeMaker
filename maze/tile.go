package maze

//############################################################
//Name: Tile
//Date: 31.07.22
//Time: 15.54
//############################################################
//Description: Used to describe each cell in a maze.
//Each cell will be able to have walls on the right, left, top and/or bottom.
//CASE1: left, right, top, bottom = null or false:
//The tile will have no walls.
//CASE2: left, right = null and top, bottom = true:
//The tile will have walls on the top and bottom, and will have no walls on the right and left
//############################################################
type Tile struct {
	Left, Right, Top, Bottom bool
	IsInspected              bool
}

//############################################################
//Name: TilePrint
//Date: 04.08.22
//Time: 10.52
//############################################################
//Dependencies: Tile
//Description : Prints a representation of how a given Tile looks.
//Usage				: tile.TilePrint()
//############################################################
func (tile Tile) TilePrint() {
	array := [9][9]bool{}

	//Set sides of box to either dotted or solid
	for i := 0; i < 9; i++ {
		if tile.Top || i%2 == 0 {
			array[0][i] = true
		}

		if tile.Bottom || i%2 == 0 {
			array[8][i] = true
		}

		if tile.Left || i%2 == 0 {
			array[i][0] = true
		}

		if tile.Right || i%2 == 0 {
			array[i][8] = true
		}
	}

	//printing box
	print("The Tile looks like this:\n#####################################\n\n")
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if array[i][j] {
				print("\033[32m██\033[0m")
			} else {
				print("  ")
			}
		}
		println()
	}
}
