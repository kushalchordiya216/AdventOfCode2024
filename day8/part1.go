package day8

type Part1Solver struct {
	nodes map[rune][]Coord
	r     int
	c     int
}

func (s *Part1Solver) Read(path string) error {
	var err error
	s.nodes, s.r, s.c, _, err = readInput(path)
	if err != nil {
		return err
	}
	return nil
}

func IsAntiNode(c1 Coord, c2 Coord, c3 Coord, l Line) bool {
	if c3.liesOn(l, c1) && c3.liesOn(l, c2) {
		d1 := c3.distance(c1)
		d2 := c3.distance(c2)
		if d1 == 2*d2 || d2 == 2*d1 {
			return true
		}
	}
	return false
}

func (s *Part1Solver) Solve() int {
	// antiNodes := getAntiNodes(s.r, s.c, s.nodes, isAntiNode)
	return 10
}
