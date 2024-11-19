package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
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

func part2(input string) int {
	parsed := parseInput(input)
	total := 0

	prefixes := map[string]int{
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
				for prefix, val := range prefixes {
					if checkPrefix(line[i:], prefix) {
						if first == -1 {
							first = val
						}
						last = val

						// jump forward to the last letter of the prefix
						i += len(prefix) - 2
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

func checkPrefix(str string, prefix string) bool {
	if len(str) < len(prefix) {
		return false
	}
	return str[:len(prefix)] == prefix
}
