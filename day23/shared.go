package day23

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	"github.com/kushalchordiya216/AOC2024/common"
	mapset "github.com/deckarep/golang-set/v2"
)

func ReadGraphFromFile(path string) (map[string]mapset.Set[string], error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("Error closing file: %v\n", err)
		}
	}(file)

	graph := make(map[string]mapset.Set[string])
	lineRegex := regexp.MustCompile(`^[a-z]{2}-[a-z]{2}$`)
	
	scanner := bufio.NewScanner(file)
	lineNum := 0
	for scanner.Scan() {
		lineNum++
		line := scanner.Text()
		
		if !lineRegex.MatchString(line) {
			return nil, &common.CustomError{
				Msg: fmt.Sprintf("line %d does not match expected format ^[a-z]{2}-[a-z]{2}$: %s", lineNum, line),
			}
		}
		
		node1 := line[:2]
		node2 := line[3:5]
		
		if graph[node1] == nil {
			graph[node1] = mapset.NewSet[string]()
		}
		if graph[node2] == nil {
			graph[node2] = mapset.NewSet[string]()
		}
		
		graph[node1].Add(node2)
		graph[node2].Add(node1)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return graph, nil
}