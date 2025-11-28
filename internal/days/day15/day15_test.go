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

const SmallExamplePart2 = `#######
#...#.#
#.....#
#..OO@#
#..O..#
#.....#
#######

<vv<<^^<<^^`

func TestScaleWarehouse(t *testing.T) {
	warehouse, _ := Parse(SmallExample)
	scaled := ScaleWarehouse(warehouse)

	// Check dimensions
	if scaled.Width != warehouse.Width*2 {
		t.Errorf("expected width %d, got %d", warehouse.Width*2, scaled.Width)
	}
	if scaled.Height != warehouse.Height {
		t.Errorf("expected height %d, got %d", warehouse.Height, scaled.Height)
	}

	// Check robot position is scaled correctly
	expectedRobotCol := warehouse.Robot.Col * 2
	if scaled.Robot.Col != expectedRobotCol {
		t.Errorf("expected robot col %d, got %d", expectedRobotCol, scaled.Robot.Col)
	}

	// Check that a wall is doubled
	if warehouse.Grid[0][0] == '#' {
		if scaled.Grid[0][0] != '#' || scaled.Grid[0][1] != '#' {
			t.Errorf("expected wall to be doubled")
		}
	}

	// Check that a box is converted to []
	for row := 0; row < warehouse.Height; row++ {
		for col := 0; col < warehouse.Width; col++ {
			if warehouse.Grid[row][col] == 'O' {
				scaledCol := col * 2
				if scaled.Grid[row][scaledCol] != '[' || scaled.Grid[row][scaledCol+1] != ']' {
					t.Errorf("expected box at (%d,%d) to become [] at (%d,%d)", row, col, row, scaledCol)
				}
			}
		}
	}
}

func TestSimulateMoveWide_HorizontalPush(t *testing.T) {
	input := `####
#@O.#
####

>`

	warehouse, moves := Parse(input)
	warehouse = ScaleWarehouse(warehouse)
	warehouse.SimulateMoveWide(moves[0])

	// After scaling: ##########  Robot @ (1,2), Box at (1,4-5), moves right
	// After move: Robot should be at (1,3), box at (1,4-5)
	if warehouse.Robot.Row != 1 || warehouse.Robot.Col != 3 {
		t.Errorf("expected robot at (1, 3), got (%d, %d)", warehouse.Robot.Row, warehouse.Robot.Col)
	}

	// Check box moved
	if warehouse.Grid[1][4] != '[' || warehouse.Grid[1][5] != ']' {
		t.Errorf("expected box at (1, 4-5), got '%c%c'", warehouse.Grid[1][4], warehouse.Grid[1][5])
	}
}

func TestSimulateMoveWide_VerticalPushAligned(t *testing.T) {
	input := `####
#..#
#.O#
#.O#
#.@#
####

^`

	warehouse, moves := Parse(input)
	warehouse = ScaleWarehouse(warehouse)
	warehouse.SimulateMoveWide(moves[0])

	// After scaling robot is at (4,4), boxes at (2,4-5) and (3,4-5)
	// After move up: robot at (3,4), boxes at (1,4-5) and (2,4-5)
	if warehouse.Robot.Row != 3 || warehouse.Robot.Col != 4 {
		t.Errorf("expected robot at (3, 4), got (%d, %d)", warehouse.Robot.Row, warehouse.Robot.Col)
	}

	// Boxes should move up
	if warehouse.Grid[1][4] != '[' || warehouse.Grid[1][5] != ']' {
		t.Errorf("expected top box at (1, 4-5), got '%c%c'", warehouse.Grid[1][4], warehouse.Grid[1][5])
	}
	if warehouse.Grid[2][4] != '[' || warehouse.Grid[2][5] != ']' {
		t.Errorf("expected second box at (2, 4-5), got '%c%c'", warehouse.Grid[2][4], warehouse.Grid[2][5])
	}
}

func TestSimulateMoveWide_VerticalPushOffset(t *testing.T) {
	// Test when boxes are offset vertically
	// Original: Box at (2,2) and (3,1), robot at (4,1)
	// Scaled: Box at (2,4-5) and (3,2-3), robot at (4,2)
	// Robot pushes left edge of box at (3,2-3), which doesn't overlap with box at (2,4-5)
	input := `#####
#...#
#.O.#
#O..#
#@..#
#####

^`

	warehouse, moves := Parse(input)
	warehouse = ScaleWarehouse(warehouse)
	warehouse.SimulateMoveWide(moves[0])

	// Robot should move up from (4,2) to (3,2)
	if warehouse.Robot.Row != 3 || warehouse.Robot.Col != 2 {
		t.Errorf("expected robot at (3, 2), got (%d, %d)", warehouse.Robot.Row, warehouse.Robot.Col)
	}

	// Bottom box should move up from (3,2-3) to (2,2-3)
	if warehouse.Grid[2][2] != '[' || warehouse.Grid[2][3] != ']' {
		t.Errorf("expected bottom box at (2, 2-3), got '%c%c'", warehouse.Grid[2][2], warehouse.Grid[2][3])
	}

	// Top box should NOT move (it doesn't overlap with the pushed box)
	if warehouse.Grid[2][4] != '[' || warehouse.Grid[2][5] != ']' {
		t.Errorf("expected top box to stay at (2, 4-5), got '%c%c'", warehouse.Grid[2][4], warehouse.Grid[2][5])
	}
}

func TestSimulateMoveWide_VerticalBlocked(t *testing.T) {
	input := `####
#.##
#.O#
#.@#
####

^`

	warehouse, moves := Parse(input)
	warehouse = ScaleWarehouse(warehouse)
	initialRobot := warehouse.Robot

	warehouse.SimulateMoveWide(moves[0])

	// Robot should not move (blocked by wall above box)
	if warehouse.Robot.Row != initialRobot.Row || warehouse.Robot.Col != initialRobot.Col {
		t.Errorf("expected robot to stay at (%d, %d), got (%d, %d)",
			initialRobot.Row, initialRobot.Col, warehouse.Robot.Row, warehouse.Robot.Col)
	}
}

func TestPart2_LargeExample(t *testing.T) {
	result := Part2(LargeExample)
	expected := 9021
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}
