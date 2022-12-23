package arg

import "time"

type DurationPtrAssertor struct {
	v    *time.Duration
	name string
}

func (asr *DurationPtrAssertor) Assert(validators ...DurationPtrValidator) error {
	return DurationPtr.Assert(asr.v, asr.name, validators...)
}
