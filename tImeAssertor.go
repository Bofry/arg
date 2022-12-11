package arg

import "time"

type TimeAssertor struct {
	v    time.Time
	name string
}

func (asr *TimeAssertor) Assert(validators ...TimeValidator) error {
	return Times.Assert(asr.v, asr.name, validators...)
}
