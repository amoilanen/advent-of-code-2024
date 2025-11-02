package day02

import (
	"github.com/amoilanen/advent-of-code-2024/internal/utils"
)

const ExampleInput = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

// Report represents a list of levels (numbers) in a single report
type Report []int

// Parse parses the input into a slice of reports
// Each line represents one report with space-separated levels
func Parse(input string) []Report {
	lines := utils.AsLines(input)
	reports := make([]Report, 0, len(lines))

	for _, line := range lines {
		if line == "" {
			continue
		}
		levels := utils.MustParseInts(line)
		if len(levels) > 0 {
			reports = append(reports, Report(levels))
		}
	}

	return reports
}

// isSafe checks if a report is safe according to the rules:
// - All levels must be either increasing or decreasing
// - Any two adjacent levels must differ by at least 1 and at most 3
func isSafe(report Report) bool {
	if len(report) < 2 {
		return true
	}

	// Determine if we should be increasing or decreasing based on first pair
	firstDiff := report[1] - report[0]
	if firstDiff == 0 {
		return false // No change is not allowed
	}

	isIncreasing := firstDiff > 0

	// Check all adjacent pairs
	for i := 0; i < len(report)-1; i++ {
		diff := report[i+1] - report[i]
		absDiff := utils.Abs(diff)

		// Check if difference is within range [1, 3]
		if absDiff < 1 || absDiff > 3 {
			return false
		}

		// Check if direction is consistent
		if isIncreasing && diff < 0 {
			return false
		}
		if !isIncreasing && diff > 0 {
			return false
		}
	}

	return true
}

// Part1 counts how many reports are safe
func Part1(reports []Report) int {
	safeCount := 0
	for _, report := range reports {
		if isSafe(report) {
			safeCount++
		}
	}
	return safeCount
}
