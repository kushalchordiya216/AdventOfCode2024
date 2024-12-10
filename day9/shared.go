package day9

import (
	"bufio"
	"os"
	"strconv"
)

type BlockCategory int

const (
	File BlockCategory = iota
	Space
)

type Block struct {
	size     int
	category BlockCategory
	id       int
}

func readInput(path string) ([]Block, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	blocks := make([]Block, 0)
	for scanner.Scan() {
		line := scanner.Text()
		for idx, ch := range line {
			num, err := strconv.Atoi(string(ch))
			if err != nil {
				return nil, err
			}
			if idx%2 == 0 {
				blocks = append(blocks, Block{size: num, category: File, id: idx / 2})
			} else {
				blocks = append(blocks, Block{size: num, category: Space})
			}

		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return blocks, nil
}

func (b Block) String() string {
	result := ""
	if b.category == Space {
		for i := 0; i < b.size; i++ {
			result += "."
		}
	} else {
		for i := 0; i < b.size; i++ {
			result += strconv.Itoa(b.id)
		}
	}
	return result
}
