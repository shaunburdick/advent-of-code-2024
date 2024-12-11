package grid

import (
	"strconv"
)

type Grid struct {
	Data []string
}

// Returns the rune at the given coordinates
func (g Grid) CharAt(c Coords) rune {
	return rune(g.Data[c.Y][c.X])
}

// Updates the grid, setting the char at the coords
func (g Grid) SetCharAt(c Coords, char string) {
	g.Data[c.Y] = g.Data[c.Y][:c.X] + char + g.Data[c.Y][c.X+1:]
}

// Checks if a set of Coords is on the grid
func (g Grid) InBounds(c Coords) bool {
	return c.X >= 0 && c.Y >= 0 && c.Y < len(g.Data) && c.X < len(g.Data[c.Y])
}

// Returns a list of Coords where the char is found
func (g Grid) FindChar(char rune) []Coords {
	coords := []Coords{}

	for y, row := range g.Data {
		for x, letter := range row {
			if letter == char {
				coords = append(coords, Coords{x, y})
			}
		}
	}

	return coords
}

type Coords struct {
	X int
	Y int
}

type Direction int

const (
	NorthWest Direction = iota
	North
	NorthEast
	West
	East
	SouthWest
	South
	SouthEast
)

var AllDirections []Direction = []Direction{
	NorthWest,
	North,
	NorthEast,
	West,
	East,
	SouthWest,
	South,
	SouthEast,
}

// Gets the next Coords in the given direction
func (c Coords) Next(d Direction) Coords {
	switch d {
	case NorthWest:
		return Coords{X: c.X - 1, Y: c.Y - 1}
	case North:
		return Coords{X: c.X, Y: c.Y - 1}
	case NorthEast:
		return Coords{X: c.X + 1, Y: c.Y - 1}
	case West:
		return Coords{X: c.X - 1, Y: c.Y}
	case East:
		return Coords{X: c.X + 1, Y: c.Y}
	case SouthWest:
		return Coords{X: c.X - 1, Y: c.Y + 1}
	case South:
		return Coords{X: c.X, Y: c.Y + 1}
	case SouthEast:
		return Coords{X: c.X + 1, Y: c.Y + 1}
	default:
		panic("Unknown Direction!")
	}
}

func (c Coords) String() string {
	return strconv.Itoa(c.X) + "-" + strconv.Itoa(c.Y)
}
