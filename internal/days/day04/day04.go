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
func (g Grid) hasWordAtPositionDirection(row int, col int, dir Direction, word string) bool {
	rows := len(g)
	if rows == 0 {
		return false
	}
	cols := len(g[0])

	for i, char := range word {
		newRow := row + i*dir.dr
		newCol := col + i*dir.dc

		// Check bounds
		if newRow < 0 || newRow >= rows || newCol < 0 || newCol >= cols {
			return false
		}

		// Check if character matches
		if g[newRow][newCol] != char {
			return false
		}
	}

	return true
}

// countWordOccurrences counts all occurrences of a word in the grid
func (g Grid) countWordOccurrences(word string) int {
	count := 0
	rows := len(g)
	if rows == 0 {
		return 0
	}
	cols := len(g[0])

	// Try starting from each position
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			// Try each direction
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

// Part2 placeholder for part 2
func Part2(grid Grid) int {
	// TODO: Implement Part 2 when we get the problem description
	return 0
}
