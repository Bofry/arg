package arg

import "net"

type IPAssertor struct {
	v    net.IP
	name string
}

func (asr *IPAssertor) Assert(validators ...IPValidator) error {
	return IPs.Assert(asr.v, asr.name, validators...)
}
