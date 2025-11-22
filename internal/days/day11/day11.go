package day11

import (
	"strconv"
	"strings"
)

const ExampleInput = `125 17`

// Parse converts the input string into a slice of stone values
func Parse(input string) []int {
	fields := strings.Fields(strings.TrimSpace(input))
	stones := make([]int, len(fields))
	for i, field := range fields {
		val, _ := strconv.Atoi(field)
		stones[i] = val
	}
	return stones
}

// countDigits returns the number of digits in a number
func countDigits(n int) int {
	if n == 0 {
		return 1
	}
	count := 0
	for n > 0 {
		count++
		n /= 10
	}
	return count
}

// splitNumber splits a number into left and right halves based on digit count
// For example: 1234 with 4 digits -> (12, 34)
func splitNumber(n int, digits int) (int, int) {
	// Calculate the divisor (10^(digits/2))
	divisor := 1
	for i := 0; i < digits/2; i++ {
		divisor *= 10
	}
	left := n / divisor
	right := n % divisor
	return left, right
}

// transformStone applies the transformation rules to a single stone value
// Returns the resulting stone value(s)
func transformStone(value int) []int {
	// Rule 1: If stone is 0, it becomes 1
	if value == 0 {
		return []int{1}
	}

	// Rule 2: If even number of digits, split into two stones
	digits := countDigits(value)
	if digits%2 == 0 {
		left, right := splitNumber(value, digits)
		return []int{left, right}
	}

	// Rule 3: Otherwise, multiply by 2024
	return []int{value * 2024}
}

func initialCounts(stones []int) map[int]int {
	stoneCount := make(map[int]int)
	for _, stone := range stones {
		stoneCount[stone]++
	}
	return stoneCount
}

func nextCounts(currentCounts map[int]int) map[int]int {
	newStoneCount := make(map[int]int)

	// Process each unique stone value
	for value, count := range currentCounts {
		// Transform the stone and add results to new map
		results := transformStone(value)
		for _, result := range results {
			newStoneCount[result] += count
		}
	}
	return newStoneCount
}

func countStones(currentCounts map[int]int) int {
	total := 0
	for _, count := range currentCounts {
		total += count
	}
	return total
}

// simulateBlinks simulates the stone transformations for a given number of blinks
// Uses a frequency map for efficiency: stone_value -> count
func simulateBlinks(stones []int, blinks int) int {
	stoneCount := initialCounts(stones)
	for i := 0; i < blinks; i++ {
		stoneCount = nextCounts(stoneCount)
	}
	return countStones(stoneCount)
}

// Part1 solves part 1: count stones after 25 blinks
// Algorithm:
// 1. Use frequency map to track unique stone values and their counts
// 2. For each blink, transform each unique value and update counts
// 3. This avoids storing duplicate stones and is much more efficient
//
// Time complexity: O(B × U) where B=blinks, U=unique stone values
// Space complexity: O(U)
func Part1(stones []int) int {
	return simulateBlinks(stones, 25)
}

// Part2 solves part 2: count stones after 75 blinks
// The same efficient frequency map algorithm works perfectly for 75 blinks
// because we only track unique values, not individual stones
//
// Time complexity: O(B × U) where B=blinks, U=unique stone values
// Space complexity: O(U)
func Part2(stones []int) int {
	return simulateBlinks(stones, 75)
}
