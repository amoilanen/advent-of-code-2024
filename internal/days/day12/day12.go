package day12

import (
	"strings"
)

const ExampleInput = `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`

// Grid represents the garden map
type Grid struct {
	Cells []string
	Rows  int
	Cols  int
}

// Parse converts the input string into a Grid
func Parse(input string) Grid {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	rows := len(lines)
	if rows == 0 {
		return Grid{}
	}

	cols := len(lines[0])
	return Grid{
		Cells: lines,
		Rows:  rows,
		Cols:  cols,
	}
}

// exploreRegion uses BFS to find all cells in a region and calculate area/perimeter
// Returns area and perimeter for the region starting at (startR, startC)
func exploreRegion(grid Grid, visited [][]bool, startRow, startColumn int) (int, int) {
	plantType := grid.Cells[startRow][startColumn]

	area := 0
	perimeter := 0

	queue := [][2]int{{startRow, startColumn}}
	visited[startRow][startColumn] = true

	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		currentRow, currentColumn := current[0], current[1]

		area++

		// Check each of the 4 sides
		for _, dir := range directions {
			neighborRow, neighborColumn := currentRow+dir[0], currentColumn+dir[1]

			// Edge contributes to perimeter if out of bounds or different plant type
			if neighborRow < 0 || neighborRow >= grid.Rows || neighborColumn < 0 || neighborColumn >= grid.Cols || grid.Cells[neighborRow][neighborColumn] != plantType {
				perimeter++
			} else if !visited[neighborRow][neighborColumn] {
				// Same plant type and unvisited - add to region
				visited[neighborRow][neighborColumn] = true
				queue = append(queue, [2]int{neighborRow, neighborColumn})
			}
		}
	}

	return area, perimeter
}

// Part1 calculates the total fencing cost for all regions
// Algorithm:
// 1. Use flood fill (BFS) to identify connected regions of same plant type
// 2. For each region, calculate area (cell count) and perimeter (edges not touching same type)
// 3. Price per region = area * perimeter
// 4. Total price = sum of all region prices
//
// Time complexity: O(rows * cols) - each cell visited once
// Space complexity: O(rows * cols) - for visited tracking
func Part1(grid Grid) int {
	visited := make([][]bool, grid.Rows)
	for i := range visited {
		visited[i] = make([]bool, grid.Cols)
	}

	totalPrice := 0

	// Find all regions using flood fill
	for r := 0; r < grid.Rows; r++ {
		for c := 0; c < grid.Cols; c++ {
			if !visited[r][c] {
				area, perimeter := exploreRegion(grid, visited, r, c)
				totalPrice += area * perimeter
			}
		}
	}

	return totalPrice
}

// Part2 placeholder for part 2
func Part2(grid Grid) int {
	return 0
}
