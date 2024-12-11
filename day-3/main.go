package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"
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
	sum := 0
	for _, command := range parsed {
		sum += command.Execute()
	}

	return sum
}

func part2(input string) int {
	parsed := parseInput(input)
	sum := 0
	enabled := true
	for _, command := range parsed {
		enabled = command.Enabled(enabled)

		if enabled {
			sum += command.Execute()
		}
	}

	return sum
}

type Command struct {
	name string
	args []string
}

func (c Command) Execute() int {
	switch c.name {
	case "mul":
		product, _ := strconv.Atoi(c.args[0])
		for _, multiplier := range c.args[1:] {
			mVal, _ := strconv.Atoi(multiplier)
			product *= mVal
		}

		return product
	}

	return 0
}

func (c Command) Enabled(current bool) bool {
	if c.name == "do" {
		return true
	} else if c.name == "don't" {
		return false
	} else {
		return current
	}
}

func parseInput(input string) []Command {
	commands := []Command{}

	commandPattern := regexp.MustCompile(`(do(?:n't)?)\(\)|(mul)\((\d+)\,(\d+)\)`)
	matches := commandPattern.FindAllStringSubmatch(input, -1)
	for _, match := range matches {
		if match[2] == "mul" {
			commands = append(commands, Command{match[2], []string{match[3], match[4]}})
		} else {
			commands = append(commands, Command{match[1], []string{}})
		}
	}

	return commands
}
