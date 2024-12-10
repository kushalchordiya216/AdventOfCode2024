package day9

import (
	"fmt"
	"slices"
)

type Part2Solver struct {
	blocks []Block
}

func (s *Part2Solver) Read(path string) error {
	var err error
	s.blocks, err = readInput(path)
	return err
}

func (s *Part2Solver) getBlockIndexById(id int) int {
	for idx, block := range s.blocks {
		if block.id == id {
			return idx
		}
	}
	return -1
}

func (s *Part2Solver) Solve() int {
	result := 0
	multiplier := 0
	target := (len(s.blocks) - 1) / 2
	targetIdx := len(s.blocks) - 1
	for target > 0 {
		for idx, block := range s.blocks[:targetIdx] {
			if block.category == Space && block.size >= s.blocks[targetIdx].size {
				s.blocks[idx].size -= s.blocks[targetIdx].size
				s.blocks[targetIdx].category = Space
				newBlock := Block{size: s.blocks[targetIdx].size, id: s.blocks[targetIdx].id, category: File}
				s.blocks = slices.Insert(s.blocks, idx, newBlock)
				break
			}
		}
		target--
		targetIdx = s.getBlockIndexById(target)
		if targetIdx == -1 {
			fmt.Printf("Failed to get block with id :%d", target)
			return 0
		}
	}
	for _, block := range s.blocks {
		if block.category == Space {
			multiplier += block.size
		} else {
			for i := 0; i < block.size; i++ {
				result += multiplier * block.id
				multiplier++
			}
		}
	}
	return result
}
