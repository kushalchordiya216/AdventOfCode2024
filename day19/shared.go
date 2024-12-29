package day19

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type Solver struct {
	towels  map[string]bool
	designs []string
}

func readTowels(scanner *bufio.Scanner) map[string]bool {
	scanner.Scan()
	line := scanner.Text()
	towels := make(map[string]bool)
	splits := strings.Split(line, ", ")
	for _, split := range splits {
		towels[split] = true
	}
	return towels
}

func readDesigns(scanner *bufio.Scanner) []string {
	designs := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		designs = append(designs, line)
	}
	return designs
}

func readInput(path string) (Solver, error) {
	file, err := os.Open(path)
	s := Solver{}

	if err != nil {
		return s, err
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)
	scanner := bufio.NewScanner(file)
	s.towels = readTowels(scanner)
	scanner.Scan()
	s.designs = readDesigns(scanner)
	return s, nil
}
