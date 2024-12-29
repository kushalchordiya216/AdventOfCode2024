package day17

import (
	"fmt"
	"strconv"
	"strings"
)

type Part1Solver Solver

func (s *Part1Solver) Read(path string) error {
	solver, err := readInput(path)
	if err != nil {
		return err
	}
	*s = Part1Solver(solver)
	return nil
}

func (s *Part1Solver) Solve() int {
	(*Solver)(s).run()
	outputStrings := make([]string, len(s.Output))
	for i, v := range s.Output {
		outputStrings[i] = strconv.Itoa(v)
	}
	fmt.Println(strings.Join(outputStrings, ","))
	return 0
}
