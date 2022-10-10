package arg

import "reflect"

var (
	_SliceAssertor = SliceAssertor("")

	_ ValueValidator = _SliceAssertor.NonEmpty
)

type SliceAssertor string

func (SliceAssertor) NonEmpty(v interface{}, name string) error {
	if reflect.TypeOf(v).Kind() != reflect.Slice {
		return &InvalidArgumentError{
			Name:   name,
			Reason: ERR_NOT_ARRAY,
		}
	}

	if reflect.ValueOf(v).Len() == 0 {
		return &InvalidArgumentError{
			Name:   name,
			Reason: ERR_EMPTY_ARRAY,
		}
	}
	return nil
}
