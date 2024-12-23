package day18

import (
	"fmt"
	"github.com/kushalchordiya216/AOC2024/common/utils"
)

type Part2Solver Solver

func (s *Part2Solver) Read(path string) error {
	var err error
	s.size = 71
	s.grid, s.blocks, err = readInput(path, s.size)
	return err
}

func (s *Part2Solver) resetGrid() {
	s.grid = utils.Grid[rune]{}
	for i := 0; i < s.size; i++ {
		row := make([]rune, s.size)
		for j := 0; j < s.size; j++ {
			row[j] = '.'
		}
		s.grid = append(s.grid, row)
	}
}

func (s *Part2Solver) canEscape(idx int) bool {
	s.resetGrid()
	i := 0
	for i <= idx {
		s.grid[s.blocks[i].Y][s.blocks[i].X] = '#'
		i++
	}
	visited := make(map[utils.Coord]bool)
	queue := []Node{Node{
		pos:  utils.Coord{X: 0, Y: 0},
		dist: 0,
	}}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if curr.pos.X == s.size-1 && curr.pos.Y == s.size-1 {
			return true
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
	return false
}

func (s *Part2Solver) Solve() int {
	lo := 0
	hi := len(s.blocks) - 1
	for lo < hi {
		mid := (lo + hi + 1) / 2
		if s.canEscape(mid) {
			lo = mid
			if !s.canEscape(mid + 1) {
				fmt.Println(s.blocks[mid+1])
				return 0
			}
		} else {
			if s.canEscape(mid - 1) {
				fmt.Println(s.blocks[mid])
				return 0
			}
			hi = mid - 1
		}
	}
	return -1
}
