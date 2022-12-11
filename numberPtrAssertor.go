package arg

type NumberPtrAssertor struct {
	v    *Number
	name string
}

func (asr *NumberPtrAssertor) Assert(validators ...NumberPtrValidator) error {
	return NumberPtr.Assert(asr.v, asr.name, validators...)
}
