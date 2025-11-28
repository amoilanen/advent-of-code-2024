package main

import (
	"fmt"
	"os"

	"github.com/amoilanen/advent-of-code-2024/internal/days/day01"
	"github.com/amoilanen/advent-of-code-2024/internal/days/day02"
	"github.com/amoilanen/advent-of-code-2024/internal/days/day03"
	"github.com/amoilanen/advent-of-code-2024/internal/days/day04"
	"github.com/amoilanen/advent-of-code-2024/internal/days/day05"
	"github.com/amoilanen/advent-of-code-2024/internal/days/day06"
	"github.com/amoilanen/advent-of-code-2024/internal/days/day07"
	"github.com/amoilanen/advent-of-code-2024/internal/days/day08"
	"github.com/amoilanen/advent-of-code-2024/internal/days/day09"
	"github.com/amoilanen/advent-of-code-2024/internal/days/day10"
	"github.com/amoilanen/advent-of-code-2024/internal/days/day11"
	"github.com/amoilanen/advent-of-code-2024/internal/days/day12"
	"github.com/amoilanen/advent-of-code-2024/internal/days/day13"
	"github.com/amoilanen/advent-of-code-2024/internal/days/day14"
	"github.com/amoilanen/advent-of-code-2024/internal/days/day15"
)

func main() {
	if len(os.Args) > 1 {
		runSpecificDay(os.Args[1])
	} else {
		runAllDays()
	}
}

func runAllDays() {
	fmt.Println("Advent of Code 2024 - Solutions")
	fmt.Println("================================")
	fmt.Println()

	runDay01()
	runDay02()
	runDay03()
	runDay04()
	runDay05()
	runDay06()
	runDay07()
	runDay08()
	runDay09()
	runDay10()
	runDay11()
	runDay12()
	runDay13()
	runDay14()
	runDay15()
}

func runSpecificDay(day string) {
	switch day {
	case "1", "day01", "day1":
		runDay01()
	case "2", "day02", "day2":
		runDay02()
	case "3", "day03", "day3":
		runDay03()
	case "4", "day04", "day4":
		runDay04()
	case "5", "day05", "day5":
		runDay05()
	case "6", "day06", "day6":
		runDay06()
	case "7", "day07", "day7":
		runDay07()
	case "8", "day08", "day8":
		runDay08()
	case "9", "day09", "day9":
		runDay09()
	case "10", "day10":
		runDay10()
	case "11", "day11":
		runDay11()
	case "12", "day12":
		runDay12()
	case "13", "day13":
		runDay13()
	case "14", "day14":
		runDay14()
	case "15", "day15":
		runDay15()
	default:
		fmt.Fprintf(os.Stderr, "Unknown day: %s\n", day)
		fmt.Fprintln(os.Stderr, "Usage: aoc2024 [day]")
		fmt.Fprintln(os.Stderr, "Example: aoc2024 1")
		os.Exit(1)
	}
}

func runDay01() {
	fmt.Println("Day 1:")
	input := day01.DayInput
	parsed := day01.Parse(input)
	fmt.Printf("  Part 1: %d\n", day01.Part1(parsed))
	fmt.Printf("  Part 2: %d\n", day01.Part2(parsed))
	fmt.Println()
}

func runDay02() {
	fmt.Println("Day 2:")
	input := day02.DayInput
	parsed := day02.Parse(input)
	fmt.Printf("  Part 1: %d\n", day02.Part1(parsed))
	fmt.Printf("  Part 2: %d\n", day02.Part2(parsed))
	fmt.Println()
}

func runDay03() {
	fmt.Println("Day 3:")
	input := day03.DayInput
	instructions := day03.Parse(input)
	fmt.Printf("  Part 1: %d\n", day03.Part1(instructions))
	fmt.Printf("  Part 2: %d\n", day03.Part2(instructions))
	fmt.Println()
}

func runDay04() {
	fmt.Println("Day 4:")
	input := day04.DayInput
	grid := day04.Parse(input)
	fmt.Printf("  Part 1: %d\n", day04.Part1(grid))
	fmt.Printf("  Part 2: %d\n", day04.Part2(grid))
	fmt.Println()
}

func runDay05() {
	fmt.Println("Day 5:")
	input := day05.DayInput
	parsed := day05.Parse(input)
	fmt.Printf("  Part 1: %d\n", day05.Part1(parsed))
	fmt.Printf("  Part 2: %d\n", day05.Part2(parsed))
	fmt.Println()
}

func runDay06() {
	fmt.Println("Day 6:")
	input := day06.DayInput
	grid, guard := day06.Parse(input)
	fmt.Printf("  Part 1: %d\n", day06.Part1(grid, guard))
	fmt.Printf("  Part 2: %d\n", day06.Part2(grid, guard))
	fmt.Println()
}

func runDay07() {
	fmt.Println("Day 7:")
	input := day07.DayInput
	equations := day07.Parse(input)
	fmt.Printf("  Part 1: %d\n", day07.Part1(equations))
	fmt.Printf("  Part 2: %d\n", day07.Part2(equations))
	fmt.Println()
}

func runDay08() {
	fmt.Println("Day 8:")
	input := day08.DayInput
	grid := day08.Parse(input)
	fmt.Printf("  Part 1: %d\n", day08.Part1(grid))
	fmt.Printf("  Part 2: %d\n", day08.Part2(grid))
	fmt.Println()
}

func runDay09() {
	fmt.Println("Day 9:")
	input := day09.DayInput
	diskMap := day09.Parse(input)
	fmt.Printf("  Part 1: %d\n", day09.Part1(diskMap))
	fmt.Printf("  Part 2: %d\n", day09.Part2(diskMap))
	fmt.Println()
}

func runDay10() {
	fmt.Println("Day 10:")
	input := day10.DayInput
	topoMap := day10.Parse(input)
	fmt.Printf("  Part 1: %d\n", day10.Part1(topoMap))
	fmt.Printf("  Part 2: %d\n", day10.Part2(topoMap))
	fmt.Println()
}

func runDay11() {
	fmt.Println("Day 11:")
	input := day11.DayInput
	stones := day11.Parse(input)
	fmt.Printf("  Part 1: %d\n", day11.Part1(stones))
	fmt.Printf("  Part 2: %d\n", day11.Part2(stones))
	fmt.Println()
}

func runDay12() {
	fmt.Println("Day 12:")
	input := day12.DayInput
	grid := day12.Parse(input)
	fmt.Printf("  Part 1: %d\n", day12.Part1(grid))
	fmt.Printf("  Part 2: %d\n", day12.Part2(grid))
	fmt.Println()
}

func runDay13() {
	fmt.Println("Day 13:")
	input := day13.DayInput
	machines := day13.Parse(input)
	fmt.Printf("  Part 1: %d\n", day13.Part1(machines))
	fmt.Printf("  Part 2: %d\n", day13.Part2(machines))
	fmt.Println()
}

func runDay14() {
	fmt.Println("Day 14:")
	input := day14.DayInput
	robots := day14.Parse(input)
	fmt.Printf("  Part 1: %d\n", day14.Part1(robots, 101, 103))
	fmt.Printf("  Part 2: %d\n", day14.Part2(robots, 101, 103))
	fmt.Println()
}

func runDay15() {
	fmt.Println("Day 15:")
	input := day15.Input
	fmt.Printf("  Part 1: %d\n", day15.Part1(input))
	fmt.Printf("  Part 2: %d\n", day15.Part2(input))
	fmt.Println()
}
