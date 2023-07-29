package arg

import (
	"fmt"

	"github.com/Bofry/arg/internal"
	"github.com/cstockton/go-conv"
)

var (
	_ FloatPtrValidator = new(FloatValidator).AssertPtr
	_ NumberValidator   = new(FloatValidator).AssertNumber
	_ ValueValidator    = new(FloatValidator).AssertValue
)

func (fn FloatValidator) AssertPtr(v *float64, name string) error {
	if v != nil {
		return fn(*v, name)
	}
	return nil
}

func (fn FloatValidator) AssertNumber(v Number, name string) error {
	float, err := v.Float64()
	if err != nil {
		return &InvalidArgumentError{
			Name:   name,
			Reason: fmt.Sprintf(internal.ERR_UNSUPPORTED_CAST_FLOAT, v),
			Err:    err,
		}
	}
	/* NOTE: normalize the float64. avoid the "-0" be treated as Signbit() carried value
	 * e.g:
	 *   var s json.Number = "-0"
	 *   f, _ := s.Float64()
	 *   math.Signbit(f)   // return true
	 * but:
	 *   var f float64 = -0
	 *   math.Signbit(f)   // return false
	 */
	if float == 0 {
		float = 0
	}
	return fn(float, name)
}

func (fn FloatValidator) AssertValue(v interface{}, name string) error {
	float, err := conv.Float64(v)
	if err != nil {
		return &InvalidArgumentError{
			Name:   name,
			Reason: fmt.Sprintf(internal.ERR_UNSUPPORTED_CAST_FLOAT, v),
			Err:    err,
		}
	}
	// NOTE: normalize the float64. avoid the "-0" be treated as Signbit() carried value
	if float == 0 {
		float = 0
	}
	return fn(float, name)
}
