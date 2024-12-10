package day8

type Part2Solver struct {
	nodes map[rune][]Coord
	r     int
	c     int
	grid  [][]rune
}

func (s *Part2Solver) Read(path string) error {
	var err error
	s.nodes, s.r, s.c, s.grid, err = readInput(path)
	if err != nil {
		return err
	}
	return nil
}

func isAntiNode2(c1 Coord, c2 Coord, c3 Coord, l Line) bool {
	return c3.liesOn(l, c1) || c3.liesOn(l, c2)
}

func (s *Part2Solver) Solve() int {
	antiNodes := getAntiNodes(s.r, s.c, s.nodes, s.grid, isAntiNode2)
	return len(antiNodes)
}
