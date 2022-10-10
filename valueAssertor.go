package arg

var (
	_ValueAssertor = ValueAssertor("")

	_ ValueValidator = _ValueAssertor.NotNil
)

type ValueAssertor string

func (ValueAssertor) Assert(v interface{}, name string, validators ...ValueValidator) error {
	for _, validate := range validators {
		if err := validate(v, name); err != nil {
			return err
		}
	}
	return nil
}

func (ValueAssertor) NotNil(v interface{}, name string) error {
	if v == nil {
		return &InvalidArgumentError{
			Name:   name,
			Reason: ERR_NIL,
		}
	}
	return nil
}
