package day6

import (
	"bufio"
	"fmt"
	"os"

	"github.com/kushalchordiya216/AOC2024/common"
)

type Direction int

const (
	Up Direction = iota
	Down
	Right
	Left
)

func getOffset(d Direction) ([]int, error) {
	switch d {
	case Up:
		return []int{-1, 0}, nil
	case Down:
		return []int{1, 0}, nil
	case Right:
		return []int{0, 1}, nil
	case Left:
		return []int{0, -1}, nil
	default:
		return nil, &common.CustomError{Msg: fmt.Sprintf("Invalid direction: %d", d)}
	}
}

func rotateRight(d Direction) (Direction, error) {
	switch d {
	case Up:
		return Right, nil
	case Down:
		return Left, nil
	case Right:
		return Down, nil
	case Left:
		return Up, nil
	default:
		return d, &common.CustomError{Msg: fmt.Sprintf("Invalid direction: %d", d)}
	}
}

type Part1Solver struct {
	grid [][]rune
	x, y int
	d    Direction
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
		row := []rune(line)
		s.grid = append(s.grid, row)

		// Look for starting position in this row
		for y, ch := range row {
			if ch != '#' && ch != '.' {
				s.x = x
				s.y = y
				switch ch {
				case '^':
					s.d = Up
				case 'v':
					s.d = Down
				case '>':
					s.d = Right
				case '<':
					s.d = Left
				}
			}
		}
		x++
	}

	return scanner.Err()
}

func (s *Part1Solver) Solve() int {
	x, y := s.x, s.y
	d := s.d
	result := 0
	for {
		if s.grid[x][y] != 'X' {
			result += 1
		}
		s.grid[x][y] = 'X'
		offset, err := getOffset(d)
		if err != nil {
			fmt.Println(err)
			return 0
		}
		newX, newY := x+offset[0], y+offset[1]
		if newX < 0 || newY < 0 || newX >= len(s.grid) || newY >= len(s.grid[0]) {
			// guard has left the grid
			return result
		}
		if s.grid[newX][newY] == '#' {
			d, err = rotateRight(d)
			if err != nil {
				fmt.Println(err)
				return 0
			}
		} else {
			x, y = newX, newY
		}
	}
}
