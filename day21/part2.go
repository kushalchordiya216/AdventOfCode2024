package day21

import (
	"os"
	"strconv"
	"strings"
)

type Part2Solver struct {
	codes []string
}

func (s *Part2Solver) Read(path string) error {
	// read list of strings from the file and store it in the codes variable
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	s.codes = strings.Split(string(content), "\n")
	return nil
}

func (s *Part2Solver) Solve() int {
	preComputeTransitions()
	final := 0
	for _, code := range s.codes {
		result := processSequences([]string{code}, 26)
		multiplier, _ := strconv.Atoi(code[0 : len(code)-1])
		result *= multiplier
		final += result
	}
	return final
}
