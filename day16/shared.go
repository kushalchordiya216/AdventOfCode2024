package day16

import (
	"bufio"
	"github.com/kushalchordiya216/AOC2024/common/utils"
	"log"
	"os"
)

type Solver struct {
	grid  utils.Grid[rune]
	start utils.Coord
	end   utils.Coord
}

type State struct {
	Position utils.Coord
	Dir      utils.Direction
	Cost     int
	Path     []utils.Coord
}

func (s *State) Hash() int {
	return (s.Position.Y << 16) | (s.Position.X << 2) | int(s.Dir)
}

type PriorityQueue []*State

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Cost < pq[j].Cost
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*State)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func readInput(path string) (utils.Grid[rune], utils.Coord, utils.Coord, error) {
	file, err := os.Open(path)
	grid := utils.Grid[rune]{}
	start := utils.Coord{0, 0}
	end := utils.Coord{0, 0}
	if err != nil {
		return utils.Grid[rune]{}, utils.Coord{}, utils.Coord{}, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)
	scanner := bufio.NewScanner(file)
	lineNum := 0
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
		for x, chr := range line {
			if chr == 'S' {
				start = utils.Coord{
					X: x,
					Y: lineNum,
				}
			}
			if chr == 'E' {
				end = utils.Coord{
					X: x,
					Y: lineNum,
				}
			}
		}
		lineNum++
	}
	return grid, start, end, nil
}

func getNeighbours(grid utils.Grid[rune], state State) []State {
	neighbours := make([]State, 0)
	for _, direction := range []utils.Direction{utils.Up, utils.Right, utils.Down, utils.Left} {
		if direction == state.Dir.Reverse() {
			continue
		}
		newPos := state.Position.PushForward(direction.GetOffset())
		if grid[newPos.Y][newPos.X] == '#' {
			continue
		}
		newPath := make([]utils.Coord, 0)
		newPath = append(newPath, state.Path...)
		newPath = append(newPath, newPos)

		neighbour := State{
			Position: newPos,
			Dir:      direction,
			Cost:     state.Cost + 1,
			Path:     newPath,
		}
		if direction != state.Dir {
			neighbour.Cost += 1000
		}
		neighbours = append(neighbours, neighbour)
	}
	return neighbours
}
