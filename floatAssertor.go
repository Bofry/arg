package arg

type FloatAssertor struct {
	v    float64
	name string
}

func (asr *FloatAssertor) Assert(validators ...FloatValidator) error {
	return Floats.Assert(asr.v, asr.name, validators...)
}
