package common

type Solver interface {
	Read(path string) error
	Solve() int
}

type CustomError struct {
	Msg string
}

func (e *CustomError) Error() string {
	return e.Msg
}
