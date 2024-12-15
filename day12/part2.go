package day12

import (
	"bufio"
	"fmt"
	"os"
)

type Part2Solver struct {
	grid [][]rune
}

func (s *Part2Solver) Read(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := []rune(line)
		s.grid = append(s.grid, row)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return err
	}
	return nil
}

func isOutOfBounds(c Coord, grid [][]rune) bool {
	if c.x < 0 || c.x >= len(grid) || c.y < 0 || c.y >= len(grid[0]) {
		return true
	}
	return false
}

// Returns number of corners in current block
func getCornerCount(c Coord, grid [][]rune) int {
	/* Region corners are also represented as co-ordinates, but there are different from the co-ordinates used to represent plots in a region
	   consider a region like so
	   +----+----+
	   | A  | A  |
	   +----+----+
	   | A  | A  |
	   +----+----+
		Each A represents a region plot
		Co-ordinates of the region plots are (0,0), (0,1), (1,0), (1,1)
		+ represent all potential region corners. They are also represented with a co-ordinate system, however, their co-ordinates are slightly different
		the + in the above diagram have the co-ordinates
		(0,0), (0,1), (0,2), (1,0), (1,1), (1,2), (2,0), (2,1), (2,2)
		of these, the ones which classify as region corners are
		(0,0), (0,2), (2,0), (2,2)
		A particular point is a outerCorner if it lies on one plot from a given region and 2 adjacent plots are of some other regions
		It's an innser corner if it lies on 3 plots of a given region and one plot of something else
		The above diagram only illustrates outercorners but inner corners are just as important for the code
		For a polygon that is constructed using plots/blocks as such, number of corners is always equal to number of sides
	*/
	counter := 0
	offsets := [][]int{{-1, -1}, {-1, 1}, {1, -1}, {1, 1}}
	for _, offset := range offsets {
		counter += isOuterCorner(c, grid, offset[0], offset[1])
		counter += isInnerCorner(c, grid, offset[0], offset[1])
	}
	return counter
}

func isOuterCorner(c Coord, grid [][]rune, d1, d2 int) int {
	/*
		We're aiming for this case
		+----+----+
		| ?  | B  |
		+----+----+
		| B  | A  |
		+----+----+
		the vertice/corner at (1,1) in this case is a an inner corner
		The vale of the block at (0,0) doesn't really matter
	*/
	adjacents := []Coord{{x: c.x + d1, y: c.y}, {x: c.x, y: c.y + d2}}
	for _, adjacent := range adjacents {
		if !isOutOfBounds(adjacent, grid) && grid[adjacent.x][adjacent.y] == grid[c.x][c.y] {
			return 0
		}
	}
	return 1
}

func isInnerCorner(c Coord, grid [][]rune, d1, d2 int) int {
	/*
		We're aiming for this case
		+----+----+
		| B  | A  |
		+----+----+
		| A  | A  |
		+----+----+
		the vertice/corner at (1,1) in this case is a an inner corner
	*/
	opposite := Coord{x: c.x + d1, y: c.y + d2}
	adjacents := []Coord{{x: c.x, y: c.y + d2}, {x: c.x + d1, y: c.y}}
	if isOutOfBounds(opposite, grid) || grid[opposite.x][opposite.y] == grid[c.x][c.y] {
		return 0
	}
	for _, adjacent := range adjacents {
		if isOutOfBounds(adjacent, grid) || grid[adjacent.x][adjacent.y] != grid[c.x][c.y] {
			return 0
		}
	}
	return 1
}

func (s *Part2Solver) dfs2(explored [][]bool, c Coord) int {
	var area, corners int
	stack := []Coord{c}
	visited := make([][]bool, 0)
	for i := 0; i < len(s.grid); i++ {
		v_row := make([]bool, len(s.grid[0]))
		visited = append(visited, v_row)
	}
	offsets := []Coord{{x: 0, y: 1}, {x: 0, y: -1}, {x: 1, y: 0}, {x: -1, y: 0}}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		explored[node.x][node.y] = true
		if visited[node.x][node.y] {
			continue
		}
		area += 1
		visited[node.x][node.y] = true
		corners += getCornerCount(node, s.grid)
		for _, offset := range offsets {
			newNode := Coord{node.x + offset.x, node.y + offset.y}
			if !isOutOfBounds(newNode, s.grid) && s.grid[newNode.x][newNode.y] == s.grid[c.x][c.y] {
				stack = append(stack, newNode)
			}
		}
	}
	return corners * area
}

func (s *Part2Solver) Solve() int {
	explored := make([][]bool, 0)
	for i := 0; i < len(s.grid); i++ {
		v_row := make([]bool, len(s.grid[0]))
		explored = append(explored, v_row)
	}
	result := 0
	for x, row := range s.grid {
		for y := range row {
			if !explored[x][y] {
				result += s.dfs2(explored, Coord{x: x, y: y})
			}
		}
	}
	return result
}
