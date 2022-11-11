package arg

import (
	"github.com/cstockton/go-conv"
)

func (fn FloatValidator) AssertNumber(v Number, name string) error {
	float, err := v.Float64()
	if err != nil {
		return err
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
		return err
	}
	// NOTE: normalize the float64. avoid the "-0" be treated as Signbit() carried value
	if float == 0 {
		float = 0
	}
	return fn(float, name)
}
