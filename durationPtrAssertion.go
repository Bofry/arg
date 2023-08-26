package arg

import (
	"fmt"
	"time"

	"github.com/Bofry/arg/internal"
)

var (
	_DurationPtrAssertion = DurationPtrAssertion("")

	_ DurationPtrValidator = _DurationPtrAssertion.NotNil
	_ DurationPtrValidator = _DurationPtrAssertion.NonNegative
	_ DurationPtrValidator = _DurationPtrAssertion.NonZero
)

type DurationPtrAssertion string

func (DurationPtrAssertion) Assert(v *time.Duration, name string, validators ...DurationPtrValidator) error {
	for _, validate := range validators {
		if err := validate(v, name); err != nil {
			return err
		}
	}
	return nil
}

func (DurationPtrAssertion) Assertor(v *time.Duration, name string) *DurationPtrAssertor {
	return &DurationPtrAssertor{v, name}
}

func (DurationPtrAssertion) NotNil(v *time.Duration, name string) error {
	if v == nil {
		return &InvalidArgumentError{
			Name:   name,
			Reason: internal.ERR_NIL,
		}
	}
	return nil
}

func (DurationPtrAssertion) NonNegative(ptr *time.Duration, name string) error {
	if ptr != nil {
		var (
			v time.Duration = *ptr
		)
		return _DurationAssertion.NonNegative(v, name)
	}
	return nil
}

func (DurationPtrAssertion) NonZero(ptr *time.Duration, name string) error {
	if ptr != nil {
		var (
			v time.Duration = *ptr
		)
		return _DurationAssertion.NonZero(v, name)
	}
	return nil
}

func (DurationPtrAssertion) MustNil(v *time.Duration, name string) error {
	if v != nil {
		return &InvalidArgumentError{
			Name:   name,
			Reason: internal.ERR_NON_NIL,
		}
	}
	return nil
}

func (DurationPtrAssertion) Must(fn DurationPtrPredicate) DurationPtrValidator {
	return func(v *time.Duration, name string) error {
		if !fn(v) {
			return &InvalidArgumentError{
				Name:   name,
				Reason: fmt.Sprintf(internal.ERR_INVALID_INTEGER, v),
			}
		}
		return nil
	}
}

func (DurationPtrAssertion) LessOrEqual(boundary time.Duration) DurationPtrValidator {
	return _DurationAssertion.LessOrEqual(boundary).AssertPtr
}

func (DurationPtrAssertion) GreaterOrEqual(boundary time.Duration) DurationPtrValidator {
	return _DurationAssertion.GreaterOrEqual(boundary).AssertPtr
}

// BetweenRange checks if given time.Duration is between the specified minimum and maximum values (both inclusive).
func (DurationPtrAssertion) BetweenRange(min, max time.Duration) DurationPtrValidator {
	return _DurationAssertion.BetweenRange(min, max).AssertPtr
}
