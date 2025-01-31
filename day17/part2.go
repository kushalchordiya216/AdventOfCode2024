package day17

type Part2Solver Solver

func (s *Part2Solver) Read(path string) error {
	solver, err := readInput(path)
	if err != nil {
		return err
	}
	*s = Part2Solver(solver)
	return nil
}

func (s *Part2Solver) dfs(curr int, pos int) int {
	/*
			Let's understand what the target program is doing.
			There's only two critical instructions to understand
			Fourth instruction divides the value in register A with 2**3 and writes it to register A
			Eight instruction loops back to 0 if the value in register A is 0 else ends the program

			Ultimately, we need to find values of A for which the program outputs the required values.
			We know the program ends when the value in register A is 0.
			So we start solving backwards. Starting from value of 0,
			we try out incremental values for A until we get a value
			which gives the desired output at the desired position. For example, output 0 at position 15
			Once we have this value, we multiply it by 2**3 and take this as the new base value of A
		    and start looking at the previous position, i.e. position 14.
			We repeat the process of incrementing the value of A until we get the desired output for the previous position
			until we get to the start of the program.
	*/
	if pos < 0 {
		return curr
	}

	// whatever the current value is A is already satisfied the output for pos+1
	// right shift by 3 (same as multiplying by 2**3) to get new base value
	curr <<= 3

	for {
		(*Solver)(s).reset()
		s.A = curr
		(*Solver)(s).run()
		if s.Output[0] == s.Program[pos] {
			// check if the first output generated by current value of A is equal to the desired output at the current position
			return s.dfs(curr, pos-1)
		}
		curr++
	}
}

func (s *Part2Solver) Solve() int {
	return s.dfs(0, len(s.Program)-1)
}
