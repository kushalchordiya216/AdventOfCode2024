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

func readInput(path string) ([][]rune, int, int, Direction, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, 0, 0, 0, err
	}
	defer file.Close()

	grid := make([][]rune, 0)

	scanner := bufio.NewScanner(file)
	lineNumber := 0
	x := 0
	y := 0
	d := Up
	for scanner.Scan() {
		line := scanner.Text()
		row := []rune(line)
		grid = append(grid, row)

		// Look for starting position in this row
		for idx, ch := range row {
			if ch != '#' && ch != '.' {
				y = idx
				x = lineNumber
				switch ch {
				case '^':
					d = Up
				case 'v':
					d = Down
				case '>':
					d = Right
				case '<':
					d = Left
				}
			}
		}
		lineNumber++
	}

	return grid, x, y, d, scanner.Err()
}

func traceGuardPath(grid [][]rune, x int, y int, d Direction) error {
	result := 0
	for {
		if grid[x][y] != 'X' {
			result += 1
		}
		grid[x][y] = 'X'
		offset, err := getOffset(d)
		if err != nil {
			return err
		}
		newX, newY := x+offset[0], y+offset[1]
		if newX < 0 || newY < 0 || newX >= len(grid) || newY >= len(grid[0]) {
			// guard has left the grid
			return nil
		}
		if grid[newX][newY] == '#' {
			d, err = rotateRight(d)
			if err != nil {
				return err
			}
		} else {
			x, y = newX, newY
		}
	}
}
