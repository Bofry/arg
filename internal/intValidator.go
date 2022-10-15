package internal

import (
	"encoding/json"

	"github.com/cstockton/go-conv"
)

func (fn IntValidator) AssertJsonNumber(v json.Number, name string) error {
	integer, err := v.Int64()
	if err != nil {
		return &InvalidArgumentError{
			Name:   name,
			Reason: err.Error(),
		}
	}
	return fn(integer, name)
}

func (fn IntValidator) Assert(v interface{}, name string) error {
	integer, err := conv.Int64(v)
	if err != nil {
		return &InvalidArgumentError{
			Name:   name,
			Reason: err.Error(),
		}
	}
	return fn(integer, name)
}
