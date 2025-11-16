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
