package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/kushalchordiya216/AOC2024/common"
	"github.com/kushalchordiya216/AOC2024/day1"
	"github.com/kushalchordiya216/AOC2024/day10"
	"github.com/kushalchordiya216/AOC2024/day11"
	"github.com/kushalchordiya216/AOC2024/day12"
	"github.com/kushalchordiya216/AOC2024/day13"
	"github.com/kushalchordiya216/AOC2024/day2"
	"github.com/kushalchordiya216/AOC2024/day3"
	"github.com/kushalchordiya216/AOC2024/day4"
	"github.com/kushalchordiya216/AOC2024/day5"
	"github.com/kushalchordiya216/AOC2024/day6"
	"github.com/kushalchordiya216/AOC2024/day7"
	"github.com/kushalchordiya216/AOC2024/day8"
	"github.com/kushalchordiya216/AOC2024/day9"
)

func selectSolver(day, part int) common.Solver {
	solverMap := map[int]map[int]common.Solver{
		1:  {1: &day1.Part1Solver{}, 2: &day1.Part2Solver{}},
		2:  {1: &day2.Part1Solver{}, 2: &day2.Part2Solver{}},
		3:  {1: &day3.Part1Solver{}, 2: &day3.Part2Solver{}},
		4:  {1: &day4.Part1Solver{}, 2: &day4.Part2Solver{}},
		5:  {1: &day5.Part1Solver{}, 2: &day5.Part2Solver{}},
		6:  {1: &day6.Part1Solver{}, 2: &day6.Part2Solver{}},
		7:  {1: &day7.Part1Solver{}, 2: &day7.Part2Solver{}},
		8:  {1: &day8.Part1Solver{}, 2: &day8.Part2Solver{}},
		9:  {1: &day9.Part1Solver{}, 2: &day9.Part2Solver{}},
		10: {1: &day10.Part1Solver{}, 2: &day10.Part2Solver{}},
		11: {1: &day11.Part1Solver{}, 2: &day11.Part2Solver{}},
		12: {1: &day12.Part1Solver{}, 2: &day12.Part2Solver{}},
		13: {1: &day13.Part1Solver{}, 2: &day13.Part2Solver{}},
	}

	if dayMap, exists := solverMap[day]; exists {
		if solver, exists := dayMap[part]; exists {
			return solver
		}
	}
	return nil
}

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
	solver := selectSolver(day, part)
	if solver == nil {
		fmt.Printf("Day %d, Part %d solution: not yet implemented", day, part)
		os.Exit(1)
	}
	if err := solver.Read(path); err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}
	result := solver.Solve()
	fmt.Printf("Day %d, Part %d solution: %d\n", day, part, result)

}
