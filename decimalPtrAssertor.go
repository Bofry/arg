package arg

type DecimalPtrAssertor struct {
	v    *Decimal
	name string
}

func (asr *DecimalPtrAssertor) Assert(validators ...DecimalPtrValidator) error {
	return DecimalPtr.Assert(asr.v, asr.name, validators...)
}
