package day11

type Part1Solver struct {
	stones []int
}

func (s *Part1Solver) Read(path string) error {
	var err error
	s.stones, err = readInput(path)
	return err
}

func (s *Part1Solver) Solve() int {
	cache := make(map[int]IterationDict)
	result := 0
	for _, stone := range s.stones {
		result += memoizedRecursiveIteration(stone, 25, cache)
	}
	return result
}
