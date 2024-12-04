package main

import (
	"reflect"
	"testing"

	file "github.com/shaunburdick/advent-of-code-2024/lib"
)

type TestDeclaration struct {
	name  string
	input string
	want  int
	run   bool
}

var example1 = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

func Test_day4_part1(t *testing.T) {
	tests := []TestDeclaration{
		{
			name:  "example",
			input: example1,
			want:  18,
			run:   true,
		},
		{
			name:  "actual",
			input: input,
			want:  2358,
			run:   file.ExistsRelativeFile("input.txt"),
		},
	}
	for _, tt := range tests {
		if tt.run {
			t.Run(tt.name, func(t *testing.T) {
				if got := part1(tt.input); got != tt.want {
					t.Errorf("part1() = %v, want %v", got, tt.want)
				}
			})
		}
	}
}

func Benchmark_day4_part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1(example1)
	}
}

var example2 = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

func Test_day4_part2(t *testing.T) {
	tests := []TestDeclaration{
		{
			name:  "example",
			input: example2,
			want:  9,
			run:   true,
		},
		{
			name:  "actual",
			input: input,
			want:  1737,
			run:   file.ExistsRelativeFile("input.txt"),
		},
	}
	for _, tt := range tests {
		if tt.run {
			t.Run(tt.name, func(t *testing.T) {
				if got := part2(tt.input); got != tt.want {
					t.Errorf("part2() = %v, want %v", got, tt.want)
				}
			})
		}
	}
}

func Benchmark_day4_part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part2(example2)
	}
}

func Test_Coords_Directions(t *testing.T) {
	t.Run("0,0", func(t *testing.T) {
		expected := []Coords{
			{
				x: 1,
				y: 0,
			},
			{
				x: 0,
				y: 1,
			},
			{
				x: 1,
				y: 1,
			},
		}
		result := Coords{0, 0}.Directions()
		if !reflect.DeepEqual(expected, result) {
			t.Error(result)
		}
	})

	t.Run("5,3", func(t *testing.T) {
		expected := []Coords{
			{
				x: 4,
				y: 2,
			},
			{
				x: 5,
				y: 2,
			},
			{
				x: 6,
				y: 2,
			},
			{
				x: 4,
				y: 3,
			},
			{
				x: 6,
				y: 3,
			},
			{
				x: 4,
				y: 4,
			},
			{
				x: 5,
				y: 4,
			},
			{
				x: 6,
				y: 4,
			},
		}
		result := Coords{5, 3}.Directions()
		if !reflect.DeepEqual(expected, result) {
			t.Error(result)
		}
	})
}
