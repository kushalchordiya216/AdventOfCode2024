package day22

import (
	"bufio"
	"os"
	"strconv"
)

type Part1Solver struct {
	numbers []int
}

func (s *Part1Solver) Read(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return err
		}
		s.numbers = append(s.numbers, num)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func iterate(number int) int {
	numberx64 := number << 6
	number ^= numberx64
	number &= 16777215

	numberby32 := number >> 5
	number ^= numberby32
	number &= 16777215

	numberx2048 := number << 11
	number ^= numberx2048
	number &= 16777215
	return number
}

func (s *Part1Solver) Solve() int {
	result := 0
	for _, num := range s.numbers {
		for range 2000 {
			num = iterate(num)
		}
		result += num
	}
	return result
}
