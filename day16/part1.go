package day16

import (
	"bufio"
	"log"
	"os"

	"github.com/kushalchordiya216/AOC2024/common/utils"
)

type Part1Solver struct {
	grid  utils.Grid[rune]
	start utils.Coord
	end   utils.Coord
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
	for scanner.Scan() {
		line := scanner.Text()
		s.grid = append(s.grid, []rune(line))
		for x, chr := range line {
			if chr == 'S' {
				s.start = utils.Coord{
					X: x,
					Y: lineNum,
				}
			}
			if chr == 'E' {
				s.end = utils.Coord{
					X: x,
					Y: lineNum,
				}
			}
		}
		lineNum++
	}
	return nil
}

func (s *Part1Solver) Solve() int {

	return 0
}
