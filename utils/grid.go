package utils

type Grid struct {
	Values        [][]byte
	Length, Width int
}

func GridFromBytes(data []byte) Grid {
	var grid Grid
	var row []byte
	for _, b := range data {
		if b == '\n' {
			grid.Values = append(grid.Values, row)
			row = nil
		} else {
			row = append(row, b)
		}
	}
	if row != nil {
		grid.Values = append(grid.Values, row)
	}
	grid.Length = len(grid.Values)
	grid.Width = len(grid.Values[0])
	return grid
}

type Coord struct {
	Y, X int
}

func (g Grid) GetEntry(c Coord) byte {
	if c.Y < 0 || c.Y >= g.Length || c.X < 0 || c.X >= g.Width {
		return 255
	}
	return g.Values[c.Y][c.X]
}

var Directions = []Coord{
	{0, 1},
	{-1, 0},
	{0, -1},
	{1, 0},
}

func AddCoords(a, b Coord) Coord {
	return Coord{Y: a.Y + b.Y, X: a.X + b.X}
}
