package day10

import (
	"testing"
)

func TestPart1Example(t *testing.T) {
	topoMap := Parse(ExampleInput)
	result := Part1(topoMap)
	expected := 36

	if result != expected {
		t.Errorf("Part1(ExampleInput) = %d; expected %d", result, expected)
	}
}

func TestFindTrailheads(t *testing.T) {
	topoMap := Parse(ExampleInput)
	trailheads := topoMap.FindTrailheads()

	if len(trailheads) != 9 {
		t.Errorf("FindTrailheads() found %d trailheads; expected 9", len(trailheads))
	}
}

func TestPart2Example(t *testing.T) {
	topoMap := Parse(ExampleInput)
	result := Part2(topoMap)
	expected := 81

	if result != expected {
		t.Errorf("Part2(ExampleInput) = %d; expected %d", result, expected)
	}
}

func TestPart2SmallExample1(t *testing.T) {
	input := `.....0.
..4321.
..5..2.
..6543.
..7..4.
..8765.
..9....`
	topoMap := Parse(input)
	result := Part2(topoMap)
	expected := 3

	if result != expected {
		t.Errorf("Part2(small example 1) = %d; expected %d", result, expected)
	}
}

func TestPart2SmallExample2(t *testing.T) {
	input := `..90..9
...1.98
...2..7
6543456
765.987
876....
987....`
	topoMap := Parse(input)
	result := Part2(topoMap)
	expected := 13

	if result != expected {
		t.Errorf("Part2(small example 2) = %d; expected %d", result, expected)
	}
}

func TestPart2SmallExample3(t *testing.T) {
	input := `012345
123456
234567
345678
4.6789
56789.`
	topoMap := Parse(input)
	result := Part2(topoMap)
	expected := 227

	if result != expected {
		t.Errorf("Part2(small example 3) = %d; expected %d", result, expected)
	}
}
