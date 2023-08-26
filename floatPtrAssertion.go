package arg

import (
	"fmt"

	"github.com/Bofry/arg/internal"
)

var (
	_FloatPtrAssertion = FloatPtrAssertion("")

	_ FloatPtrValidator = _FloatPtrAssertion.NotNil
	_ FloatPtrValidator = _FloatPtrAssertion.NonNanNorInf
	_ FloatPtrValidator = _FloatPtrAssertion.NonNegativeNumber
	_ FloatPtrValidator = _FloatPtrAssertion.NonZero
)

type FloatPtrAssertion string

func (FloatPtrAssertion) Assert(v *float64, name string, validators ...FloatPtrValidator) error {
	for _, validate := range validators {
		if err := validate(v, name); err != nil {
			return err
		}
	}
	return nil
}

func (FloatPtrAssertion) Assertor(v *float64, name string) *FloatPtrAssertor {
	return &FloatPtrAssertor{v, name}
}

func (FloatPtrAssertion) NotNil(v *float64, name string) error {
	if v == nil {
		return &InvalidArgumentError{
			Name:   name,
			Reason: internal.ERR_NIL,
		}
	}
	return nil
}

func (FloatPtrAssertion) NonNanNorInf(ptr *float64, name string) error {
	if ptr != nil {
		var (
			v float64 = *ptr
		)
		return _FloatAssertion.NonNanNorInf(v, name)
	}
	return nil
}

func (FloatPtrAssertion) NonNegativeNumber(ptr *float64, name string) error {
	if ptr != nil {
		var (
			v float64 = *ptr
		)
		return _FloatAssertion.NonNegativeNumber(v, name)
	}
	return nil
}

func (FloatPtrAssertion) NonZero(ptr *float64, name string) error {
	if ptr != nil {
		var (
			v float64 = *ptr
		)
		return _FloatAssertion.NonZero(v, name)
	}
	return nil
}

func (FloatPtrAssertion) MustNil(v *float64, name string) error {
	if v != nil {
		return &InvalidArgumentError{
			Name:   name,
			Reason: internal.ERR_NON_NIL,
		}
	}
	return nil
}

func (FloatPtrAssertion) Must(fn FloatPtrPredicate) FloatPtrValidator {
	return func(v *float64, name string) error {
		if !fn(v) {
			return &InvalidArgumentError{
				Name:   name,
				Reason: fmt.Sprintf(internal.ERR_INVALID_FLOAT, v),
			}
		}
		return nil
	}
}

func (FloatPtrAssertion) Less(comparand float64) FloatPtrValidator {
	return _FloatAssertion.Less(comparand).AssertPtr
}

func (FloatPtrAssertion) LessOrEqual(comparand float64) FloatPtrValidator {
	return _FloatAssertion.LessOrEqual(comparand).AssertPtr
}

func (FloatPtrAssertion) Greater(comparand float64) FloatPtrValidator {
	return _FloatAssertion.Greater(comparand).AssertPtr
}

func (FloatPtrAssertion) GreaterOrEqual(comparand float64) FloatPtrValidator {
	return _FloatAssertion.GreaterOrEqual(comparand).AssertPtr
}

// BetweenRange checks if given number is between the specified minimum and maximum values (both inclusive).
func (FloatPtrAssertion) BetweenRange(min, max float64) FloatPtrValidator {
	return _FloatAssertion.BetweenRange(min, max).AssertPtr
}
