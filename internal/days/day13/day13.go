package day13

import (
	"regexp"
	"strconv"
	"strings"
)

const ExampleInput = `Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279`

// Vector represents a 2D coordinate or movement
type Vector struct {
	X int
	Y int
}

// Machine represents a claw machine configuration
type Machine struct {
	ButtonA Vector
	ButtonB Vector
	Prize   Vector
}

// Solution represents the solution for a machine
type Solution struct {
	Valid    bool // Whether a valid solution exists
	APresses int  // Number of A button presses
	BPresses int  // Number of B button presses
	Cost     int  // Total cost in tokens
}

// Parse converts the input string into a slice of Machine configurations
func Parse(input string) []Machine {
	var machines []Machine

	// Regular expressions for parsing
	buttonRegex := regexp.MustCompile(`Button [AB]: X\+(\d+), Y\+(\d+)`)
	prizeRegex := regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)

	lines := strings.Split(strings.TrimSpace(input), "\n")

	var currentMachine Machine
	lineInMachine := 0

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			// Empty line separates machines
			if lineInMachine > 0 {
				machines = append(machines, currentMachine)
				currentMachine = Machine{}
				lineInMachine = 0
			}
			continue
		}

		if strings.HasPrefix(line, "Button A:") {
			matches := buttonRegex.FindStringSubmatch(line)
			x, _ := strconv.Atoi(matches[1])
			y, _ := strconv.Atoi(matches[2])
			currentMachine.ButtonA = Vector{X: x, Y: y}
			lineInMachine++
		} else if strings.HasPrefix(line, "Button B:") {
			matches := buttonRegex.FindStringSubmatch(line)
			x, _ := strconv.Atoi(matches[1])
			y, _ := strconv.Atoi(matches[2])
			currentMachine.ButtonB = Vector{X: x, Y: y}
			lineInMachine++
		} else if strings.HasPrefix(line, "Prize:") {
			matches := prizeRegex.FindStringSubmatch(line)
			x, _ := strconv.Atoi(matches[1])
			y, _ := strconv.Atoi(matches[2])
			currentMachine.Prize = Vector{X: x, Y: y}
			lineInMachine++
		}
	}

	// Don't forget the last machine
	if lineInMachine > 0 {
		machines = append(machines, currentMachine)
	}

	return machines
}

// SolveMachineWithConstraints finds the optimal solution using Cramer's rule
// Algorithm:
// We need to solve the system of linear equations:
//   a * A_x + b * B_x = Prize_x
//   a * A_y + b * B_y = Prize_y
//
// Using Cramer's rule:
//   D = A_x * B_y - A_y * B_x  (determinant)
//   a = (Prize_x * B_y - Prize_y * B_x) / D
//   b = (A_x * Prize_y - A_y * Prize_x) / D
//
// Parameters:
//   maxPresses: Maximum button presses allowed (-1 for no limit)
//
// Time complexity: O(1)
// Space complexity: O(1)
func SolveMachineWithConstraints(machine Machine, maxPresses int) Solution {
	const (
		ButtonACost = 3
		ButtonBCost = 1
	)

	ax, ay := machine.ButtonA.X, machine.ButtonA.Y
	bx, by := machine.ButtonB.X, machine.ButtonB.Y
	px, py := machine.Prize.X, machine.Prize.Y

	// Calculate determinant
	det := ax*by - ay*bx

	// If determinant is 0, the buttons are parallel (no unique solution)
	if det == 0 {
		return Solution{Valid: false}
	}

	// Apply Cramer's rule
	numeratorA := px*by - py*bx
	numeratorB := ax*py - ay*px

	// Check if solutions are integers
	if numeratorA%det != 0 || numeratorB%det != 0 {
		return Solution{Valid: false}
	}

	a := numeratorA / det
	b := numeratorB / det

	// Check if solutions are non-negative
	if a < 0 || b < 0 {
		return Solution{Valid: false}
	}

	// Check press limit if specified
	if maxPresses >= 0 && (a > maxPresses || b > maxPresses) {
		return Solution{Valid: false}
	}

	// Verify the solution (sanity check)
	if a*ax+b*bx != px || a*ay+b*by != py {
		return Solution{Valid: false}
	}

	cost := a*ButtonACost + b*ButtonBCost

	return Solution{
		Valid:    true,
		APresses: a,
		BPresses: b,
		Cost:     cost,
	}
}

// SolveMachine finds the optimal solution with the standard 100-press limit
func SolveMachine(machine Machine) Solution {
	return SolveMachineWithConstraints(machine, 100)
}

// Part1 calculates the minimum tokens needed to win all possible prizes
// Algorithm:
// 1. Parse input to get all machine configurations
// 2. For each machine, solve using Cramer's rule with 100-press limit
// 3. Sum up costs for all solvable machines
//
// Time complexity: O(n) where n is number of machines
// Space complexity: O(1) additional space
func Part1(machines []Machine) int {
	totalCost := 0

	for _, machine := range machines {
		solution := SolveMachine(machine)
		if solution.Valid {
			totalCost += solution.Cost
		}
	}

	return totalCost
}

// Part2 calculates the minimum tokens with corrected prize coordinates
// Algorithm:
// 1. Add 10000000000000 to each prize's X and Y coordinates
// 2. Solve without the 100-press limit (units were miscalculated)
// 3. Sum up costs for all solvable machines
//
// Key insight: The mathematical solution using Cramer's rule works for any
// coordinate values. Large numbers don't affect the O(1) complexity.
//
// Time complexity: O(n) where n is number of machines
// Space complexity: O(1) additional space
func Part2(machines []Machine) int {
	const PrizeOffset = 10000000000000

	totalCost := 0

	for _, machine := range machines {
		// Create a corrected machine with offset prize coordinates
		correctedMachine := Machine{
			ButtonA: machine.ButtonA,
			ButtonB: machine.ButtonB,
			Prize: Vector{
				X: machine.Prize.X + PrizeOffset,
				Y: machine.Prize.Y + PrizeOffset,
			},
		}

		// Solve without press limit
		solution := SolveMachineWithConstraints(correctedMachine, -1)
		if solution.Valid {
			totalCost += solution.Cost
		}
	}

	return totalCost
}
