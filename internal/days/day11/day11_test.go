package day11

import (
	"testing"
)

func TestPart1Example(t *testing.T) {
	stones := Parse(ExampleInput)
	result := Part1(stones)
	expected := 55312

	if result != expected {
		t.Errorf("Part1(ExampleInput) = %d; expected %d", result, expected)
	}
}

func TestPart2Example(t *testing.T) {
	stones := Parse(ExampleInput)
	result := Part2(stones)
	// We don't have the expected value from the problem description,
	// but we can verify it runs without error and produces a reasonable result
	if result <= 0 {
		t.Errorf("Part2(ExampleInput) = %d; expected positive value", result)
	}
	// The result should be larger than Part1
	part1Result := Part1(stones)
	if result <= part1Result {
		t.Errorf("Part2(ExampleInput) = %d; expected > Part1 result %d", result, part1Result)
	}
}

func TestTransformStone(t *testing.T) {
	tests := []struct {
		input    int
		expected []int
	}{
		{0, []int{1}},                // Rule 1: 0 -> 1
		{1, []int{2024}},             // Rule 3: odd digits, multiply by 2024
		{10, []int{1, 0}},            // Rule 2: 2 digits, split
		{99, []int{9, 9}},            // Rule 2: 2 digits, split
		{999, []int{2021976}},        // Rule 3: 3 digits (odd), multiply
		{1000, []int{10, 0}},         // Rule 2: 4 digits, split (no leading zeros)
		{125, []int{253000}},         // Rule 3: 3 digits (odd)
		{17, []int{1, 7}},            // Rule 2: 2 digits
		{2024, []int{20, 24}},        // Rule 2: 4 digits
		{14168, []int{28676032}},     // Rule 3: 5 digits (odd)
	}

	for _, test := range tests {
		result := transformStone(test.input)
		if len(result) != len(test.expected) {
			t.Errorf("transformStone(%d) returned %d values; expected %d",
				test.input, len(result), len(test.expected))
			continue
		}
		for i := range result {
			if result[i] != test.expected[i] {
				t.Errorf("transformStone(%d) = %v; expected %v",
					test.input, result, test.expected)
				break
			}
		}
	}
}

func TestSimulateBlinksSmall(t *testing.T) {
	tests := []struct {
		input    []int
		blinks   int
		expected int
	}{
		{[]int{0, 1, 10, 99, 999}, 1, 7},    // First example from problem
		{[]int{125, 17}, 1, 3},               // After 1 blink: 253000 1 7
		{[]int{125, 17}, 2, 4},               // After 2 blinks: 253 0 2024 14168
		{[]int{125, 17}, 3, 5},               // After 3 blinks
		{[]int{125, 17}, 4, 9},               // After 4 blinks
		{[]int{125, 17}, 5, 13},              // After 5 blinks
		{[]int{125, 17}, 6, 22},              // After 6 blinks: 22 stones
	}

	for _, test := range tests {
		result := simulateBlinks(test.input, test.blinks)
		if result != test.expected {
			t.Errorf("simulateBlinks(%v, %d) = %d; expected %d",
				test.input, test.blinks, result, test.expected)
		}
	}
}

func TestCountDigits(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{0, 1},
		{1, 1},
		{9, 1},
		{10, 2},
		{99, 2},
		{100, 3},
		{999, 3},
		{1000, 4},
		{9999, 4},
		{10000, 5},
	}

	for _, test := range tests {
		result := countDigits(test.input)
		if result != test.expected {
			t.Errorf("countDigits(%d) = %d; expected %d",
				test.input, result, test.expected)
		}
	}
}

func TestSplitNumber(t *testing.T) {
	tests := []struct {
		value    int
		digits   int
		left     int
		right    int
	}{
		{10, 2, 1, 0},
		{99, 2, 9, 9},
		{1000, 4, 10, 0},
		{1234, 4, 12, 34},
		{2024, 4, 20, 24},
	}

	for _, test := range tests {
		left, right := splitNumber(test.value, test.digits)
		if left != test.left || right != test.right {
			t.Errorf("splitNumber(%d, %d) = (%d, %d); expected (%d, %d)",
				test.value, test.digits, left, right, test.left, test.right)
		}
	}
}
