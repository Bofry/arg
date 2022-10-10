package arg

import "math"

var (
	_FloatAssertor = FloatAssertor("")

	_ FloatValidator = _FloatAssertor.NonNanNorInf
	_ FloatValidator = _FloatAssertor.NonNegativeNumber
	_ FloatValidator = _FloatAssertor.NonZero
)

type FloatAssertor string

func (FloatAssertor) Assert(v float64, name string, validators ...FloatValidator) error {
	for _, validate := range validators {
		if err := validate(v, name); err != nil {
			return err
		}
	}
	return nil
}

func (FloatAssertor) NonNanNorInf(v float64, name string) error {
	if isInfinity(v) || isNan(v) {
		return &InvalidArgumentError{
			Name:   name,
			Reason: ERR_NAN_OR_INFINITY,
		}
	}
	return nil
}

func (FloatAssertor) NonNegativeNumber(v float64, name string) error {
	if math.Signbit(v) {
		return &InvalidArgumentError{
			Name:   name,
			Reason: ERR_NON_NEGATIVE_NUMBER,
		}
	}
	return nil
}

func (FloatAssertor) NonZero(v float64, name string) error {
	if v == 0 {
		return &InvalidArgumentError{
			Name:   name,
			Reason: ERR_NON_ZERO,
		}
	}
	return nil
}

func (FloatAssertor) BetweenRange(min, max float64) FloatValidator {
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
