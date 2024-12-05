package day5

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Part1Solver struct {
	after   map[int][]int
	manuals [][]int
}

func (s *Part1Solver) Read(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	s.after = make(map[int][]int)
	s.manuals = make([][]int, 0)
	readingManuals := false

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			readingManuals = true
			continue
		}

		if readingManuals {
			// Parse manual numbers
			var numbers []int
			numStrs := strings.Split(line, ",")
			for _, numStr := range numStrs {
				num, err := strconv.Atoi(numStr)
				if err != nil {
					return err
				}
				numbers = append(numbers, num)
			}
			s.manuals = append(s.manuals, numbers)
		} else {
			// Parse order mappings
			parts := strings.Split(line, "|")
			if len(parts) != 2 {
				return fmt.Errorf("invalid line format: %s", line)
			}

			key, err := strconv.Atoi(strings.TrimSpace(parts[0]))
			if err != nil {
				return err
			}

			value, err := strconv.Atoi(strings.TrimSpace(parts[1]))
			if err != nil {
				return err
			}

			s.after[key] = append(s.after[key], value)
		}
	}

	return scanner.Err()
}

func (s *Part1Solver) Solve() int {
	result := 0
	for _, manual := range s.manuals {
		valid := true
		for i, current := range manual {

			// Check all previous elements in this manual
			for j := 0; j < i; j++ {
				previous := manual[j]
				// Check if previous element appears in current element's order list
				for _, comesAfter := range s.after[current] {
					if comesAfter == previous {
						valid = false
						break
					}
				}
			}
		}
		if valid {
			result += manual[(len(manual)-1)/2]
		}
	}
	return result
}
