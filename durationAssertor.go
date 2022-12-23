package arg

import "time"

type DurationAssertor struct {
	v    time.Duration
	name string
}

func (asr *DurationAssertor) Assert(validators ...DurationValidator) error {
	return Durations.Assert(asr.v, asr.name, validators...)
}
