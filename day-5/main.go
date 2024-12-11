package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"slices"
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
	rules, updates := parseInput(input)

	middlePages := 0
	por := PageOrderingRules{rules}
	for _, update := range updates {
		if por.ValidateUpdate(update) {
			// truncates toward zero
			middleIndex := len(update) / 2
			pageNum, err := strconv.Atoi(update[middleIndex])
			if err != nil {
				log.Fatalf("Unable to parse %s", update[middleIndex])
			}

			middlePages += pageNum
		}
	}

	return middlePages
}

func part2(input string) int {
	rules, updates := parseInput(input)

	middlePages := 0
	por := PageOrderingRules{rules}
	for _, update := range updates {
		if !por.ValidateUpdate(update) {
			newUpdate := por.FixUpdate(update)

			// truncates toward zero
			middleIndex := len(newUpdate) / 2
			pageNum, err := strconv.Atoi(newUpdate[middleIndex])
			if err != nil {
				log.Fatalf("Unable to parse %s", newUpdate[middleIndex])
			}

			middlePages += pageNum
		}
	}

	return middlePages
}

type PageOrderRule struct {
	Before string
	After  string
}

type PageOrderingRules struct {
	Rules []PageOrderRule
}

func (p PageOrderingRules) RuleFindIndex(rule PageOrderRule, pages []string) (int, int) {
	beforeIndex := -1
	afterIndex := -1
	for i, page := range pages {
		if page == rule.Before {
			beforeIndex = i
		} else if page == rule.After {
			afterIndex = i
		}

		if beforeIndex != -1 && afterIndex != -1 {
			// we found both, stop looking
			break
		}
	}

	return beforeIndex, afterIndex
}

func (p PageOrderingRules) ValidateUpdate(pages []string) bool {
	for _, rule := range p.Rules {
		beforeIndex, afterIndex := p.RuleFindIndex(rule, pages)

		// if we found both values, and they break the rule
		// we can stop
		if beforeIndex != -1 && afterIndex != -1 && beforeIndex > afterIndex {
			return false
		}
	}

	return true
}

func (p PageOrderingRules) FixUpdate(pages []string) []string {
	slices.SortStableFunc(pages, func(a, b string) int {
		for _, rule := range p.Rules {
			if rule.Before == a && rule.After == b {
				return -1
			} else if rule.After == a && rule.Before == b {
				return 1
			}
		}

		return 0
	})

	return pages
}

func parseInput(input string) ([]PageOrderRule, [][]string) {
	rows := strings.Split(input, "\n")
	rules := []PageOrderRule{}
	updates := [][]string{}

	for _, row := range rows {
		if strings.Contains(row, "|") {
			items := strings.Split(row, "|")
			rules = append(rules, PageOrderRule{items[0], items[1]})
		} else if strings.Contains(row, ",") {
			pages := strings.Split(row, ",")
			updates = append(updates, pages)
		}
	}

	return rules, updates
}
