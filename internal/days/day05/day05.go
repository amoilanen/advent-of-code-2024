package day05

import (
	"strconv"
	"strings"
)

const ExampleInput = `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

// OrderingRule represents a page ordering rule (before|after)
type OrderingRule struct {
	Before int
	After  int
}

// IntSet represents a set of integers (Go's idiomatic set using map with empty struct)
type IntSet map[int]struct{}

// RuleSet is an efficient structure for checking ordering rules
// It maps each page to the set of pages that must come after it
type RuleSet struct {
	mustComeAfter map[int]IntSet
}

// newRuleSet builds an efficient rule lookup structure from ordering rules
func newRuleSet(rules []OrderingRule) *RuleSet {
	rs := &RuleSet{
		mustComeAfter: make(map[int]IntSet),
	}

	for _, rule := range rules {
		if rs.mustComeAfter[rule.Before] == nil {
			rs.mustComeAfter[rule.Before] = make(IntSet)
		}
		rs.mustComeAfter[rule.Before][rule.After] = struct{}{}
	}

	return rs
}

// Update represents a list of page numbers to print
type Update []int

// Input contains the parsed rules and updates
type Input struct {
	Rules   []OrderingRule
	RuleSet *RuleSet
	Updates []Update
}

// Parse parses the input into rules and updates
func Parse(input string) Input {
	sections := strings.Split(strings.TrimSpace(input), "\n\n")
	if len(sections) != 2 {
		return Input{}
	}

	// Parse rules
	ruleLines := strings.Split(strings.TrimSpace(sections[0]), "\n")
	rules := make([]OrderingRule, 0, len(ruleLines))
	for _, line := range ruleLines {
		parts := strings.Split(line, "|")
		if len(parts) != 2 {
			continue
		}
		before, err1 := strconv.Atoi(parts[0])
		after, err2 := strconv.Atoi(parts[1])
		if err1 == nil && err2 == nil {
			rules = append(rules, OrderingRule{Before: before, After: after})
		}
	}

	// Parse updates
	updateLines := strings.Split(strings.TrimSpace(sections[1]), "\n")
	updates := make([]Update, 0, len(updateLines))
	for _, line := range updateLines {
		parts := strings.Split(line, ",")
		update := make(Update, 0, len(parts))
		for _, part := range parts {
			num, err := strconv.Atoi(strings.TrimSpace(part))
			if err == nil {
				update = append(update, num)
			}
		}
		if len(update) > 0 {
			updates = append(updates, update)
		}
	}

	// Build efficient rule set
	ruleSet := newRuleSet(rules)

	return Input{Rules: rules, RuleSet: ruleSet, Updates: updates}
}

// isValid checks if an update follows all applicable ordering rules using a RuleSet
// Algorithm: For each page, check if any preceding page should actually come after it
func (u Update) isValid(ruleSet *RuleSet) bool {
	// Iterate through each page in the update
	for i, currentPage := range u {
		// Get the set of pages that must come after the current page
		if mustFollowCurrent, exists := ruleSet.mustComeAfter[currentPage]; exists {
			// Check all pages that came before the current page
			for j := 0; j < i; j++ {
				precedingPage := u[j]
				// If a preceding page should actually come after current page, it's a violation
				if _, shouldFollow := mustFollowCurrent[precedingPage]; shouldFollow {
					return false
				}
			}
		}
	}

	return true
}

// middlePage returns the middle page number of an update
func (u Update) middlePage() int {
	if len(u) == 0 {
		return 0
	}
	return u[len(u)/2]
}

// reorder sorts the update pages according to the ordering rules
// Uses a comparison function based on the RuleSet
func (u Update) reorder(ruleSet *RuleSet) Update {
	// Create a copy to avoid modifying the original
	reordered := make(Update, len(u))
	copy(reordered, u)

	// Sort based on ordering rules
	// We use a simple bubble sort to ensure stability and correctness
	// For each pair of adjacent elements, check if they violate any rule
	changed := true
	for changed {
		changed = false
		for i := 0; i < len(reordered)-1; i++ {
			pageI := reordered[i]
			pageJ := reordered[i+1]

			// Check if pageI must come after pageJ (wrong order)
			if mustFollowJ, exists := ruleSet.mustComeAfter[pageJ]; exists {
				if _, shouldFollow := mustFollowJ[pageI]; shouldFollow {
					// Swap them
					reordered[i], reordered[i+1] = reordered[i+1], reordered[i]
					changed = true
				}
			}
		}
	}

	return reordered
}

// Part1 finds the sum of middle page numbers from correctly-ordered updates
func Part1(input Input) int {
	sum := 0

	for _, update := range input.Updates {
		if update.isValid(input.RuleSet) {
			sum += update.middlePage()
		}
	}

	return sum
}

// Part2 finds the sum of middle page numbers after reordering incorrectly-ordered updates
func Part2(input Input) int {
	sum := 0

	for _, update := range input.Updates {
		if !update.isValid(input.RuleSet) {
			// Update is incorrect, reorder it
			corrected := update.reorder(input.RuleSet)
			sum += corrected.middlePage()
		}
	}

	return sum
}
