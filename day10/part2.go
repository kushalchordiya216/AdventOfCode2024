package day10

import (
	"bufio"
	"os"
	"strconv"
)

type Part2Solver struct {
	grid       [][]int
	trailheads []Coord
}

func (s *Part2Solver) Read(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	x := 0
	for scanner.Scan() {
		line := scanner.Text()
		var numbers []int
		for y, chr := range line {
			num, err := strconv.Atoi(string(chr))
			if err != nil {
				return err
			}
			numbers = append(numbers, num)
			if num == 0 {
				s.trailheads = append(s.trailheads, Coord{x: x, y: y})
			}
		}
		s.grid = append(s.grid, numbers)
		x++
	}
	return nil
}

func dfs2(head Coord, grid [][]int) int {
	stack := make([]Coord, 0)
	stack = append(stack, head)
	visited := make([][]bool, 0)
	for i := range grid {
		row := make([]bool, len(grid[i]))
		visited = append(visited, row)
	}
	result := 0
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		visited[node.x][node.y] = true
		stack = stack[:len(stack)-1]
		newNodes := explore(node, grid)
		for _, newNode := range newNodes {
			if grid[newNode.x][newNode.y] == 9 {
				result += 1
			} else {
				stack = append(stack, newNode)
			}
			visited[newNode.x][newNode.y] = true
		}
	}
	return result
}

func (s *Part2Solver) Solve() int {
	result := 0
	for _, head := range s.trailheads {
		result += dfs2(head, s.grid)
	}
	return result
}
