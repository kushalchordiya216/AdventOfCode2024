package day5

import "slices"

type Part2Solver struct {
	after   map[int][]int
	manuals [][]int
}

func (s *Part2Solver) Read(path string) error {
	solver := &Part1Solver{after: s.after, manuals: s.manuals}
	err := solver.Read(path)
	if err != nil {
		return err
	}
	s.after = solver.after
	s.manuals = solver.manuals
	return nil
}

func (s *Part2Solver) Solve() int {
	result := 0
	for _, manual := range s.manuals {
		flag := 0
		for i := 0; i < len(manual); i++ {
			for j := i; j < len(manual); j++ {
				if i == j {
					continue
				}
				if slices.Contains(s.after[manual[j]], manual[i]) {
					manual[j], manual[i] = manual[i], manual[j]
					flag = 1
				}
			}
		}
		if flag == 1 {
			result += manual[(len(manual)-1)/2]
		}
	}
	return result
}
