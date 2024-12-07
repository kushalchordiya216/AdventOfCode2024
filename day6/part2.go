package day6

type Part2Solver struct {
	grid [][]rune
	x, y int
	d    Direction
}

func (s *Part2Solver) Read(path string) error {
	solver := Part1Solver{grid: s.grid, x: s.x, y: s.y, d: s.d}
	err := solver.Read(path)
	if err != nil {
		return err
	}
	s.grid = solver.grid
	s.x = solver.x
	s.y = solver.y
	s.d = solver.d
	return nil
}

func checkIfLooping(grid [][]rune, x int, y int, d Direction) bool {
	visited := make([][]map[Direction]bool, len(grid))
	for i := range visited {
		visited[i] = make([]map[Direction]bool, len(grid[0]))
		for j := range visited[i] {
			visited[i][j] = make(map[Direction]bool)
		}
	}

	for {
		// Check if we've been here before with same direction
		if visited[x][y][d] {
			return true
		}
		// Mark current position and direction as visited
		visited[x][y][d] = true

		offset, err := getOffset(d)
		if err != nil {
			return false
		}

		newX, newY := x+offset[0], y+offset[1]
		if newX < 0 || newY < 0 || newX >= len(grid) || newY >= len(grid[0]) {
			return false // Left the grid
		}

		if grid[newX][newY] == '#' { // Wall
			d, err = rotateRight(d)
			if err != nil {
				return false
			}
		} else {
			x, y = newX, newY
		}
	}
}

func (s *Part2Solver) Solve() int {
	solver := Part1Solver{grid: s.grid, x: s.x, y: s.y, d: s.d}
	solver.Solve() // modify the board to mark guards visited path
	s.grid = solver.grid
	result := 0
	for i := range s.grid {
		for j := range s.grid[i] {
			if s.grid[i][j] == 'X' {
				s.grid[i][j] = '#'
				if checkIfLooping(s.grid, s.x, s.y, s.d) {
					result++
				}
				s.grid[i][j] = 'X'
			}
		}
	}
	return result
}