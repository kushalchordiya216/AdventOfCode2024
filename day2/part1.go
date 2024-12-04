package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/kushalchordiya216/AOC2024/common"
)

type Part1Solver struct {
	levels [][]int
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
		nums := strings.Split(line, " ")
		var level []int
		for _, num := range nums {
			val, err := strconv.Atoi(num)
			if err != nil {
				return &common.CustomError{Msg: fmt.Sprintf("Expected all elements to be integers, found: %s", num)}
			}
			level = append(level, val)
		}
		s.levels = append(s.levels, level)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return err
	}
	return nil
}

func (s *Part1Solver) Solve() int {
	counter := 0
	for _, level := range s.levels {
		flag := 0
		safe := 1
		for i := 0; i < len(level)-1; i++ {
			diff := level[i+1] - level[i]
			if diff < -3 || diff > 3 || diff == 0 {
				safe = 0
				break
			}
			if diff < 0 {
				if flag == 0 {
					flag = -1
				}
				if flag == 1 {
					safe = 0
					break
				}
			} else if diff > 0 {
				if flag == 0 {
					flag = 1
				}
				if flag == -1 {
					safe = 0
					break
				}
			}
		}
		if safe == 1 {
			counter += 1
		}
	}
	return counter
}
