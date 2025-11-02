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

func TestPart2WithExample(t *testing.T) {
	instructions := Parse(ExampleInputPart2)
	result := Part2(instructions)
	expected := 48 // 2*4 + 8*5 = 8 + 40 = 48
	// mul(5,5) and mul(11,8) are disabled by don't()

	if result != expected {
		t.Errorf("Part2(ExampleInputPart2) = %d, want %d", result, expected)
	}
}

func TestParseInstructions(t *testing.T) {
	instructions := Parse(ExampleInputPart2)

	// Expected sequence: mul(2,4), don't(), mul(5,5), mul(11,8), do(), mul(8,5)
	if len(instructions) != 6 {
		t.Fatalf("ParseInstructions returned %d instructions, want 6", len(instructions))
	}

	// Check first instruction: mul(2,4)
	mul0, ok := instructions[0].(MulInstruction)
	if !ok {
		t.Errorf("Instruction 0: got type %T, want MulInstruction", instructions[0])
	} else if mul0.Mul.X != 2 || mul0.Mul.Y != 4 {
		t.Errorf("Instruction 0: got Mul{X: %d, Y: %d}, want Mul{X: 2, Y: 4}", mul0.Mul.X, mul0.Mul.Y)
	}

	// Check second instruction: don't()
	if _, ok := instructions[1].(DontInstruction); !ok {
		t.Errorf("Instruction 1: got type %T, want DontInstruction", instructions[1])
	}

	// Check fifth instruction: do()
	if _, ok := instructions[4].(DoInstruction); !ok {
		t.Errorf("Instruction 4: got type %T, want DoInstruction", instructions[4])
	}

	// Check last instruction: mul(8,5)
	mul5, ok := instructions[5].(MulInstruction)
	if !ok {
		t.Errorf("Instruction 5: got type %T, want MulInstruction", instructions[5])
	} else if mul5.Mul.X != 8 || mul5.Mul.Y != 5 {
		t.Errorf("Instruction 5: got Mul{X: %d, Y: %d}, want Mul{X: 8, Y: 5}", mul5.Mul.X, mul5.Mul.Y)
	}
}
