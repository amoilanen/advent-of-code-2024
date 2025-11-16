package day10

import (
	"strings"
)

const ExampleInput = `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

// Position represents a coordinate on the topographic map
type Position struct {
	Row, Col int
}

// TopoMap represents the topographic map with heights 0-9
type TopoMap struct {
	Grid [][]int
	Rows int
	Cols int
}

// Parse converts the input string into a TopoMap
// Each character represents a height from 0-9
func Parse(input string) TopoMap {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	rows := len(lines)
	if rows == 0 {
		return TopoMap{}
	}

	cols := len(lines[0])
	grid := make([][]int, rows)

	for i, line := range lines {
		grid[i] = make([]int, cols)
		for j, ch := range line {
			grid[i][j] = int(ch - '0')
		}
	}

	return TopoMap{
		Grid: grid,
		Rows: rows,
		Cols: cols,
	}
}

// FindTrailheads returns all positions with height 0
func (tm TopoMap) FindTrailheads() []Position {
	trailheads := []Position{}

	for r := 0; r < tm.Rows; r++ {
		for c := 0; c < tm.Cols; c++ {
			if tm.Grid[r][c] == 0 {
				trailheads = append(trailheads, Position{r, c})
			}
		}
	}

	return trailheads
}

func (tm TopoMap) GetNeighbors(pos Position) []Position {
	directions := []Position{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	neighbors := []Position{}
	for _, dir := range directions {
		newRow := pos.Row + dir.Row
		newCol := pos.Col + dir.Col
		if newRow >= 0 && newRow < tm.Rows && newCol >= 0 && newCol < tm.Cols {
			neighbors = append(neighbors, Position{newRow, newCol})
		}
	}
	return neighbors
}

// GetTrailContinuations returns valid neighboring positions (4-directional)
// Only returns neighbors where height increases by exactly 1
func (tm TopoMap) GetTrailContinuations(pos Position, currentHeight int) []Position {
	continuations := []Position{}
	neighbors := tm.GetNeighbors(pos)
	for _, neighbor := range neighbors {
		if tm.Grid[neighbor.Row][neighbor.Col] == currentHeight+1 {
			continuations = append(continuations, neighbor)
		}
	}
	return continuations
}

// ScoreTrailhead calculates the score for a single trailhead
// Score = number of distinct height-9 positions reachable via valid hiking trails
// Uses BFS to explore all reachable positions
func (tm TopoMap) ScoreTrailhead(start Position) int {
	// Set to track unique height-9 positions reached
	reachedNines := make(map[Position]bool)

	// BFS queue - stores positions to explore
	queue := []Position{start}

	// Track visited positions to avoid cycles
	visited := make(map[Position]bool)
	visited[start] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		currentHeight := tm.Grid[current.Row][current.Col]

		// If we reached height 9, record it
		if currentHeight == 9 {
			reachedNines[current] = true
			continue
		}

		// Explore neighbors with height = currentHeight + 1
		for _, neighbor := range tm.GetTrailContinuations(current, currentHeight) {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}

	return len(reachedNines)
}

// Part1 calculates the sum of scores for all trailheads
// Algorithm:
// 1. Find all trailheads (height 0 positions)
// 2. For each trailhead, use BFS to find all reachable height-9 positions
// 3. Sum the scores
//
// Time complexity: O(R × C × T) where R=rows, C=cols, T=number of trailheads
// Space complexity: O(R × C) for visited sets
func Part1(topoMap TopoMap) int {
	trailheads := topoMap.FindTrailheads()
	totalScore := 0

	for _, trailhead := range trailheads {
		score := topoMap.ScoreTrailhead(trailhead)
		totalScore += score
	}

	return totalScore
}

// CountPaths calculates the number of distinct hiking trails from a position
// to any height-9 position using dynamic programming with memoization
// Returns: number of distinct paths from pos to any height-9 position
func (tm TopoMap) CountPaths(pos Position, memo map[Position]int) int {
	// Check if already computed
	if count, exists := memo[pos]; exists {
		return count
	}

	currentHeight := tm.Grid[pos.Row][pos.Col]

	// Base case: reached height 9 - this is one complete path
	if currentHeight == 9 {
		memo[pos] = 1
		return 1
	}

	// Recursive case: sum paths from all valid continuations
	totalPaths := 0
	for _, neighbor := range tm.GetTrailContinuations(pos, currentHeight) {
		totalPaths += tm.CountPaths(neighbor, memo)
	}

	memo[pos] = totalPaths
	return totalPaths
}

// RateTrailhead calculates the rating for a single trailhead
// Rating = number of distinct hiking trails that begin at this trailhead
func (tm TopoMap) RateTrailhead(start Position, memo map[Position]int) int {
	return tm.CountPaths(start, memo)
}

// Part2 calculates the sum of ratings for all trailheads
// Algorithm:
// 1. Find all trailheads (height 0 positions)
// 2. For each trailhead, count distinct paths to any height-9 using DP
// 3. Sum the ratings
//
// Time complexity: O(R × C) - each position computed at most once via memoization
// Space complexity: O(R × C) for memoization map
func Part2(topoMap TopoMap) int {
	trailheads := topoMap.FindTrailheads()
	totalRating := 0

	// Shared memoization map across all trailheads for efficiency
	memo := make(map[Position]int)

	for _, trailhead := range trailheads {
		rating := topoMap.RateTrailhead(trailhead, memo)
		totalRating += rating
	}

	return totalRating
}
