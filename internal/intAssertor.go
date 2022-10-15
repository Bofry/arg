package internal

type IntAssertor struct {
	v    int64
	name string
}

func (arg *IntAssertor) Assert(validators ...IntValidator) error {
	return Int.Assert(arg.v, arg.name, validators...)
}
