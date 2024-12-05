package day4

type Part2Solver struct {
	grid [][]rune
}

func (s *Part2Solver) Read(path string) error {
	solver := &Part1Solver{grid: s.grid}
	err := solver.Read(path)
	if err != nil {
		return err
	}
	s.grid = solver.grid
	return nil
}

func check(s *Part2Solver, x int, y int) bool {
	MCount := 0
	SCount := 0
	offsets := [][]int{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
	corners := [][]int{}
	for _, offset := range offsets {
		newX := x + offset[0]
		newY := y + offset[1]
		if newX < 0 || newY < 0 || newX >= len(s.grid) || newY >= len(s.grid[0]) {
			return false
		}
		if s.grid[newX][newY] == 'M' {
			MCount += 1
		} else if s.grid[newX][newY] == 'S' {
			SCount += 1
		} else {
			return false
		}
		corners = append(corners, []int{newX, newY})
	}
	// this ensures we have the right amount of M and S
	// now only need to check their relative positioning at the corner of the 'X' shapre
	if MCount != 2 || SCount != 2 {
		return false
	}
	/*
		checks for the existence of the such shapes
		M        S
		    A
		S      	 M
	*/
	if s.grid[corners[0][0]][corners[0][1]] == s.grid[corners[3][0]][corners[3][1]] || s.grid[corners[1][0]][corners[1][1]] == s.grid[corners[2][0]][corners[2][1]] {
		return false
	}
	return true
}

func (s *Part2Solver) Solve() int {
	result := 0
	for x := range s.grid {
		for y := range s.grid[x] {
			if s.grid[x][y] == 'A' {
				if check(s, x, y) {
					result += 1
				}
			}
		}
	}
	return result
}
