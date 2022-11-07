package arg

type ValueAssertor struct {
	v    interface{}
	name string
}

func (asr *ValueAssertor) Assert(validators ...ValueValidator) error {
	return Values.Assert(asr.v, asr.name, validators...)
}
