package day07

import (
	"strconv"
	"strings"
)

const ExampleInput = `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

// Operator represents a mathematical operator
type Operator int

const (
	Add Operator = iota
	Multiply
	Concatenate
)

// Equation represents a calibration equation
type Equation struct {
	TestValue int
	Numbers   []int
}

// Parse parses the input into a slice of equations
func Parse(input string) []Equation {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	equations := make([]Equation, 0, len(lines))

	for _, line := range lines {
		if line == "" {
			continue
		}

		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			continue
		}

		testValue, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			continue
		}

		numStrs := strings.Fields(parts[1])
		numbers := make([]int, 0, len(numStrs))
		for _, numStr := range numStrs {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				continue
			}
			numbers = append(numbers, num)
		}

		equations = append(equations, Equation{
			TestValue: testValue,
			Numbers:   numbers,
		})
	}

	return equations
}

// evaluate evaluates the numbers with the given operators (left-to-right)
func evaluate(numbers []int, operators []Operator) int {
	if len(numbers) == 0 {
		return 0
	}

	result := numbers[0]
	for i := 0; i < len(operators) && i < len(numbers)-1; i++ {
		switch operators[i] {
		case Add:
			result += numbers[i+1]
		case Multiply:
			result *= numbers[i+1]
		case Concatenate:
			concatenated := strconv.Itoa(result) + strconv.Itoa(numbers[i+1])
			result, _ = strconv.Atoi(concatenated)
		}
	}

	return result
}

// canBeMadeTrueWithOperators checks if equation can be made true using given operators
func canBeMadeTrueWithOperators(eq Equation, allowedOperators []Operator) bool {
	if len(eq.Numbers) == 0 {
		return false
	}

	if len(eq.Numbers) == 1 {
		return eq.Numbers[0] == eq.TestValue
	}

	// Number of operator positions
	numOps := len(eq.Numbers) - 1
	numOperatorTypes := len(allowedOperators)

	// Calculate total combinations: numOperatorTypes^numOps (exponentiation)
	totalCombinations := 1
	for i := 0; i < numOps; i++ {
		totalCombinations *= numOperatorTypes
	}

	// Try all combinations using base-n representation
	for combo := 0; combo < totalCombinations; combo++ {
		operators := make([]Operator, numOps)
		temp := combo
		for i := 0; i < numOps; i++ {
			operatorIndex := temp % numOperatorTypes
			operators[i] = allowedOperators[operatorIndex]
			temp /= numOperatorTypes
		}

		if evaluate(eq.Numbers, operators) == eq.TestValue {
			return true
		}
	}

	return false
}

// canBeMadeTrue checks if the equation can be made true with Add and Multiply operators
func canBeMadeTrue(eq Equation) bool {
	return canBeMadeTrueWithOperators(eq, []Operator{Add, Multiply})
}

// canBeMadeTrueWithConcat checks if the equation can be made true with Add, Multiply, and Concatenate
func canBeMadeTrueWithConcat(eq Equation) bool {
	return canBeMadeTrueWithOperators(eq, []Operator{Add, Multiply, Concatenate})
}

// Part1 calculates the sum of test values from equations that can be made true
func Part1(equations []Equation) int {
	sum := 0
	for _, eq := range equations {
		if canBeMadeTrue(eq) {
			sum += eq.TestValue
		}
	}
	return sum
}

// Part2 calculates the sum using Add, Multiply, and Concatenate operators
func Part2(equations []Equation) int {
	sum := 0
	for _, eq := range equations {
		if canBeMadeTrueWithConcat(eq) {
			sum += eq.TestValue
		}
	}
	return sum
}
