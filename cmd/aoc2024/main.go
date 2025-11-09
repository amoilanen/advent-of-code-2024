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
	fmt.Println()
}
