package arg

import (
	"fmt"

	"github.com/Bofry/arg/internal"
)

var (
	_NumberPtrAssertion = NumberPtrAssertion("")

	_ NumberPtrValidator = _NumberPtrAssertion.NotNil
	_ NumberPtrValidator = _NumberPtrAssertion.IsNumber
	_ NumberPtrValidator = _NumberPtrAssertion.NonNegativeNumber
	_ NumberPtrValidator = _NumberPtrAssertion.NonZero
	_ NumberPtrValidator = _NumberPtrAssertion.NonNanNorInf
)

type NumberPtrAssertion string

func (NumberPtrAssertion) Assert(v *Number, name string, validators ...NumberPtrValidator) error {
	for _, validate := range validators {
		if err := validate(v, name); err != nil {
			return err
		}
	}
	return nil
}

func (NumberPtrAssertion) Assertor(v *Number, name string) *NumberPtrAssertor {
	return &NumberPtrAssertor{v, name}
}

func (NumberPtrAssertion) NotNil(v *Number, name string) error {
	if v == nil {
		return &InvalidArgumentError{
			Name:   name,
			Reason: internal.ERR_NIL,
		}
	}
	return nil
}

func (NumberPtrAssertion) IsNumber(ptr *Number, name string) error {
	if ptr != nil {
		var (
			v Number = *ptr
		)
		return _NumberAssertion.IsNumber(v, name)
	}
	return nil
}

func (NumberPtrAssertion) NonNegativeNumber(ptr *Number, name string) error {
	if ptr != nil {
		var (
			v Number = *ptr
		)
		return _NumberAssertion.NonNegativeNumber(v, name)
	}
	return nil
}

func (NumberPtrAssertion) NonZero(ptr *Number, name string) error {
	if ptr != nil {
		var (
			v Number = *ptr
		)
		return _NumberAssertion.NonZero(v, name)
	}
	return nil
}

func (NumberPtrAssertion) NonNanNorInf(ptr *Number, name string) error {
	if ptr != nil {
		var (
			v Number = *ptr
		)
		return _NumberAssertion.NonNanNorInf(v, name)
	}
	return nil
}

func (NumberPtrAssertion) MustNil(v *Number, name string) error {
	if v != nil {
		return &InvalidArgumentError{
			Name:   name,
			Reason: internal.ERR_NON_NIL,
		}
	}
	return nil
}

func (NumberPtrAssertion) Must(fn NumberPtrPredicate) NumberPtrValidator {
	return func(v *Number, name string) error {
		if !fn(v) {
			return &InvalidArgumentError{
				Name:   name,
				Reason: fmt.Sprintf(internal.ERR_INVALID_NUMBER, v.String()),
			}
		}
		return nil
	}
}

func (NumberPtrAssertion) MustInt(fn IntPredicate) NumberPtrValidator {
	return _NumberAssertion.MustInt(fn).AssertPtr
}

func (NumberPtrAssertion) MustFloat(fn FloatPredicate) NumberPtrValidator {
	return _NumberAssertion.MustFloat(fn).AssertPtr
}

func (NumberPtrAssertion) NotIn(values ...int64) NumberPtrValidator {
	return _NumberAssertion.NotIn(values...).AssertPtr
}

func (NumberPtrAssertion) Less(comparand float64) NumberPtrValidator {
	return _NumberAssertion.Less(comparand).AssertPtr
}

func (NumberPtrAssertion) LessOrEqual(comparand float64) NumberPtrValidator {
	return _NumberAssertion.LessOrEqual(comparand).AssertPtr
}

func (NumberPtrAssertion) Greater(comparand float64) NumberPtrValidator {
	return _NumberAssertion.Greater(comparand).AssertPtr
}

func (NumberPtrAssertion) GreaterOrEqual(comparand float64) NumberPtrValidator {
	return _NumberAssertion.GreaterOrEqual(comparand).AssertPtr
}

// BetweenRange checks if given number is between the specified minimum and maximum values (both inclusive).
func (NumberPtrAssertion) BetweenRange(min, max float64) NumberPtrValidator {
	return _NumberAssertion.BetweenRange(min, max).AssertPtr
}
