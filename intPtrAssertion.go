package arg

import (
	"fmt"

	"github.com/Bofry/arg/internal"
)

var (
	_IntPtrAssertion = IntPtrAssertion("")

	_ IntPtrValidator = _IntPtrAssertion.NotNil
	_ IntPtrValidator = _IntPtrAssertion.NonNegativeInteger
	_ IntPtrValidator = _IntPtrAssertion.NonZero
)

type IntPtrAssertion string

func (IntPtrAssertion) Assert(v *int64, name string, validators ...IntPtrValidator) error {
	for _, validate := range validators {
		if err := validate(v, name); err != nil {
			return err
		}
	}
	return nil
}

func (IntPtrAssertion) Assertor(v *int64, name string) *IntPtrAssertor {
	return &IntPtrAssertor{v, name}
}

func (IntPtrAssertion) NotNil(v *int64, name string) error {
	if v == nil {
		return &InvalidArgumentError{
			Name:   name,
			Reason: internal.ERR_NIL,
		}
	}
	return nil
}

func (IntPtrAssertion) NonNegativeInteger(ptr *int64, name string) error {
	if ptr != nil {
		var (
			v int64 = *ptr
		)
		return _IntAssertion.NonNegativeInteger(v, name)
	}
	return nil
}

func (IntPtrAssertion) NonZero(ptr *int64, name string) error {
	if ptr != nil {
		var (
			v int64 = *ptr
		)
		return _IntAssertion.NonZero(v, name)
	}
	return nil
}

func (IntPtrAssertion) Must(fn IntPtrPredicate) IntPtrValidator {
	return func(v *int64, name string) error {
		if !fn(v) {
			return &InvalidArgumentError{
				Name:   name,
				Reason: fmt.Sprintf(internal.ERR_INVALID_INTEGER, v),
			}
		}
		return nil
	}
}

func (IntPtrAssertion) NotIn(values ...int64) IntPtrValidator {
	return _IntAssertion.NotIn(values...).AssertPtr
}

func (IntPtrAssertion) In(values ...int64) IntPtrValidator {
	return _IntAssertion.In(values...).AssertPtr
}

func (IntPtrAssertion) LessOrEqual(boundary int64) IntPtrValidator {
	return _IntAssertion.LessOrEqual(boundary).AssertPtr
}

func (IntPtrAssertion) GreaterOrEqual(boundary int64) IntPtrValidator {
	return _IntAssertion.GreaterOrEqual(boundary).AssertPtr
}

// BetweenRange checks if given integer is between the specified minimum and maximum values (both inclusive).
func (IntPtrAssertion) BetweenRange(min, max int64) IntPtrValidator {
	return _IntAssertion.BetweenRange(min, max).AssertPtr
}
