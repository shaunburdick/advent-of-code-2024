package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/shaunburdick/advent-of-code-2024/lib/file"
	"github.com/shaunburdick/advent-of-code-2024/lib/grid"
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
	antinodes := make(map[string]struct{})
	antennaMap := grid.Grid{Data: parsed}
	nodes := UniqueNodes(antennaMap)

	for _, coords := range nodes {
		// if there are more than one instance of the node
		if len(coords) > 1 {
			for i, coordA := range coords {
				// apply to every other coord
				for _, coordB := range coords[i+1:] {
					antiNodeA := grid.Coords{X: 2*coordA.X - coordB.X, Y: 2*coordA.Y - coordB.Y}
					antiNodeB := grid.Coords{X: 2*coordB.X - coordA.X, Y: 2*coordB.Y - coordA.Y}

					if antennaMap.InBounds(antiNodeA) {
						antinodes[antiNodeA.String()] = struct{}{}
					}

					if antennaMap.InBounds(antiNodeB) {
						antinodes[antiNodeB.String()] = struct{}{}
					}
				}
			}
		}
	}

	return len(antinodes)
}

func UniqueNodes(g grid.Grid) map[rune][]grid.Coords {
	nodes := make(map[rune][]grid.Coords)

	for y, row := range g.Data {
		for x, char := range row {
			if char != '.' {
				if _, found := nodes[char]; !found {
					nodes[char] = []grid.Coords{}
				}

				nodes[char] = append(nodes[char], grid.Coords{X: x, Y: y})
			}
		}
	}

	return nodes
}

func part2(input string) int {
	parsed := parseInput(input)
	_ = parsed

	return 0
}

func parseInput(input string) (ans []string) {
	return strings.Split(input, "\n")
}
