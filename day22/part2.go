package day22

import (
	"bufio"
	"os"
	"strconv"
)

type Part2Solver struct {
	numbers []int
}

func (s *Part2Solver) Read(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return err
		}
		s.numbers = append(s.numbers, num)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

// organise the whole thing as a tree where each node has 20 children, accessed via keys from range -9 to 9
type NTree struct {
	children map[int]*NTree
	val      int
	flag     bool
}

func (n *NTree) set(key int) *NTree {
	child := &NTree{
		val:      0,
		flag:     true,
		children: make(map[int]*NTree),
	}
	n.children[key] = child
	return child
}

func (n *NTree) get(key int) *NTree {
	return n.children[key]
}

// Store the value for sequence of the numbers in the queue
// Only store the first value, i.e., if value for this sequence hasn't already been set for this iteration
func (n *NTree) deepSet(queue []int, value int) {
	leaf := n.get(queue[0]).get(queue[1]).get(queue[2]).get(queue[3])
	if leaf.flag {
		leaf.val += value
		leaf.flag = false
	}
}

func (n *NTree) deepGet(queue []int) int {
	leaf := n.get(queue[0]).get(queue[1]).get(queue[2]).get(queue[3])
	if leaf.flag {
		return leaf.val
	}
	return -1
}

func (n *NTree) initTree(depth int) {
	if depth == 0 {
		return
	}
	for i := -9; i <= 9; i++ {
		child := n.set(i)
		child.initTree(depth - 1)
	}
}

func (n *NTree) resetFlags(depth int) {
	if depth == 0 {
		n.flag = true
		return
	}
	for i := -9; i <= 9; i++ {
		child := n.get(i)
		child.resetFlags(depth - 1)
	}
}

func addToQueue(queue *[]int, val int) {
	if len(*queue) == 4 {
		(*queue)[0] = (*queue)[1]
		(*queue)[1] = (*queue)[2]
		(*queue)[2] = (*queue)[3]
		(*queue)[3] = val
		// with constant length of queue
		// shifting the entire thing is more efficient than
		// mem alloc/de-alloc for slicing
	} else {
		*queue = append(*queue, val)
	}
}

func (n *NTree) getMaxValue(depth int) int {
	if depth == 0 {
		return n.val
	}
	result := -1
	for i := -9; i <= 9; i++ {
		child := n.get(i)
		temp := child.getMaxValue(depth - 1)
		if temp > result {
			result = temp
		}
	}
	return result
}

func (s *Part2Solver) Solve() int {
	tree := NTree{
		val:      0,
		children: make(map[int]*NTree),
		flag:     true,
	}
	tree.initTree(4)
	for _, num := range s.numbers {
		queue := make([]int, 0, 4)
		prev := -1
		for range 2000 {
			unitPlace := num % 10
			if prev != -1 {
				diff := unitPlace - prev
				addToQueue(&queue, diff)
			}
			prev = unitPlace
			if len(queue) >= 4 {
				tree.deepSet(queue, unitPlace)
			}
			num = iterate(num)
		}
		tree.resetFlags(4)
	}
	return tree.getMaxValue(4)
}
