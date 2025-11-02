package utils

import (
	"reflect"
	"testing"
)

func TestAsLines(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []string
	}{
		{
			name:  "simple lines",
			input: "line1\nline2\nline3",
			want:  []string{"line1", "line2", "line3"},
		},
		{
			name:  "lines with whitespace",
			input: "  line1  \n  line2  \n  line3  ",
			want:  []string{"line1", "line2", "line3"},
		},
		{
			name:  "empty lines in middle preserved",
			input: "line1\n\nline2",
			want:  []string{"line1", "", "line2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AsLines(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AsLines() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseInts(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    []int
		wantErr bool
	}{
		{
			name:    "space-separated integers",
			input:   "1 2 3 4 5",
			want:    []int{1, 2, 3, 4, 5},
			wantErr: false,
		},
		{
			name:    "multiple spaces",
			input:   "1    2   3",
			want:    []int{1, 2, 3},
			wantErr: false,
		},
		{
			name:    "negative numbers",
			input:   "-1 2 -3",
			want:    []int{-1, 2, -3},
			wantErr: false,
		},
		{
			name:    "invalid input",
			input:   "1 abc 3",
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseInts(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseInts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseInts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMustParseInts(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		want      []int
		wantPanic bool
	}{
		{
			name:      "valid input",
			input:     "1 2 3",
			want:      []int{1, 2, 3},
			wantPanic: false,
		},
		{
			name:      "invalid input panics",
			input:     "1 abc 3",
			want:      nil,
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.wantPanic {
					t.Errorf("MustParseInts() panic = %v, wantPanic %v", r != nil, tt.wantPanic)
				}
			}()
			got := MustParseInts(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MustParseInts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseInt(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    int
		wantErr bool
	}{
		{
			name:    "valid positive",
			input:   "42",
			want:    42,
			wantErr: false,
		},
		{
			name:    "valid negative",
			input:   "-42",
			want:    -42,
			wantErr: false,
		},
		{
			name:    "with whitespace",
			input:   "  123  ",
			want:    123,
			wantErr: false,
		},
		{
			name:    "invalid",
			input:   "abc",
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseInt(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
