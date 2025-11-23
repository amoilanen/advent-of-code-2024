package day13

import (
	"testing"
)

func TestParse(t *testing.T) {
	machines := Parse(ExampleInput)

	if len(machines) != 4 {
		t.Errorf("Expected 4 machines, got %d", len(machines))
	}

	// Test first machine
	if machines[0].ButtonA.X != 94 || machines[0].ButtonA.Y != 34 {
		t.Errorf("Machine 0 Button A: got (%d, %d), want (94, 34)",
			machines[0].ButtonA.X, machines[0].ButtonA.Y)
	}

	if machines[0].ButtonB.X != 22 || machines[0].ButtonB.Y != 67 {
		t.Errorf("Machine 0 Button B: got (%d, %d), want (22, 67)",
			machines[0].ButtonB.X, machines[0].ButtonB.Y)
	}

	if machines[0].Prize.X != 8400 || machines[0].Prize.Y != 5400 {
		t.Errorf("Machine 0 Prize: got (%d, %d), want (8400, 5400)",
			machines[0].Prize.X, machines[0].Prize.Y)
	}
}

func TestSolveMachine(t *testing.T) {
	tests := []struct {
		name        string
		machine     Machine
		expectValid bool
		expectA     int
		expectB     int
		expectCost  int
	}{
		{
			name: "First machine - solvable",
			machine: Machine{
				ButtonA: Vector{X: 94, Y: 34},
				ButtonB: Vector{X: 22, Y: 67},
				Prize:   Vector{X: 8400, Y: 5400},
			},
			expectValid: true,
			expectA:     80,
			expectB:     40,
			expectCost:  280,
		},
		{
			name: "Second machine - unsolvable",
			machine: Machine{
				ButtonA: Vector{X: 26, Y: 66},
				ButtonB: Vector{X: 67, Y: 21},
				Prize:   Vector{X: 12748, Y: 12176},
			},
			expectValid: false,
		},
		{
			name: "Third machine - solvable",
			machine: Machine{
				ButtonA: Vector{X: 17, Y: 86},
				ButtonB: Vector{X: 84, Y: 37},
				Prize:   Vector{X: 7870, Y: 6450},
			},
			expectValid: true,
			expectA:     38,
			expectB:     86,
			expectCost:  200,
		},
		{
			name: "Fourth machine - unsolvable",
			machine: Machine{
				ButtonA: Vector{X: 69, Y: 23},
				ButtonB: Vector{X: 27, Y: 71},
				Prize:   Vector{X: 18641, Y: 10279},
			},
			expectValid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solution := SolveMachine(tt.machine)

			if solution.Valid != tt.expectValid {
				t.Errorf("Valid = %v, want %v", solution.Valid, tt.expectValid)
			}

			if solution.Valid {
				if solution.APresses != tt.expectA {
					t.Errorf("APresses = %d, want %d", solution.APresses, tt.expectA)
				}
				if solution.BPresses != tt.expectB {
					t.Errorf("BPresses = %d, want %d", solution.BPresses, tt.expectB)
				}
				if solution.Cost != tt.expectCost {
					t.Errorf("Cost = %d, want %d", solution.Cost, tt.expectCost)
				}
			}
		})
	}
}

func TestPart1(t *testing.T) {
	machines := Parse(ExampleInput)
	result := Part1(machines)

	// Expected: 2 prizes won (machines 0 and 2), costing 280 + 200 = 480 tokens
	if result != 480 {
		t.Errorf("Part1() = %d, want 480", result)
	}
}

func TestPart2(t *testing.T) {
	machines := Parse(ExampleInput)
	result := Part2(machines)

	// After adding 10000000000000 to prize coordinates,
	// only machines 1 and 3 are solvable (according to problem description)
	if result != 875318608908 {
		t.Errorf("Part1() = %d, want 480", result)
	}
}

func TestSolveMachineWithoutPressLimit(t *testing.T) {
	// Test a machine that requires more than 100 presses
	machine := Machine{
		ButtonA: Vector{X: 26, Y: 66},
		ButtonB: Vector{X: 67, Y: 21},
		Prize:   Vector{X: 10000000012748, Y: 10000000012176},
	}

	solution := SolveMachineWithConstraints(machine, -1) // No press limit

	if !solution.Valid {
		t.Errorf("Expected valid solution for machine with offset prize")
	}

	// Verify the solution is correct
	if solution.Valid {
		ax, ay := machine.ButtonA.X, machine.ButtonA.Y
		bx, by := machine.ButtonB.X, machine.ButtonB.Y
		px, py := machine.Prize.X, machine.Prize.Y

		resultX := solution.APresses*ax + solution.BPresses*bx
		resultY := solution.APresses*ay + solution.BPresses*by

		if resultX != px || resultY != py {
			t.Errorf("Solution doesn't reach prize: got (%d, %d), want (%d, %d)",
				resultX, resultY, px, py)
		}
	}
}
