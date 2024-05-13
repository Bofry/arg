package arg

import (
	"fmt"

	"github.com/Bofry/arg/internal"
)

var (
	_DecimalAssertion = DecimalAssertion("")

	_ DecimalValidator = _DecimalAssertion.NonNegativeNumber
	_ DecimalValidator = _DecimalAssertion.NonZero
)

type DecimalAssertion string

func (DecimalAssertion) Assert(v Decimal, name string, validators ...DecimalValidator) error {
	for _, validate := range validators {
		if err := validate(v, name); err != nil {
			return err
		}
	}
	return nil
}

func (DecimalAssertion) Assertor(v Decimal, name string) *DecimalAssertor {
	return &DecimalAssertor{v, name}
}

func (DecimalAssertion) NonNegativeNumber(v Decimal, name string) error {
	if v.Sign() < 0 {
		return &InvalidArgumentError{
			Name:   name,
			Reason: internal.ERR_NON_NEGATIVE_NUMBER,
		}
	}
	return nil
}

func (DecimalAssertion) NonZero(v Decimal, name string) error {
	if v.Sign() == 0 {
		return &InvalidArgumentError{
			Name:   name,
			Reason: internal.ERR_NON_ZERO,
		}
	}
	return nil
}

func (DecimalAssertion) Must(fn DecimalPredicate) DecimalValidator {
	return func(v Decimal, name string) error {
		if !fn(v) {
			return &InvalidArgumentError{
				Name:   name,
				Reason: fmt.Sprintf(internal.ERR_INVALID_NUMBER, v),
			}
		}
		return nil
	}
}

func (DecimalAssertion) Less(comparand Decimal) DecimalValidator {
	return func(v Decimal, name string) error {
		if comparand.LessThanOrEqual(v) {
			return &InvalidArgumentError{
				Name:   name,
				Reason: internal.ERR_OUT_OF_RANGE,
			}
		}
		return nil
	}
}

func (DecimalAssertion) LessOrEqual(comparand Decimal) DecimalValidator {
	return func(v Decimal, name string) error {
		if comparand.LessThan(v) {
			return &InvalidArgumentError{
				Name:   name,
				Reason: internal.ERR_OUT_OF_RANGE,
			}
		}
		return nil
	}
}

func (DecimalAssertion) Greater(comparand Decimal) DecimalValidator {
	return func(v Decimal, name string) error {
		if comparand.GreaterThanOrEqual(v) {
			return &InvalidArgumentError{
				Name:   name,
				Reason: internal.ERR_OUT_OF_RANGE,
			}
		}
		return nil
	}
}

func (DecimalAssertion) GreaterOrEqual(comparand Decimal) DecimalValidator {
	return func(v Decimal, name string) error {
		if comparand.GreaterThan(v) {
			return &InvalidArgumentError{
				Name:   name,
				Reason: internal.ERR_OUT_OF_RANGE,
			}
		}
		return nil
	}
}

// BetweenRange checks if given number is between the specified minimum and maximum values (both inclusive).
func (DecimalAssertion) BetweenRange(min, max Decimal) DecimalValidator {
	return func(v Decimal, name string) error {
		if min.GreaterThan(v) || v.GreaterThan(max) {
			return &InvalidArgumentError{
				Name:   name,
				Reason: internal.ERR_OUT_OF_RANGE,
			}
		}
		return nil
	}
}
