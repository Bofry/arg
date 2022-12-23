package arg

import (
	"fmt"
	"time"

	"github.com/Bofry/arg/internal"
)

var (
	_DurationAssertion = DurationAssertion("")

	_ DurationValidator = _DurationAssertion.NonNegative
	_ DurationValidator = _DurationAssertion.NonZero
)

type DurationAssertion string

func (DurationAssertion) Assert(v time.Duration, name string, validators ...DurationValidator) error {
	for _, validate := range validators {
		if err := validate(v, name); err != nil {
			return err
		}
	}
	return nil
}

func (DurationAssertion) Assertor(v time.Duration, name string) *DurationAssertor {
	return &DurationAssertor{v, name}
}

func (DurationAssertion) NonNegative(v time.Duration, name string) error {
	if v < 0 {
		return &InvalidArgumentError{
			Name:   name,
			Reason: internal.ERR_NON_NEGATIVE_DURATION,
		}
	}
	return nil
}

func (DurationAssertion) NonZero(v time.Duration, name string) error {
	if v == 0 {
		return &InvalidArgumentError{
			Name:   name,
			Reason: internal.ERR_NON_ZERO,
		}
	}
	return nil
}

func (DurationAssertion) Must(fn DurationPredicate) DurationValidator {
	return func(v time.Duration, name string) error {
		if !fn(v) {
			return &InvalidArgumentError{
				Name:   name,
				Reason: fmt.Sprintf(internal.ERR_INVALID_DURATION, v),
			}
		}
		return nil
	}
}

func (DurationAssertion) LessOrEqual(boundary time.Duration) DurationValidator {
	return func(v time.Duration, name string) error {
		if boundary < v {
			return &InvalidArgumentError{
				Name:   name,
				Reason: internal.ERR_OUT_OF_RANGE,
			}
		}
		return nil
	}
}

func (DurationAssertion) LessOrEqualByString(durationBoundary string) DurationValidator {
	v, err := time.ParseDuration(durationBoundary)
	if err != nil {
		panic(err)
	}
	return _DurationAssertion.LessOrEqual(v)
}

func (DurationAssertion) GreaterOrEqual(boundary time.Duration) DurationValidator {
	return func(v time.Duration, name string) error {
		if boundary > v {
			return &InvalidArgumentError{
				Name:   name,
				Reason: internal.ERR_OUT_OF_RANGE,
			}
		}
		return nil
	}
}

func (DurationAssertion) GreaterOrEqualByString(durationBoundary string) DurationValidator {
	v, err := time.ParseDuration(durationBoundary)
	if err != nil {
		panic(err)
	}
	return _DurationAssertion.GreaterOrEqual(v)
}

// BetweenRange checks if given time.Duration is between the specified minimum and maximum values (both inclusive).
func (DurationAssertion) BetweenRange(min, max time.Duration) DurationValidator {
	return func(v time.Duration, name string) error {
		if min > v || v > max {
			return &InvalidArgumentError{
				Name:   name,
				Reason: internal.ERR_OUT_OF_RANGE,
			}
		}
		return nil
	}
}

// BetweenRange checks if given time.Duration is between the specified minimum and maximum values (both inclusive).
func (DurationAssertion) BetweenRangeByString(min, max string) DurationValidator {
	lbound, err := time.ParseDuration(min)
	if err != nil {
		panic(err)
	}
	ubound, err := time.ParseDuration(max)
	if err != nil {
		panic(err)
	}
	return _DurationAssertion.BetweenRange(lbound, ubound)
}
