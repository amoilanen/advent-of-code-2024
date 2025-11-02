package day01

import (
	"sort"
	"strings"

	"github.com/amoilanen/advent-of-code-2024/internal/utils"
)

const ExampleInput = `3   4
4   3
2   5
1   3
3   9
3   3`

// LocationLists represents the two lists of location IDs
type LocationLists struct {
	Left  []int
	Right []int
}

// Parse parses the input into two lists of location IDs
func Parse(input string) LocationLists {
	lines := utils.AsLines(input)
	lists := LocationLists{
		Left:  make([]int, 0, len(lines)),
		Right: make([]int, 0, len(lines)),
	}

	for _, line := range lines {
		if line == "" {
			continue
		}
		// Split by whitespace and parse two numbers
		parts := strings.Fields(line)
		if len(parts) != 2 {
			continue
		}
		lists.Left = append(lists.Left, utils.MustParseInt(parts[0]))
		lists.Right = append(lists.Right, utils.MustParseInt(parts[1]))
	}

	return lists
}

// Part1 calculates the total distance between the two lists
// by pairing up sorted numbers and summing their distances
func Part1(lists LocationLists) int {
	// Make copies to avoid modifying the originals
	left := make([]int, len(lists.Left))
	right := make([]int, len(lists.Right))
	copy(left, lists.Left)
	copy(right, lists.Right)

	// Sort both lists
	sort.Ints(left)
	sort.Ints(right)

	// Calculate total distance
	totalDistance := 0
	for i := 0; i < len(left) && i < len(right); i++ {
		distance := utils.Abs(left[i] - right[i])
		totalDistance += distance
	}

	return totalDistance
}

// Part2 calculates the similarity score between the two lists
// For each number in the left list, multiply it by the number of times
// it appears in the right list, then sum all products
func Part2(lists LocationLists) int {
	// Build frequency map of right list
	rightFreq := make(map[int]int)
	for _, num := range lists.Right {
		rightFreq[num]++
	}

	// Calculate similarity score
	similarityScore := 0
	for _, num := range lists.Left {
		count := rightFreq[num]
		similarityScore += num * count
	}

	return similarityScore
}
