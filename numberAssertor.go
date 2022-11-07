package arg

type NumberAssertor struct {
	v    Number
	name string
}

func (asr *NumberAssertor) Assert(validators ...NumberValidator) error {
	return Numbers.Assert(asr.v, asr.name, validators...)
}
