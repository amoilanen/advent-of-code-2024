package day03

import (
	"regexp"
	"strconv"
)

const ExampleInput = `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`
const ExampleInputPart2 = `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`

// Mul represents a multiplication instruction with two operands
type Mul struct {
	X int
	Y int
}

// Instruction is the interface that all instructions implement
type Instruction interface {
	Position() int
}

// MulInstruction represents a multiplication instruction
type MulInstruction struct {
	Pos int
	Mul Mul
}

func (m MulInstruction) Position() int { return m.Pos }

// DoInstruction represents a do() instruction that enables mul
type DoInstruction struct {
	Pos int
}

func (d DoInstruction) Position() int { return d.Pos }

// DontInstruction represents a don't() instruction that disables mul
type DontInstruction struct {
	Pos int
}

func (d DontInstruction) Position() int { return d.Pos }

// Pattern to match all instructions: mul(X,Y), do(), or don't()
var instructionPattern = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)

// Parse extracts all instructions (mul, do, don't) from the corrupted memory
func Parse(input string) []Instruction {
	matches := instructionPattern.FindAllStringSubmatch(input, -1)
	indices := instructionPattern.FindAllStringSubmatchIndex(input, -1)
	instructions := make([]Instruction, 0, len(matches))

	for i, match := range matches {
		position := indices[i][0]

		// Determine instruction type based on the matched string
		if match[0] == "do()" {
			instructions = append(instructions, DoInstruction{Pos: position})
		} else if match[0] == "don't()" {
			instructions = append(instructions, DontInstruction{Pos: position})
		} else {
			x, _ := strconv.Atoi(match[1])
			y, _ := strconv.Atoi(match[2])
			instructions = append(instructions, MulInstruction{
				Pos: position,
				Mul: Mul{X: x, Y: y},
			})
		}
	}

	return instructions
}

// Part1 calculates the sum of all multiplication results
// Ignores do() and don't() instructions
func Part1(instructions []Instruction) int {
	sum := 0
	for _, instruction := range instructions {
		if mul, ok := instruction.(MulInstruction); ok {
			sum += mul.Mul.X * mul.Mul.Y
		}
	}
	return sum
}

// Part2 calculates the sum of enabled multiplication results
// mul instructions are enabled at start, and can be toggled by do() and don't()
func Part2(instructions []Instruction) int {
	sum := 0
	enabled := true // mul instructions are enabled at the beginning

	for _, instruction := range instructions {
		switch inst := instruction.(type) {
		case DoInstruction:
			enabled = true
		case DontInstruction:
			enabled = false
		case MulInstruction:
			if enabled {
				sum += inst.Mul.X * inst.Mul.Y
			}
		}
	}

	return sum
}
