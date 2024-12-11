package grid

import (
	"strconv"
)

type Grid struct {
	Data []string
}

// Returns the rune at the given coordinates
func (g Grid) CharAt(c Coords) rune {
	return rune(g.Data[c.y][c.x])
}

// Updates the grid, setting the char at the coords
func (g Grid) SetCharAt(c Coords, char string) {
	g.Data[c.y] = g.Data[c.y][:c.x] + char + g.Data[c.y][c.x+1:]
}

// Checks if a set of Coords is on the grid
func (g Grid) InBounds(c Coords) bool {
	return c.x >= 0 && c.y >= 0 && c.y < len(g.Data) && c.x < len(g.Data[c.y])
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
	x int
	y int
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
		return Coords{x: c.x - 1, y: c.y - 1}
	case North:
		return Coords{x: c.x, y: c.y - 1}
	case NorthEast:
		return Coords{x: c.x + 1, y: c.y - 1}
	case West:
		return Coords{x: c.x - 1, y: c.y}
	case East:
		return Coords{x: c.x + 1, y: c.y}
	case SouthWest:
		return Coords{x: c.x - 1, y: c.y + 1}
	case South:
		return Coords{x: c.x, y: c.y + 1}
	case SouthEast:
		return Coords{x: c.x + 1, y: c.y + 1}
	default:
		panic("Unknown Direction!")
	}
}

func (c Coords) String() string {
	return strconv.Itoa(c.x) + "-" + strconv.Itoa(c.y)
}
