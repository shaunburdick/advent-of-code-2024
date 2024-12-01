package main

import (
	"testing"
)

var example1 = `3   4
4   3
2   5
1   3
3   9
3   3`

func Test_day1_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example1,
			want:  11,
		},
		{
			name:  "actual",
			input: input,
			want:  2367773,
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

func Benchmark_day1_part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1(input)
	}
}

var example2 = `3   4
4   3
2   5
1   3
3   9
3   3`

func Test_day1_part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example2,
			want:  31,
		},
		{
			name:  "actual",
			input: input,
			want:  21271939,
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

func Benchmark_day1_part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part2(input)
	}
}
