package day01

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	input := `3   4
4   3
2   5
1   3
3   9
3   3`

	want := LocationLists{
		Left:  []int{3, 4, 2, 1, 3, 3},
		Right: []int{4, 3, 5, 3, 9, 3},
	}

	got := Parse(input)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Parse() = %v, want %v", got, want)
	}
}

func TestParseEmptyInput(t *testing.T) {
	input := ""
	got := Parse(input)
	if len(got.Left) != 0 || len(got.Right) != 0 {
		t.Errorf("Parse() = %v, want empty lists", got)
	}
}

func TestParseSingleLine(t *testing.T) {
	input := `10   20`
	want := LocationLists{
		Left:  []int{10},
		Right: []int{20},
	}
	got := Parse(input)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Parse() = %v, want %v", got, want)
	}
}

func TestParseWithVariableSpacing(t *testing.T) {
	input := `1    2
3  4
5      6`
	want := LocationLists{
		Left:  []int{1, 3, 5},
		Right: []int{2, 4, 6},
	}
	got := Parse(input)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Parse() = %v, want %v", got, want)
	}
}

func TestPart1(t *testing.T) {
	tests := []struct {
		name  string
		lists LocationLists
		want  int
	}{
		{
			name: "example from problem description",
			lists: LocationLists{
				Left:  []int{3, 4, 2, 1, 3, 3},
				Right: []int{4, 3, 5, 3, 9, 3},
			},
			want: 11, // 2 + 1 + 0 + 1 + 2 + 5 = 11
		},
		{
			name: "identical lists",
			lists: LocationLists{
				Left:  []int{1, 2, 3},
				Right: []int{1, 2, 3},
			},
			want: 0,
		},
		{
			name: "single pair",
			lists: LocationLists{
				Left:  []int{5},
				Right: []int{10},
			},
			want: 5,
		},
		{
			name: "all same distance",
			lists: LocationLists{
				Left:  []int{1, 2, 3},
				Right: []int{2, 3, 4},
			},
			want: 3, // 1 + 1 + 1
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part1(tt.lists); got != tt.want {
				t.Errorf("Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithExampleInput(t *testing.T) {
	parsed := Parse(ExampleInput)

	t.Run("Part1 with example input", func(t *testing.T) {
		want := 11
		if got := Part1(parsed); got != want {
			t.Errorf("Part1() = %v, want %v", got, want)
		}
	})

	t.Run("Part2 with example input", func(t *testing.T) {
		// Part 2 not yet implemented
		want := 0
		if got := Part2(parsed); got != want {
			t.Errorf("Part2() = %v, want %v", got, want)
		}
	})
}

// BenchmarkPart1 benchmarks the Part1 solution
func BenchmarkPart1(b *testing.B) {
	parsed := Parse(ExampleInput)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Part1(parsed)
	}
}

// BenchmarkPart2 benchmarks the Part2 solution
func BenchmarkPart2(b *testing.B) {
	parsed := Parse(ExampleInput)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Part2(parsed)
	}
}
