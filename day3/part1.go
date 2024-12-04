package day3

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type Part1Solver struct {
	text string
}

func (s *Part1Solver) Read(path string) error {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		s.text += line
	}
	return nil
}

func (s *Part1Solver) Solve() int {
	pattern := `mul\((\d+),(\d+)\)`
	matches := regexp.MustCompile(pattern).FindAllStringSubmatch(s.text, -1)
	sum := 0
	for _, match := range matches {
		// match[0] is the full match, match[1] and match[2] are the capture groups
		num1 := 0
		num2 := 0
		fmt.Sscanf(match[1], "%d", &num1)
		fmt.Sscanf(match[2], "%d", &num2)
		sum += num1 * num2
	}
	return sum
}
