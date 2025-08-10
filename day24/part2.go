package day24

import (
	"fmt"
	"slices"
	"strings"
)

type Part2Solver struct {
	nodes          map[string]*Node
	operationAlias map[string]string
	carryAlias     map[string]string // key is carry, such as c00, value is identifier of the node that is the carry
	swappedNodes   []string
}

/*
* This was a tricky one. The core idea here, is to look at the entire operation
* as one big full adder or ripple carry adder. So, supposing we have two binary numbers,
* represented with x0x1x2x3...xi y0y1y2y3...yi, which need to be added to get a result of z0z1z2z3...zi
* when we repsent the entire operations as a full adder, we get the following boolean expressions for each result bit
* zi = x_i XOR y_i XOR c_i-1
* c_i = (x_i AND y_i) OR (c_i-1 AND (x_i XOR y_i))
* where ci is the carry after adding x_i and y_i, which is carried over to the next bit
* We there fore, start by verifying lower bits first and then eventually move up
* At each point, if we notice that the expected aliases are not mapped together,
* we initiate a swap
 */

func (s *Part2Solver) Read(path string) error {
	nodes, operationAlias, err := parseInput(path)
	s.carryAlias = make(map[string]string)
	if err != nil {
		return err
	}
	s.nodes = nodes
	s.operationAlias = operationAlias
	return nil
}

func (s *Part2Solver) halfAdderVerification() {
	// verify z0 = x0 XOR y0
	key := "x00-XOR-y00"
	if s.operationAlias[key] != "z00" {
		// z00 is not properly mapped, initiate swap
		tgt := s.nodes[s.operationAlias[key]]
		src := s.nodes["z00"]

		s.swap(tgt, src)
	}
	s.carryAlias["c00"] = s.operationAlias["x00-AND-y00"]
}

func (s *Part2Solver) getExpectedOperationAlias(i int) string {
	// for a binary adder, result of zi = x_i XOR y_i XOR c_i-1
	// by the time we're evaluating zi, c_i-1 is already known and verified
	xi := fmt.Sprintf("x%02d", i)
	yi := fmt.Sprintf("y%02d", i)
	node1 := s.operationAlias[fmt.Sprintf("%s-XOR-%s", xi, yi)]
	node2 := s.carryAlias[fmt.Sprintf("c%02d", i-1)]
	return fmt.Sprintf("%s-XOR-%s", node1, node2)
}

func (s *Part2Solver) getExpectedCarryAlias(i int) string {
	// for binary adder, carry of zi = x_i AND y_i OR c_i-1 AND (x_i XOR y_i)
	xi := fmt.Sprintf("x%02d", i)
	yi := fmt.Sprintf("y%02d", i)
	node1 := s.operationAlias[fmt.Sprintf("%s-AND-%s", xi, yi)]
	node2 := s.carryAlias[fmt.Sprintf("c%02d", i-1)]
	node3 := s.operationAlias[fmt.Sprintf("%s-XOR-%s", xi, yi)]
	node4 := s.operationAlias[fmt.Sprintf("%s-AND-%s", node2, node3)]
	return fmt.Sprintf("%s-OR-%s", node1, node4)
}

func (s *Part2Solver) swap(tgt, src *Node) {
	// Make a note of swaps
	s.swappedNodes = append(s.swappedNodes, tgt.Identifier)
	s.swappedNodes = append(s.swappedNodes, src.Identifier)

	s.operationAlias[tgt.Dependency.getAliases()[0]] = src.Identifier
	s.operationAlias[tgt.Dependency.getAliases()[1]] = src.Identifier
	s.operationAlias[src.Dependency.getAliases()[0]] = tgt.Identifier
	s.operationAlias[src.Dependency.getAliases()[1]] = tgt.Identifier

	// Swap the dependencies of candidates
	src.Dependency, tgt.Dependency = tgt.Dependency, src.Dependency
}

func (s *Part2Solver) fullAdderVerification() {
	for i := 1; i < 45; i++ {
		sum := fmt.Sprintf("z%02d", i)
		carry := fmt.Sprintf("c%02d", i)
		expectedAlias := s.getExpectedOperationAlias(i)
		if val, ok := s.operationAlias[expectedAlias]; ok {
			// expected alias for operation exists but is not mapped to zi and needs to be swapped
			if val != sum {
				tgt := s.nodes[s.operationAlias[expectedAlias]]
				src := s.nodes[sum]

				s.swap(tgt, src)
			}
		} else {
			// expected alias for operation does not exist, one of the building blocks is wrong
			// carry has already been verified in the last iteration of the loop
			// so we need to loop at the other building block, i.e., x_i XOR y_i
			expression := fmt.Sprintf("x%02d-XOR-y%02d", i, i)
			var tgt, src *Node
			tgt = s.nodes[s.operationAlias[expression]]
			carry := s.getExpectedCarryAlias(i - 1)
			sumNode := s.nodes[sum]
			input1 := sumNode.Dependency.Input1
			input2 := sumNode.Dependency.Input2
			if input1 == carry {
				// whichever dependency is equal to the carry, is the one that's correct while the other one needs to be swapped
				src = s.nodes[input2]
			} else {
				src = s.nodes[input1]
			}
			s.swap(tgt, src)
		}
		s.carryAlias[carry] = s.operationAlias[s.getExpectedCarryAlias(i)]
	}
}

func (s *Part2Solver) Solve() int {
	s.halfAdderVerification()
	s.fullAdderVerification()

	// Sort swappedNodes and join them with a comma
	slices.Sort(s.swappedNodes)
	swappedNodesStr := strings.Join(s.swappedNodes, ",")
	fmt.Println(swappedNodesStr)
	return 0
}
