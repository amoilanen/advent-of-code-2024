package day07

import "testing"

func TestPart1Example(t *testing.T) {
	equations := Parse(ExampleInput)
	result := Part1(equations)
	expected := 3749
	if result != expected {
		t.Errorf("Part1(ExampleInput) = %d; want %d", result, expected)
	}
}

func TestParse(t *testing.T) {
	equations := Parse(ExampleInput)

	expected := []Equation{
		{TestValue: 190, Numbers: []int{10, 19}},
		{TestValue: 3267, Numbers: []int{81, 40, 27}},
		{TestValue: 83, Numbers: []int{17, 5}},
	}

	for i, ex := range expected {
		if ex.TestValue != equations[i].TestValue {
			t.Errorf("Equation %d: expected test value %d, got %d", i, ex.TestValue, equations[i].TestValue)
		}

		if len(ex.Numbers) != len(equations[i].Numbers) {
			t.Errorf("Equation %d: expected %d numbers, got %d", i, len(ex.Numbers), len(equations[i].Numbers))
			continue
		}

		for j, num := range ex.Numbers {
			if num != equations[i].Numbers[j] {
				t.Errorf("Equation %d, number %d: expected %d, got %d", i, j, num, equations[i].Numbers[j])
			}
		}
	}
}

func TestEvaluate(t *testing.T) {
	tests := []struct {
		name      string
		numbers   []int
		operators []Operator
		expected  int
	}{
		{
			name:      "81 * 40 + 27",
			numbers:   []int{81, 40, 27},
			operators: []Operator{Multiply, Add},
			expected:  3267,
		},
		{
			name:      "11 + 6 * 16 + 20",
			numbers:   []int{11, 6, 16, 20},
			operators: []Operator{Add, Multiply, Add},
			expected:  292,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := evaluate(tt.numbers, tt.operators)
			if result != tt.expected {
				t.Errorf("evaluate(%v, %v) = %d; want %d", tt.numbers, tt.operators, result, tt.expected)
			}
		})
	}
}

func TestCanBeMadeTrue(t *testing.T) {
	tests := []struct {
		name     string
		equation Equation
		expected bool
	}{
		{
			name:     "190: 10 19 (can be made true)",
			equation: Equation{TestValue: 190, Numbers: []int{10, 19}},
			expected: true,
		},
		{
			name:     "3267: 81 40 27 (can be made true)",
			equation: Equation{TestValue: 3267, Numbers: []int{81, 40, 27}},
			expected: true,
		},
		{
			name:     "83: 17 5 (cannot be made true)",
			equation: Equation{TestValue: 83, Numbers: []int{17, 5}},
			expected: false,
		},
		{
			name:     "292: 11 6 16 20 (can be made true)",
			equation: Equation{TestValue: 292, Numbers: []int{11, 6, 16, 20}},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := canBeMadeTrue(tt.equation)
			if result != tt.expected {
				t.Errorf("canBeMadeTrue(%v) = %v; want %v", tt.equation, result, tt.expected)
			}
		})
	}
}

func TestConcatenation(t *testing.T) {
	tests := []struct {
		name      string
		numbers   []int
		operators []Operator
		expected  int
	}{
		{
			name:      "15 || 6",
			numbers:   []int{15, 6},
			operators: []Operator{Concatenate},
			expected:  156,
		},
		{
			name:      "6 * 8 || 6 * 15",
			numbers:   []int{6, 8, 6, 15},
			operators: []Operator{Multiply, Concatenate, Multiply},
			expected:  7290,
		},
		{
			name:      "17 || 8 + 14",
			numbers:   []int{17, 8, 14},
			operators: []Operator{Concatenate, Add},
			expected:  192,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := evaluate(tt.numbers, tt.operators)
			if result != tt.expected {
				t.Errorf("evaluate(%v, %v) = %d; want %d", tt.numbers, tt.operators, result, tt.expected)
			}
		})
	}
}

func TestPart2Example(t *testing.T) {
	equations := Parse(ExampleInput)
	result := Part2(equations)
	expected := 11387
	if result != expected {
		t.Errorf("Part2(ExampleInput) = %d; want %d", result, expected)
	}
}

func TestCanBeMadeTrueWithConcat(t *testing.T) {
	tests := []struct {
		name     string
		equation Equation
		expected bool
	}{
		{
			name:     "156: 15 6 (can with concat)",
			equation: Equation{TestValue: 156, Numbers: []int{15, 6}},
			expected: true,
		},
		{
			name:     "7290: 6 8 6 15 (can with concat)",
			equation: Equation{TestValue: 7290, Numbers: []int{6, 8, 6, 15}},
			expected: true,
		},
		{
			name:     "192: 17 8 14 (can with concat)",
			equation: Equation{TestValue: 192, Numbers: []int{17, 8, 14}},
			expected: true,
		},
		{
			name:     "83: 17 5 (still cannot)",
			equation: Equation{TestValue: 83, Numbers: []int{17, 5}},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := canBeMadeTrueWithConcat(tt.equation)
			if result != tt.expected {
				t.Errorf("canBeMadeTrueWithConcat(%v) = %v; want %v", tt.equation, result, tt.expected)
			}
		})
	}
}
