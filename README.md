# Advent of Code 2024

Solutions to [Advent of Code 2024](https://adventofcode.com/2024) in Go.

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

## License

MIT License - See LICENSE file for details
