package day04

import "testing"

func TestPart1Example(t *testing.T) {
	grid := Parse(ExampleInput)
	result := Part1(grid)
	expected := 18
	if result != expected {
		t.Errorf("Part1(ExampleInput) = %d; want %d", result, expected)
	}
}

func TestPart2Example(t *testing.T) {
	grid := Parse(ExampleInput)
	result := Part2(grid)
	expected := 9
	if result != expected {
		t.Errorf("Part2(ExampleInput) = %d; want %d", result, expected)
	}
}

func TestParse(t *testing.T) {
	grid := Parse(ExampleInput)
	if len(grid) != 10 {
		t.Errorf("Expected 10 rows, got %d", len(grid))
	}
	if len(grid[0]) != 10 {
		t.Errorf("Expected 10 columns, got %d", len(grid[0]))
	}
}

func TestSearchWord(t *testing.T) {
	grid := Parse(ExampleInput)

	tests := []struct {
		name     string
		row      int
		col      int
		dir      Direction
		word     string
		expected bool
	}{
		{
			name:     "XMAS horizontal at (0,5)",
			row:      0,
			col:      5,
			dir:      Direction{0, 1},
			word:     "XMAS",
			expected: true,
		},
		{
			name:     "XMAS not found",
			row:      0,
			col:      0,
			dir:      Direction{0, 1},
			word:     "XMAS",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := grid.hasWordAtPositionDirection(tt.row, tt.col, tt.dir, tt.word)
			if result != tt.expected {
				t.Errorf("searchWord(%d, %d, %v, %q) = %v; want %v",
					tt.row, tt.col, tt.dir, tt.word, result, tt.expected)
			}
		})
	}
}
