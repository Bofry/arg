package arg

import (
	"fmt"
	"reflect"
	"time"

	"github.com/Bofry/arg/internal"
	"github.com/cstockton/go-conv"
)

func (fn DurationValidator) AssertPtr(v *time.Duration, name string) error {
	if v != nil {
		return fn(*v, name)
	}
	return nil
}

func (fn DurationValidator) AssertValue(v interface{}, name string) error {
	var d time.Duration

	switch v.(type) {
	case time.Duration:
		{
			d = v.(time.Duration)
		}
	case string:
		{
			var err error
			d, err = conv.Duration(v)
			if err != nil {
				return &InvalidArgumentError{
					Name:   name,
					Reason: fmt.Sprintf(internal.ERR_UNSUPPORTED_CAST_DURATION, v),
				}
			}
		}
	default:
		rv := reflect.ValueOf(v)
		kind := rv.Kind()

		if internal.IsKindNumeric(kind) {
			var err error
			d, err = conv.Duration(v)
			if err != nil {
				return &InvalidArgumentError{
					Name:   name,
					Reason: fmt.Sprintf(internal.ERR_UNSUPPORTED_CAST_DURATION, v),
				}
			}
		} else {
			return &InvalidArgumentError{
				Name:   name,
				Reason: fmt.Sprintf(internal.ERR_UNSUPPORTED_CAST_DURATION, v),
			}
		}
	}

	return fn(d, name)
}
