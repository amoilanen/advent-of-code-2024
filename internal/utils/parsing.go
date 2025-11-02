package utils

import (
	"strconv"
	"strings"
)

// AsLines splits input into lines and trims whitespace
func AsLines(input string) []string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	result := make([]string, 0, len(lines))
	for _, line := range lines {
		result = append(result, strings.TrimSpace(line))
	}
	return result
}

// ParseInts parses space-separated integers from a string
func ParseInts(input string) ([]int, error) {
	fields := strings.Fields(input)
	nums := make([]int, 0, len(fields))
	for _, field := range fields {
		num, err := strconv.Atoi(field)
		if err != nil {
			return nil, err
		}
		nums = append(nums, num)
	}
	return nums, nil
}

// MustParseInts parses space-separated integers, panics on error
func MustParseInts(input string) []int {
	nums, err := ParseInts(input)
	if err != nil {
		panic(err)
	}
	return nums
}

// ParseInt parses a single integer
func ParseInt(s string) (int, error) {
	return strconv.Atoi(strings.TrimSpace(s))
}

// MustParseInt parses a single integer, panics on error
func MustParseInt(s string) int {
	num, err := ParseInt(s)
	if err != nil {
		panic(err)
	}
	return num
}
