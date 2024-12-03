package main

import (
	"testing"

	file "github.com/shaunburdick/advent-of-code-2024/lib"
)

type TestDeclaration struct {
	name  string
	input string
	want  int
	run   bool
}

var example1 = `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`

func Test_day3_part1(t *testing.T) {
	tests := []TestDeclaration{
		{
			name:  "example",
			input: example1,
			want:  161,
			run:   true,
		},
		{
			name:  "actual",
			input: input,
			want:  182780583,
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

func Benchmark_day3_part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1(example1)
	}
}

var example2 = `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`

func Test_day3_part2(t *testing.T) {
	tests := []TestDeclaration{
		{
			name:  "example",
			input: example2,
			want:  48,
			run:   true,
		},
		{
			name:  "actual",
			input: input,
			want:  90772405,
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

func Benchmark_day3_part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part2(example2)
	}
}
