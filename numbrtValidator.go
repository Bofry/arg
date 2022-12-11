package arg

import (
	"fmt"

	"github.com/Bofry/arg/internal"
)

var (
	_ NumberPtrValidator = new(NumberValidator).AssertPtr
	_ ValueValidator     = new(NumberValidator).AssertValue
)

var (
	dummyNumber Number
)

func (fn NumberValidator) AssertPtr(v *Number, name string) error {
	if v != nil {
		return fn(*v, name)
	}
	return nil
}

func (fn NumberValidator) AssertValue(v interface{}, name string) error {
	switch real := v.(type) {
	case Number:
		return fn(real, name)
	case *Number:
		if real != nil {
			return fn(*real, name)
		}
		return nil
	}
	return &InvalidArgumentError{
		Name:   name,
		Reason: fmt.Sprintf(internal.ERR_INVALID_CAST_TYPE, v, dummyNumber),
	}
}
