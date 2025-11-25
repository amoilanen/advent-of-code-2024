package day14

import (
	"regexp"
	"strconv"
	"strings"
)

const ExampleInput = `p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3`

// Vector represents a 2D coordinate or movement
type Vector struct {
	X int
	Y int
}

// Robot represents a robot with position and velocity
type Robot struct {
	Position Vector
	Velocity Vector
}

// Parse converts the input string into a slice of Robot configurations
// Input format: "p=x,y v=vx,vy" one per line
func Parse(input string) []Robot {
	var robots []Robot

	// Regular expression for parsing: p=x,y v=vx,vy
	robotRegex := regexp.MustCompile(`p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)`)

	lines := strings.Split(strings.TrimSpace(input), "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		matches := robotRegex.FindStringSubmatch(line)
		if len(matches) == 5 {
			px, _ := strconv.Atoi(matches[1])
			py, _ := strconv.Atoi(matches[2])
			vx, _ := strconv.Atoi(matches[3])
			vy, _ := strconv.Atoi(matches[4])

			robots = append(robots, Robot{
				Position: Vector{X: px, Y: py},
				Velocity: Vector{X: vx, Y: vy},
			})
		}
	}

	return robots
}

// CalculatePosition calculates where a robot will be after N seconds
// The space wraps around (teleportation at edges)
// Algorithm: new_pos = (initial_pos + velocity * seconds) mod dimensions
// Time complexity: O(1)
// Space complexity: O(1)
func CalculatePosition(robot Robot, seconds int, width int, height int) Vector {
	// Calculate new position with wrapping
	newX := (robot.Position.X + robot.Velocity.X*seconds) % width
	newY := (robot.Position.Y + robot.Velocity.Y*seconds) % height

	// Handle negative modulo results (Go's % can return negative values)
	if newX < 0 {
		newX += width
	}
	if newY < 0 {
		newY += height
	}

	return Vector{X: newX, Y: newY}
}

// CountQuadrants counts how many robots are in each quadrant
// Robots exactly in the middle (horizontally or vertically) are not counted
// Returns: (top-left, top-right, bottom-left, bottom-right)
// Time complexity: O(n) where n is number of positions
// Space complexity: O(1)
func CountQuadrants(positions []Vector, width int, height int) (int, int, int, int) {
	midX := width / 2
	midY := height / 2

	var topLeft, topRight, bottomLeft, bottomRight int

	for _, pos := range positions {
		// Skip robots on the middle lines
		if pos.X == midX || pos.Y == midY {
			continue
		}

		if pos.X < midX && pos.Y < midY {
			topLeft++
		} else if pos.X > midX && pos.Y < midY {
			topRight++
		} else if pos.X < midX && pos.Y > midY {
			bottomLeft++
		} else if pos.X > midX && pos.Y > midY {
			bottomRight++
		}
	}

	return topLeft, topRight, bottomLeft, bottomRight
}

// Part1 calculates the safety factor after 100 seconds
// Algorithm:
// 1. Calculate each robot's position after the given time
// 2. Count robots in each quadrant (excluding middle lines)
// 3. Multiply the counts to get the safety factor
//
// Time complexity: O(n) where n is number of robots
// Space complexity: O(n) for storing positions
func Part1(robots []Robot, width int, height int) int {
	const seconds = 100

	// Calculate positions after 100 seconds
	var positions []Vector
	for _, robot := range robots {
		pos := CalculatePosition(robot, seconds, width, height)
		positions = append(positions, pos)
	}

	// Count robots in each quadrant
	q1, q2, q3, q4 := CountQuadrants(positions, width, height)

	// Calculate safety factor
	safetyFactor := q1 * q2 * q3 * q4

	return safetyFactor
}

// sortInts performs a simple bubble sort on a slice of integers
// Time complexity: O(n^2)
// Space complexity: O(1)
func sortInts(coords []int) {
	for i := 0; i < len(coords)-1; i++ {
		for j := i + 1; j < len(coords); j++ {
			if coords[i] > coords[j] {
				coords[i], coords[j] = coords[j], coords[i]
			}
		}
	}
}

// countContiguousSequences counts how many contiguous sequences of at least minLength exist
// in a sorted slice of coordinates
// Time complexity: O(n)
// Space complexity: O(1)
func countContiguousSequences(sortedCoords []int, minLength int) int {
	if len(sortedCoords) < minLength {
		return 0
	}

	foundSequences := 0
	currentLength := 1

	for i := 1; i < len(sortedCoords); i++ {
		if sortedCoords[i] == sortedCoords[i-1]+1 {
			currentLength++
		} else {
			if currentLength >= minLength {
				foundSequences++
			}
			currentLength = 1
		}
	}

	// Don't forget the last sequence
	if currentLength >= minLength {
		foundSequences++
	}

	return foundSequences
}

// countLinesInDirection counts lines of at least minLength in a specific direction
// groupBy extracts the coordinate to group by (Y for horizontal, X for vertical)
// lineCoord extracts the coordinate along the line (X for horizontal, Y for vertical)
// Time complexity: O(n log n) for sorting + O(n) for counting
// Space complexity: O(n) for grouping
func countLinesInDirection(positions []Vector, minLength int, groupBy func(Vector) int, lineCoord func(Vector) int) int {
	// Group positions by the grouping coordinate
	byGroup := make(map[int][]int)
	for _, pos := range positions {
		key := groupBy(pos)
		value := lineCoord(pos)
		byGroup[key] = append(byGroup[key], value)
	}

	lineCount := 0

	// For each group, sort coordinates and count contiguous sequences
	for _, coords := range byGroup {
		if len(coords) < minLength {
			continue
		}

		sortInts(coords)
		lineCount += countContiguousSequences(coords, minLength)
	}

	return lineCount
}

// CountHorizontalLines counts how many horizontal lines of at least minLength exist
// A horizontal line is a contiguous sequence of robots on the same Y coordinate
// Time complexity: O(n log n) for sorting + O(n) for counting
// Space complexity: O(n) for grouping by Y coordinate
func CountHorizontalLines(positions []Vector, minLength int) int {
	return countLinesInDirection(positions, minLength,
		func(v Vector) int { return v.Y }, // Group by Y
		func(v Vector) int { return v.X }, // Line along X
	)
}

// CountVerticalLines counts how many vertical lines of at least minLength exist
// A vertical line is a contiguous sequence of robots on the same X coordinate
// Time complexity: O(n log n) for sorting + O(n) for counting
// Space complexity: O(n) for grouping by X coordinate
func CountVerticalLines(positions []Vector, minLength int) int {
	return countLinesInDirection(positions, minLength,
		func(v Vector) int { return v.X }, // Group by X
		func(v Vector) int { return v.Y }, // Line along Y
	)
}

// HasChristmasTreePattern detects if robot positions form a Christmas tree pattern
// Heuristic: Look for significant horizontal and vertical line formations
// The thresholds are adjusted based on typical Christmas tree structures
// Time complexity: O(n log n)
// Space complexity: O(n)
func HasChristmasTreePattern(positions []Vector) bool {
	// Try multiple heuristics - a Christmas tree should have:
	// 1. Many horizontal lines (branches)
	// 2. Some vertical lines (trunk/center)

	horizontalLines := CountHorizontalLines(positions, 10)
	verticalLines := CountVerticalLines(positions, 15)

	// Original strict heuristic
	if horizontalLines > 15 && verticalLines >= 3 {
		return true
	}

	// More relaxed heuristics for different input variations
	if horizontalLines > 10 && verticalLines >= 2 {
		return true
	}

	// Alternative: look for many consecutive positions (high structure)
	// If we have many lines total, it's likely a pattern
	if horizontalLines >= 8 && verticalLines >= 2 {
		return true
	}

	return false
}

// Part2 finds the minimum number of seconds for robots to display the Easter egg
// Algorithm:
// 1. Simulate robots over time (up to width * height iterations max)
// 2. For each time step, check if positions form a Christmas tree pattern
// 3. Return the first time when the pattern is detected
//
// Time complexity: O(T * n log n) where T is time to find pattern, n is robot count
// Space complexity: O(n) for storing positions
func Part2(robots []Robot, width int, height int) int {
	// Maximum cycle length is width * height
	maxIterations := width * height

	for seconds := 1; seconds < maxIterations; seconds++ {
		// Calculate positions at this time
		var positions []Vector
		for _, robot := range robots {
			pos := CalculatePosition(robot, seconds, width, height)
			positions = append(positions, pos)
		}

		// Check if this forms a Christmas tree pattern
		if HasChristmasTreePattern(positions) {
			return seconds
		}
	}

	// Pattern not found
	return -1
}
