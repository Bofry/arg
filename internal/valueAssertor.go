package internal

type ValueAssertor struct {
	v    interface{}
	name string
}

func (arg *ValueAssertor) Assert(validators ...ValueValidator) error {
	return Value.Assert(arg.v, arg.name, validators...)
}
