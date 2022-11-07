package arg

import (
	"fmt"
	"sort"

	"github.com/Bofry/arg/internal"
)

var (
	_IntAssertion = IntAssertion("")

	_ IntValidator = _IntAssertion.NonNegativeInteger
	_ IntValidator = _IntAssertion.NonZero
)

type IntAssertion string

func (IntAssertion) Assert(v int64, name string, validators ...IntValidator) error {
	for _, validate := range validators {
		if err := validate(v, name); err != nil {
			return err
		}
	}
	return nil
}

func (IntAssertion) Assertor(v int64, name string) *IntAssertor {
	return &IntAssertor{v, name}
}

func (IntAssertion) NonNegativeInteger(v int64, name string) error {
	if v < 0 {
		return &InvalidArgumentError{
			Name:   name,
			Reason: internal.ERR_NON_NEGATIVE_INTENGER,
		}
	}
	return nil
}

func (IntAssertion) NonZero(v int64, name string) error {
	if v == 0 {
		return &InvalidArgumentError{
			Name:   name,
			Reason: internal.ERR_NON_ZERO,
		}
	}
	return nil
}

func (IntAssertion) NotIn(values ...int64) IntValidator {
	sort.Slice(values, func(i, j int) bool { return values[i] < values[j] })
	return func(v int64, name string) error {
		i := sort.Search(len(values), func(i int) bool { return values[i] >= v })
		if i < len(values) && values[i] == v {
			return &InvalidArgumentError{
				Name:   name,
				Reason: fmt.Sprintf(internal.ERR_INVALID_INTEGER, v),
			}
		}
		return nil
	}
}

func (IntAssertion) In(values ...int64) IntValidator {
	sort.Slice(values, func(i, j int) bool { return values[i] < values[j] })
	return func(v int64, name string) error {
		i := sort.Search(len(values), func(i int) bool { return values[i] >= v })
		if i >= len(values) || values[i] != v {
			return &InvalidArgumentError{
				Name:   name,
				Reason: fmt.Sprintf(internal.ERR_INVALID_INTEGER, v),
			}
		}
		return nil
	}
}

func (IntAssertion) Must(fn IntPredicate) IntValidator {
	return func(v int64, name string) error {
		if !fn(v) {
			return &InvalidArgumentError{
				Name:   name,
				Reason: fmt.Sprintf(internal.ERR_INVALID_INTEGER, v),
			}
		}
		return nil
	}
}

func (IntAssertion) LessOrEqual(boundary int64) IntValidator {
	return func(v int64, name string) error {
		if boundary < v {
			return &InvalidArgumentError{
				Name:   name,
				Reason: internal.ERR_OUT_OF_RANGE,
			}
		}
		return nil
	}
}

func (IntAssertion) GreaterOrEqual(boundary int64) IntValidator {
	return func(v int64, name string) error {
		if boundary > v {
			return &InvalidArgumentError{
				Name:   name,
				Reason: internal.ERR_OUT_OF_RANGE,
			}
		}
		return nil
	}
}

// BetweenRange checks if given integer is between the specified minimum and maximum values (both inclusive).
func (IntAssertion) BetweenRange(min, max int64) IntValidator {
	return func(v int64, name string) error {
		if min > v || v > max {
			return &InvalidArgumentError{
				Name:   name,
				Reason: internal.ERR_OUT_OF_RANGE,
			}
		}
		return nil
	}
}
