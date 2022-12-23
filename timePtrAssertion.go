package arg

import (
	"fmt"
	"time"

	"github.com/Bofry/arg/internal"
)

var (
	_TimePtrAssertion = TimePtrAssertion("")

	_ TimePtrValidator = _TimePtrAssertion.NotNil
	_ TimePtrValidator = _TimePtrAssertion.IsUTC
	_ TimePtrValidator = _TimePtrAssertion.NonZero
)

type TimePtrAssertion string

func (TimePtrAssertion) Assert(v *time.Time, name string, validators ...TimePtrValidator) error {
	for _, validate := range validators {
		if err := validate(v, name); err != nil {
			return err
		}
	}
	return nil
}

func (TimePtrAssertion) Assertor(v *time.Time, name string) *TimePtrAssertor {
	return &TimePtrAssertor{v, name}
}

func (TimePtrAssertion) NotNil(v *time.Time, name string) error {
	if v == nil {
		return &InvalidArgumentError{
			Name:   name,
			Reason: internal.ERR_NIL,
		}
	}
	return nil
}

func (TimePtrAssertion) IsUTC(ptr *time.Time, name string) error {
	if ptr != nil {
		var (
			v time.Time = *ptr
		)
		return _TimeAssertion.IsUTC(v, name)
	}
	return nil
}

func (TimePtrAssertion) NonZero(ptr *time.Time, name string) error {
	if ptr != nil {
		var (
			v time.Time = *ptr
		)
		return _TimeAssertion.NonZero(v, name)
	}
	return nil
}

func (TimePtrAssertion) Must(fn TimePtrPredicate) TimePtrValidator {
	return func(v *time.Time, name string) error {
		if !fn(v) {
			return &InvalidArgumentError{
				Name:   name,
				Reason: fmt.Sprintf(internal.ERR_INVALID_NUMBER, v.String()),
			}
		}
		return nil
	}
}

func (TimePtrAssertion) NotInWeekday(values ...time.Weekday) TimePtrValidator {
	return _TimeAssertion.NotInWeekday(values...).AssertPtr
}

func (TimePtrAssertion) InWeekday(values ...time.Weekday) TimePtrValidator {
	return _TimeAssertion.InWeekday(values...).AssertPtr
}

func (TimePtrAssertion) AfterOrEqual(boundary time.Time) TimePtrValidator {
	return _TimeAssertion.AfterOrEqual(boundary).AssertPtr
}

func (TimePtrAssertion) BeforeOrEqual(boundary time.Time) TimePtrValidator {
	return _TimeAssertion.BeforeOrEqual(boundary).AssertPtr
}

// BetweenRange checks if given time.TIme is between the specified minimum and maximum values (both inclusive).
func (TimePtrAssertion) BetweenRange(min, max time.Time) TimePtrValidator {
	return _TimeAssertion.BetweenRange(min, max).AssertPtr
}
