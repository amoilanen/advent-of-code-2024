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

// Part2 solves part 2 of the day's problem
// (Placeholder until part 2 is revealed)
func Part2(lists LocationLists) int {
	// Part 2 will be implemented once the problem is revealed
	return 0
}
