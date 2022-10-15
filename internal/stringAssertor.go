package internal

type StringAssertor struct {
	v    string
	name string
}

func (arg *StringAssertor) Assert(validators ...StringValidator) error {
	return String.Assert(arg.v, arg.name, validators...)
}
