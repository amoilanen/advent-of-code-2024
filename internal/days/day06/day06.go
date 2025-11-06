package day06

import "strings"

const ExampleInput = `
....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

// Direction represents the guard's facing direction
type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

// Position represents a coordinate on the grid
type Position struct {
	Row, Col int
}

// Grid represents the lab map
type Grid struct {
	obstacles map[Position]bool
	rows      int
	cols      int
}

// Guard represents the guard's state
type Guard struct {
	pos Position
	dir Direction
}

// Parse parses the input into a grid and guard
func Parse(input string) (*Grid, *Guard) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	rows := len(lines)
	cols := 0
	if rows > 0 {
		cols = len(lines[0])
	}

	grid := &Grid{
		obstacles: make(map[Position]bool),
		rows:      rows,
		cols:      cols,
	}

	var guard *Guard

	for r, line := range lines {
		for c, char := range line {
			pos := Position{Row: r, Col: c}
			switch char {
			case '#':
				grid.obstacles[pos] = true
			case '^':
				guard = &Guard{pos: pos, dir: Up}
			case '>':
				guard = &Guard{pos: pos, dir: Right}
			case 'v':
				guard = &Guard{pos: pos, dir: Down}
			case '<':
				guard = &Guard{pos: pos, dir: Left}
			}
		}
	}

	return grid, guard
}

// turnRight turns the guard 90 degrees clockwise
func (g *Guard) turnRight() {
	g.dir = (g.dir + 1) % 4
}

// nextPosition returns the position directly in front of the guard
func (g *Guard) nextPosition() Position {
	pos := g.pos
	switch g.dir {
	case Up:
		pos.Row--
	case Right:
		pos.Col++
	case Down:
		pos.Row++
	case Left:
		pos.Col--
	}
	return pos
}

func (g *Guard) move() *Guard {
	nextPos := g.nextPosition()
	g.pos = nextPos
	return g
}

func (g *Guard) moveOnGrid(grid *Grid) (*Guard, bool) {
	next := g.nextPosition()

	// If next position is out of bounds, guard leaves the area
	if !grid.isInBounds(next) {
		return g, true
	}

	// If there's an obstacle, turn right
	if grid.hasObstacle(next) {
		g.turnRight()
	} else {
		g.move()
	}
	return g, false
}

// isInBounds checks if a position is within the grid
func (grid *Grid) isInBounds(pos Position) bool {
	return pos.Row >= 0 && pos.Row < grid.rows && pos.Col >= 0 && pos.Col < grid.cols
}

// hasObstacle checks if there's an obstacle at the given position
func (grid *Grid) hasObstacle(pos Position) bool {
	return grid.obstacles[pos]
}

// simulatePatrol simulates the guard's patrol and returns visited positions
// Creates a copy of the guard to avoid modifying the original
func simulatePatrol(grid *Grid, guard *Guard) map[Position]bool {
	if guard == nil {
		return make(map[Position]bool)
	}

	// Create a copy of the guard to avoid modifying the original
	simulatedGuard := &Guard{
		pos: guard.pos,
		dir: guard.dir,
	}

	visited := make(map[Position]bool)
	visited[simulatedGuard.pos] = true

	for {
		_, movedOffGrid := simulatedGuard.moveOnGrid(grid)
		if !movedOffGrid {
			visited[simulatedGuard.pos] = true
		}
		if movedOffGrid {
			break
		}
	}

	return visited
}

// State represents a guard's position and direction
type State struct {
	pos Position
	dir Direction
}

// simulateWithLoopDetection simulates the guard's patrol and detects if a loop occurs
// Returns true if a loop is detected, false if guard leaves the area
func simulateWithLoopDetection(grid *Grid, guard *Guard) bool {
	if guard == nil {
		return false
	}

	// Create a copy of the guard
	simulatedGuard := &Guard{
		pos: guard.pos,
		dir: guard.dir,
	}

	// Track states (position + direction)
	states := make(map[State]bool)
	currentState := State{pos: simulatedGuard.pos, dir: simulatedGuard.dir}
	states[currentState] = true

	for {
		_, movedOffGrid := simulatedGuard.moveOnGrid(grid)
		if movedOffGrid {
			return false
		}

		// Check for loop: if we've seen this (position, direction) state before
		currentState = State{pos: simulatedGuard.pos, dir: simulatedGuard.dir}
		if states[currentState] {
			return true // Loop detected
		}
		states[currentState] = true
	}
}

// Part1 counts the number of distinct positions visited by the guard
func Part1(grid *Grid, guard *Guard) int {
	visited := simulatePatrol(grid, guard)
	return len(visited)
}

// Part2 counts how many positions could have a new obstruction to create a loop
func Part2(grid *Grid, guard *Guard) int {
	if guard == nil {
		return 0
	}

	// Get the original patrol path - only these positions are candidates
	originalPath := simulatePatrol(grid, guard)

	count := 0
	startPos := guard.pos

	// Try placing an obstruction at each position on the original path
	for pos := range originalPath {
		// Skip the starting position
		if pos == startPos {
			continue
		}

		// Temporarily add an obstruction
		grid.obstacles[pos] = true

		// Check if this creates a loop
		if simulateWithLoopDetection(grid, guard) {
			count++
		}

		// Remove the obstruction
		delete(grid.obstacles, pos)
	}

	return count
}
