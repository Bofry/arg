package arg

type IntPtrAssertor struct {
	v    *int64
	name string
}

func (asr *IntPtrAssertor) Assert(validators ...IntPtrValidator) error {
	return IntPtr.Assert(asr.v, asr.name, validators...)
}
