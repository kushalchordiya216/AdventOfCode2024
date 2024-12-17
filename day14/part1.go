package day14

type Part1Solver Solver

func (s *Part1Solver) Read(path string) error {
	var err error
	s.points, err = readInput(path)
	if err != nil {
		return err
	}
	return nil
}

func (s *Part1Solver) Solve() int {
	quarters := []int{0, 0, 0, 0}
	for idx := range s.points {
		s.points[idx].p.x += 100 * s.points[idx].v.x
		s.points[idx].p.y += 100 * s.points[idx].v.y
		s.points[idx].p.x = Pmod(s.points[idx].p.x, MaxCol)
		s.points[idx].p.y = Pmod(s.points[idx].p.y, MaxRow)
		if s.points[idx].p.x < MiddleCol && s.points[idx].p.y < MiddleRow {
			quarters[0]++
		} else if s.points[idx].p.x > MiddleCol && s.points[idx].p.y < MiddleRow {
			quarters[1]++
		} else if s.points[idx].p.x > MiddleCol && s.points[idx].p.y > MiddleRow {
			quarters[2]++
		} else if s.points[idx].p.x < MiddleCol && s.points[idx].p.y > MiddleRow {
			quarters[3]++
		}
	}
	result := 1
	for i := 0; i < 4; i++ {
		result *= quarters[i]
	}
	return result
}
