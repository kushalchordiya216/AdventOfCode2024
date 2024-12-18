package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/kushalchordiya216/AOC2024/common"
)

type Part2Solver struct {
	lookup1 map[int]int
	lookup2 map[int]int
}

func (s *Part2Solver) Read(path string) error {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return nil
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("Error closing file: %v\n", err)
		}
	}(file)

	s.lookup1 = make(map[int]int)
	s.lookup2 = make(map[int]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, " ")
		if len(nums) < 2 {
			return &common.CustomError{Msg: fmt.Sprintf("Expected at least 2 numbers per line in input file, received: %v", nums)}
		}
		num1, err1 := strconv.Atoi(nums[0])
		if err1 != nil {
			return err1
		}
		if val, exists := s.lookup1[num1]; exists {
			s.lookup1[num1] = val + 1
		} else {
			s.lookup1[num1] = 1
		}
		num2, err2 := strconv.Atoi(nums[len(nums)-1])
		if err2 != nil {
			return err2
		}
		if val, exists := s.lookup2[num2]; exists {
			s.lookup2[num2] = val + 1
		} else {
			s.lookup2[num2] = 1
		}
	}
	return nil
}

func (s *Part2Solver) Solve() int {
	similarityScore := 0
	for key, val1 := range s.lookup1 {
		val2 := 0
		if v, exists := s.lookup2[key]; exists {
			val2 = v
		}
		similarityScore += key * val2 * val1
	}
	return similarityScore
}
