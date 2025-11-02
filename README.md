# Advent of Code 2024

Solutions to [Advent of Code 2024](https://adventofcode.com/2024) in Go.

## Project Structure

```
advent-of-code-2024/
├── cmd/
│   └── aoc2024/          # Main executable
│       └── main.go
├── internal/
│   ├── days/             # Solutions for each day
│   │   ├── day01/
│   │   │   ├── day01.go
│   │   │   └── day01_test.go
│   │   ├── day02/
│   │   └── ...
│   └── utils/            # Common utilities
│       ├── parsing.go    # Input parsing helpers
│       ├── math.go       # Math utilities
│       └── ...
├── go.mod
└── README.md
```

## Usage

### Running Solutions

Run all implemented days:
```bash
go run cmd/aoc2024/main.go
```

Run a specific day:
```bash
go run cmd/aoc2024/main.go 1
# or
go run cmd/aoc2024/main.go day01
```

Build the executable:
```bash
go build -o aoc2024 cmd/aoc2024/main.go
./aoc2024
```

### Running Tests

Run all tests:
```bash
go test ./...
```

Run tests for a specific day:
```bash
go test ./internal/days/day01
```

Run tests with coverage:
```bash
go test -cover ./...
```

Generate detailed coverage report:
```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

Run benchmarks:
```bash
go test -bench=. ./...
```

### Development

Each day's solution follows this structure:

1. **day01.go** - Contains:
   - `Parse(input string)` - Parses the input into appropriate data structures
   - `Part1(data)` - Solves part 1
   - `Part2(data)` - Solves part 2
   - `ExampleInput` - Example input for testing

2. **day01_test.go** - Contains:
   - Unit tests for parsing
   - Unit tests for both parts
   - Benchmarks for performance testing

## Common Utilities

The `internal/utils` package provides helper functions:

- **parsing.go**: Input parsing utilities
  - `AsLines(input string)` - Split input into lines
  - `ParseInts(input string)` - Parse space-separated integers
  - `MustParseInt(s string)` - Parse single integer (panics on error)

- **math.go**: Mathematical utilities
  - `Abs(x int)` - Absolute value
  - `Min/Max(a, b int)` - Min/max of two integers
  - `GCD/LCM(a, b int)` - Greatest common divisor / Least common multiple
  - `Sum/Product(nums []int)` - Sum/product of slice

## Testing Strategy

All solutions are developed with test-driven development:

1. Parse the example input from the problem
2. Write tests for the example cases
3. Implement the solution
4. Verify with actual puzzle input
5. Add edge case tests as needed

## License

MIT License - See LICENSE file for details
