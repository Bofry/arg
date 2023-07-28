package arg

import (
	"fmt"
	"strconv"

	"github.com/Bofry/arg/internal"
	"github.com/cstockton/go-conv"
)

var (
	_ UIntPtrValidator = new(UIntValidator).AssertPtr
	_ NumberValidator  = new(UIntValidator).AssertNumber
	_ ValueValidator   = new(UIntValidator).AssertValue
)

func (fn UIntValidator) AssertPtr(v *uint64, name string) error {
	if v != nil {
		return fn(*v, name)
	}
	return nil
}

func (fn UIntValidator) AssertNumber(v Number, name string) error {
	integer, err := strconv.ParseUint(v.String(), 10, 64)
	if err != nil {
		return err
	}
	return fn(integer, name)
}

func (fn UIntValidator) AssertValue(v interface{}, name string) error {
	integer, err := conv.Uint64(v)
	if err != nil {
		return &InvalidArgumentError{
			Name:   name,
			Reason: fmt.Sprintf(internal.ERR_UNSUPPORTED_CAST_UINT, v),
			Err:    err,
		}
	}
	return fn(integer, name)
}
