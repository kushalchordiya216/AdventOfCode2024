package day21

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/kushalchordiya216/AOC2024/common/utils"
)

var transitions map[byte]map[byte][]string
var cache sync.Map

func yPosition(num byte) int {
	switch num {
	case '0', 'A':
		return 0
	case '1', '2', '3':
		return 1
	case '4', '5', '6':
		return 2
	case '7', '8', '9':
		return 3
	}
	return -1
}

func xPosition(num byte) int {
	switch num {
	case '1', '4', '7':
		return 0
	case '0', '2', '5', '8':
		return 1
	case 'A', '3', '6', '9':
		return 2
	}
	return -1
}

// returns all the possible paths from current to next
func addTransitionsForNumeric(current byte, next byte) []string {
	if current == next {
		return []string{"A"}
	}

	curr_pos := utils.Coord{X: xPosition(current), Y: yPosition(current)}
	next_pos := utils.Coord{X: xPosition(next), Y: yPosition(next)}
	x_diff := next_pos.X - curr_pos.X
	y_diff := next_pos.Y - curr_pos.Y

	forbid := utils.Coord{X: 0, Y: 0}
	paths := []string{}
	intermediate_pos := utils.Coord{X: curr_pos.X + x_diff, Y: curr_pos.Y}
	if x_diff != 0 && intermediate_pos != forbid {
		// if intermediate position is not forbidden
		// add the path which makes the horizontal steps first
		path := ""
		if x_diff > 0 {
			path += strings.Repeat(">", x_diff) // next position is to the right
		} else {
			path += strings.Repeat("<", -x_diff)
		}
		if y_diff > 0 {
			path += strings.Repeat("^", y_diff)
		} else {
			path += strings.Repeat("v", -y_diff)
		}
		path += "A"
		paths = append(paths, path)
	}

	intermediate_pos = utils.Coord{X: curr_pos.X, Y: curr_pos.Y + y_diff}
	if y_diff != 0 && intermediate_pos != forbid {
		// if intermediate position is not forbidden
		// add the path which makes the vertical steps first
		path := ""
		if y_diff > 0 {
			path += strings.Repeat("^", y_diff)
		} else {
			path += strings.Repeat("v", -y_diff)
		}
		if x_diff > 0 {
			path += strings.Repeat(">", x_diff)
		} else {
			path += strings.Repeat("<", -x_diff)
		}
		path += "A"
		paths = append(paths, path)
	}
	return paths
}

func preComputeTransitions() {
	transitions = map[byte]map[byte][]string{
		'A': {
			'^': {"<A"},
			'v': {"v<A", "<vA"},
			'<': {"v<<A"},
			'>': {"vA"},
			'A': {"A"},
		},
		'^': {
			'^': {"A"},
			'v': {"vA"},
			'<': {"v<A"},
			'>': {"v>A", "v>A"},
			'A': {">A"},
		},
		'<': {
			'^': {">^A"},
			'v': {">A"},
			'<': {"A"},
			'>': {">>A"},
			'A': {">>^A"},
		},
		'>': {
			'^': {"<^A", "^<A"},
			'v': {"<A"},
			'<': {"<<A"},
			'>': {"A"},
			'A': {"^A"},
		},
		'v': {
			'^': {"^A"},
			'v': {"A"},
			'<': {"<A"},
			'>': {">A"},
			'A': {">^A", "^>A"},
		},
		'0': {}, '1': {}, '2': {}, '3': {}, '4': {}, '5': {}, '6': {}, '7': {}, '8': {}, '9': {},
	}
	numericKeys := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A'}
	for _, key := range numericKeys {
		for _, key2 := range numericKeys {
			transitions[key][key2] = addTransitionsForNumeric(key, key2)
		}
	}
}

// Process all given sequences and returns the shortest possible length for given depth
func processSequences(sequences []string, depth int) int {
	lengths := []int{}
	for _, sequence := range sequences {
		sequence = "A" + sequence
		length := 0
		for j := 0; j < len(sequence)-1; j++ {
			length += processTransition(sequence[j], sequence[j+1], depth-1)
		}
		lengths = append(lengths, length)
	}
	sort.Ints(lengths)
	return lengths[0]
}

// processes a single transition from one character to the next
// returns the length of the shortest possible sequence for given depth
func processTransition(current, next byte, depth int) int {
	if depth == 0 {
		//print(transitions[current][next][0])
		return len(transitions[current][next][0])
	}
	cacheKey := fmt.Sprintf("%c%c%d", current, next, depth)
	if val, ok := cache.Load(cacheKey); ok {
		return val.(int)
	}
	sequences := transitions[current][next]
	result := processSequences(sequences, depth)
	cache.Store(cacheKey, result)
	return result
}
