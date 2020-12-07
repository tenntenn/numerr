package numerr

import "fmt"

// New creates a numbered error.
// The error can get number by Num and unwrap by errors.Unwrap.
func New(n int, err error) error {
	if err == nil {
		return nil
	}
	return &numberedErr{
		n:   n,
		err: err,
	}
}

// Num gets number of a numbered error.
// If err is not numbered error, Num return 0 and false.
func Num(err error) (n int, ok bool) {
	nerr, ok := err.(*numberedErr)
	if ok {
		return nerr.n, true
	}
	return 0, false
}

type numberedErr struct {
	n   int
	err error
}

func (nerr *numberedErr) Error() string {
	return fmt.Sprintf("[%d]: %v", nerr.n, nerr.err)
}

func (nerr *numberedErr) Unwrap() error {
	return nerr.err
}
