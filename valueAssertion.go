package arg

var (
	_ValueAssertion = ValueAssertion("")

	_ ValueValidator = _ValueAssertion.NotNil
)

type ValueAssertion string

func (ValueAssertion) Assert(v interface{}, name string, validators ...ValueValidator) error {
	for _, validate := range validators {
		if err := validate(v, name); err != nil {
			return err
		}
	}
	return nil
}

func (ValueAssertion) Assertor(v interface{}, name string) *ValueAssertor {
	return &ValueAssertor{v, name}
}

func (ValueAssertion) NotNil(v interface{}, name string) error {
	if v == nil {
		return &InvalidArgumentError{
			Name:   name,
			Reason: ERR_NIL,
		}
	}
	return nil
}
