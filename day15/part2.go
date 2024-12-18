package day15

import (
	"bufio"
	"github.com/kushalchordiya216/AOC2024/common/utils"
	"log"
	"os"
)

type Part2Solver struct {
	grid     utils.Grid[rune]
	position utils.Coord
	moves    []utils.Direction
}

func (s *Part2Solver) Read(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)
	scanner := bufio.NewScanner(file)
	lineNum := 0
	flag := false
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			flag = true
		}
		if flag {
			for _, chr := range line {
				switch chr {
				case '<':
					s.moves = append(s.moves, utils.Left)
					break
				case '>':
					s.moves = append(s.moves, utils.Right)
					break
				case '^':
					s.moves = append(s.moves, utils.Up)
					break
				case 'v':
					s.moves = append(s.moves, utils.Down)
					break
				}
			}
		} else {
			chars := []rune(line)
			row := make([]rune, 0)
			for _, chr := range chars {
				switch chr {
				case '#':
					row = append(row, '#', '#')
					break
				case '.':
					row = append(row, '.', '.')
					break
				case 'O':
					row = append(row, '[', ']')
					break
				case '@':
					row = append(row, '@', '.')
					break
				}
			}
			s.grid = append(s.grid, row)
			for x, chr := range row {
				if chr == '@' {
					s.position = utils.Coord{
						X: x,
						Y: lineNum,
					}
				}
			}
		}
		lineNum++
	}
	return nil
}

func (s *Part2Solver) canMove(move utils.Direction) bool {
	offset := move.GetOffset()
	current := s.position
	queue := []utils.Coord{current}
	visited := make(map[utils.Coord]bool)
	for len(queue) > 0 {
		current = queue[0]
		queue = queue[1:]
		if visited[current] {
			continue
		}
		visited[current] = true
		current = current.PushForward(offset)
		if s.grid.WithinBounds(current) {
			if s.grid[current.Y][current.X] == '[' {
				queue = append(queue, current)
				if move == utils.Up || move == utils.Down {
					neighbor := current.PushForward(utils.Coord{X: 1, Y: 0})
					queue = append(queue, neighbor)
				}
			} else if s.grid[current.Y][current.X] == ']' {
				queue = append(queue, current)
				if move == utils.Up || move == utils.Down {
					neighbor := current.PushForward(utils.Coord{X: -1, Y: 0})
					queue = append(queue, neighbor)
				}
			} else if s.grid[current.Y][current.X] == '#' {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func (s *Part2Solver) recursiveVerticalShift(move utils.Direction, current utils.Coord, visited map[utils.Coord]bool) {
	var next, adjacentNext utils.Coord
	if s.grid[current.Y][current.X] == '.' || visited[current] {
		return
	}
	visited[current] = true
	next = current.PushForward(move.GetOffset())
	if s.grid[next.Y][next.X] == '[' {
		adjacentNext = next.PushForward(utils.Coord{X: 1, Y: 0})
		s.recursiveVerticalShift(move, adjacentNext, visited)
	} else if s.grid[next.Y][next.X] == ']' {
		adjacentNext = next.PushForward(utils.Coord{X: -1, Y: 0})
		s.recursiveVerticalShift(move, adjacentNext, visited)
	}
	s.recursiveVerticalShift(move, next, visited)
	s.grid[next.Y][next.X] = s.grid[current.Y][current.X]
	s.grid[current.Y][current.X] = '.'
	return
}

// Shift shifts all grid values by offset
func (s *Part2Solver) shift(move utils.Direction) {
	if move == utils.Left || move == utils.Right { // Horizontal shift is unaffected by larger boxes
		offset := move.GetOffset()
		current := s.position
		for {
			current = current.PushForward(offset)
			if s.grid[current.Y][current.X] == '.' {
				break
			}
		}
		for current != s.position {
			prev := current.PushReverse(offset)
			s.grid[current.Y][current.X] = s.grid[prev.Y][prev.X]
			current = prev
		}
	} else {
		visited := make(map[utils.Coord]bool)
		s.recursiveVerticalShift(move, s.position, visited)
	}
	s.grid[s.position.Y][s.position.X] = '.'
}

func (s *Part2Solver) makeMove(move utils.Direction) {
	if !s.canMove(move) {
		return
	}
	s.shift(move)
	s.position = s.position.PushForward(move.GetOffset())
}

func (s *Part2Solver) calcGPS() int {
	result := 0
	for y := range len(s.grid) {
		for x := range s.grid[y] {
			if s.grid[y][x] == '[' {
				result += 100*y + x
			}
		}
	}
	return result
}

func (s *Part2Solver) Solve() int {
	for _, move := range s.moves {
		s.makeMove(move)
	}
	return s.calcGPS()
}
