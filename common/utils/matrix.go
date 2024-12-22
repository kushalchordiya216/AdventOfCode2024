package utils

import "fmt"

type Coord struct {
	X int
	Y int
}

func (c Coord) PushForward(offset Coord) Coord {
	return Coord{c.X + offset.X, c.Y + offset.Y}
}

func (c Coord) PushReverse(offset Coord) Coord {
	return Coord{c.X - offset.X, c.Y - offset.Y}
}

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

func (d Direction) String() string {
	switch d {
	case Up:
		return "Up"
	case Down:
		return "Down"
	case Left:
		return "Left"
	case Right:
		return "Right"
	default:
		return fmt.Sprint("Unknown")
	}
}

func (d Direction) RotateRight() Direction {
	switch d {
	case Up:
		return Right
	case Down:
		return Left
	case Left:
		return Up
	case Right:
		return Down
	default:
		return Up
	}
}

func (d Direction) RotateLeft() Direction {
	switch d {
	case Up:
		return Left
	case Down:
		return Right
	case Left:
		return Down
	case Right:
		return Up
	default:
		return Up
	}
}

func (d Direction) Reverse() Direction {
	switch d {
	case Up:
		return Down
	case Down:
		return Up
	case Left:
		return Right
	case Right:
		return Left
	default:
		return Up
	}
}

func (d Direction) GetOffset() Coord {
	switch d {
	case Up:
		return Coord{
			X: 0,
			Y: -1,
		}
	case Down:
		return Coord{
			X: 0,
			Y: 1,
		}
	case Left:
		return Coord{
			X: -1,
			Y: 0,
		}
	case Right:
		return Coord{
			X: 1,
			Y: 0,
		}
	default:
		return Coord{
			X: 0,
			Y: 0,
		}
	}
}

type Grid[T int | rune] [][]T

func (g Grid[T]) WithinBounds(coord Coord) bool {
	return coord.X >= 0 && coord.X < len(g[0]) && coord.Y >= 0 && coord.Y < len(g)
}

func (g Grid[T]) Print() {
	for _, row := range g {
		for _, col := range row {
			fmt.Printf("%c", col)
		}
		fmt.Println()
	}
}
