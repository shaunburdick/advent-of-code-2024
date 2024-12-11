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

var example1 = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

func Test_day0_part1(t *testing.T) {
	tests := []TestDeclaration{
		{
			name:  "example",
			input: example1,
			want:  142,
			run:   true,
		},
		{
			name:  "actual",
			input: input,
			want:  52974,
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

func Benchmark_day0_part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1(example1)
	}
}

var example2 = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`

func Test_day0_part2(t *testing.T) {
	tests := []TestDeclaration{
		{
			name:  "example",
			input: example2,
			want:  281,
			run:   true,
		},
		{
			name:  "actual",
			input: input,
			want:  53340,
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

func Benchmark_day0_part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part2(example2)
	}
}
