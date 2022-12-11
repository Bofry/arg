package arg

import "time"

type TimePtrAssertor struct {
	v    *time.Time
	name string
}

func (asr *TimePtrAssertor) Assert(validators ...TimePtrValidator) error {
	return TimePtr.Assert(asr.v, asr.name, validators...)
}
