package day14

import (
	"testing"
)

func TestParse(t *testing.T) {
	robots := Parse(ExampleInput)

	if len(robots) != 12 {
		t.Errorf("Expected 12 robots, got %d", len(robots))
	}

	// Test first robot: p=0,4 v=3,-3
	if robots[0].Position.X != 0 || robots[0].Position.Y != 4 {
		t.Errorf("Robot 0 Position: got (%d, %d), want (0, 4)",
			robots[0].Position.X, robots[0].Position.Y)
	}

	if robots[0].Velocity.X != 3 || robots[0].Velocity.Y != -3 {
		t.Errorf("Robot 0 Velocity: got (%d, %d), want (3, -3)",
			robots[0].Velocity.X, robots[0].Velocity.Y)
	}

	// Test last robot: p=9,5 v=-3,-3
	if robots[11].Position.X != 9 || robots[11].Position.Y != 5 {
		t.Errorf("Robot 11 Position: got (%d, %d), want (9, 5)",
			robots[11].Position.X, robots[11].Position.Y)
	}

	if robots[11].Velocity.X != -3 || robots[11].Velocity.Y != -3 {
		t.Errorf("Robot 11 Velocity: got (%d, %d), want (-3, -3)",
			robots[11].Velocity.X, robots[11].Velocity.Y)
	}
}

func TestCalculatePosition(t *testing.T) {
	tests := []struct {
		name     string
		robot    Robot
		seconds  int
		width    int
		height   int
		expectedX int
		expectedY int
	}{
		{
			name: "Robot p=2,4 v=2,-3 after 1 second",
			robot: Robot{
				Position: Vector{X: 2, Y: 4},
				Velocity: Vector{X: 2, Y: -3},
			},
			seconds:  1,
			width:    11,
			height:   7,
			expectedX: 4,
			expectedY: 1,
		},
		{
			name: "Robot p=2,4 v=2,-3 after 2 seconds",
			robot: Robot{
				Position: Vector{X: 2, Y: 4},
				Velocity: Vector{X: 2, Y: -3},
			},
			seconds:  2,
			width:    11,
			height:   7,
			expectedX: 6,
			expectedY: 5,
		},
		{
			name: "Robot p=2,4 v=2,-3 after 5 seconds (wrapping)",
			robot: Robot{
				Position: Vector{X: 2, Y: 4},
				Velocity: Vector{X: 2, Y: -3},
			},
			seconds:  5,
			width:    11,
			height:   7,
			expectedX: 1,
			expectedY: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pos := CalculatePosition(tt.robot, tt.seconds, tt.width, tt.height)

			if pos.X != tt.expectedX || pos.Y != tt.expectedY {
				t.Errorf("CalculatePosition() = (%d, %d), want (%d, %d)",
					pos.X, pos.Y, tt.expectedX, tt.expectedY)
			}
		})
	}
}

func TestCountQuadrants(t *testing.T) {
	// After 100 seconds in example, we should have:
	// Top-left: 1, Top-right: 3, Bottom-left: 4, Bottom-right: 1
	robots := Parse(ExampleInput)

	// Calculate positions after 100 seconds
	var positions []Vector
	for _, robot := range robots {
		pos := CalculatePosition(robot, 100, 11, 7)
		positions = append(positions, pos)
	}

	q1, q2, q3, q4 := CountQuadrants(positions, 11, 7)

	if q1 != 1 {
		t.Errorf("Top-left quadrant: got %d, want 1", q1)
	}
	if q2 != 3 {
		t.Errorf("Top-right quadrant: got %d, want 3", q2)
	}
	if q3 != 4 {
		t.Errorf("Bottom-left quadrant: got %d, want 4", q3)
	}
	if q4 != 1 {
		t.Errorf("Bottom-right quadrant: got %d, want 1", q4)
	}
}

func TestPart1(t *testing.T) {
	robots := Parse(ExampleInput)
	result := Part1(robots, 11, 7)

	// Expected safety factor: 1 * 3 * 4 * 1 = 12
	if result != 12 {
		t.Errorf("Part1() = %d, want 12", result)
	}
}

func TestCountHorizontalLines(t *testing.T) {
	tests := []struct {
		name           string
		positions      []Vector
		minLength      int
		expectedCount  int
	}{
		{
			name: "Single horizontal line of length 5",
			positions: []Vector{
				{X: 0, Y: 0},
				{X: 1, Y: 0},
				{X: 2, Y: 0},
				{X: 3, Y: 0},
				{X: 4, Y: 0},
			},
			minLength:     5,
			expectedCount: 1,
		},
		{
			name: "Two horizontal lines",
			positions: []Vector{
				{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 0}, {X: 3, Y: 0}, {X: 4, Y: 0},
				{X: 0, Y: 2}, {X: 1, Y: 2}, {X: 2, Y: 2}, {X: 3, Y: 2}, {X: 4, Y: 2},
			},
			minLength:     5,
			expectedCount: 2,
		},
		{
			name: "Line too short",
			positions: []Vector{
				{X: 0, Y: 0},
				{X: 1, Y: 0},
				{X: 2, Y: 0},
			},
			minLength:     5,
			expectedCount: 0,
		},
		{
			name: "Broken line doesn't count",
			positions: []Vector{
				{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 0},
				{X: 4, Y: 0}, {X: 5, Y: 0}, // Gap at x=3
			},
			minLength:     5,
			expectedCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			count := CountHorizontalLines(tt.positions, tt.minLength)
			if count != tt.expectedCount {
				t.Errorf("CountHorizontalLines() = %d, want %d", count, tt.expectedCount)
			}
		})
	}
}

func TestCountVerticalLines(t *testing.T) {
	tests := []struct {
		name           string
		positions      []Vector
		minLength      int
		expectedCount  int
	}{
		{
			name: "Single vertical line of length 5",
			positions: []Vector{
				{X: 0, Y: 0},
				{X: 0, Y: 1},
				{X: 0, Y: 2},
				{X: 0, Y: 3},
				{X: 0, Y: 4},
			},
			minLength:     5,
			expectedCount: 1,
		},
		{
			name: "Two vertical lines",
			positions: []Vector{
				{X: 0, Y: 0}, {X: 0, Y: 1}, {X: 0, Y: 2}, {X: 0, Y: 3}, {X: 0, Y: 4},
				{X: 2, Y: 0}, {X: 2, Y: 1}, {X: 2, Y: 2}, {X: 2, Y: 3}, {X: 2, Y: 4},
			},
			minLength:     5,
			expectedCount: 2,
		},
		{
			name: "Line too short",
			positions: []Vector{
				{X: 0, Y: 0},
				{X: 0, Y: 1},
				{X: 0, Y: 2},
			},
			minLength:     5,
			expectedCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			count := CountVerticalLines(tt.positions, tt.minLength)
			if count != tt.expectedCount {
				t.Errorf("CountVerticalLines() = %d, want %d", count, tt.expectedCount)
			}
		})
	}
}

func TestHasChristmasTreePattern(t *testing.T) {
	// Create a pattern with enough lines to trigger detection
	var positions []Vector

	// Create 16 horizontal lines of length 10
	for y := 0; y < 16; y++ {
		for x := 0; x < 10; x++ {
			positions = append(positions, Vector{X: x, Y: y * 2})
		}
	}

	// Create 3 vertical lines of length 15
	for x := 20; x < 23; x++ {
		for y := 0; y < 15; y++ {
			positions = append(positions, Vector{X: x, Y: y})
		}
	}

	if !HasChristmasTreePattern(positions) {
		t.Error("Expected HasChristmasTreePattern() = true for valid pattern")
	}

	// Test with insufficient pattern
	smallPattern := []Vector{
		{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 0},
	}

	if HasChristmasTreePattern(smallPattern) {
		t.Error("Expected HasChristmasTreePattern() = false for small pattern")
	}
}
