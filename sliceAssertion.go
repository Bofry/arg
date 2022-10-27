package arg

import (
	"reflect"

	"github.com/Bofry/arg/internal"
)

var (
	_SliceAssertion = SliceAssertion("")

	_ ValueValidator = _SliceAssertion.NonEmpty
)

type SliceAssertion string

func (SliceAssertion) NonEmpty(v interface{}, name string) error {
	if reflect.TypeOf(v).Kind() != reflect.Slice {
		return &InvalidArgumentError{
			Name:   name,
			Reason: internal.ERR_NOT_ARRAY,
		}
	}

	if reflect.ValueOf(v).Len() == 0 {
		return &InvalidArgumentError{
			Name:   name,
			Reason: internal.ERR_EMPTY_ARRAY,
		}
	}
	return nil
}
