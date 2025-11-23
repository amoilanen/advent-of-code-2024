package day12

import (
	"testing"
)

func TestPart1WithExamples(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name: "Small example",
			input: `AAAA
BBCD
BBCC
EEEC`,
			expected: 140,
		},
		{
			name: "Example with nested regions",
			input: `OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`,
			expected: 772,
		},
		{
			name:     "Large example",
			input:    ExampleInput,
			expected: 1930,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			grid := Parse(tt.input)
			result := Part1(grid)
			if result != tt.expected {
				t.Errorf("Part1() = %d, want %d", result, tt.expected)
			}
		})
	}
}

func TestPart2WithExamples(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name: "Small example - 4 simple regions",
			input: `AAAA
BBCD
BBCC
EEEC`,
			expected: 80, // A=16, B=16, C=32, D=4, E=12
		},
		{
			name: "Example with nested regions",
			input: `OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`,
			expected: 436,
		},
		{
			name: "E-shaped region",
			input: `EEEEE
EXXXX
EEEEE
EXXXX
EEEEE`,
			expected: 236, // E region: 17 area * 12 sides = 204, plus X regions
		},
		{
			name: "Nested region with inner holes",
			input: `AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA`,
			expected: 368, // A has outer+inner sides, B regions are simple
		},
		{
			name:     "Large example",
			input:    ExampleInput,
			expected: 1206,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			grid := Parse(tt.input)
			result := Part2(grid)
			if result != tt.expected {
				t.Errorf("Part2() = %d, want %d", result, tt.expected)
			}
		})
	}
}
