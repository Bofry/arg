package internal

import "math"

var (
	_FloatAssertion = FloatAssertion("")

	_ FloatValidator = _FloatAssertion.NonNanNorInf
	_ FloatValidator = _FloatAssertion.NonNegativeNumber
	_ FloatValidator = _FloatAssertion.NonZero
)

type FloatAssertion string

func (FloatAssertion) Assert(v float64, name string, validators ...FloatValidator) error {
	for _, validate := range validators {
		if err := validate(v, name); err != nil {
			return err
		}
	}
	return nil
}

func (FloatAssertion) Assertor(v float64, name string) *FloatAssertor {
	return &FloatAssertor{v, name}
}

func (FloatAssertion) NonNanNorInf(v float64, name string) error {
	if isInfinity(v) || isNan(v) {
		return &InvalidArgumentError{
			Name:   name,
			Reason: ERR_NAN_OR_INFINITY,
		}
	}
	return nil
}

func (FloatAssertion) NonNegativeNumber(v float64, name string) error {
	if math.Signbit(v) {
		return &InvalidArgumentError{
			Name:   name,
			Reason: ERR_NON_NEGATIVE_NUMBER,
		}
	}
	return nil
}

func (FloatAssertion) NonZero(v float64, name string) error {
	if v == 0 {
		return &InvalidArgumentError{
			Name:   name,
			Reason: ERR_NON_ZERO,
		}
	}
	return nil
}

func (FloatAssertion) Less(comparand float64) FloatValidator {
	if isInfinity(comparand) || isNan(comparand) {
		panic("specified arguemnt 'comparand' cannot be NaN or Infinity")
	}

	return func(v float64, name string) error {
		if comparand <= v {
			return &InvalidArgumentError{
				Name:   name,
				Reason: ERR_OUT_OF_RANGE,
			}
		}
		return nil
	}
}

func (FloatAssertion) LessOrEqual(comparand float64) FloatValidator {
	if isInfinity(comparand) || isNan(comparand) {
		panic("specified arguemnt 'comparand' cannot be NaN or Infinity")
	}

	return func(v float64, name string) error {
		if comparand < v {
			return &InvalidArgumentError{
				Name:   name,
				Reason: ERR_OUT_OF_RANGE,
			}
		}
		return nil
	}
}

func (FloatAssertion) Greater(comparand float64) FloatValidator {
	if isInfinity(comparand) || isNan(comparand) {
		panic("specified arguemnt 'comparand' cannot be NaN or Infinity")
	}

	return func(v float64, name string) error {
		if comparand >= v {
			return &InvalidArgumentError{
				Name:   name,
				Reason: ERR_OUT_OF_RANGE,
			}
		}
		return nil
	}
}

func (FloatAssertion) GreaterOrEqual(comparand float64) FloatValidator {
	if isInfinity(comparand) || isNan(comparand) {
		panic("specified arguemnt 'comparand' cannot be NaN or Infinity")
	}

	return func(v float64, name string) error {
		if comparand > v {
			return &InvalidArgumentError{
				Name:   name,
				Reason: ERR_OUT_OF_RANGE,
			}
		}
		return nil
	}
}

// check if given value is between the specified minimum and maximum values (both inclusive).
func (FloatAssertion) BetweenRange(min, max float64) FloatValidator {
	if isInfinity(min) || isNan(min) {
		panic("specified arguemnt 'min' cannot be NaN or Infinity")
	}
	if isInfinity(max) || isNan(max) {
		panic("specified arguemnt 'max' cannot be NaN or Infinity")
	}

	return func(v float64, name string) error {
		if min > v || v > max {
			return &InvalidArgumentError{
				Name:   name,
				Reason: ERR_OUT_OF_RANGE,
			}
		}
		return nil
	}
}
