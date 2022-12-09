package arg

type StringPtrAssertor struct {
	v    *string
	name string
}

func (asr *StringPtrAssertor) Assert(validators ...StringPtrValidator) error {
	return StringPtr.Assert(asr.v, asr.name, validators...)
}