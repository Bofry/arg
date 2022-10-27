package arg

import "fmt"

type InvalidArgumentError struct {
	Name   string
	Reason string
	Err    error
}

func (e *InvalidArgumentError) Error() string {
	if len(e.Reason) > 0 {
		return fmt.Sprintf("invalid argument %q; %s", e.Name, e.Reason)
	}
	if e.Err != nil {
		return fmt.Sprintf("invalid argument %q; %s", e.Name, e.Err.Error())
	}
	return fmt.Sprintf("invalid argument %q", e.Name)
}

// Unwrap returns the underlying error.
func (e *InvalidArgumentError) Unwrap() error { return e.Err }
