package arg

type ValueAssertor struct {
	v    interface{}
	name string
}

func (arg *ValueAssertor) Assert(validators ...ValueValidator) error {
	return Values.Assert(arg.v, arg.name, validators...)
}
