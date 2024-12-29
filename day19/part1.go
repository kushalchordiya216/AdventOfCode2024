package day19

type Part1Solver Solver

func (s *Part1Solver) Read(path string) error {
	solver, err := readInput(path)
	if err != nil {
		return err
	}
	*s = Part1Solver(solver)
	return nil
}

func (s *Part1Solver) isValidDesign(design string, maxTowelLength int) bool {
	dp := make([]bool, len(design))
	for i := range dp {
		prefix := design[:i+1]
		if s.towels[prefix] {
			dp[i] = true
			continue
		}
		for j := 1; j <= i && j <= maxTowelLength; j++ {
			if dp[i-j] && s.towels[design[i-j+1:i+1]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(design)-1]
}

func (s *Part1Solver) Solve() int {
	var maxTowelLength int
	for towel := range s.towels {
		maxTowelLength = max(maxTowelLength, len(towel))
	}
	counter := 0
	for _, design := range s.designs {
		if s.isValidDesign(design, maxTowelLength) {
			counter++
		}
	}
	return counter
}
