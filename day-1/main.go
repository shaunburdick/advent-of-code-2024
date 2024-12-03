package main

import (
	"flag"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	file "github.com/shaunburdick/advent-of-code-2024/lib"
)

var input string

func init() {
	// do this in init (not main) so test file has same input
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
	totalDistance := 0
	lists := struct {
		left  []int
		right []int
	}{
		left:  []int{},
		right: []int{},
	}

	// convert to separate numbers list
	for _, row := range parsed {
		words := strings.Fields(row)
		left, leftErr := strconv.Atoi(words[0])
		right, rightErr := strconv.Atoi(words[1])

		if leftErr != nil || rightErr != nil {
			log.Fatalf("Error parsing number: %s -> %s, %s", row, leftErr, rightErr)
		}

		lists.left = append(lists.left, left)
		lists.right = append(lists.right, right)
	}

	// sort them
	sort.Ints(lists.left)
	sort.Ints(lists.right)

	// calculate distances
	for rowNum := range lists.left {
		distance := lists.left[rowNum] - lists.right[rowNum]
		if distance < 0 {
			distance = -distance
		}
		totalDistance += distance
	}

	return totalDistance
}

func part2(input string) int {
	parsed := parseInput(input)
	similarityScore := 0
	lists := struct {
		left  []int
		right map[int]int
	}{
		left:  []int{},
		right: make(map[int]int),
	}

	// convert to separate numbers list
	for _, row := range parsed {
		words := strings.Fields(row)
		left, leftErr := strconv.Atoi(words[0])
		right, rightErr := strconv.Atoi(words[1])

		if leftErr != nil || rightErr != nil {
			log.Fatalf("Error parsing number: %s -> %s, %s", row, leftErr, rightErr)
		}

		lists.left = append(lists.left, left)
		if _, ok := lists.right[right]; !ok {
			lists.right[right] = 1
		} else {
			lists.right[right] += 1
		}
	}

	// calculate distances
	for _, leftVal := range lists.left {
		if rightVal, ok := lists.right[leftVal]; ok {
			similarityScore += leftVal * rightVal
		}
	}

	return similarityScore
}

func parseInput(input string) (ans []string) {
	return strings.Split(input, "\n")
}
