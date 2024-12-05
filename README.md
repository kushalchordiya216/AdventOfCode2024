# Advent of Code 2024 

This repository contains solutions for [Advent of Code 2024](https://adventofcode.com/2024) in golang. The code is written with the goal of being able to compile a single executable binary which has solutions for all 25 days of the challenge.

## How to Run
If you have golang installed on your machine, simply clone the repository and run `go build`
Alternatively, if you're on an M-series Mac you can directly use the compiled binary in the repo (AOC2024)

The executable requires the following commandline arguments 
- day: Integer value between 1 to 25 indicating which day's challenge you want to solve 
- part: Either 1 or 2, indicating which part of that day's challenge you want to solve

Example, 
`./AOC2024 3 2`
Will give the solution for the 2nd part of the day 3 challenge

Optionally, you can also include the file path for the input fle which has the data required to test the solution. If a path is not provided, it will default to the actual `input.txt` file for that day, already included in the repo

## Navigating the code
Some common utilities and interfaces are declared in the [common.go](common/commons.go) file. This includes the `Solver` interface which is what is implemented to solve a challenge. 

It's a basic interface with only 2 methods, `Read`, which reads the expected input format for 
the challenge and `Solve` which solves the challenge and returns the answer

Since AOC challenges always have an integer answer, it's possible to implement a common interface for all challenges and decide which one to use based on runtime arguments

Then, each day's solution goes into it's own folder, split into `part1.go` and `part2.go` files, which implement the respective Solvers for each part

