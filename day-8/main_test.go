package main

import (
	"testing"

	file "github.com/shaunburdick/advent-of-code-2024/lib/file"
)

type TestDeclaration struct {
	name  string
	input string
	want  int
	run   bool
}

var example1 = `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`

func Test_day8_part1(t *testing.T) {
	tests := []TestDeclaration{
		{
			name:  "example",
			input: example1,
			want:  14,
			run:   true,
		},
		{
			name:  "actual",
			input: input,
			want:  327,
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

func Benchmark_day8_part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1(example1)
	}
}

var example2 = `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`

func Test_day8_part2(t *testing.T) {
	tests := []TestDeclaration{
		{
			name:  "example",
			input: example2,
			want:  34,
			run:   true,
		},
		{
			name:  "actual",
			input: input,
			want:  1233,
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

func Benchmark_day8_part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part2(example2)
	}
}
