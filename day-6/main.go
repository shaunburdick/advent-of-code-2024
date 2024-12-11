package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"strconv"
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

func part1(input string) int {
	parsed := parseInput(input)
	grid := Grid{parsed}
	guard := NewGuard(grid)

	locations := make(map[string]struct{})
	// include guard's starting position
	locations[guard.Location.MakeId()] = struct{}{}

	guard.Patrol()

	for location := range guard.MovementUpdates {
		locations[location.MakeId()] = struct{}{}
	}

	return len(locations)
}

func part2(input string) int {
	parsed := parseInput(input)
	_ = parsed

	return 0
}

type Guard struct {
	Grid            Grid
	Location        Coords
	Direction       Direction
	MovementUpdates chan Coords
	Stop            chan struct{}
}

func NewGuard(grid Grid) *Guard {
	return &Guard{
		Grid:            grid,
		Location:        grid.FindChar("^"),
		Direction:       North,
		MovementUpdates: make(chan Coords),
		Stop:            make(chan struct{}),
	}
}

func (g *Guard) Move(newCoords Coords) {
	g.Location = newCoords
	g.MovementUpdates <- newCoords
}

func (g *Guard) Patrol() {
	go func() {
		for {
			select {
			case <-g.Stop:
				return
			default:
				// get the next step
				nextCoords := g.Location.Direction(g.Direction)
				// if we are still on the grid
				if g.Grid.InBounds(nextCoords) {
					nextStep := g.Grid.CharAt(nextCoords)
					// if something is in the way
					if nextStep == '#' {
						// turn right
						g.Direction = g.Direction.TurnRight()
					} else if nextStep == '.' || nextStep == 'X' {
						// take next step
						g.Move(nextCoords)
					} else if nextStep == '^' {
						if g.Direction == North {
							log.Fatal("WE ARE IN AN INFINITE LOOP?!")
						} else {
							g.Move(nextCoords)
						}
					} else {
						log.Fatalf("WHAT DID I STEP IN?! %s", string(nextStep))
					}
				} else {
					// Stop Patrolling
					close(g.MovementUpdates)
					close(g.Stop)
				}
			}
		}
	}()
}

type Grid struct {
	Data []string
}

func (g Grid) CharAt(c Coords) rune {
	return rune(g.Data[c.y][c.x])
}

func (g Grid) SetCharAt(c Coords, char string) {
	g.Data[c.y] = g.Data[c.y][:c.x] + char + g.Data[c.y][c.x+1:]
}

func (g Grid) InBounds(c Coords) bool {
	return c.x >= 0 && c.y >= 0 && c.y < len(g.Data) && c.x < len(g.Data[c.y])
}

func (g Grid) FindChar(char string) Coords {
	coords := Coords{-1, -1}

	for y, row := range g.Data {
		index := strings.Index(row, char)
		if index != -1 {
			coords.x = index
			coords.y = y
			return coords
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
	North Direction = iota
	West
	East
	South
)

func (d Direction) TurnRight() Direction {
	switch d {
	case North:
		return East
	case East:
		return South
	case South:
		return West
	case West:
		return North
	}

	panic("Unknown Direction!")
}

func (c Coords) Direction(d Direction) Coords {
	switch d {
	case North:
		return Coords{x: c.x, y: c.y - 1}
	case West:
		return Coords{x: c.x - 1, y: c.y}
	case East:
		return Coords{x: c.x + 1, y: c.y}
	case South:
		return Coords{x: c.x, y: c.y + 1}
	}

	panic("Unknown Direction!")
}

func (c Coords) MakeId() string {
	return strconv.Itoa(c.x) + "-" + strconv.Itoa(c.y)
}

func parseInput(input string) (ans []string) {
	return strings.Split(input, "\n")
}
