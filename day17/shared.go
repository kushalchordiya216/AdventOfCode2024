package day17

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Solver struct {
	A, B, C            int
	InstructionPointer int
	Program            []int
	Output             []int
}

func readRegisterValues(line string) (int, error) {
	splits := strings.Split(line, " ")
	num, err := strconv.Atoi(splits[len(splits)-1])
	if err != nil {
		return 0, err
	}
	return num, nil
}

func readTape(line string) ([]int, error) {
	splits := strings.Split(line, " ")
	if len(splits) != 2 {
		return nil, fmt.Errorf("invalid tape: %s", line)
	}
	splits = strings.Split(splits[1], ",")
	tape := make([]int, 0)
	for i := range splits {
		num, err := strconv.Atoi(splits[i])
		if err != nil {
			return nil, err
		}
		tape = append(tape, num)
	}
	return tape, nil
}

func readInput(path string) (Solver, error) {
	file, err := os.Open(path)
	s := Solver{}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "Register") {
			if strings.Contains(line, "A") {
				s.A, err = readRegisterValues(line)
				if err != nil {
					return s, err
				}
			} else if strings.Contains(line, "B") {
				s.B, err = readRegisterValues(line)
				if err != nil {
					return s, err
				}
			} else {
				s.C, err = readRegisterValues(line)
				if err != nil {
					return s, err
				}
			}
		}

		if strings.Contains(line, "Program") {
			s.Program, err = readTape(line)
			if err != nil {
				return s, err
			}
		}
	}
	return s, nil
}

func (s *Solver) getOperandValue(o int) int {
	switch o {
	case 0, 1, 2, 3:
		return o
	case 4:
		return s.A
	case 5:
		return s.B
	case 6:
		return s.C
	default:
		return -1
	}
}

func (s *Solver) compute(operator int, operand int) {
	switch operator {
	case 0:
		s.A = s.A >> s.getOperandValue(operand)
		break
	case 1:
		s.B = s.B ^ operand
		break
	case 2:
		s.B = s.getOperandValue(operand) & 7 // ANDing with 7 conserves only last 3 bits, which is same as modulo 8
		break
	case 3:
		if s.A != 0 {
			s.InstructionPointer = operand
		}
		break
	case 4:
		s.B = s.B ^ s.C
		break
	case 5:
		s.Output = append(s.Output, s.getOperandValue(operand)&7)
		break
	case 6:
		s.B = s.A >> s.getOperandValue(operand)
		break
	case 7:
		s.C = s.A >> s.getOperandValue(operand)
		break
	default:
		_ = fmt.Errorf("unsupported operator %d", operator)
		break
	}
}

func (s *Solver) reset() {
	s.InstructionPointer = 0
	s.Output = make([]int, 0)
}

func (s *Solver) run() {
	for s.InstructionPointer < len(s.Program) {
		opcode := s.Program[s.InstructionPointer]
		operand := s.Program[s.InstructionPointer+1]
		s.InstructionPointer += 2
		s.compute(opcode, operand)
	}
}
