package day09

import (
	"testing"
)

func TestParse(t *testing.T) {
	diskMap := Parse("12345")
	expected := []int{0, -1, -1, 1, 1, 1, -1, -1, -1, -1, 2, 2, 2, 2, 2}

	if len(diskMap.Blocks) != len(expected) {
		t.Fatalf("Parse() length = %d; want %d", len(diskMap.Blocks), len(expected))
	}

	for i, val := range diskMap.Blocks {
		if val != expected[i] {
			t.Errorf("Parse() block[%d] = %d; want %d", i, val, expected[i])
		}
	}
}

func TestCompact(t *testing.T) {
	diskMap := Parse("12345")
	diskMap.Compact()

	// After compaction: 022111222
	expected := []int{0, 2, 2, 1, 1, 1, 2, 2, 2, -1, -1, -1, -1, -1, -1}

	if len(diskMap.Blocks) != len(expected) {
		t.Fatalf("Compact() length = %d; want %d", len(diskMap.Blocks), len(expected))
	}

	for i, val := range diskMap.Blocks {
		if val != expected[i] {
			t.Errorf("Compact() block[%d] = %d; want %d", i, val, expected[i])
		}
	}
}

func TestPart1(t *testing.T) {
	diskMap := Parse(ExampleInput)
	result := Part1(diskMap)
	expected := 1928

	if result != expected {
		t.Errorf("Part1() = %d; want %d", result, expected)
	}
}

func TestPart2(t *testing.T) {
	diskMap := Parse(ExampleInput)
	result := Part2(diskMap)
	expected := 2858

	if result != expected {
		t.Errorf("Part2() = %d; want %d", result, expected)
	}
}
