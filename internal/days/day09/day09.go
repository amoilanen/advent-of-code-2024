package day09

import (
	"strings"
)

const ExampleInput = `2333133121414131402`

// DiskMap represents the parsed disk structure
type DiskMap struct {
	Blocks []int // -1 represents free space, >= 0 represents file ID
}

// Parse converts the compact disk map string into an expanded block representation
// The input alternates between file lengths and free space lengths
// Example: "12345" -> file(1 block, ID=0), free(2), file(3, ID=1), free(4), file(5, ID=2)
func Parse(input string) DiskMap {
	input = strings.TrimSpace(input)
	blocks := []int{}
	fileID := 0

	for i, ch := range input {
		length := int(ch - '0')

		if i%2 == 0 {
			// Even index: file blocks
			for j := 0; j < length; j++ {
				blocks = append(blocks, fileID)
			}
			fileID++
		} else {
			// Odd index: free space blocks
			for j := 0; j < length; j++ {
				blocks = append(blocks, -1)
			}
		}
	}

	return DiskMap{Blocks: blocks}
}

// Compact performs disk compaction by moving file blocks from the end
// to the leftmost free space positions until no gaps remain between files
//
// Algorithm:
// 1. Use two pointers: left (scans for free space) and right (scans for file blocks)
// 2. Find next free space from left
// 3. Find next file block from right
// 4. Swap them
// 5. Repeat until pointers meet
//
// Time complexity: O(n) where n is the number of blocks
func (dm *DiskMap) Compact() {
	left := 0
	right := len(dm.Blocks) - 1

	for left < right {
		// Find next free space from left
		for left < right && dm.Blocks[left] != -1 {
			left++
		}

		// Find next file block from right
		for left < right && dm.Blocks[right] == -1 {
			right--
		}

		// Swap if both pointers are valid and haven't crossed
		if left < right {
			dm.Blocks[left], dm.Blocks[right] = dm.Blocks[right], dm.Blocks[left]
			left++
			right--
		}
	}
}

// Checksum calculates the filesystem checksum
// For each block position, multiply the position by the file ID
// Free space blocks (ID = -1) are skipped
//
// Formula: sum of (position × file_ID) for all file blocks
// Time complexity: O(n)
func (dm DiskMap) Checksum() int {
	checksum := 0
	for pos, fileID := range dm.Blocks {
		if fileID != -1 {
			checksum += pos * fileID
		}
	}
	return checksum
}

// Part1 solves part 1: compact the disk and calculate checksum
func Part1(diskMap DiskMap) int {
	// Make a copy to avoid modifying the original
	dm := DiskMap{Blocks: make([]int, len(diskMap.Blocks))}
	copy(dm.Blocks, diskMap.Blocks)

	// Compact the disk
	dm.Compact()

	// Calculate and return checksum
	return dm.Checksum()
}

// Part2 placeholder for part 2
func Part2(diskMap DiskMap) int {
	// To be implemented
	return 0
}
