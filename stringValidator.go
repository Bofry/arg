package arg

import (
	"github.com/cstockton/go-conv"
)

var (
	_ NumberValidator = new(StringValidator).AssertNumber
	_ ValueValidator  = new(StringValidator).AssertValue
)

func (fn StringValidator) AssertNumber(v Number, name string) error {
	return fn(v.String(), name)
}

func (fn StringValidator) AssertValue(v interface{}, name string) error {
	str, err := conv.String(v)
	if err != nil {
		return &InvalidArgumentError{
			Name:   name,
			Reason: err.Error(),
		}
	}
	return fn(str, name)
}
