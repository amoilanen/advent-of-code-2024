package day05

import "testing"

func TestPart1Example(t *testing.T) {
	input := Parse(ExampleInput)
	result := Part1(input)
	expected := 143
	if result != expected {
		t.Errorf("Part1(ExampleInput) = %d; want %d", result, expected)
	}
}

func TestPart2Example(t *testing.T) {
	input := Parse(ExampleInput)
	result := Part2(input)
	expected := 123
	if result != expected {
		t.Errorf("Part2(ExampleInput) = %d; want %d", result, expected)
	}
}

func TestParse(t *testing.T) {
	input := Parse(ExampleInput)

	if len(input.Rules) != 21 {
		t.Errorf("Expected 21 rules, got %d", len(input.Rules))
	}

	if len(input.Updates) != 6 {
		t.Errorf("Expected 6 updates, got %d", len(input.Updates))
	}

	// Check first rule
	if input.Rules[0].Before != 47 || input.Rules[0].After != 53 {
		t.Errorf("First rule should be 47|53, got %d|%d", input.Rules[0].Before, input.Rules[0].After)
	}

	// Check first update
	if len(input.Updates[0]) != 5 {
		t.Errorf("First update should have 5 pages, got %d", len(input.Updates[0]))
	}
}

func TestUpdateIsValid(t *testing.T) {
	input := Parse(ExampleInput)

	tests := []struct {
		name      string
		updateIdx int
		expected  bool
	}{
		{"First update (75,47,61,53,29)", 0, true},
		{"Second update (97,61,53,29,13)", 1, true},
		{"Third update (75,29,13)", 2, true},
		{"Fourth update (75,97,47,61,53)", 3, false},
		{"Fifth update (61,13,29)", 4, false},
		{"Sixth update (97,13,75,29,47)", 5, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := input.Updates[tt.updateIdx].isValid(input.RuleSet)
			if result != tt.expected {
				t.Errorf("Update %v isValid = %v; want %v",
					input.Updates[tt.updateIdx], result, tt.expected)
			}
		})
	}
}

func TestMiddlePage(t *testing.T) {
	tests := []struct {
		name     string
		update   Update
		expected int
	}{
		{"Update with 5 pages", Update{75, 47, 61, 53, 29}, 61},
		{"Update with 5 pages", Update{97, 61, 53, 29, 13}, 53},
		{"Update with 3 pages", Update{75, 29, 13}, 29},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.update.middlePage()
			if result != tt.expected {
				t.Errorf("middlePage(%v) = %d; want %d", tt.update, result, tt.expected)
			}
		})
	}
}

func TestReorder(t *testing.T) {
	input := Parse(ExampleInput)

	tests := []struct {
		name     string
		original Update
		expected Update
	}{
		{
			name:     "Fourth update (75,97,47,61,53)",
			original: Update{75, 97, 47, 61, 53},
			expected: Update{97, 75, 47, 61, 53},
		},
		{
			name:     "Fifth update (61,13,29)",
			original: Update{61, 13, 29},
			expected: Update{61, 29, 13},
		},
		{
			name:     "Sixth update (97,13,75,29,47)",
			original: Update{97, 13, 75, 29, 47},
			expected: Update{97, 75, 47, 29, 13},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.original.reorder(input.RuleSet)
			if len(result) != len(tt.expected) {
				t.Errorf("reorder(%v) length = %d; want %d", tt.original, len(result), len(tt.expected))
				return
			}
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("reorder(%v) = %v; want %v", tt.original, result, tt.expected)
					break
				}
			}
		})
	}
}
