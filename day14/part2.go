package day14

type Part2Solver Solver

func (s *Part2Solver) Read(path string) error {
	var err error
	s.points, err = readInput(path)
	if err != nil {
		return err
	}
	return nil
}

func hasStraightLine(points []Point, length int) bool {
	lookup := make(map[Position]bool)
	for _, point := range points {
		lookup[point.p] = true
	}
	for _, point := range points {
		for i := 1; i <= length; i++ {
			if _, ok := lookup[Position{point.p.x + i, point.p.y}]; !ok {
				break
			}
			if i == length {
				return true
			}
		}
	}
	return false
}

func (s *Part2Solver) Solve() int {
	steps := 0
	for {
		for idx := range s.points {
			s.points[idx].p.x += s.points[idx].v.x
			s.points[idx].p.y += s.points[idx].v.y
			s.points[idx].p.x = Pmod(s.points[idx].p.x, MaxCol)
			s.points[idx].p.y = Pmod(s.points[idx].p.y, MaxRow)
		}
		steps++
		if hasStraightLine(s.points, 10) {
			//PrintBoard(s.points)
			return steps
		}
	}
}
