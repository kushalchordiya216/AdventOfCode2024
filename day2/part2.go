package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/kushalchordiya216/AOC2024/common"
)

type Part2Solver struct {
	levels [][]int
}

func (s *Part2Solver) Read(path string) error {
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

// returns if a given array is safe or not. If the array is not safe, returns the index at which a break was discovered
func isSafe(level []int) (bool, int) {
	flag := 0
	for i := 0; i < len(level)-1; i++ {
		diff := level[i+1] - level[i]
		if diff < -3 || diff > 3 || diff == 0 {
			return false, i
		}
		if diff < 0 {
			if flag == 0 {
				flag = -1
			}
			if flag == 1 {
				return false, i
			}
		} else if diff > 0 {
			if flag == 0 {
				flag = 1
			}
			if flag == -1 {
				return false, i
			}
		}
	}
	return true, -1
}

func (s *Part2Solver) Solve() int {
	counter := 0
	for _, level := range s.levels {
		safe, idx := isSafe(level)
		if safe {
			counter += 1
		} else {
			// check is array is safe after removing either, the element where the fault was detected, or the element directly before or after it
			levelCopy := make([]int, len(level))
			copy(levelCopy, level)
			levelCopy = append(levelCopy[:idx], levelCopy[idx+1:]...)
			safe, _ = isSafe(levelCopy)
			if safe {
				counter += 1
				continue
			}
			if idx > 0 {
				levelCopy := make([]int, len(level))
				copy(levelCopy, level)
				levelCopy = append(levelCopy[:idx-1], levelCopy[idx:]...)
				safe, _ = isSafe(levelCopy)
				if safe {
					counter += 1
					continue
				}
			}
			if idx < len(level)-1 {
				levelCopy := make([]int, len(level))
				copy(levelCopy, level)
				levelCopy = append(levelCopy[:idx+1], levelCopy[idx+2:]...)
				safe, _ = isSafe(levelCopy)
				if safe {
					counter += 1
					continue
				}
			}
		}
	}
	return counter
}
