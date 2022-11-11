package arg

import (
	"fmt"
	"sort"

	"github.com/Bofry/arg/internal"
)

var (
	_UIntAssertion = UIntAssertion("")

	_ UIntValidator = _UIntAssertion.NonZero
)

type UIntAssertion string

func (UIntAssertion) Assert(v uint64, name string, validators ...UIntValidator) error {
	for _, validate := range validators {
		if err := validate(v, name); err != nil {
			return err
		}
	}
	return nil
}

func (UIntAssertion) Assertor(v uint64, name string) *UIntAssertor {
	return &UIntAssertor{v, name}
}

func (UIntAssertion) NonZero(v uint64, name string) error {
	if v == 0 {
		return &InvalidArgumentError{
			Name:   name,
			Reason: internal.ERR_NON_ZERO,
		}
	}
	return nil
}

func (UIntAssertion) NotIn(values ...uint64) UIntValidator {
	sort.Slice(values, func(i, j int) bool { return values[i] < values[j] })
	return func(v uint64, name string) error {
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

func (UIntAssertion) In(values ...uint64) UIntValidator {
	sort.Slice(values, func(i, j int) bool { return values[i] < values[j] })
	return func(v uint64, name string) error {
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

func (UIntAssertion) Must(fn UIntPredicate) UIntValidator {
	return func(v uint64, name string) error {
		if !fn(v) {
			return &InvalidArgumentError{
				Name:   name,
				Reason: fmt.Sprintf(internal.ERR_INVALID_INTEGER, v),
			}
		}
		return nil
	}
}

func (UIntAssertion) LessOrEqual(boundary uint64) UIntValidator {
	return func(v uint64, name string) error {
		if boundary < v {
			return &InvalidArgumentError{
				Name:   name,
				Reason: internal.ERR_OUT_OF_RANGE,
			}
		}
		return nil
	}
}

func (UIntAssertion) GreaterOrEqual(boundary uint64) UIntValidator {
	return func(v uint64, name string) error {
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
func (UIntAssertion) BetweenRange(min, max uint64) UIntValidator {
	return func(v uint64, name string) error {
		if min > v || v > max {
			return &InvalidArgumentError{
				Name:   name,
				Reason: internal.ERR_OUT_OF_RANGE,
			}
		}
		return nil
	}
}
