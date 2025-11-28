package day15

import (
	"testing"
)

const SmallExample = `########
#..O.O.#
##@.O..#
#...O..#
#.#.O..#
#...O..#
#......#
########

<^^>>>vv<v>>v<<`

const LargeExample = `##########
#..O..O.O#
#......O.#
#.OO..O.O#
#..O@..O.#
#O#..O...#
#O..O..O.#
#.OO.O.OO#
#....O...#
##########

<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^
vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v
><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<
<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^
^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><
^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^
>^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^
<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>
^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>
v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^`

func TestParse(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		wantHeight   int
		wantWidth    int
		wantRobotRow int
		wantRobotCol int
		wantMoves    int
	}{
		{"small example", SmallExample, 8, 8, 2, 2, 15},
		{"large example", LargeExample, 10, 10, 4, 4, 700},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			warehouse, moves := Parse(tt.input)

			// Check warehouse dimensions
			if warehouse.Height != tt.wantHeight {
				t.Errorf("expected height %d, got %d", tt.wantHeight, warehouse.Height)
			}
			if warehouse.Width != tt.wantWidth {
				t.Errorf("expected width %d, got %d", tt.wantWidth, warehouse.Width)
			}

			// Check robot position
			if warehouse.Robot.Row != tt.wantRobotRow || warehouse.Robot.Col != tt.wantRobotCol {
				t.Errorf("expected robot at (%d, %d), got (%d, %d)",
					tt.wantRobotRow, tt.wantRobotCol, warehouse.Robot.Row, warehouse.Robot.Col)
			}

			// Check moves count
			if len(moves) != tt.wantMoves {
				t.Errorf("expected %d moves, got %d", tt.wantMoves, len(moves))
			}
		})
	}
}

func TestParse_SmallExample(t *testing.T) {
	warehouse, moves := Parse(SmallExample)

	// Check dimensions
	if warehouse.Height != 8 {
		t.Errorf("expected height 8, got %d", warehouse.Height)
	}
	if warehouse.Width != 8 {
		t.Errorf("expected width 8, got %d", warehouse.Width)
	}

	// Check robot initial position (row 2, col 2)
	if warehouse.Robot.Row != 2 || warehouse.Robot.Col != 2 {
		t.Errorf("expected robot at (2, 2), got (%d, %d)", warehouse.Robot.Row, warehouse.Robot.Col)
	}

	// Check moves
	expectedMoves := "<^^>>>vv<v>>v<<"
	if len(moves) != len(expectedMoves) {
		t.Errorf("expected %d moves, got %d", len(expectedMoves), len(moves))
	}
}

func TestSimulateMove_NoObstacle(t *testing.T) {
	input := `####
#@.#
####

>`

	warehouse, moves := Parse(input)
	warehouse.SimulateMove(moves[0])

	// Robot should move from (1,1) to (1,2)
	if warehouse.Robot.Row != 1 || warehouse.Robot.Col != 2 {
		t.Errorf("expected robot at (1, 2), got (%d, %d)", warehouse.Robot.Row, warehouse.Robot.Col)
	}
}

func TestSimulateMove_Wall(t *testing.T) {
	input := `####
#@.#
####

<`

	warehouse, moves := Parse(input)
	warehouse.SimulateMove(moves[0])

	// Robot should not move (wall to the left)
	if warehouse.Robot.Row != 1 || warehouse.Robot.Col != 1 {
		t.Errorf("expected robot to stay at (1, 1), got (%d, %d)", warehouse.Robot.Row, warehouse.Robot.Col)
	}
}

func TestSimulateMove_PushOneBox(t *testing.T) {
	input := `#####
#@O.#
#####

>`

	warehouse, moves := Parse(input)
	warehouse.SimulateMove(moves[0])

	// Robot should move to (1,2), box should move to (1,3)
	if warehouse.Robot.Row != 1 || warehouse.Robot.Col != 2 {
		t.Errorf("expected robot at (1, 2), got (%d, %d)", warehouse.Robot.Row, warehouse.Robot.Col)
	}

	// Check box moved
	if warehouse.Grid[1][3] != 'O' {
		t.Errorf("expected box at (1, 3), got '%c'", warehouse.Grid[1][3])
	}
	if warehouse.Grid[1][1] != '.' {
		t.Errorf("expected empty at (1, 1), got '%c'", warehouse.Grid[1][1])
	}
}

func TestSimulateMove_PushMultipleBoxes(t *testing.T) {
	input := `######
#@OO.#
######

>`

	warehouse, moves := Parse(input)
	warehouse.SimulateMove(moves[0])

	// Robot should move to (1,2), boxes should shift right
	if warehouse.Robot.Row != 1 || warehouse.Robot.Col != 2 {
		t.Errorf("expected robot at (1, 2), got (%d, %d)", warehouse.Robot.Row, warehouse.Robot.Col)
	}

	// Check boxes moved
	if warehouse.Grid[1][3] != 'O' || warehouse.Grid[1][4] != 'O' {
		t.Errorf("expected boxes at (1, 3) and (1, 4)")
	}
}

func TestSimulateMove_PushBoxWall(t *testing.T) {
	input := `#####
#@O.#
#####

<<<`

	warehouse, moves := Parse(input)

	// First move left - should not move (wall)
	warehouse.SimulateMove(moves[0])
	if warehouse.Robot.Row != 1 || warehouse.Robot.Col != 1 {
		t.Errorf("expected robot to stay at (1, 1), got (%d, %d)", warehouse.Robot.Row, warehouse.Robot.Col)
	}
}

func TestCalculateGPS(t *testing.T) {
	tests := []struct {
		name     string
		row      int
		col      int
		expected int
	}{
		{"example from description", 1, 4, 104},
		{"top-left corner", 0, 0, 0},
		{"second row start", 1, 0, 100},
		{"bottom-right area", 5, 7, 507},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CalculateGPS(tt.row, tt.col)
			if result != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, result)
			}
		})
	}
}

func TestPart1_SmallExample(t *testing.T) {
	result := Part1(SmallExample)
	expected := 2028
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func TestPart1_LargeExample(t *testing.T) {
	result := Part1(LargeExample)
	expected := 10092
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}
