package day03

import (
	"regexp"
	"strconv"
)

const ExampleInput = `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`

// Mul represents a multiplication instruction with two operands
type Mul struct {
	X int
	Y int
}

// Pattern to match valid mul(X,Y) instructions where X and Y are 1-3 digit numbers
var mulPattern = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

// Parse extracts all valid mul instructions from the corrupted memory
func Parse(input string) []Mul {
	matches := mulPattern.FindAllStringSubmatch(input, -1)
	results := make([]Mul, 0, len(matches))

	for _, match := range matches {
		// match[0] is the full match, match[1] is X, match[2] is Y
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])
		results = append(results, Mul{X: x, Y: y})
	}

	return results
}

// Part1 calculates the sum of all multiplication results
func Part1(instructions []Mul) int {
	sum := 0
	for _, mul := range instructions {
		sum += mul.X * mul.Y
	}
	return sum
}
