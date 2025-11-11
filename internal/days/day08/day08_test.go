package day08

import (
	"testing"
)

func TestPart1(t *testing.T) {
	grid := Parse(ExampleInput)
	result := Part1(grid)
	expected := 14

	if result != expected {
		t.Errorf("Part1() = %d; want %d", result, expected)
	}
}

func TestPart2(t *testing.T) {
	grid := Parse(ExampleInput)
	result := Part2(grid)
	expected := 34

	if result != expected {
		t.Errorf("Part2() = %d; want %d", result, expected)
	}
}
