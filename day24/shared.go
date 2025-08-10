package day24

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Dependency struct {
	Input1    string
	Input2    string
	Operation string
}

func (d *Dependency) getAliases() []string {
	aliases := make([]string, 0)
	aliases = append(aliases, fmt.Sprintf("%s-%s-%s", d.Input1, d.Operation, d.Input2))
	aliases = append(aliases, fmt.Sprintf("%s-%s-%s", d.Input2, d.Operation, d.Input1))
	return aliases
}

type Node struct {
	Identifier string
	Value      *int
	Dependency *Dependency
}

func parseInput(path string) (map[string]*Node, map[string]string, error) {
	nodes := make(map[string]*Node)
	operationAlias := make(map[string]string)

	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Parse initial nodes (x00: 1 or y00: 0 pattern)
	initialNodePattern := regexp.MustCompile(`^[xy]\d{2}:\s[01]$`)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			break
		}

		if initialNodePattern.MatchString(line) {
			parts := strings.Split(line, ": ")
			identifier := parts[0]
			value, _ := strconv.Atoi(parts[1])

			nodes[identifier] = &Node{
				Identifier: identifier,
				Value:      &value,
				Dependency: nil,
			}
		}
	}

	// Parse gate definitions
	gatePattern := regexp.MustCompile(`^(\w+)\s+(AND|OR|XOR)\s+(\w+)\s+->\s+(\w+)$`)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}

		matches := gatePattern.FindStringSubmatch(line)
		if len(matches) == 5 {
			input1 := matches[1]
			operation := matches[2]
			input2 := matches[3]
			output := matches[4]

			// Ensure input nodes exist
			if nodes[input1] == nil {
				nodes[input1] = &Node{
					Identifier: input1,
					Value:      nil,
					Dependency: nil,
				}
			}
			if nodes[input2] == nil {
				nodes[input2] = &Node{
					Identifier: input2,
					Value:      nil,
					Dependency: nil,
				}
			}

			if nodes[output] != nil {
				nodes[output].Dependency = &Dependency{
					Input1:    input1,
					Input2:    input2,
					Operation: operation,
				}
			} else {
				nodes[output] = &Node{
					Identifier: output,
					Value:      nil,
					Dependency: &Dependency{
						Input1:    input1,
						Input2:    input2,
						Operation: operation,
					},
				}
			}

			// Build reverse dependency graph
			key1 := fmt.Sprintf("%s-%s-%s", input1, operation, input2)
			key2 := fmt.Sprintf("%s-%s-%s", input2, operation, input1)
			operationAlias[key1] = output
			operationAlias[key2] = output
		}
	}

	return nodes, operationAlias, nil
}
