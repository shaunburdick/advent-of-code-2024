package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"math"
	"slices"
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

func part1(input string) int64 {
	parsed := parseInput(input)

	var total int64
	total = 0

	for _, calibration := range parsed {
		testValue, numbers := ParseCalibration(calibration)
		iter := IterateOperators(numbers[0], numbers[1:])
		if slices.Contains(iter, testValue) {
			total += testValue
		}
	}

	return total
}

func part2(input string) int64 {
	parsed := parseInput(input)
	var total int64
	total = 0

	for _, calibration := range parsed {
		testValue, numbers := ParseCalibration(calibration)
		iter := IterateOperatorsWithConcat(numbers[0], numbers[1:])
		if slices.Contains(iter, testValue) {
			total += testValue
		}
	}

	return total
}

func IterateOperators(carry int64, numbers []int64) (results []int64) {
	addResult := ApplyOperator(Add, carry, numbers[0])
	mulResult := ApplyOperator(Multiply, carry, numbers[0])

	// base case
	if len(numbers) == 1 {
		results = append(results, addResult, mulResult)
	} else {
		nextAdd := IterateOperators(addResult, numbers[1:])
		nextMul := IterateOperators(mulResult, numbers[1:])

		results = append(results, nextAdd...)
		results = append(results, nextMul...)
	}

	return results
}

func IterateOperatorsWithConcat(carry int64, numbers []int64) (results []int64) {
	addResult := ApplyOperator(Add, carry, numbers[0])
	mulResult := ApplyOperator(Multiply, carry, numbers[0])
	conResult := ApplyOperator(Concat, carry, numbers[0])

	// base case
	if len(numbers) == 1 {
		results = append(results, addResult, mulResult, conResult)
	} else {
		nextAdd := IterateOperatorsWithConcat(addResult, numbers[1:])
		nextMul := IterateOperatorsWithConcat(mulResult, numbers[1:])
		nextCon := IterateOperatorsWithConcat(conResult, numbers[1:])

		results = append(results, nextAdd...)
		results = append(results, nextMul...)
		results = append(results, nextCon...)
	}

	return results
}

func ParseCalibration(c string) (testValue int64, numbers []int64) {
	calSplit := strings.Split(c, ":")
	tv, tvErr := strconv.Atoi(calSplit[0])
	if tvErr != nil {
		log.Fatalf("Unable to parse test value: %s", calSplit[0])
	}
	testValue = int64(tv)

	for i, num := range strings.Fields(calSplit[1]) {
		numInt, numErr := strconv.Atoi(num)
		if numErr != nil {
			log.Fatalf("Unable to parse number at position  %d: %s", i, num)
		}

		numbers = append(numbers, int64(numInt))
	}

	return testValue, numbers
}

type Operator int

const (
	Add Operator = iota
	Multiply
	Concat
)

func ApplyOperator(op Operator, a int64, b int64) int64 {
	switch op {
	case Add:
		return a + b
	case Multiply:
		return a * b
	case Concat:
		digitsB := int(math.Log10(float64(b))) + 1
		shiftedA := (a * int64(math.Pow10(digitsB)))
		return shiftedA + b
	default:
		panic("Unknown Operator")
	}
}

func ApplyOperators(op Operator, numbers []int64) int64 {
	total := ApplyOperator(op, numbers[0], numbers[1])

	if len(numbers) > 2 {
		for _, number := range numbers[2:] {
			total = ApplyOperator(op, total, number)
		}
	}

	return total
}

func parseInput(input string) (ans []string) {
	return strings.Split(input, "\n")
}
