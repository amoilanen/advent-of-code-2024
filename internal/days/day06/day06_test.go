package day06

import "testing"

func TestPart1Example(t *testing.T) {
	grid, guard := Parse(ExampleInput)
	result := Part1(grid, guard)
	expected := 41
	if result != expected {
		t.Errorf("Part1(ExampleInput) = %d; want %d", result, expected)
	}
}

func TestPart2Example(t *testing.T) {
	grid, guard := Parse(ExampleInput)
	result := Part2(grid, guard)
	expected := 6
	if result != expected {
		t.Errorf("Part2(ExampleInput) = %d; want %d", result, expected)
	}
}

func TestParse(t *testing.T) {
	grid, guard := Parse(ExampleInput)

	if grid.rows != 10 {
		t.Errorf("Expected 10 rows, got %d", grid.rows)
	}

	if grid.cols != 10 {
		t.Errorf("Expected 10 columns, got %d", grid.cols)
	}

	if guard == nil {
		t.Fatal("Expected guard to be found")
	}

	if guard.pos.Row != 6 || guard.pos.Col != 4 {
		t.Errorf("Expected guard at (6, 4), got (%d, %d)", guard.pos.Row, guard.pos.Col)
	}

	if guard.dir != Up {
		t.Errorf("Expected guard facing Up, got %v", guard.dir)
	}

	// Check some obstacles
	if !grid.hasObstacle(Position{Row: 0, Col: 4}) {
		t.Error("Expected obstacle at (0, 4)")
	}
}

func TestTurnRight(t *testing.T) {
	guard := &Guard{pos: Position{Row: 0, Col: 0}, dir: Up}

	guard.turnRight()
	if guard.dir != Right {
		t.Errorf("After turning right from Up, expected Right, got %v", guard.dir)
	}

	guard.turnRight()
	if guard.dir != Down {
		t.Errorf("After turning right from Right, expected Down, got %v", guard.dir)
	}

	guard.turnRight()
	if guard.dir != Left {
		t.Errorf("After turning right from Down, expected Left, got %v", guard.dir)
	}

	guard.turnRight()
	if guard.dir != Up {
		t.Errorf("After turning right from Left, expected Up, got %v", guard.dir)
	}
}

func TestNextPosition(t *testing.T) {
	tests := []struct {
		name     string
		pos      Position
		dir      Direction
		expected Position
	}{
		{"Up", Position{5, 5}, Up, Position{4, 5}},
		{"Right", Position{5, 5}, Right, Position{5, 6}},
		{"Down", Position{5, 5}, Down, Position{6, 5}},
		{"Left", Position{5, 5}, Left, Position{5, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			guard := &Guard{pos: tt.pos, dir: tt.dir}
			next := guard.nextPosition()
			if next != tt.expected {
				t.Errorf("nextPosition() = %v; want %v", next, tt.expected)
			}
		})
	}
}
