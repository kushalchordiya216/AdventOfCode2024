package day20

import "github.com/kushalchordiya216/AOC2024/common/utils"

type Part2Solver Solver

func (s *Part2Solver) Read(path string) error {
	solver, err := readInput(path)
	if err != nil {
		return err
	}
	*s = Part2Solver(solver)
	return nil
}

func (s *Part2Solver) Solve() int {
	costs := make(map[utils.Coord]int)
	(*Solver)(s).dfs(s.start, 0, costs)
	return (*Solver)(s).findAllCheats(20, costs, 100)
}
