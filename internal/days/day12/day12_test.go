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
