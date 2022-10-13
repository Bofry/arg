package arg

type FloatAssertor struct {
	v    float64
	name string
}

func (arg *FloatAssertor) Assert(validators ...FloatValidator) error {
	return Float.Assert(arg.v, arg.name, validators...)
}
