package day1

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/kushalchordiya216/AOC2024/common"
)

type Part1Solver struct {
	list1, list2 IntHeap
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (s *Part1Solver) Read(path string) error {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return err
	}
	defer file.Close()
	heap.Init(&s.list1)
	heap.Init(&s.list2)

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
		num2, err2 := strconv.Atoi(nums[len(nums)-1])
		if err2 != nil {
			return err2
		}
		heap.Push(&s.list1, num1)
		heap.Push(&s.list2, num2)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return err
	}

	return nil
}

func (s *Part1Solver) Solve() int {
	totalSum := 0
	n := s.list1.Len()
	for i := 0; i < n; i++ {
		val1 := heap.Pop(&s.list1).(int)
		val2 := heap.Pop(&s.list2).(int)
		diff := val1 - val2
		if diff < 0 {
			diff = -diff
		}
		totalSum += diff
	}
	return totalSum
}
