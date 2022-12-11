package arg

import (
	"fmt"
	"sort"
	"time"

	"github.com/Bofry/arg/internal"
)

var (
	_TimeAssertion = TimeAssertion("")

	_ TimeValidator = _TimeAssertion.IsUTC
	_ TimeValidator = _TimeAssertion.NonZero
)

type TimeAssertion string

func (TimeAssertion) Assert(v time.Time, name string, validators ...TimeValidator) error {
	for _, validate := range validators {
		if err := validate(v, name); err != nil {
			return err
		}
	}
	return nil
}

func (TimeAssertion) Assertor(v time.Time, name string) *TimeAssertor {
	return &TimeAssertor{v, name}
}

func (TimeAssertion) IsUTC(v time.Time, name string) error {
	if v.Location() != time.UTC {
		return &InvalidArgumentError{
			Name:   name,
			Reason: internal.ERR_NON_UTC,
		}
	}
	return nil
}

func (TimeAssertion) NonZero(v time.Time, name string) error {
	if v.IsZero() {
		return &InvalidArgumentError{
			Name:   name,
			Reason: internal.ERR_NON_ZERO,
		}
	}
	return nil
}

func (TimeAssertion) Must(fn TimePredicate) TimeValidator {
	return func(v time.Time, name string) error {
		if !fn(v) {
			return &InvalidArgumentError{
				Name:   name,
				Reason: fmt.Sprintf(internal.ERR_INVALID_TIME, v.String()),
			}
		}
		return nil
	}
}

func (TimeAssertion) InWeekday(values ...time.Weekday) TimeValidator {
	sort.Slice(values, func(i, j int) bool { return values[i] < values[j] })
	return func(v time.Time, name string) error {
		weekday := v.Weekday()
		i := sort.Search(len(values), func(i int) bool { return values[i] >= weekday })
		if i >= len(values) || values[i] != weekday {
			return &InvalidArgumentError{
				Name:   name,
				Reason: internal.ERR_OUT_OF_RANGE,
			}
		}
		return nil
	}
}

func (TimeAssertion) NotInWeekday(values ...time.Weekday) TimeValidator {
	sort.Slice(values, func(i, j int) bool { return values[i] < values[j] })
	return func(v time.Time, name string) error {
		weekday := v.Weekday()
		i := sort.Search(len(values), func(i int) bool { return values[i] >= weekday })
		if i < len(values) && values[i] == weekday {
			return &InvalidArgumentError{
				Name:   name,
				Reason: internal.ERR_OUT_OF_RANGE,
			}
		}
		return nil
	}
}

func (TimeAssertion) AfterOrEqual(boundary time.Time) TimeValidator {
	return func(v time.Time, name string) error {
		if !v.After(boundary) {
			return &InvalidArgumentError{
				Name:   name,
				Reason: internal.ERR_OUT_OF_RANGE,
			}
		}
		return nil
	}
}

func (TimeAssertion) BeforeOrEqual(boundary time.Time) TimeValidator {
	return func(v time.Time, name string) error {
		if !v.Before(boundary) {
			return &InvalidArgumentError{
				Name:   name,
				Reason: internal.ERR_OUT_OF_RANGE,
			}
		}
		return nil
	}
}

// BetweenRange checks if given integer is between the specified minimum and maximum values (both inclusive).
func (TimeAssertion) BetweenRange(min, max time.Time) TimeValidator {
	return func(v time.Time, name string) error {
		if !v.After(min) || !v.Before(max) {
			return &InvalidArgumentError{
				Name:   name,
				Reason: internal.ERR_OUT_OF_RANGE,
			}
		}
		return nil
	}
}
