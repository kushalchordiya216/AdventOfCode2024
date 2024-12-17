package day14

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

const MaxRow = 103
const MaxCol = 101
const MiddleRow = 51
const MiddleCol = 50

type Position struct {
	x, y int
}

type Velocity struct {
	x, y int
}

type Point struct {
	p Position
	v Velocity
}

func readInput(path string) ([]Point, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	points := make([]Point, 0)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)
	scanner := bufio.NewScanner(file)
	pattern := regexp.MustCompile(`-?\d+`)
	for scanner.Scan() {
		line := scanner.Text()
		numbers := pattern.FindAllString(line, -1)
		x, err := strconv.Atoi(numbers[0])
		if err != nil {
			return nil, err
		}
		y, err := strconv.Atoi(numbers[1])
		if err != nil {
			return nil, err
		}
		x2, err := strconv.Atoi(numbers[2])
		if err != nil {
			return nil, err
		}
		y2, err := strconv.Atoi(numbers[3])
		if err != nil {
			return nil, err
		}
		points = append(points, Point{Position{x, y}, Velocity{x2, y2}})
	}
	return points, nil
}

func Pmod(a, b int) int {
	m := a % b
	if m == 0 {
		return 0
	}
	if a < 0 && b < 0 {
		m -= b
	}
	if a < 0 && b > 0 {
		m += b
	}
	return m
}

type Solver struct {
	points []Point
}
