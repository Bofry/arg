package arg

import "github.com/Bofry/arg/internal"

func Assert(errs ...error) error {
	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}

func ThrowError(name, reason string) error {
	return &internal.InvalidArgumentError{
		Name:   name,
		Reason: reason,
	}
}
