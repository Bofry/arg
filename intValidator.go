package arg

import (
	"encoding/json"

	"github.com/cstockton/go-conv"
)

func (fn IntValidator) AssertNumber(v json.Number, name string) error {
	integer, err := v.Int64()
	if err != nil {
		return err
	}
	return fn(integer, name)
}

func (fn IntValidator) Assert(v interface{}, name string) error {
	integer, err := conv.Int64(v)
	if err != nil {
		return err
	}
	return fn(integer, name)
}
