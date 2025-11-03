package day04

import "strings"

const ExampleInput = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

// Grid represents the word search grid
type Grid [][]rune

// Parse converts the input string into a 2D grid
func Parse(input string) Grid {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make(Grid, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}
	return grid
}

// dimensions returns the number of rows and columns in the grid
func (g Grid) dimensions() (rows, cols int) {
	rows = len(g)
	if rows == 0 {
		return 0, 0
	}
	cols = len(g[0])
	return rows, cols
}

// isInBounds checks if the given position is within the grid bounds
func (g Grid) isInBounds(row, col int) bool {
	rows, cols := g.dimensions()
	return row >= 0 && row < rows && col >= 0 && col < cols
}

// at safely gets the character at the given position
func (g Grid) at(row, col int) rune {
	if !g.isInBounds(row, col) {
		return 0
	}
	return g[row][col]
}

// Direction represents a search direction (row delta, col delta)
type Direction struct {
	dr, dc int
}

// All 8 possible directions to search
var directions = []Direction{
	{0, 1},   // Right
	{0, -1},  // Left
	{1, 0},   // Down
	{-1, 0},  // Up
	{1, 1},   // Down-Right
	{1, -1},  // Down-Left
	{-1, 1},  // Up-Right
	{-1, -1}, // Up-Left
}

// hasWordAtPositionDirection checks if a word exists starting from (row, col) in a given direction
func (g Grid) hasWordAtPositionDirection(row, col int, dir Direction, word string) bool {
	for i, char := range word {
		checkRow := row + i*dir.dr
		checkCol := col + i*dir.dc

		if g.at(checkRow, checkCol) != char {
			return false
		}
	}
	return true
}

// countMatches counts positions in the grid where the predicate returns true
func (g Grid) countMatches(predicate func(row, col int) bool) int {
	count := 0
	rows, cols := g.dimensions()

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if predicate(row, col) {
				count++
			}
		}
	}
	return count
}

// countWordOccurrences counts all occurrences of a word in the grid in all directions
func (g Grid) countWordOccurrences(word string) int {
	count := 0
	rows, cols := g.dimensions()

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			for _, dir := range directions {
				if g.hasWordAtPositionDirection(row, col, dir, word) {
					count++
				}
			}
		}
	}
	return count
}

// Part1 counts how many times "XMAS" appears in the word search
func Part1(grid Grid) int {
	return grid.countWordOccurrences("XMAS")
}

// isDiagonalMAS checks if two positions form a valid MAS diagonal (can be forward or backward)
func isDiagonalMAS(char1, char2 rune) bool {
	return (char1 == 'M' && char2 == 'S') || (char1 == 'S' && char2 == 'M')
}

// isXMAS checks if position (row, col) is the center 'A' of an X-MAS pattern
// An X-MAS is two "MAS" forming an X shape, where each MAS can be forward or backward
func (g Grid) isXMAS(row, col int) bool {
	// Check if center is 'A'
	if g.at(row, col) != 'A' {
		return false
	}

	// Get the 4 corners (at() returns 0 for out-of-bounds, which won't match M or S)
	topLeft := g.at(row-1, col-1)
	topRight := g.at(row-1, col+1)
	bottomLeft := g.at(row+1, col-1)
	bottomRight := g.at(row+1, col+1)

	// Both diagonals must form MAS (forward or backward)
	diag1Valid := isDiagonalMAS(topLeft, bottomRight)
	diag2Valid := isDiagonalMAS(topRight, bottomLeft)

	return diag1Valid && diag2Valid
}

// Part2 counts how many X-MAS patterns appear (two MAS in X shape)
func Part2(grid Grid) int {
	return grid.countMatches(grid.isXMAS)
}
