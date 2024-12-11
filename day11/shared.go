package day11

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type IterationDict map[int]int

func readInput(path string) ([]int, error) {
	stones := make([]int, 0)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, " ")
		for _, numStr := range nums {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				return nil, err
			}
			stones = append(stones, num)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return stones, nil
}

func splitEvenDigitNumber(num int) (int, int, bool) {
	digitCount := 0
	originalNum := num
	for num > 0 {
		num /= 10
		digitCount++
	}

	if digitCount%2 != 0 {
		return 0, 0, false
	}

	num = originalNum
	divisor := 1
	for i := 0; i < digitCount/2; i++ {
		divisor *= 10
	}

	leftHalf := num / divisor
	rightHalf := num % divisor

	return leftHalf, rightHalf, true
}

func memoizedRecursiveIteration(num int, iterations int, cache map[int]IterationDict) int {
	if val, ok := cache[num][iterations]; ok {
		return val
	}
	if _, ok := cache[num]; !ok {
		cache[num] = make(map[int]int, 0)
	}
	if iterations == 0 {
		return 1
	}
	if num == 0 {
		val := memoizedRecursiveIteration(1, iterations-1, cache)
		cache[num][iterations] = val
		return val
	}
	left, right, ok := splitEvenDigitNumber(num)
	if ok {
		leftVal := memoizedRecursiveIteration(left, iterations-1, cache)
		rightVal := memoizedRecursiveIteration(right, iterations-1, cache)
		cache[num][iterations] = leftVal + rightVal
		return leftVal + rightVal
	}
	val := memoizedRecursiveIteration(2024*num, iterations-1, cache)
	cache[num][iterations] = val
	return val
}
