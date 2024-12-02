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
	safeCount := 0
	for _, row := range parsed {
		if safeReport(row) {
			safeCount++
		}
	}

	return safeCount
}

func part2(input string) int {
	parsed := parseInput(input)
	safeCount := 0
	for _, row := range parsed {
		if safeReport(row) {
			safeCount++
		} else {
			// find an iteration that does work, removing one number at a time
			for i := 0; i < len(row); i++ {
				// make a new slice and append the items around i
				// I hate this...
				newRow := make([]int, 0)
				newRow = append(newRow, row[:i]...)
				newRow = append(newRow, row[i+1:]...)
				if safeReport(newRow) {
					safeCount++
					break
				}
			}
		}
	}

	return safeCount
}

func absDiff(a int, b int) int {
	diff := a - b
	if diff < 0 {
		diff = -diff
	}

	return diff
}

func safeRange(diff int) bool {
	return diff < 4 && diff > 0
}

func checkLevel(increasing bool, a int, b int) bool {
	return (increasing == (a < b)) && safeRange(absDiff(a, b))
}

func safeReport(report []int) bool {
	safe := true
	increasing := report[0] < report[1]
	if !safeRange(absDiff(report[0], report[1])) {
		safe = false
	}

	for i := 2; i < len(report); i++ {
		if !checkLevel(increasing, report[i-1], report[i]) {
			safe = false
			break
		}
	}

	return safe
}

func parseInput(input string) (ans [][]int) {
	ans = [][]int{}
	for _, row := range strings.Split(input, "\n") {
		words := strings.Fields(row)
		numbers := make([]int, len(words))
		for i, num := range words {
			parsedNumber, err := strconv.Atoi(num)
			if err != nil {
				log.Fatalf("Unable to parse number: %s", num)
			}
			numbers[i] = parsedNumber
		}
		ans = append(ans, numbers)
	}

	return ans
}
