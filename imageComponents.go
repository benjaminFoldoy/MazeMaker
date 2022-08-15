package main

type Shape struct {
	Start []Position
	End   []Position
}

func (shape Shape) Walls(tile Tile) Shape {
	if tile.Top {
		shape.Start = append(shape.Start, Position{0, 0})
		shape.End = append(shape.End, Position{12, 3})
	}
	if tile.Left {
		shape.Start = append(shape.Start, Position{0, 0})
		shape.End = append(shape.End, Position{3, 12})
	}
	if tile.Bottom {
		shape.Start = append(shape.Start, Position{0, 9})
		shape.End = append(shape.End, Position{12, 12})
	}
	if tile.Right {
		shape.Start = append(shape.Start, Position{9, 0})
		shape.End = append(shape.End, Position{12, 12})
	}
	return shape
}

func (shape Shape) Length() int {
	return len(shape.Start)
}

func (shape *Shape) PlaceShape(position Position) {
	for i := 0; i < shape.Length(); i++ {
		shape.Start[i].OffsetFromPosition(position)
	}
}
