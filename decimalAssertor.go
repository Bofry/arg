package arg

type DecimalAssertor struct {
	v    Decimal
	name string
}

func (asr *DecimalAssertor) Assert(validators ...DecimalValidator) error {
	return Decimals.Assert(asr.v, asr.name, validators...)
}
