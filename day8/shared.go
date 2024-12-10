package day8

import (
	"bufio"
	"math"
	"os"
)

type Coord struct {
	x float64
	y float64
}

type Line struct {
	m float64
	b float64
}

func (c *Coord) distance(c1 Coord) float64 {
	//return distance between c and c1 using co-ordinate geometry
	dx := c.x - c1.x
	dy := c.y - c1.y
	return math.Sqrt(dx*dx + dy*dy)
}

func (c *Coord) liesOn(l Line, c1 Coord) bool {
	// check is a given co-ordinate point (x, y) lies on a given line y = mx + c and return if true
	return (c.y-c1.y)/(c.x-c1.x) == l.m
}

func getLine(c1 Coord, c2 Coord) Line {
	// return a line passing through both coords
	if c2.x == c1.x {
		return Line{m: math.Inf(1), b: float64(c1.x)}
	} else {
		m := (c2.y - c1.y) / (c2.x - c1.x)
		b := -m*c1.x + c1.y
		return Line{m: m, b: b}
	}
}

func readInput(path string) (map[rune][]Coord, int, int, [][]rune, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, 0, 0, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	nodes := make(map[rune][]Coord)
	grid := make([][]rune, 0)
	x := 0
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]rune, 0)
		for y, char := range line {
			if char != '.' {
				nodes[char] = append(nodes[char], Coord{x: float64(x), y: float64(y)})
			}
			row = append(row, char)
		}
		grid = append(grid, row)
		x++
	}
	r := len(grid)
	c := len(grid[0])
	if err := scanner.Err(); err != nil {
		return nil, 0, 0, nil, err
	}

	return nodes, r, c, grid, nil
}

type Checker func(Coord, Coord, Coord, Line) bool

func getAntiNodes(r int, c int, nodes map[rune][]Coord, grid [][]rune, checker Checker) []Coord {
	antiNodes := make(map[Coord]bool, 0)
	for _, coords := range nodes {
		for i := 0; i < len(coords); i++ {
			for j := i; j < len(coords); j++ {
				if i == j {
					continue
				}
				c1 := coords[i]
				c2 := coords[j]
				l := getLine(c1, c2)
				for x := 0; x < r; x++ {
					for y := 0; y < c; y++ {
						c3 := Coord{x: float64(x), y: float64(y)}
						if checker(c1, c2, c3, l) {
							antiNodes[c3] = true
							grid[x][y] = '#'
						}
					}
				}
			}
		}
	}
	keys := make([]Coord, len(antiNodes))

	i := 0
	for k := range antiNodes {
		keys[i] = k
		i++
	}
	return keys
}
