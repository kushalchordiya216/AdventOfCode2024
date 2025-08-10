package day24

import (
	"fmt"
)

type Part1Solver struct {
	nodes map[string]*Node
}

func (s *Part1Solver) Read(path string) error {
	nodes, _, err := parseInput(path)
	if err != nil {
		return err
	}
	s.nodes = nodes
	return nil
}

func (s *Part1Solver) resolveValue(identifier string) int {
	node := s.nodes[identifier]
	if node.Value != nil {
		return *node.Value
	}

	if node.Dependency == nil {
		return 0
	}

	val1 := s.resolveValue(node.Dependency.Input1)
	val2 := s.resolveValue(node.Dependency.Input2)

	var result int
	switch node.Dependency.Operation {
	case "AND":
		result = val1 & val2
	case "OR":
		result = val1 | val2
	case "XOR":
		result = val1 ^ val2
	}

	node.Value = &result
	return result
}

func (s *Part1Solver) Solve() int {
	binaryBits := make([]int, 100)

	for i := range 100 {
		zIdentifier := fmt.Sprintf("z%02d", i)
		if _, exists := s.nodes[zIdentifier]; exists {
			binaryBits[i] = s.resolveValue(zIdentifier)
		} else {
			binaryBits[i] = 0
		}
	}

	result := 0
	for i := range binaryBits {
		if binaryBits[i] == 1 {
			result |= (1 << i)
		}
	}
	fmt.Println(len(s.nodes))
	return result
}
