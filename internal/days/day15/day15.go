package day15

import (
	"strings"
)

// Position represents a coordinate in the warehouse
type Position struct {
	Row int
	Col int
}

// Warehouse represents the state of the warehouse
type Warehouse struct {
	Grid   [][]rune
	Width  int
	Height int
	Robot  Position
}

// Parse converts the input string into a Warehouse and list of moves
// Input format: grid map followed by blank line, then movement commands
// Time complexity: O(n) where n is input size
// Space complexity: O(w*h) for grid storage
func Parse(input string) (*Warehouse, []rune) {
	parts := strings.Split(strings.TrimSpace(input), "\n\n")
	if len(parts) != 2 {
		return nil, nil
	}

	gridLines := strings.Split(parts[0], "\n")
	moveLines := strings.ReplaceAll(parts[1], "\n", "")

	// Build grid and find robot
	var grid [][]rune
	var robot Position
	robot.Row = -1
	robot.Col = -1

	for row, line := range gridLines {
		rowRunes := []rune(line)
		grid = append(grid, rowRunes)

		// Find robot position
		for col, ch := range rowRunes {
			if ch == '@' {
				robot.Row = row
				robot.Col = col
			}
		}
	}

	warehouse := &Warehouse{
		Grid:   grid,
		Width:  len(grid[0]),
		Height: len(grid),
		Robot:  robot,
	}

	moves := []rune(moveLines)

	return warehouse, moves
}

// GetDirection converts a move character to a direction vector
func GetDirection(move rune) Position {
	switch move {
	case '^':
		return Position{-1, 0}
	case 'v':
		return Position{1, 0}
	case '<':
		return Position{0, -1}
	case '>':
		return Position{0, 1}
	default:
		return Position{0, 0}
	}
}

// isInBounds checks if a position is within the warehouse bounds
func (w *Warehouse) isInBounds(row, col int) bool {
	return row >= 0 && row < w.Height && col >= 0 && col < w.Width
}

// findPushTarget finds the first non-box cell in the given direction
// Returns the position and cell type, or (0, 0, 0) if out of bounds
func (w *Warehouse) findPushTarget(startRow, startCol int, dir Position) (int, int, rune) {
	row := startRow
	col := startCol

	for {
		row += dir.Row
		col += dir.Col

		if !w.isInBounds(row, col) {
			return 0, 0, 0 // Out of bounds
		}

		cell := w.Grid[row][col]
		if cell != 'O' {
			return row, col, cell // Found non-box (wall or empty)
		}
		// Continue if it's a box
	}
}

// moveRobot updates the robot position on the grid
func (w *Warehouse) moveRobot(newRow, newCol int) {
	w.Grid[w.Robot.Row][w.Robot.Col] = '.'
	w.Grid[newRow][newCol] = '@'
	w.Robot.Row = newRow
	w.Robot.Col = newCol
}

// SimulateMove simulates a single robot move
// Algorithm:
// 1. Determine direction from move character
// 2. Look ahead in that direction to find first non-box cell
// 3. If it's a wall, don't move
// 4. If it's empty, shift everything (robot and any boxes) in that direction
//
// Time complexity: O(max(width, height)) - scanning in one direction
// Space complexity: O(1)
func (w *Warehouse) SimulateMove(move rune) {
	dir := GetDirection(move)
	if dir.Row == 0 && dir.Col == 0 {
		return // Invalid move
	}

	nextRow := w.Robot.Row + dir.Row
	nextCol := w.Robot.Col + dir.Col

	if !w.isInBounds(nextRow, nextCol) {
		return
	}

	nextCell := w.Grid[nextRow][nextCol]

	switch nextCell {
	case '#':
		// Wall - can't move
		return

	case '.':
		// Empty space - just move robot
		w.moveRobot(nextRow, nextCol)
		return

	case 'O':
		// Box - check if we can push
		targetRow, targetCol, targetCell := w.findPushTarget(nextRow, nextCol, dir)

		if targetCell == '#' || targetCell == 0 {
			// Wall or out of bounds - can't push
			return
		}

		if targetCell == '.' {
			// Found empty space - push boxes and move robot
			w.Grid[targetRow][targetCol] = 'O'
			w.moveRobot(nextRow, nextCol)
		}
	}
}

// CalculateGPS calculates the GPS coordinate for a position
// GPS = 100 * row + col
// Time complexity: O(1)
// Space complexity: O(1)
func CalculateGPS(row, col int) int {
	return 100*row + col
}

// SumBoxGPS calculates the sum of GPS coordinates for all boxes
// Time complexity: O(width * height)
// Space complexity: O(1)
func (w *Warehouse) SumBoxGPS() int {
	sum := 0
	for row := 0; row < w.Height; row++ {
		for col := 0; col < w.Width; col++ {
			if w.Grid[row][col] == 'O' {
				sum += CalculateGPS(row, col)
			}
		}
	}
	return sum
}

// Part1 simulates all robot moves and returns sum of GPS coordinates
// Algorithm:
// 1. Parse input to get warehouse state and moves
// 2. Simulate each move in sequence
// 3. Calculate sum of GPS coordinates for all boxes
//
// Time complexity: O(M * N) where M = number of moves, N = grid size
// Space complexity: O(N) for grid storage
func Part1(input string) int {
	warehouse, moves := Parse(input)

	// Handle empty input
	if warehouse == nil {
		return 0
	}

	// Simulate all moves
	for _, move := range moves {
		warehouse.SimulateMove(move)
	}

	// Calculate sum of GPS coordinates
	return warehouse.SumBoxGPS()
}
