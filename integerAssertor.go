package arg

type IntegerAssertor struct {
	v    int64
	name string
}

func (arg *IntegerAssertor) Assert(validators ...IntegerValidator) error {
	return Int.Assert(arg.v, arg.name, validators...)
}
