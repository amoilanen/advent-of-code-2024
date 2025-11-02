package day02

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	input := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1`

	want := []Report{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
	}

	got := Parse(input)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Parse() = %v, want %v", got, want)
	}
}

func TestParseEmptyInput(t *testing.T) {
	input := ""
	got := Parse(input)
	if len(got) != 0 {
		t.Errorf("Parse() = %v, want empty slice", got)
	}
}

func TestParseSingleLine(t *testing.T) {
	input := `1 2 3 4 5`
	want := []Report{{1, 2, 3, 4, 5}}
	got := Parse(input)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Parse() = %v, want %v", got, want)
	}
}

func TestIsSafe(t *testing.T) {
	tests := []struct {
		name   string
		report Report
		want   bool
		reason string
	}{
		{
			name:   "safe decreasing by 1 or 2",
			report: Report{7, 6, 4, 2, 1},
			want:   true,
			reason: "levels are all decreasing by 1 or 2",
		},
		{
			name:   "unsafe increase of 5",
			report: Report{1, 2, 7, 8, 9},
			want:   false,
			reason: "2 to 7 is an increase of 5",
		},
		{
			name:   "unsafe decrease of 4",
			report: Report{9, 7, 6, 2, 1},
			want:   false,
			reason: "6 to 2 is a decrease of 4",
		},
		{
			name:   "unsafe mixed increase and decrease",
			report: Report{1, 3, 2, 4, 5},
			want:   false,
			reason: "1 to 3 is increasing but 3 to 2 is decreasing",
		},
		{
			name:   "unsafe no change",
			report: Report{8, 6, 4, 4, 1},
			want:   false,
			reason: "4 to 4 is neither an increase nor a decrease",
		},
		{
			name:   "safe increasing by 1, 2, or 3",
			report: Report{1, 3, 6, 7, 9},
			want:   true,
			reason: "levels are all increasing by 1, 2, or 3",
		},
		{
			name:   "safe single level",
			report: Report{5},
			want:   true,
			reason: "single level is safe",
		},
		{
			name:   "safe two levels increasing",
			report: Report{1, 2},
			want:   true,
			reason: "two levels with valid increase",
		},
		{
			name:   "safe two levels decreasing",
			report: Report{5, 3},
			want:   true,
			reason: "two levels with valid decrease",
		},
		{
			name:   "unsafe two levels no change",
			report: Report{5, 5},
			want:   false,
			reason: "no change between levels",
		},
		{
			name:   "unsafe increase too large at start",
			report: Report{1, 5, 6, 7},
			want:   false,
			reason: "first increase is 4",
		},
		{
			name:   "safe consistent increase by 3",
			report: Report{1, 4, 7, 10},
			want:   true,
			reason: "all increases are exactly 3",
		},
		{
			name:   "safe consistent decrease by 3",
			report: Report{10, 7, 4, 1},
			want:   true,
			reason: "all decreases are exactly 3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isSafe(tt.report)
			if got != tt.want {
				t.Errorf("isSafe(%v) = %v, want %v; reason: %s", tt.report, got, tt.want, tt.reason)
			}
		})
	}
}

func TestPart1(t *testing.T) {
	tests := []struct {
		name    string
		reports []Report
		want    int
	}{
		{
			name: "example from problem description",
			reports: []Report{
				{7, 6, 4, 2, 1}, // safe
				{1, 2, 7, 8, 9}, // unsafe
				{9, 7, 6, 2, 1}, // unsafe
				{1, 3, 2, 4, 5}, // unsafe
				{8, 6, 4, 4, 1}, // unsafe
				{1, 3, 6, 7, 9}, // safe
			},
			want: 2,
		},
		{
			name:    "no reports",
			reports: []Report{},
			want:    0,
		},
		{
			name: "all safe",
			reports: []Report{
				{1, 2, 3, 4, 5},
				{10, 8, 6, 4, 2},
				{1, 3, 6, 7, 9},
			},
			want: 3,
		},
		{
			name: "all unsafe",
			reports: []Report{
				{1, 2, 7, 8, 9},
				{9, 7, 6, 2, 1},
				{8, 6, 4, 4, 1},
			},
			want: 0,
		},
		{
			name: "single report safe",
			reports: []Report{
				{1, 2, 3},
			},
			want: 1,
		},
		{
			name: "single report unsafe",
			reports: []Report{
				{1, 1, 1},
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Part1(tt.reports)
			if got != tt.want {
				t.Errorf("Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithExampleInput(t *testing.T) {
	parsed := Parse(ExampleInput)

	t.Run("Part1 with example input", func(t *testing.T) {
		want := 2
		got := Part1(parsed)
		if got != want {
			t.Errorf("Part1() = %v, want %v", got, want)
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
