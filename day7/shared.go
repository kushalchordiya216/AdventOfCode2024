package day7

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

func readInput(path string) ([]int, [][]int, error) {
	file, err := os.Open(path)
	targets := make([]int, 0)
	numLists := make([][]int, 0)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
		parts := strings.Split(line, ":")

		if len(strings.TrimSpace(line)) == 0 {
			fmt.Printf("Line %d is empty or contains only whitespace\n", lineNumber)
			continue
		}

		if len(parts) != 2 {
			fmt.Printf("Line %d has invalid format: '%s'\n", lineNumber, line)
			return nil, nil, fmt.Errorf("invalid line format: %s", line)
		}

		target, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			return nil, nil, fmt.Errorf("invalid result number: %v", err)
		}
		targets = append(targets, target)

		numStrs := strings.Fields(strings.TrimSpace(parts[1]))
		numbers := make([]int, 0, len(numStrs))
		for _, numStr := range numStrs {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				return nil, nil, fmt.Errorf("invalid number in sequence: %v", err)
			}
			numbers = append(numbers, num)
		}

		numLists = append(numLists, numbers)
	}

	// Check if there was an error during scanning
	if err := scanner.Err(); err != nil {
		fmt.Printf("Scanner error: %v\n", err)
		return nil, nil, err
	}

	return targets, numLists, nil
}

type CheckerFunc func(int, int, []int) bool

type Job struct {
	target int
	nums   []int
	result int
}

func worker(jobs <-chan Job, results chan<- Job, wg *sync.WaitGroup, checker CheckerFunc) {
	defer wg.Done()
	for job := range jobs {
		if checker(job.target, job.nums[0], job.nums[1:]) {
			job.result = job.target
		}
		results <- job
	}
}

func ParallelSolver(targets []int, numLists [][]int, checker CheckerFunc) int {
	var finalResult int
	numWorkers := runtime.NumCPU()

	jobs := make(chan Job, len(targets))
	results := make(chan Job, len(targets))
	var wg sync.WaitGroup

	// Start worker pool
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(jobs, results, &wg, checker)
	}

	for i := range targets {
		jobs <- Job{target: targets[i], nums: numLists[i], result: 0}
	}
	close(jobs)

	// Collect results in another goroutine
	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		finalResult += result.result
	}

	return finalResult
}
