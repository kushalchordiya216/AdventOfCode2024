package day16

import (
	"container/heap"
	"github.com/kushalchordiya216/AOC2024/common/utils"
)

type Part2Solver Solver

func (s *Part2Solver) Read(path string) error {
	var err error
	s.grid, s.start, s.end, err = readInput(path)
	return err
}

func (s *Part2Solver) dijkstra(grid utils.Grid[rune], startPos utils.Coord, endPos utils.Coord, minCost int) map[utils.Coord]bool {
	startState := State{Position: startPos, Dir: utils.Right, Cost: 0, Path: []utils.Coord{startPos}}
	visited := make(map[int]int)
	pq := make(PriorityQueue, 0)
	heap.Push(&pq, &startState)
	uniquePositions := make(map[utils.Coord]bool)

	for pq.Len() > 0 {
		state := heap.Pop(&pq).(*State)

		if cost, ok := visited[state.Hash()]; ok && cost < state.Cost {
			continue
		}
		visited[state.Hash()] = state.Cost
		if state.Position == endPos {
			if state.Cost > minCost {
				break
			}
			minCost = state.Cost
			for _, pos := range state.Path {
				uniquePositions[pos] = true
			}
		}
		for _, neighbour := range getNeighbours(grid, *state) {
			heap.Push(&pq, &neighbour)
		}
	}
	return uniquePositions
}

func (s *Part2Solver) Solve() int {
	result := s.dijkstra(s.grid, s.start, s.end, 1<<31-1)
	for pos := range result {
		s.grid[pos.Y][pos.X] = 'O'
	}
	s.grid.Print()
	return len(result)
}
