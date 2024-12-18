package day3

import (
	"bufio"
	"fmt"
	"os"
)

type Solver struct {
	text string
}

func readInput(path string) (string, error) {
	file, err := os.Open(path)
	var text string
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return "", err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("Error closing file: %v\n", err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		text += line
	}
	return text, nil
}
