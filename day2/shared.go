package day2

import (
	"bufio"
	"fmt"
	"github.com/kushalchordiya216/AOC2024/common"
	"github.com/kushalchordiya216/AOC2024/common/utils"
	"os"
	"strconv"
	"strings"
)

type Solver struct {
	levels utils.Grid[int]
}

func readInput(path string) (utils.Grid[int], error) {
	file, err := os.Open(path)
	var levels utils.Grid[int]
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("Error closing file: %v\n", err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, " ")
		var level []int
		for _, num := range nums {
			val, err := strconv.Atoi(num)
			if err != nil {
				return nil, &common.CustomError{Msg: fmt.Sprintf("Expected all elements to be integers, found: %s", num)}
			}
			level = append(level, val)
		}
		levels = append(levels, level)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return nil, err
	}
	return levels, nil
}

func IsSafe(level []int) (bool, int) {
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
		} else {
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
