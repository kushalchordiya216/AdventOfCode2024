package day19

type Part2Solver Solver

func (s *Part2Solver) Read(path string) error {
	solver, err := readInput(path)
	if err != nil {
		return err
	}
	*s = Part2Solver(solver)
	return nil
}

func (s *Part2Solver) getValidDesigns(design string, maxTowelLength int) int {
	dp := make([]int, len(design))
	for i := range dp {
		prefix := design[:i+1]
		if s.towels[prefix] {
			dp[i] += 1
		}
		for j := 1; j <= i && j <= maxTowelLength; j++ {
			if s.towels[design[i-j+1:i+1]] {
				dp[i] += dp[i-j]
			}
		}
	}
	return dp[len(design)-1]
}

func (s *Part2Solver) Solve() int {
	var maxTowelLength int
	for towel := range s.towels {
		maxTowelLength = max(maxTowelLength, len(towel))
	}
	counter := 0
	for _, design := range s.designs {
		counter += s.getValidDesigns(design, maxTowelLength)
	}
	return counter
}
