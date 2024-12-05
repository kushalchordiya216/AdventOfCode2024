package day4

import (
	"bufio"
	"fmt"
	"os"
)

type Part1Solver struct {
	grid [][]rune
}

func (s *Part1Solver) Read(path string) error {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := []rune(line)
		s.grid = append(s.grid, row)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return err
	}
	return nil
}

func getNextState(state rune) rune {
	switch state {
	case 'X':
		return 'M'
	case 'M':
		return 'A'
	case 'A':
		return 'S'
	case 'S':
		return '\n'
	default:
		return ' '
	}
}

func explore(s *Part1Solver, x int, y int) int {
	offsets := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
	counter := 0
	for _, offset := range offsets {
		state := 'M'
		newX, newY := x, y
		for {
			newX = newX + offset[0]
			newY = newY + offset[1]
			if newX < 0 || newY < 0 || newX >= len(s.grid) || newY >= len(s.grid[0]) {
				break
			}
			if s.grid[newX][newY] == state {
				state = getNextState(state)
				if state == '\n' {
					counter += 1
					break
				}
			} else {
				break
			}

		}
	}
	return counter
}

func (s *Part1Solver) Solve() int {
	result := 0
	for x := range s.grid {
		for y := range s.grid[x] {
			if s.grid[x][y] == 'X' {
				result += explore(s, x, y)
			}
		}
	}
	return result
}
