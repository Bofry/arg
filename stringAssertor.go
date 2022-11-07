package arg

type StringAssertor struct {
	v    string
	name string
}

func (asr *StringAssertor) Assert(validators ...StringValidator) error {
	return Strings.Assert(asr.v, asr.name, validators...)
}
