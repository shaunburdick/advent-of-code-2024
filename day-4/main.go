package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"strings"

	file "github.com/shaunburdick/advent-of-code-2024/lib/file"
)

var input string

func init() {
	// do this in init (not main) so test file has same input
	inputFile, err := file.LoadRelativeFile("input.txt")
	if err != nil {
		log.Println(err)
	}

	input = strings.TrimRight(inputFile, "\n")
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(input)
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		fmt.Println("Output:", ans)
	}
}

const XMAS = "XMAS"

func part1(input string) int {
	parsed := parseInput(input)
	grid := Grid{parsed}
	foundWords := 0
	// start traversing the grid
	for y, row := range parsed {
		for x := range row {
			coord := Coords{x, y}
			foundWords += xmasSearch(grid, "", coord, AllDirections)
		}
	}

	return foundWords
}

func part2(input string) int {
	parsed := parseInput(input)
	grid := Grid{parsed}
	foundMas := 0
	// start traversing the grid
	for y, row := range grid.Data {
		if y == 0 || y == len(grid.Data)-1 {
			// skip edges
			continue
		}

		for x, char := range row {
			if x == 0 || x == len(row)-1 {
				// skip edges
				continue
			}
			if char == 'A' {
				coord := Coords{x, y}
				wing1 := string(grid.CharAt(coord.Direction(NorthWest))) + "A" + string(grid.CharAt(coord.Direction(SouthEast)))
				wing2 := string(grid.CharAt(coord.Direction(NorthEast))) + "A" + string(grid.CharAt(coord.Direction(SouthWest)))

				if (wing1 == "MAS" || wing1 == "SAM") && (wing2 == "MAS" || wing2 == "SAM") {
					foundMas += 1
				}
			}
		}
	}

	return foundMas
}

type Grid struct {
	Data []string
}

func (g Grid) CharAt(c Coords) rune {
	return rune(g.Data[c.y][c.x])
}

func (g Grid) InBounds(c Coords) bool {
	return c.x >= 0 && c.y >= 0 && c.y < len(g.Data) && c.x < len(g.Data[c.y])
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

// Generate a list of valid coordinates from the existing one
// Will only return valid positive values
func (c Coords) Directions() []Coords {
	directions := []Coords{}

	for _, dir := range AllDirections {
		newCoord := c.Direction(dir)
		// check for bounds
		if newCoord.x >= 0 && newCoord.y >= 0 {
			directions = append(directions, newCoord)
		}
	}

	return directions
}

func (c Coords) Direction(d Direction) Coords {
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
		return c
	}
}

func xmasSearch(grid Grid, currentWord string, coords Coords, directions []Direction) int {
	foundWords := 0
	newWord := currentWord + string(grid.CharAt(coords))
	if newWord == XMAS {
		// we match!
		return 1
	} else if len(newWord) >= len(XMAS) || !strings.HasPrefix(XMAS, newWord) {
		// we don't match!
		return 0
	} else {
		// we partially match, keep looking!
		for _, direction := range directions {
			newCoord := coords.Direction(direction)
			// check for bounds
			if grid.InBounds(newCoord) {
				foundWords += xmasSearch(grid, newWord, newCoord, []Direction{direction})
			}
		}
	}

	return foundWords
}

func parseInput(input string) (ans []string) {
	return strings.Split(input, "\n")
}
