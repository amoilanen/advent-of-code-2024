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

func TestCompactWholeFiles(t *testing.T) {
	diskMap := Parse("12345")
	diskMap.CompactWholeFiles()

	// After whole-file compaction: 0..111....22222
	// File 2 (size 5) cannot move left (no space)
	// File 1 (size 3) moves to position 1-3
	// File 0 (size 1) stays at position 0
	// Result: 0..111....22222
	expected := []int{0, -1, -1, 1, 1, 1, -1, -1, -1, -1, 2, 2, 2, 2, 2}

	if len(diskMap.Blocks) != len(expected) {
		t.Fatalf("CompactWholeFiles() length = %d; want %d", len(diskMap.Blocks), len(expected))
	}

	for i, val := range diskMap.Blocks {
		if val != expected[i] {
			t.Errorf("CompactWholeFiles() block[%d] = %d; want %d", i, val, expected[i])
		}
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

func TestPart2_SmallExample(t *testing.T) {
	// "12345" -> 0..111....22222
	// After whole-file compaction, file 1 can't move (no space of size 3 to the left)
	// File 2 can't move (no space of size 5 to the left)
	// Final: 0..111....22222
	diskMap := Parse("12345")
	result := Part2(diskMap)
	// Checksum: 0*0 + 3*1 + 4*1 + 5*1 + 10*2 + 11*2 + 12*2 + 13*2 + 14*2 = 0 + 3 + 4 + 5 + 20 + 22 + 24 + 26 + 28 = 132
	expected := 132

	if result != expected {
		t.Errorf("Part2('12345') = %d; want %d", result, expected)
	}
}

func TestPart2_FileCanMove(t *testing.T) {
	// "1313" -> 0...1...
	// File 1 (size 1) can move to position 1
	// Final: 01......
	diskMap := Parse("1313")
	result := Part2(diskMap)
	// Checksum: 0*0 + 1*1 = 1
	expected := 1

	if result != expected {
		t.Errorf("Part2('1313') = %d; want %d", result, expected)
	}
}

func TestPart2_MultipleFilesMoveLeft(t *testing.T) {
	// "131213" -> 0...1..2...
	// File 2 (size 1) can move to position 1
	// File 1 (size 1) can move to position 2
	// Final: 021........
	diskMap := Parse("131213")
	result := Part2(diskMap)
	// After compaction: 021........
	// Checksum: 0*0 + 1*2 + 2*1 = 0 + 2 + 2 = 4
	expected := 4

	if result != expected {
		t.Errorf("Part2('131213') = %d; want %d", result, expected)
	}
}

func TestPart2_NoFragmentation(t *testing.T) {
	// "101010" -> 012 (no free space)
	// Files are packed together, no moves possible
	diskMap := Parse("101010")
	result := Part2(diskMap)
	// Checksum: 0*0 + 1*1 + 2*2 = 0 + 1 + 4 = 5
	expected := 5

	if result != expected {
		t.Errorf("Part2('101010') = %d; want %d", result, expected)
	}
}
