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
// skipIndices: indices to skip when checking (for Problem Dampener); pass empty slice or nil for no skipping
func isSafe(report Report, skipIndices []int) bool {
	// Build list of valid indices (excluding skipIndices)
	validIndices := buildValidIndices(report, skipIndices)

	if len(validIndices) < 2 {
		return true
	}

	// Determine if we should be increasing or decreasing based on first valid pair
	firstDiff := report[validIndices[1]] - report[validIndices[0]]
	if firstDiff == 0 {
		return false // No change is not allowed
	}

	isIncreasing := firstDiff > 0

	// Check all adjacent pairs of valid indices
	for i := 0; i < len(validIndices)-1; i++ {
		curr := validIndices[i]
		next := validIndices[i+1]

		diff := report[next] - report[curr]
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

// buildValidIndices creates a list of valid indices, excluding those in skipIndices
// Optimized for the common case where skipIndices is empty
func buildValidIndices(report Report, skipIndices []int) []int {
	// Optimize for the common case: no skipping
	if len(skipIndices) == 0 {
		// Create a simple sequential list
		validIndices := make([]int, len(report))
		for i := range report {
			validIndices[i] = i
		}
		return validIndices
	}

	// Build a map of indices to skip for O(1) lookup
	skipMap := make(map[int]bool, len(skipIndices))
	for _, idx := range skipIndices {
		skipMap[idx] = true
	}

	// Build list of valid indices (excluding skipIndices)
	validIndices := make([]int, 0, len(report)-len(skipIndices))
	for i := 0; i < len(report); i++ {
		if !skipMap[i] {
			validIndices = append(validIndices, i)
		}
	}

	return validIndices
}

// isSafeWithDampener checks if a report is safe, either as-is or by removing a single level
func isSafeWithDampener(report Report) bool {
	// First check if it's already safe (no skipping)
	if isSafe(report, nil) {
		return true
	}

	// Try skipping each level one at a time
	for i := 0; i < len(report); i++ {
		if isSafe(report, []int{i}) {
			return true
		}
	}

	return false
}

// Part1 counts how many reports are safe
func Part1(reports []Report) int {
	safeCount := 0
	for _, report := range reports {
		if isSafe(report, nil) {
			safeCount++
		}
	}
	return safeCount
}

// Part2 counts how many reports are safe with the Problem Dampener
// The Problem Dampener allows removing a single level to make an unsafe report safe
func Part2(reports []Report) int {
	safeCount := 0
	for _, report := range reports {
		if isSafeWithDampener(report) {
			safeCount++
		}
	}
	return safeCount
}
