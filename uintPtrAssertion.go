package arg

import (
	"fmt"

	"github.com/Bofry/arg/internal"
)

var (
	_UIntPtrAssertion = UIntPtrAssertion("")

	_ UIntPtrValidator = _UIntPtrAssertion.NotNil
	_ UIntPtrValidator = _UIntPtrAssertion.NonZero
)

type UIntPtrAssertion string

func (UIntPtrAssertion) Assert(v *uint64, name string, validators ...UIntPtrValidator) error {
	for _, validate := range validators {
		if err := validate(v, name); err != nil {
			return err
		}
	}
	return nil
}

func (UIntPtrAssertion) Assertor(v *uint64, name string) *UIntPtrAssertor {
	return &UIntPtrAssertor{v, name}
}

func (UIntPtrAssertion) NotNil(v *uint64, name string) error {
	if v == nil {
		return &InvalidArgumentError{
			Name:   name,
			Reason: internal.ERR_NIL,
		}
	}
	return nil
}

func (UIntPtrAssertion) NonZero(ptr *uint64, name string) error {
	if ptr != nil {
		var (
			v uint64 = *ptr
		)
		return _UIntAssertion.NonZero(v, name)
	}
	return nil

}

func (UIntPtrAssertion) Must(fn UIntPtrPredicate) UIntPtrValidator {
	return func(v *uint64, name string) error {
		if !fn(v) {
			return &InvalidArgumentError{
				Name:   name,
				Reason: fmt.Sprintf(internal.ERR_INVALID_INTEGER, v),
			}
		}
		return nil
	}
}

func (UIntPtrAssertion) NotIn(values ...uint64) UIntPtrValidator {
	return _UIntAssertion.NotIn(values...).AssertPtr
}

func (UIntPtrAssertion) In(values ...uint64) UIntPtrValidator {
	return _UIntAssertion.In(values...).AssertPtr
}

func (UIntPtrAssertion) LessOrEqual(boundary uint64) UIntPtrValidator {
	return _UIntAssertion.LessOrEqual(boundary).AssertPtr
}

func (UIntPtrAssertion) GreaterOrEqual(boundary uint64) UIntPtrValidator {
	return _UIntAssertion.GreaterOrEqual(boundary).AssertPtr
}

// BetweenRange checks if given integer is between the specified minimum and maximum values (both inclusive).
func (UIntPtrAssertion) BetweenRange(min, max uint64) UIntPtrValidator {
	return _UIntAssertion.BetweenRange(min, max).AssertPtr
}
