package arg

type FloatPtrAssertor struct {
	v    *float64
	name string
}

func (asr *FloatPtrAssertor) Assert(validators ...FloatPtrValidator) error {
	return FloatPtr.Assert(asr.v, asr.name, validators...)
}
