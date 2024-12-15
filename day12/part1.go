package day12

import (
	"bufio"
	"os"
)

type Part1Solver struct {
	grid [][]rune
}

type Coord struct {
	x, y int
}

func (s *Part1Solver) Read(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := []rune(line)
		s.grid = append(s.grid, row)
	}
	return nil
}

func dfs(grid [][]rune, explored [][]bool, c Coord) int {
	var area, peri int
	stack := []Coord{c}
	visited := make([][]bool, 0)
	for i := 0; i < len(grid); i++ {
		v_row := make([]bool, len(grid[0]))
		visited = append(visited, v_row)
	}
	offsets := []Coord{{x: 0, y: 1}, {x: 0, y: -1}, {x: 1, y: 0}, {x: -1, y: 0}}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		explored[node.x][node.y] = true
		if visited[node.x][node.y] {
			continue
		}
		area += 1
		visited[node.x][node.y] = true
		for _, offset := range offsets {
			newX, newY := node.x+offset.x, node.y+offset.y
			if newX < 0 || newX >= len(grid) || newY < 0 || newY >= len(grid) {
				peri += 1
				continue
			}
			if grid[newX][newY] == grid[c.x][c.y] {
				stack = append(stack, Coord{x: newX, y: newY})
			} else {
				peri++
			}
		}
	}
	return peri * area
}

func (s *Part1Solver) Solve() int {
	explored := make([][]bool, 0)
	for i := 0; i < len(s.grid); i++ {
		v_row := make([]bool, len(s.grid[0]))
		explored = append(explored, v_row)
	}
	result := 0
	for x, row := range s.grid {
		for y := range row {
			if !explored[x][y] {
				result += dfs(s.grid, explored, Coord{x: x, y: y})
			}
		}
	}
	return result
}
