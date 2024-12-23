package day18

import (
	"github.com/kushalchordiya216/AOC2024/common/utils"
)

type Part1Solver Solver

func (s *Part1Solver) Read(path string) error {
	var err error
	s.size = 71
	s.grid, s.blocks, err = readInput(path, s.size)
	return err
}

type Node struct {
	pos  utils.Coord
	dist int
}

func (s *Part1Solver) Solve() int {
	// iterate through the first 12 entries in s.blocks and change the grid value to # for corresponding coords
	for i := 0; i < 1024; i++ {
		if !s.grid.WithinBounds(s.blocks[i]) {
			continue
		}
		s.grid[s.blocks[i].Y][s.blocks[i].X] = '#'
	}
	s.grid.Print()
	// implement a bfs search starting from 0,0 to size, size. Allowed movements are Up, Down, Right, Left,
	// same as the directions defined in utils package
	// bfs has to avoid cells with # value
	// return the number of steps required to reach size,size
	visited := make(map[utils.Coord]bool)
	queue := []Node{Node{
		pos:  utils.Coord{X: 0, Y: 0},
		dist: 0,
	}}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if curr.pos.X == s.size-1 && curr.pos.Y == s.size-1 {
			return curr.dist
		}
		if visited[curr.pos] {
			continue
		}
		visited[curr.pos] = true
		for _, dir := range []utils.Direction{utils.Up, utils.Down, utils.Left, utils.Right} {
			nextPos := curr.pos.PushForward(dir.GetOffset())
			if !s.grid.WithinBounds(nextPos) || s.grid[nextPos.Y][nextPos.X] == '#' {
				continue
			}
			newNode := Node{
				pos:  nextPos,
				dist: curr.dist + 1,
			}
			queue = append(queue, newNode)
		}
	}
	return -1
}
