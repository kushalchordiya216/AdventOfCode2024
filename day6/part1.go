package day6

import "fmt"

type Part1Solver struct {
	grid [][]rune
	x, y int
	d    Direction
}

func (s *Part1Solver) Read(path string) error {
	var err error
	s.grid, s.x, s.y, s.d, err = readInput(path)
	return err
}

func (s *Part1Solver) Solve() int {
	err := traceGuardPath(s.grid, s.x, s.y, s.d)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	counter := 0
	for i := range s.grid {
		for j := range s.grid[i] {
			if s.grid[i][j] == 'X' {
				counter += 1
			}
		}
	}
	return counter
}
