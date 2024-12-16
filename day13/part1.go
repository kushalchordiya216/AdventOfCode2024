package day13

import (
	"bufio"
	"fmt"
	"os"
)

type Part1Solver struct {
	machines []Machine
}

func (s *Part1Solver) Read(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	machines := make([]Machine, 0)
	for scanner.Scan() {
		m, err := readMachineInput(scanner)
		if err != nil {
			return err
		}
		machines = append(machines, m)
	}
	s.machines = machines
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return err
	}
	return nil
}

func (s *Part1Solver) Solve() int {
	result := 0
	for _, m := range s.machines {
		x, y := m.solveByDeterminants()
		if x == 0 && y == 0 {
			continue
		}
		result += 3*x + y
	}
	return result
}
