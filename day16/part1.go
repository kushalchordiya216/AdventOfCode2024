package day16

import (
	"container/heap"
	"github.com/kushalchordiya216/AOC2024/common/utils"
)

type Part1Solver Solver

func (s *Part1Solver) Read(path string) error {
	var err error
	s.grid, s.start, s.end, err = readInput(path)
	return err
}

func (s *Part1Solver) dijkstra(grid utils.Grid[rune], startPos utils.Coord, endPos utils.Coord) (int, bool) {
	startState := State{Position: startPos, Dir: utils.Right, Cost: 0, Path: []utils.Coord{startPos}}
	visited := make(map[int]bool)
	pq := make(PriorityQueue, 0)
	heap.Push(&pq, &startState)

	for pq.Len() > 0 {
		state := heap.Pop(&pq).(*State)
		if visited[state.Hash()] {
			continue
		}
		visited[state.Hash()] = true

		if state.Position == endPos {
			return state.Cost, true
		}

		for _, neighbour := range getNeighbours(grid, *state) {
			heap.Push(&pq, &neighbour)
		}
	}
	return 0, false // Path not found
}

func (s *Part1Solver) Solve() int {
	cost, found := s.dijkstra(s.grid, s.start, s.end)

	if found {
		return cost
	}
	return 0
}
