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

var example1 = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

func Test_day2_part1(t *testing.T) {
	tests := []TestDeclaration{
		{
			name:  "example",
			input: example1,
			want:  2,
			run:   true,
		},
		{
			name:  "actual",
			input: input,
			want:  356,
			run:   file.ExistsRelativeFile("input.txt"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_day2_part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1(example1)
	}
}

var example2 = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

func Test_day2_part2(t *testing.T) {
	tests := []TestDeclaration{
		{
			name:  "example",
			input: example2,
			want:  4,
			run:   true,
		},
		{
			name:  "actual",
			input: input,
			want:  413,
			run:   file.ExistsRelativeFile("input.txt"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_day2_outliers(t *testing.T) {
	t.Run("Bad first number", func(t *testing.T) {
		if part2("5 6 5 4 3 2") != 1 {
			t.Errorf("Bad first number failed")
		}
	})

	t.Run("Bad second number", func(t *testing.T) {
		if part2("1 3 2 4 5") != 1 {
			t.Errorf("Bad second number failed")
		}
	})

	t.Run("Bad last number", func(t *testing.T) {
		if part2("1 2 3 4 3") != 1 {
			t.Errorf("Bad last number failed")
		}
	})

	t.Run("Bad example number", func(t *testing.T) {
		if part2("75 78 79 80 84 81") != 1 {
			t.Errorf("Bad example failed")
		}
	})
}

func Benchmark_day2_part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part2(example2)
	}
}
