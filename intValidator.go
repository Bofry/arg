package arg

import (
	"fmt"

	"github.com/Bofry/arg/internal"
	"github.com/cstockton/go-conv"
)

var (
	_ IntPtrValidator = new(IntValidator).AssertPtr
	_ NumberValidator = new(IntValidator).AssertNumber
	_ ValueValidator  = new(IntValidator).AssertValue
)

func (fn IntValidator) AssertPtr(v *int64, name string) error {
	if v != nil {
		return fn(*v, name)
	}
	return nil
}

func (fn IntValidator) AssertNumber(v Number, name string) error {
	integer, err := v.Int64()
	if err != nil {
		return &InvalidArgumentError{
			Name:   name,
			Reason: fmt.Sprintf(internal.ERR_UNSUPPORTED_CAST_INT, v),
			Err:    err,
		}
	}
	return fn(integer, name)
}

func (fn IntValidator) AssertValue(v interface{}, name string) error {
	integer, err := conv.Int64(v)
	if err != nil {
		return &InvalidArgumentError{
			Name:   name,
			Reason: fmt.Sprintf(internal.ERR_UNSUPPORTED_CAST_INT, v),
			Err:    err,
		}
	}
	return fn(integer, name)
}
