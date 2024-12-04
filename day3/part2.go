package day3

import (
	"fmt"
	"regexp"
)

type Part2Solver struct {
	text string
}

func (s *Part2Solver) Read(path string) error {
	solver := &Part1Solver{text: s.text}
	err := solver.Read(path)
	if err != nil {
		return err
	}
	s.text = solver.text
	return nil
}

func (s *Part2Solver) Solve() int {
	fmt.Println("In my lane")
	pattern := `mul\((\d+),(\d+)\)|do\(\)|don't\(\)`
	matches := regexp.MustCompile(pattern).FindAllStringSubmatch(s.text, -1)
	flag := 1
	sum := 0
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
				fmt.Sscanf(match[1], "%d", &num1)
				fmt.Sscanf(match[2], "%d", &num2)
				sum += num1 * num2
			}
		}
	}
	return sum
}
