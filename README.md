# Advent of Code - 2024

An attempt at [Advent of Code](https://adventofcode.com/2024) using [Go](https://go.dev)!

![[Workflow Status](https://github.com/shaunburdick/advent-of-code-2024/actions/workflows/golang.yml)](https://github.com/shaunburdick/advent-of-code-2024/actions/workflows/golang.yml/badge.svg)

Each day will be setup as a separate folder.

-   [Day 0](/day-0/) - Trebuchet?! (2023 Day 1)

## Environment Setup

To setup your environment:

1. [Install](https://go.dev/dl/) Golang (or `brew install golang`)
2. Install [Golang CI Lint](https://github.com/golangci/golangci-lint): `brew install golangci-lint`
3. Install [Just](https://github.com/casey/just): `brew install just`
4. Install [gotestfmt](https://github.com/GoTestTools/gotestfmt): `go install github.com/gotesttools/gotestfmt/v2/cmd/gotestfmt@latest`
    - Make sure your GOPATH/bin is in your PATH: `export PATH=$PATH:$(go env GOPATH)/bin`
5. Clone this repository
6. Create a new day from template: `just create X` (where X is the day number)
7. Be Merry! 🎄

## Testing

To run tests for an individual day X, run `just test X` or to run all tests run `just test-all`

## Linting

To run linting for an individual day X, run `just lint X` or to run all linting run `just lint-all`

## Running

To run a solution, run `just run <day> <part>` where day is the day number and part is the part number. Example: `just run 1 1` to run day 1, part 1

## Benchmarks

To run benchmarks for an individual day X, run `just bench X` or to run all benchmarks run `just bench-all`
