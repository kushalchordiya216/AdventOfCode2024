package day11

type Part2Solver struct {
	stones []int
}

func (s *Part2Solver) Read(path string) error {
	var err error
	s.stones, err = readInput(path)
	return err
}

func (s *Part2Solver) Solve() int {
	cache := make(map[int]IterationDict)
	result := 0
	for _, stone := range s.stones {
		result += memoizedRecursiveIteration(stone, 75, cache)
	}
	return result
}
