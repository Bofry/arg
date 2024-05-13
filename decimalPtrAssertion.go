package arg

import (
	"fmt"

	"github.com/Bofry/arg/internal"
)

var (
	_DecimalPtrAssertion = DecimalPtrAssertion("")

	_ DecimalPtrValidator = _DecimalPtrAssertion.NotNil
	_ DecimalPtrValidator = _DecimalPtrAssertion.NonNegativeNumber
	_ DecimalPtrValidator = _DecimalPtrAssertion.NonZero
)

type DecimalPtrAssertion string

func (DecimalPtrAssertion) Assert(v *Decimal, name string, validators ...DecimalPtrValidator) error {
	for _, validate := range validators {
		if err := validate(v, name); err != nil {
			return err
		}
	}
	return nil
}

func (DecimalPtrAssertion) Assertor(v *Decimal, name string) *DecimalPtrAssertor {
	return &DecimalPtrAssertor{v, name}
}

func (DecimalPtrAssertion) NotNil(v *Decimal, name string) error {
	if v == nil {
		return &InvalidArgumentError{
			Name:   name,
			Reason: internal.ERR_NIL,
		}
	}
	return nil
}

func (DecimalPtrAssertion) NonNegativeNumber(ptr *Decimal, name string) error {
	if ptr != nil {
		var (
			v Decimal = *ptr
		)
		return _DecimalAssertion.NonNegativeNumber(v, name)
	}
	return nil
}

func (DecimalPtrAssertion) NonZero(ptr *Decimal, name string) error {
	if ptr != nil {
		var (
			v Decimal = *ptr
		)
		return _DecimalAssertion.NonZero(v, name)
	}
	return nil
}

func (DecimalPtrAssertion) MustNil(v *Decimal, name string) error {
	if v != nil {
		return &InvalidArgumentError{
			Name:   name,
			Reason: internal.ERR_NON_NIL,
		}
	}
	return nil
}

func (DecimalPtrAssertion) Must(fn DecimalPtrPredicate) DecimalPtrValidator {
	return func(v *Decimal, name string) error {
		if !fn(v) {
			return &InvalidArgumentError{
				Name:   name,
				Reason: fmt.Sprintf(internal.ERR_INVALID_FLOAT, v),
			}
		}
		return nil
	}
}

func (DecimalPtrAssertion) Less(comparand Decimal) DecimalPtrValidator {
	return _DecimalAssertion.Less(comparand).AssertPtr
}

func (DecimalPtrAssertion) LessOrEqual(comparand Decimal) DecimalPtrValidator {
	return _DecimalAssertion.LessOrEqual(comparand).AssertPtr
}

func (DecimalPtrAssertion) Greater(comparand Decimal) DecimalPtrValidator {
	return _DecimalAssertion.Greater(comparand).AssertPtr
}

func (DecimalPtrAssertion) GreaterOrEqual(comparand Decimal) DecimalPtrValidator {
	return _DecimalAssertion.GreaterOrEqual(comparand).AssertPtr
}

// BetweenRange checks if given number is between the specified minimum and maximum values (both inclusive).
func (DecimalPtrAssertion) BetweenRange(min, max Decimal) DecimalPtrValidator {
	return _DecimalAssertion.BetweenRange(min, max).AssertPtr
}
