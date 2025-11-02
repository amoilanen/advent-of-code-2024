package day03

import (
	"testing"
)

func TestPart1WithExample(t *testing.T) {
	instructions := Parse(ExampleInput)
	result := Part1(instructions)
	expected := 161 // 2*4 + 5*5 + 11*8 + 8*5 = 8 + 25 + 88 + 40 = 161

	if result != expected {
		t.Errorf("Part1(ExampleInput) = %d, want %d", result, expected)
	}
}

func TestParse(t *testing.T) {
	instructions := Parse(ExampleInput)

	// Expected: mul(2,4), mul(5,5), mul(11,8), mul(8,5)
	expected := []Mul{
		{X: 2, Y: 4},
		{X: 5, Y: 5},
		{X: 11, Y: 8},
		{X: 8, Y: 5},
	}

	if len(instructions) != len(expected) {
		t.Fatalf("Parse returned %d instructions, want %d", len(instructions), len(expected))
	}

	for i, inst := range instructions {
		if inst.X != expected[i].X || inst.Y != expected[i].Y {
			t.Errorf("Instruction %d: got Mul{X: %d, Y: %d}, want Mul{X: %d, Y: %d}",
				i, inst.X, inst.Y, expected[i].X, expected[i].Y)
		}
	}
}
