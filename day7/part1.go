package day7

type Part1Solver struct {
	targets  []int
	numLists [][]int
}

func (s *Part1Solver) Read(path string) error {
	var err error
	s.targets, s.numLists, err = readInput(path)
	return err
}

func checkIfValidEquation(target int, result int, remaining []int) bool {
	if len(remaining) == 0 {
		return result == target
	}

	return checkIfValidEquation(target, result+remaining[0], remaining[1:]) || checkIfValidEquation(target, result*remaining[0], remaining[1:])
}

func (s *Part1Solver) Solve() int {
	return ParallelSolver(s.targets, s.numLists, checkIfValidEquation)
}
