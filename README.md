# numerr

[![PkgGoDev](https://pkg.go.dev/badge/github.com/tenntenn/numerr)](https://pkg.go.dev/github.com/tenntenn/numerr)

`numerr` provides a numbered error.

```go
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
```
