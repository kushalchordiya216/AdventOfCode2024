package day2

type Part2Solver Solver

func (s *Part2Solver) Read(path string) error {
	var err error
	s.levels, err = readInput(path)
	return err
}

// PopIndexAndCheckSafe checks if the level is safe after removing the element at the given index
func PopIndexAndCheckSafe(level []int, idx int) bool {
	levelCopy := append(level[:idx], level[idx+1:]...)
	safe, _ := IsSafe(levelCopy)
	return safe
}

func (s *Part2Solver) Solve() int {
	counter := 0
	for _, level := range s.levels {
		safe, idx := IsSafe(level)
		if safe {
			counter += 1
		} else {
			// check is array is safe after removing either, the element where the fault was detected, or the element directly before or after it
			if PopIndexAndCheckSafe(level, idx) {
				counter += 1
			} else if idx > 0 && PopIndexAndCheckSafe(level, idx-1) {
				counter += 1
				continue
			} else if idx < len(level)-1 && PopIndexAndCheckSafe(level, idx+1) {
				counter += 1
				continue
			}
		}
	}
	return counter
}
