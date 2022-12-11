package arg

import (
	"fmt"
	"strconv"

	"github.com/Bofry/arg/internal"
)

var (
	_NumberAssertion = NumberAssertion("")

	_ NumberValidator = _NumberAssertion.IsNumber
	_ NumberValidator = _NumberAssertion.NonNegativeNumber
	_ NumberValidator = _NumberAssertion.NonZero
	_ NumberValidator = _NumberAssertion.NonNanNorInf
)

type NumberAssertion string

func (NumberAssertion) Assert(v Number, name string, validators ...NumberValidator) error {
	for _, validate := range validators {
		if err := validate(v, name); err != nil {
			return err
		}
	}
	return nil
}

func (NumberAssertion) Assertor(v Number, name string) *NumberAssertor {
	return &NumberAssertor{v, name}
}

func (NumberAssertion) IsNumber(v Number, name string) error {
	var err error
	_, err = v.Float64()
	if err == nil {
		return nil
	}
	_, err = v.Int64()
	if err == nil {
		return nil
	}
	if err != nil {
		return &InvalidArgumentError{
			Name:   name,
			Reason: fmt.Sprintf(internal.ERR_INVALID_NUMBER, v.String()),
			Err:    err,
		}
	}
	return nil
}

func (NumberAssertion) NonNegativeNumber(v Number, name string) error {
	validate := _NumberAssertion.warpFloatValidator(
		Floats.NonNegativeNumber, true,
	)
	return validate(v, name)
}

func (NumberAssertion) NonZero(v Number, name string) error {
	validate := _NumberAssertion.warpFloatValidator(
		Floats.NonZero, true,
	)
	return validate(v, name)
}

func (NumberAssertion) NonNanNorInf(v Number, name string) error {
	validate := _NumberAssertion.warpFloatValidator(
		Floats.NonNanNorInf, true,
	)
	return validate(v, name)
}

func (NumberAssertion) Must(fn NumberPredicate) NumberValidator {
	return func(v Number, name string) error {
		if !fn(v) {
			return &InvalidArgumentError{
				Name:   name,
				Reason: fmt.Sprintf(internal.ERR_INVALID_NUMBER, v.String()),
			}
		}
		return nil
	}
}

func (NumberAssertion) MustInt(fn IntPredicate) NumberValidator {
	return _NumberAssertion.warpIntValidator(
		Ints.Must(fn), false,
	)
}

func (NumberAssertion) MustFloat(fn FloatPredicate) NumberValidator {
	return _NumberAssertion.warpFloatValidator(
		Floats.Must(fn), false,
	)
}

func (NumberAssertion) NotIn(values ...int64) NumberValidator {
	return _NumberAssertion.warpIntValidator(
		Ints.NotIn(values...), false,
	)
}

func (NumberAssertion) Less(comparand float64) NumberValidator {
	return _NumberAssertion.warpFloatValidator(
		Floats.Less(comparand), true,
	)
}

func (NumberAssertion) LessOrEqual(comparand float64) NumberValidator {
	return _NumberAssertion.warpFloatValidator(
		Floats.LessOrEqual(comparand), true,
	)
}

func (NumberAssertion) Greater(comparand float64) NumberValidator {
	return _NumberAssertion.warpFloatValidator(
		Floats.Greater(comparand), true,
	)
}

func (NumberAssertion) GreaterOrEqual(comparand float64) NumberValidator {
	return _NumberAssertion.warpFloatValidator(
		Floats.GreaterOrEqual(comparand), true,
	)
}

// BetweenRange checks if given number is between the specified minimum and maximum values (both inclusive).
func (NumberAssertion) BetweenRange(min, max float64) NumberValidator {
	return _NumberAssertion.warpFloatValidator(
		Floats.BetweenRange(min, max), true,
	)
}

func (NumberAssertion) warpFloatValidator(validator FloatValidator, throwUnderlyingError bool) NumberValidator {
	return func(v Number, name string) error {
		err := validator.AssertNumber(v, name)
		if err != nil {
			switch _err := err.(type) {
			case *strconv.NumError:
				if _, _err := v.Int64(); _err == nil {
					return nil
				}
			case *InvalidArgumentError:
				if throwUnderlyingError {
					return _err
				}
			}
			return &InvalidArgumentError{
				Name:   name,
				Reason: fmt.Sprintf(internal.ERR_INVALID_NUMBER, v.String()),
				Err:    err,
			}
		}
		return nil
	}
}

func (NumberAssertion) warpIntValidator(validator IntValidator, throwUnderlyingError bool) NumberValidator {
	return func(v Number, name string) error {
		err := validator.AssertNumber(v, name)
		if err != nil {
			switch _err := err.(type) {
			case *strconv.NumError:
				if _, _err := v.Float64(); _err == nil {
					return nil
				}
			case *InvalidArgumentError:
				if throwUnderlyingError {
					return _err
				}
			}
			return &InvalidArgumentError{
				Name:   name,
				Reason: fmt.Sprintf(internal.ERR_INVALID_NUMBER, v.String()),
				Err:    err,
			}
		}
		return nil
	}
}
