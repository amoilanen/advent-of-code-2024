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

// findFile locates a file by ID and returns its start position and length
// Returns (-1, 0) if file not found
func (dm DiskMap) findFile(fileID int) (startPos, length int) {
	startPos = -1
	length = 0

	for i, id := range dm.Blocks {
		if id == fileID {
			if startPos == -1 {
				startPos = i
			}
			length++
		} else if startPos != -1 {
			// File blocks are contiguous, so we're done
			break
		}
	}

	return startPos, length
}

// findFreeSpan finds the leftmost free space span that can fit targetLength blocks
// Only searches up to maxPos (exclusive)
// Returns the start position of the free span, or -1 if not found
func (dm DiskMap) findFreeSpan(targetLength, maxPos int) int {
	i := 0
	for i < maxPos {
		// Skip non-free blocks
		if dm.Blocks[i] != -1 {
			i++
			continue
		}

		// Found start of free space, count consecutive free blocks
		spanStart := i
		spanLength := 0
		for i < maxPos && dm.Blocks[i] == -1 {
			spanLength++
			i++
		}

		// Check if this span is large enough
		if spanLength >= targetLength {
			return spanStart
		}
	}

	return -1 // No suitable span found
}

// moveFile moves a file from one position to another
// Copies file blocks to new position and marks old position as free
func (dm *DiskMap) moveFile(fileID, fromPos, toPos, length int) {
	// Copy file blocks to new position
	for i := 0; i < length; i++ {
		dm.Blocks[toPos+i] = fileID
	}

	// Mark old position as free
	for i := 0; i < length; i++ {
		dm.Blocks[fromPos+i] = -1
	}
}

// CompactWholeFiles performs whole-file defragmentation
// Processes files in decreasing file ID order
// Each file is moved at most once to the leftmost suitable free space
//
// Algorithm:
//  1. Find maximum file ID
//  2. For each file from max ID down to 0:
//     a. Locate the file (start position and length)
//     b. Find leftmost free span that can fit the file (must be to the left)
//     c. If found, move the entire file
//     d. If not found, file stays in place
//
// Time complexity: O(n² × m) worst case, where n = number of files, m = total blocks
func (dm *DiskMap) CompactWholeFiles() {
	// Find maximum file ID
	maxFileID := 0
	for _, id := range dm.Blocks {
		if id > maxFileID {
			maxFileID = id
		}
	}

	// Process files in decreasing ID order
	for fileID := maxFileID; fileID >= 0; fileID-- {
		// Find where this file currently is
		startPos, length := dm.findFile(fileID)
		if startPos == -1 || length == 0 {
			continue // File not found or empty
		}

		// Find leftmost free span that can fit this file
		// Only search to the left of the file's current position
		freePos := dm.findFreeSpan(length, startPos)
		if freePos == -1 {
			continue // No suitable free space found
		}

		// Move the file to the free space
		dm.moveFile(fileID, startPos, freePos, length)
	}
}

// Part2 solves part 2: compact whole files and calculate checksum
func Part2(diskMap DiskMap) int {
	// Make a copy to avoid modifying the original
	dm := DiskMap{Blocks: make([]int, len(diskMap.Blocks))}
	copy(dm.Blocks, diskMap.Blocks)

	// Compact using whole-file defragmentation
	dm.CompactWholeFiles()

	// Calculate and return checksum
	return dm.Checksum()
}
