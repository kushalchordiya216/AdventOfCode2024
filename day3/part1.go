package day3

import (
	"fmt"
	"regexp"
)

type Part1Solver Solver

func (s *Part1Solver) Read(path string) error {
	var err error
	s.text, err = readInput(path)
	return err
}

func (s *Part1Solver) Solve() int {
	pattern := `mul\((\d+),(\d+)\)`
	matches := regexp.MustCompile(pattern).FindAllStringSubmatch(s.text, -1)
	sum := 0
	var err error
	for _, match := range matches {
		// match[0] is the full match, match[1] and match[2] are the capture groups
		num1 := 0
		num2 := 0
		_, err = fmt.Sscanf(match[1], "%d", &num1)
		if err != nil {
			return 0
		}
		_, err = fmt.Sscanf(match[2], "%d", &num2)
		if err != nil {
			return 0
		}
		sum += num1 * num2
	}
	return sum
}
