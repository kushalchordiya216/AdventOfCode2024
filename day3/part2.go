package day3

import (
	"fmt"
	"regexp"
)

type Part2Solver Solver

func (s *Part2Solver) Read(path string) error {
	var err error
	s.text, err = readInput(path)
	return err
}

func (s *Part2Solver) Solve() int {
	pattern := `mul\((\d+),(\d+)\)|do\(\)|don't\(\)`
	matches := regexp.MustCompile(pattern).FindAllStringSubmatch(s.text, -1)
	flag := 1
	sum := 0
	var err error
	for _, match := range matches {
		// match[0] is the full match, match[1] and match[2] are the capture groups
		if match[0] == "do()" {
			flag = 1
		} else if match[0] == "don't()" {
			flag = 0
		} else {
			if flag == 1 {
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
		}
	}
	return sum
}
