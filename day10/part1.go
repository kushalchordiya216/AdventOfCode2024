package day10

import (
	"bufio"
	"os"
	"strconv"
)

type Coord struct {
	x, y int
}

type Part1Solver struct {
	grid       [][]int
	trailheads []Coord
}

func (s *Part1Solver) Read(path string) error {
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

func explore(c Coord, grid [][]int) []Coord {
	offsets := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	nodes := make([]Coord, 0)
	for _, offset := range offsets {
		x, y := c.x+offset[0], c.y+offset[1]
		if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) {
			continue
		}
		if grid[x][y] != grid[c.x][c.y]+1 {
			continue
		}
		nodes = append(nodes, Coord{x: x, y: y})
	}
	return nodes
}

func dfs(head Coord, grid [][]int) int {
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
			if grid[newNode.x][newNode.y] == 9 && !visited[newNode.x][newNode.y] {
				result += 1
			} else {
				stack = append(stack, newNode)
			}
			visited[newNode.x][newNode.y] = true
		}
	}
	return result
}

func (s *Part1Solver) Solve() int {
	result := 0
	for _, head := range s.trailheads {
		result += dfs(head, s.grid)
	}
	return result
}
