package arg

type IntAssertor struct {
	v    int64
	name string
}

func (asr *IntAssertor) Assert(validators ...IntValidator) error {
	return Ints.Assert(asr.v, asr.name, validators...)
}
