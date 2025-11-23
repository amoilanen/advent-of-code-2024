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

// findRegion uses BFS to find all cells in a connected region
// Returns a set of all cells in the region starting at (startRow, startColumn)
func findRegion(grid Grid, visited [][]bool, startRow, startColumn int) map[[2]int]bool {
	plantType := grid.Cells[startRow][startColumn]
	regionCells := make(map[[2]int]bool)

	queue := [][2]int{{startRow, startColumn}}
	visited[startRow][startColumn] = true

	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		currentRow, currentColumn := current[0], current[1]

		regionCells[[2]int{currentRow, currentColumn}] = true

		// Explore neighbors
		for _, dir := range directions {
			neighborRow, neighborColumn := currentRow+dir[0], currentColumn+dir[1]

			if neighborRow >= 0 && neighborRow < grid.Rows &&
				neighborColumn >= 0 && neighborColumn < grid.Cols &&
				grid.Cells[neighborRow][neighborColumn] == plantType &&
				!visited[neighborRow][neighborColumn] {
				visited[neighborRow][neighborColumn] = true
				queue = append(queue, [2]int{neighborRow, neighborColumn})
			}
		}
	}

	return regionCells
}

// calculatePerimeter calculates the perimeter of a region
// Perimeter is the count of edges that border cells outside the region
func calculatePerimeter(grid Grid, regionCells map[[2]int]bool) int {
	perimeter := 0
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for cell := range regionCells {
		currentRow, currentColumn := cell[0], cell[1]
		plantType := grid.Cells[currentRow][currentColumn]

		// Check each of the 4 sides
		for _, dir := range directions {
			neighborRow, neighborColumn := currentRow+dir[0], currentColumn+dir[1]

			// Edge contributes to perimeter if out of bounds or different plant type
			if neighborRow < 0 || neighborRow >= grid.Rows ||
				neighborColumn < 0 || neighborColumn >= grid.Cols ||
				grid.Cells[neighborRow][neighborColumn] != plantType {
				perimeter++
			}
		}
	}

	return perimeter
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
				regionCells := findRegion(grid, visited, r, c)
				area := len(regionCells)
				perimeter := calculatePerimeter(grid, regionCells)
				totalPrice += area * perimeter
			}
		}
	}

	return totalPrice
}

// isCorner checks if a corner exists at a cell's corner position
// Returns true if either an outer or inner corner is detected
func isCorner(regionCells map[[2]int]bool, row, col, dRow1, dCol1, dRow2, dCol2 int) bool {
	neighbor1 := regionCells[[2]int{row + dRow1, col + dCol1}]
	neighbor2 := regionCells[[2]int{row + dRow2, col + dCol2}]
	diagonal := regionCells[[2]int{row + dRow1 + dRow2, col + dCol1 + dCol2}]

	// Outer corner: both orthogonal neighbors are outside region
	if !neighbor1 && !neighbor2 {
		return true
	}

	// Inner corner: both orthogonal neighbors are inside, but diagonal is outside
	if neighbor1 && neighbor2 && !diagonal {
		return true
	}

	return false
}

// countCorners counts the number of corners in a region
// Key insight: number of sides = number of corners in any closed polygon
//
// For each cell, we check 4 possible corners (NW, NE, SW, SE):
// - Outer corner: both orthogonal neighbors are NOT in region
// - Inner corner: both orthogonal neighbors ARE in region, but diagonal is NOT
func countCorners(regionCells map[[2]int]bool) int {
	// Define the 4 corner configurations: [vertical offset, horizontal offset]
	// Each corner is defined by two orthogonal directions
	cornerConfigs := [][4]int{
		{-1, 0, 0, -1}, // NW: top, left
		{-1, 0, 0, 1},  // NE: top, right
		{1, 0, 0, -1},  // SW: bottom, left
		{1, 0, 0, 1},   // SE: bottom, right
	}

	totalCorners := 0

	for cell := range regionCells {
		row, col := cell[0], cell[1]

		// Check all 4 corners of this cell
		for _, config := range cornerConfigs {
			if isCorner(regionCells, row, col, config[0], config[1], config[2], config[3]) {
				totalCorners++
			}
		}
	}

	return totalCorners
}

// Part2 calculates the total fencing cost using bulk discount
// Algorithm:
// 1. Use flood fill (BFS) to identify connected regions of same plant type
// 2. For each region, calculate area and number of sides (corners)
// 3. Key insight: In any closed polygon, sides = corners
// 4. Count corners by checking each cell's 4 corner positions:
//   - Outer corners: both orthogonal neighbors outside region
//   - Inner corners: both orthogonal neighbors inside, diagonal outside
//
// 5. Price per region = area * sides
//
// Time complexity: O(rows * cols) - each cell visited once
// Space complexity: O(rows * cols) - for visited tracking and region storage
func Part2(grid Grid) int {
	visited := make([][]bool, grid.Rows)
	for i := range visited {
		visited[i] = make([]bool, grid.Cols)
	}

	totalPrice := 0

	// Find all regions using flood fill
	for r := 0; r < grid.Rows; r++ {
		for c := 0; c < grid.Cols; c++ {
			if !visited[r][c] {
				regionCells := findRegion(grid, visited, r, c)
				area := len(regionCells)
				sides := countCorners(regionCells)
				totalPrice += area * sides
			}
		}
	}

	return totalPrice
}
