package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"

	file "github.com/shaunburdick/advent-of-code-2024/lib"
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
	total := 0
	for _, line := range parsed {
		first := -1
		last := -1
		for i := 0; i < len(line); i++ {
			if line[i] >= '0' && line[i] <= '9' {
				num, err := strconv.Atoi(line[i : i+1])
				if err != nil {
					log.Fatalf("unable to parse into number: %s", err)
				}
				if first == -1 {
					first = num
				}
				last = num
			}
		}

		total += first*10 + last
	}

	return total
}

var prefixes map[string]int = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"zero":  0,
}

func part2(input string) int {
	parsed := parseInput(input)
	total := 0

	for _, line := range parsed {
		first := -1
		last := -1
		for i := 0; i < len(line); i++ {
			if line[i] >= '0' && line[i] <= '9' {
				num, err := strconv.Atoi(line[i : i+1])
				if err != nil {
					log.Fatalf("unable to parse into number: %s", err)
				}
				if first == -1 {
					first = num
				}
				last = num
			} else {
				// the number names range from 3 to 5 in length so we just need
				// to check those three lengths in the map
				for j := 3; j < 6; j++ {
					end := i + j
					// If we are over the end of the string, no point in checking
					if end > len(line) {
						break
					}
					val, ok := prefixes[line[i:end]]
					if ok {
						if first == -1 {
							first = val
						}
						last = val

						// jump forward to the last letter of the prefix
						i += j - 2
						break
					}
				}
			}
		}

		total += first*10 + last
	}

	return total
}

func parseInput(input string) (ans []string) {
	return strings.Split(input, "\n")
}
