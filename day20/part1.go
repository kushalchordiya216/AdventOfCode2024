package day20

import (
	"github.com/kushalchordiya216/AOC2024/common/utils"
)

type Part1Solver Solver

func (s *Part1Solver) Read(path string) error {
	solver, err := readInput(path)
	if err != nil {
		return err
	}
	*s = Part1Solver(solver)
	return nil
}

func (s *Part1Solver) Solve() int {
	costs := make(map[utils.Coord]int)
	(*Solver)(s).dfs(s.start, 0, costs)
	return (*Solver)(s).findAllCheats(2, costs, 100)
}
