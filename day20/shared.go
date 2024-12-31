package day20

import (
	"bufio"
	"github.com/kushalchordiya216/AOC2024/common/utils"
	"log"
	"math"
	"os"
)

type Solver struct {
	grid  utils.Grid[byte]
	start utils.Coord
	end   utils.Coord
}

func readInput(path string) (Solver, error) {
	s := Solver{}
	file, err := os.Open(path)
	if err != nil {
		return s, err
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := []byte(line)
		s.grid = append(s.grid, row)
		for i, chr := range line {
			if chr == 'S' {
				s.start = utils.Coord{X: i, Y: len(s.grid) - 1}
			}
			if chr == 'E' {
				s.end = utils.Coord{X: i, Y: len(s.grid) - 1}
			}
		}
	}
	return s, nil
}

func (s *Solver) dfs(curr utils.Coord, cost int, costs map[utils.Coord]int) {
	costs[curr] = cost
	if curr == s.end {
		return
	}
	for _, direction := range []utils.Direction{utils.Up, utils.Down, utils.Left, utils.Right} {
		next := curr.PushForward(direction.GetOffset())
		if _, ok := costs[next]; ok {
			continue
		}
		if s.grid.WithinBounds(next) && s.grid[next.Y][next.X] != '#' {
			s.dfs(next, cost+1, costs)
			break
		}
	}
}

func (s *Solver) findAllCheats(maxPathLen int, costs map[utils.Coord]int, minTimeSaved int) int {
	counter := 0
	for pos, val := range costs {
		for pos2, val2 := range costs {
			distance := int(math.Abs(float64(pos.X-pos2.X)) + math.Abs(float64(pos.Y-pos2.Y)))
			if distance <= maxPathLen && val2-(val+distance) >= minTimeSaved {
				counter++
			}
		}
	}
	return counter
}
