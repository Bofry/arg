package arg

type StringAssertor struct {
	v    string
	name string
}

func (arg *StringAssertor) Assert(validators ...StringValidator) error {
	return Strings.Assert(arg.v, arg.name, validators...)
}
