package day7

import (
	"fmt"
	"strconv"
)

type Part2Solver struct {
	targets  []int
	numLists [][]int
}

func (s *Part2Solver) Read(path string) error {
	targets, numLists, err := readInput(path)
	if err != nil {
		return err
	}
	s.targets, s.numLists = targets, numLists
	return nil
}

func concatenateNums(a, b int) int {
	// Convert both numbers to string and concatenate them
	str := fmt.Sprintf("%d%d", a, b)
	// Convert back to integer
	result, _ := strconv.Atoi(str)
	return result
}

func checkIfValidEquation2(target int, result int, remaining []int) bool {
	if len(remaining) == 0 {
		return result == target
	}
	return checkIfValidEquation2(target, result+remaining[0], remaining[1:]) || checkIfValidEquation2(target, result*remaining[0], remaining[1:]) || checkIfValidEquation2(target, concatenateNums(result, remaining[0]), remaining[1:])
}

func (s *Part2Solver) Solve() int {
	return ParallelSolver(s.targets, s.numLists, checkIfValidEquation2)
}
