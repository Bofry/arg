package arg

type UIntPtrAssertor struct {
	v    *uint64
	name string
}

func (asr *UIntPtrAssertor) Assert(validators ...UIntPtrValidator) error {
	return UIntPtr.Assert(asr.v, asr.name, validators...)
}
