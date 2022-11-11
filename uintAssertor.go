package arg

type UIntAssertor struct {
	v    uint64
	name string
}

func (asr *UIntAssertor) Assert(validators ...UIntValidator) error {
	return UInts.Assert(asr.v, asr.name, validators...)
}
