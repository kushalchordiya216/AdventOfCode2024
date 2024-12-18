package day2

type Part1Solver Solver

func (s *Part1Solver) Read(path string) error {
	var err error
	s.levels, err = readInput(path)
	return err
}

func (s *Part1Solver) Solve() int {
	counter := 0
	for _, level := range s.levels {
		safe, _ := IsSafe(level)
		if safe {
			counter += 1
		}
	}
	return counter
}
