package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/kushalchordiya216/AOC2024/common"
	"github.com/kushalchordiya216/AOC2024/day1"
	"github.com/kushalchordiya216/AOC2024/day2"
	"github.com/kushalchordiya216/AOC2024/day3"
	"github.com/kushalchordiya216/AOC2024/day4"
	"github.com/kushalchordiya216/AOC2024/day5"
	"github.com/kushalchordiya216/AOC2024/day6"
	"github.com/kushalchordiya216/AOC2024/day7"
	"github.com/kushalchordiya216/AOC2024/day8"
	"github.com/kushalchordiya216/AOC2024/day9"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Please provide the day and part to be solved")
		return
	}
	day, err1 := strconv.Atoi(os.Args[1])
	part, err2 := strconv.Atoi(os.Args[2])
	if err1 != nil || err2 != nil || day <= 0 || day > 25 || part < 1 || part > 2 {
		fmt.Println("Please provide a valid input")
		return
	}

	var path string
	if len(os.Args) >= 4 {
		path = os.Args[3]
	} else {
		path = fmt.Sprintf("day%d/input.txt", day)
	}
	var solver common.Solver

	switch day {
	case 1:
		if part == 1 {
			solver = &day1.Part1Solver{}
		} else {
			solver = &day1.Part2Solver{}
		}
	case 2:
		if part == 1 {
			solver = &day2.Part1Solver{}
		} else {
			solver = &day2.Part2Solver{}
		}
	case 3:
		if part == 1 {
			solver = &day3.Part1Solver{}
		} else {
			solver = &day3.Part2Solver{}
		}
	case 4:
		if part == 1 {
			solver = &day4.Part1Solver{}
		} else {
			solver = &day4.Part2Solver{}
		}
	case 5:
		if part == 1 {
			solver = &day5.Part1Solver{}
		} else {
			solver = &day5.Part2Solver{}
		}
	case 6:
		if part == 1 {
			solver = &day6.Part1Solver{}
		} else {
			solver = &day6.Part2Solver{}
		}
	case 7:
		if part == 1 {
			solver = &day7.Part1Solver{}
		} else {
			solver = &day7.Part2Solver{}
		}
	case 8:
		if part == 1 {
			solver = &day8.Part1Solver{}
		} else {
			solver = &day8.Part2Solver{}
		}
	case 9:
		if part == 1 {
			solver = &day9.Part1Solver{}
		} else {
			solver = &day9.Part2Solver{}
		}
	}

	if err := solver.Read(path); err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}
	result := solver.Solve()
	fmt.Printf("Day %d, Part %d solution: %d\n", day, part, result)

}
