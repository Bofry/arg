package arg

import (
	"fmt"

	"github.com/Bofry/arg/internal"
	"github.com/cstockton/go-conv"
)

var (
	_ StringPtrValidator = new(StringValidator).AssertPtr
	_ NumberValidator    = new(StringValidator).AssertNumber
	_ ValueValidator     = new(StringValidator).AssertValue
)

func (fn StringValidator) AssertPtr(v *string, name string) error {
	if v != nil {
		return fn(*v, name)
	}
	return nil
}

func (fn StringValidator) AssertNumber(v Number, name string) error {
	return fn(v.String(), name)
}

func (fn StringValidator) AssertValue(v interface{}, name string) error {
	str, err := conv.String(v)
	if err != nil {
		return &InvalidArgumentError{
			Name:   name,
			Reason: fmt.Sprintf(internal.ERR_UNSUPPORTED_CAST_STRING, v),
			Err:    err,
		}
	}
	return fn(str, name)
}
