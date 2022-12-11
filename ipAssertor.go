package arg

type IPAssertor struct {
	v    IP
	name string
}

func (asr *IPAssertor) Assert(validators ...IPValidator) error {
	return IPs.Assert(asr.v, asr.name, validators...)
}
