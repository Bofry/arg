package arg

type NumberAssertor struct {
	v    Number
	name string
}

func (arg *NumberAssertor) Assert(validators ...NumberValidator) error {
	return Numbers.Assert(arg.v, arg.name, validators...)
}
