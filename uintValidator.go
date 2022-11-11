package arg

import (
	"strconv"

	"github.com/cstockton/go-conv"
)

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
		return err
	}
	return fn(integer, name)
}
