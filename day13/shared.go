package day13

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"

	"github.com/kushalchordiya216/AOC2024/common"
)

type Machine struct {
	Ax, Ay int
	Bx, By int
	Tx, Ty int
}

func readMachineInput(scanner *bufio.Scanner) (Machine, error) {
	buttonRegex := regexp.MustCompile(`X\+(\d+), Y\+(\d+)`)
	prizeRegex := regexp.MustCompile(`X=(\d+), Y=(\d+)`)

	buttonA := scanner.Text()
	Ax, Ay, err := extractCoords(buttonA, buttonRegex)
	if err != nil {
		return Machine{}, err
	}
	scanner.Scan()

	buttonB := scanner.Text()
	Bx, By, err := extractCoords(buttonB, buttonRegex)
	if err != nil {
		return Machine{}, err
	}
	scanner.Scan()

	prize := scanner.Text()
	Tx, Ty, err := extractCoords(prize, prizeRegex)
	if err != nil {
		return Machine{}, err
	}
	scanner.Scan() // new line

	return Machine{Ax: Ax, Ay: Ay, Bx: Bx, By: By, Tx: Tx, Ty: Ty}, nil
}

func extractCoords(line string, regexp *regexp.Regexp) (int, int, error) {
	matches := regexp.FindStringSubmatch(line)
	if len(matches) == 3 {
		x, err := strconv.Atoi(matches[1])
		if err != nil {
			fmt.Printf("Error parsing X coordinate: %v\n", err)
			return 0, 0, err
		}
		y, err := strconv.Atoi(matches[2])
		if err != nil {
			fmt.Printf("Error parsing Y coordinate: %v\n", err)
			return 0, 0, err
		}
		return x, y, nil
	} else {
		fmt.Printf("Error in parsing, did not match expected regex %s\n", regexp.String())
		return 0, 0, &common.CustomError{Msg: "Did not have valid input"}
	}
}

/*
Ax = a1, Ay = a2
Bx = b1, By = b2
Tx = c1, Ty = c2

a1*b2 - a2*b1 = Ax*By - Bx*Ay
c1*b2 - c2*b1 = Tx*By - Ty*Bx
a1*c2 - a2*c1 = Ax*Ty - Ay*Tx
*/
func (m *Machine) solveByDeterminants() (int, int) {
	D := (m.Ax * m.By) - (m.Bx * m.Ay)
	if D == 0 {
		return 0, 0
	}
	Dx := (m.Tx * m.By) - (m.Ty * m.Bx)
	Dy := (m.Ax * m.Ty) - (m.Ay * m.Tx)
	if Dx%D == 0 && Dy%D == 0 {
		return Dx / D, Dy / D
	}
	return 0, 0
}
