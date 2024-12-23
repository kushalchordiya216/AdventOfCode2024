package day18

import (
	"bufio"
	"github.com/kushalchordiya216/AOC2024/common"
	"github.com/kushalchordiya216/AOC2024/common/utils"
	"log"
	"os"
	"strconv"
	"strings"
)

type Solver struct {
	grid   utils.Grid[rune]
	blocks []utils.Coord
	size   int
}

func readInput(path string, size int) (utils.Grid[rune], []utils.Coord, error) {
	file, err := os.Open(path)
	grid := utils.Grid[rune]{}
	blocks := []utils.Coord{}
	if err != nil {
		return utils.Grid[rune]{}, []utils.Coord{}, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		splits := strings.Split(line, ",")
		if len(splits) != 2 {
			return utils.Grid[rune]{}, []utils.Coord{}, &common.CustomError{Msg: "Expected 2 splits"}
		}
		x, err := strconv.Atoi(strings.TrimSpace(splits[0]))
		if err != nil {
			return utils.Grid[rune]{}, []utils.Coord{}, err
		}
		y, err := strconv.Atoi(strings.TrimSpace(splits[1]))
		if err != nil {
			return utils.Grid[rune]{}, []utils.Coord{}, err
		}
		blocks = append(blocks, utils.Coord{X: x, Y: y})
	}

	for i := 0; i < size; i++ {
		row := make([]rune, size)
		for j := 0; j < size; j++ {
			row[j] = '.'
		}
		grid = append(grid, row)
	}
	return grid, blocks, nil
}
