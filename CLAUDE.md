# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is an Advent of Code 2024 solutions repository written in Go. The project compiles to a single executable that can solve any day/part combination using command-line arguments.

## Build and Run Commands

- **Build the project**: `go build` (creates `AOC2024` executable)
- **Run a solution**: `./AOC2024 <day> <part> [input_file_path]`
  - `day`: Integer 1-25
  - `part`: 1 or 2
  - `input_file_path`: Optional, defaults to `day{N}/input.txt`
- **Example**: `./AOC2024 3 2` runs day 3, part 2
- **Format code**: `go fmt ./...`
- **Run tests**: `go test ./...` (though no tests currently exist)
- **Check for issues**: `go vet ./...`

## Architecture

### Core Interface
All solutions implement the `common.Solver` interface:
```go
type Solver interface {
    Read(path string) error  // Parse input file
    Solve() int             // Return integer solution
}
```

### Project Structure
- **main.go**: Entry point with solver selection logic via `selectSolver(day, part)`
- **common/**: Shared interfaces and utilities
  - `commons.go`: Core `Solver` interface and `CustomError` type
  - `utils/matrix.go`: Grid utilities, coordinate system, and direction handling
- **dayN/**: Individual day solutions
  - `part1.go`: Part 1 solver implementation
  - `part2.go`: Part 2 solver implementation  
  - `shared.go`: Common utilities for both parts (when present)
  - `input.txt`: Challenge input data

### Common Patterns

#### Solver Implementation
Each part implements its own solver struct:
```go
type Part1Solver struct {
    // fields for parsed input
}

func (s *Part1Solver) Read(path string) error {
    // Parse input file into struct fields
}

func (s *Part1Solver) Solve() int {
    // Implement solution logic
}
```

#### Grid-based Problems
Many solutions use the utilities in `common/utils/matrix.go`:
- `Coord{X, Y}`: 2D coordinates with helper methods
- `Direction`: Up/Down/Left/Right enum with rotation methods
- `Grid[T]`: Generic 2D grid with bounds checking

#### File Reading Pattern
Solutions consistently use:
```go
file, err := os.Open(path)
defer file.Close()
scanner := bufio.NewScanner(file)
for scanner.Scan() {
    line := scanner.Text()
    // Process line
}
```

### Adding New Solutions

1. Create `dayN/` directory
2. Implement `Part1Solver` and `Part2Solver` in respective files
3. Add day to import statements in `main.go`
4. Add solver mappings to `selectSolver()` function
5. Place input data in `dayN/input.txt`

The main.go file needs updates in two places:
- Import section: `"github.com/kushalchordiya216/AOC2024/dayN"`  
- Solver map: `N: {1: &dayN.Part1Solver{}, 2: &dayN.Part2Solver{}}`

## General Rules
- When designing data structures, opt for simpler solutions. Do not make the data structure overly complex or layered 
- Opt for re-use when possible instead of hand-crafting specific data structures per case 
- Maintain a focus on performance and efficiency. Design algorithms that are time efficient, avoid copying memory in hot paths, use concurrency where beneficial
