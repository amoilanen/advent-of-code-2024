package day08

import (
	"strings"

	"github.com/amoilanen/advent-of-code-2024/internal/utils"
)

const ExampleInput = `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`

// Point represents a coordinate on the grid
type Point struct {
	Row int
	Col int
}

// Grid represents the antenna map
type Grid struct {
	Width    int
	Height   int
	Antennas map[rune][]Point // Map from frequency to antenna positions
}

// Parse parses the input into a Grid
func Parse(input string) Grid {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := Grid{
		Height:   len(lines),
		Antennas: make(map[rune][]Point),
	}

	if len(lines) > 0 {
		grid.Width = len(lines[0])
	}

	for row, line := range lines {
		for col, ch := range line {
			if ch != '.' {
				grid.Antennas[ch] = append(grid.Antennas[ch], Point{Row: row, Col: col})
			}
		}
	}

	return grid
}

// isInBounds checks if a point is within the grid bounds
func (g Grid) isInBounds(p Point) bool {
	return p.Row >= 0 && p.Row < g.Height && p.Col >= 0 && p.Col < g.Width
}

// findAntinodes finds all antinodes for a given pair of antennas
// An antinode occurs at a point that is perfectly in line with two antennas,
// where one antenna is twice as far away as the other
func findAntinodes(a1, a2 Point) []Point {
	// The two antinodes are:
	// 1. Beyond a2: a2 + (a2 - a1) = 2*a2 - a1
	// 2. Beyond a1: a1 - (a2 - a1) = 2*a1 - a2
	return []Point{
		{Row: 2*a2.Row - a1.Row, Col: 2*a2.Col - a1.Col},
		{Row: 2*a1.Row - a2.Row, Col: 2*a1.Col - a2.Col},
	}
}

// Part1 calculates the number of unique antinode locations
func Part1(grid Grid) int {
	antinodes := make(map[Point]bool)

	// For each frequency
	for _, positions := range grid.Antennas {
		// Check all pairs of antennas with the same frequency
		for i := 0; i < len(positions); i++ {
			for j := i + 1; j < len(positions); j++ {
				// Find antinodes for this pair
				nodes := findAntinodes(positions[i], positions[j])
				for _, node := range nodes {
					// Only count antinodes within the grid bounds
					if grid.isInBounds(node) {
						antinodes[node] = true
					}
				}
			}
		}
	}

	return len(antinodes)
}

// findAllAntinodesOnLine finds all antinodes on the line through two antennas
// In part 2, any point in line with at least two antennas is an antinode
func (g Grid) findAllAntinodesOnLine(a1, a2 Point) []Point {
	dr := a2.Row - a1.Row
	dc := a2.Col - a1.Col

	// Find GCD to get the smallest step
	gcdVal := utils.GCD(utils.Abs(dr), utils.Abs(dc))
	stepRow := dr / gcdVal
	stepCol := dc / gcdVal

	antinodes := []Point{}

	// Walk forward from a1 (including a1 itself)
	current := a1
	for g.isInBounds(current) {
		antinodes = append(antinodes, current)
		current = Point{Row: current.Row + stepRow, Col: current.Col + stepCol}
	}

	// Walk backward from a1 (excluding a1 since it's already included)
	current = Point{Row: a1.Row - stepRow, Col: a1.Col - stepCol}
	for g.isInBounds(current) {
		antinodes = append(antinodes, current)
		current = Point{Row: current.Row - stepRow, Col: current.Col - stepCol}
	}

	return antinodes
}

// Part2 calculates the number of unique antinode locations using the updated model
// where any point in line with at least two antennas of the same frequency is an antinode
func Part2(grid Grid) int {
	antinodes := make(map[Point]bool)

	// For each frequency
	for _, positions := range grid.Antennas {
		// Check all pairs of antennas with the same frequency
		for i := 0; i < len(positions); i++ {
			for j := i + 1; j < len(positions); j++ {
				// Find all antinodes on the line through this pair
				nodes := grid.findAllAntinodesOnLine(positions[i], positions[j])
				for _, node := range nodes {
					antinodes[node] = true
				}
			}
		}
	}

	return len(antinodes)
}
