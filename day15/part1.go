package day15

import (
	"bufio"
	"log"
	"os"

	"github.com/kushalchordiya216/AOC2024/common/utils"
)

type Part1Solver struct {
	grid     utils.Grid[rune]
	position utils.Coord
	moves    []utils.Direction
}

func (s *Part1Solver) Read(path string) error {
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
			s.grid = append(s.grid, []rune(line))
			for x, chr := range line {
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

func (s *Part1Solver) canMove(move utils.Direction) bool {
	offset := move.GetOffset()
	current := s.position
	for {
		current = current.PushForward(offset)
		if s.grid.WithinBounds(current) {
			if s.grid[current.Y][current.X] == '#' {
				return false
			}
			if s.grid[current.Y][current.X] == '.' {
				return true
			}
		} else {
			return false
		}
	}
}

// Shift shifts all grid values by offset
func (s *Part1Solver) shift(move utils.Direction) {
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
	s.grid[s.position.Y][s.position.X] = '.'
}

func (s *Part1Solver) makeMove(move utils.Direction) {
	if !s.canMove(move) {
		return
	}
	s.shift(move)
	s.position = s.position.PushForward(move.GetOffset())
}

func (s *Part1Solver) calcGPS() int {
	result := 0
	for y := range len(s.grid) {
		for x := range s.grid[y] {
			if s.grid[y][x] == 'O' {
				result += 100*y + x
			}
		}
	}
	return result
}

func (s *Part1Solver) Solve() int {
	for _, move := range s.moves {
		s.makeMove(move)
	}
	return s.calcGPS()
}
