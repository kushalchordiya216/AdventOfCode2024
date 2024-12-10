package day9

type Part1Solver struct {
	blocks []Block
}

func (s *Part1Solver) Read(path string) error {
	var err error
	s.blocks, err = readInput(path)
	return err
}

func (s *Part1Solver) Solve() int {
	result := 0
	multiplier := 0
	current := 0
	for current < len(s.blocks) {
		block := s.blocks[current]
		if block.category == File {
			for i := 0; i < block.size; i++ {
				result += multiplier * block.id
				multiplier++
			}
		} else {
			for i := 0; i < block.size; i++ {
				end := len(s.blocks) - 1
				result += multiplier * s.blocks[end].id
				multiplier++
				s.blocks[end].size--
				if s.blocks[end].size == 0 {
					s.blocks = s.blocks[:len(s.blocks)-2]
				}
			}
		}
		current += 1
	}
	return result
}
