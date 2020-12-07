package numerr_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/tenntenn/numerr"
	"go.uber.org/multierr"
)

func Example() {
	// import "go.uber.org/multierr"

	oddErr := func(n int) error {
		if n%2 != 0 {
			return errors.New("error")
		}
		return nil
	}

	var err error
	for i := 0; i < 10; i++ {
		err = multierr.Append(err, numerr.New(i, oddErr(i)))
	}

	for _, err := range multierr.Errors(err) {
		if n, ok := numerr.Num(err); ok {
			fmt.Print(n)
		}
	}

	// Output: 13579
}

func TestNum(t *testing.T) {
	t.Parallel()
	cases := map[string]struct {
		wrap bool
		n    int
		err  error
		ok   bool
	}{
		"ok":      {true, 0, errors.New("err"), true},
		"notwrap": {false, -1, errors.New("err"), false},
		"minus":   {true, -1, errors.New("err"), true},
		"nil":     {true, 0, nil, false},
	}

	for name, tt := range cases {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			var want int
			err := tt.err
			if tt.wrap {
				err = numerr.New(tt.n, tt.err)
				want = tt.n
			}
			got, ok := numerr.Num(err)

			if ok != tt.ok {
				t.Errorf("want %v but got %v", tt.ok, ok)
			}

			if got != want {
				t.Errorf("want %v but got %v", want, got)
			}
		})
	}
}

func TestUnwrap(t *testing.T) {
	t.Parallel()
	cases := map[string]struct {
		err error
	}{
		"unwrap": {errors.New("err")},
		"nil":    {nil},
	}

	for name, tt := range cases {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			err := numerr.New(0, tt.err)
			got := errors.Unwrap(err)
			if got != tt.err {
				t.Errorf("want %v but got %v", tt.err, got)
			}
		})
	}

}
