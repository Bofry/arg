package arg

func Assert(errs ...error) error {
	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}

func ThrowError(name, reason string) error {
	return &InvalidArgumentError{
		Name:   name,
		Reason: reason,
	}
}
