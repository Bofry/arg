package internal

import (
	"encoding/json"

	"github.com/cstockton/go-conv"
)

var (
	_ JsonNumberValidator = new(StringValidator).AssertJsonNumber
	_ ValueValidator      = new(StringValidator).Assert
)

func (fn StringValidator) AssertJsonNumber(v json.Number, name string) error {
	return fn(v.String(), name)
}

func (fn StringValidator) Assert(v interface{}, name string) error {
	str, err := conv.String(v)
	if err != nil {
		return &InvalidArgumentError{
			Name:   name,
			Reason: err.Error(),
		}
	}
	return fn(str, name)
}
