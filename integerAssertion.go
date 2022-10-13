package arg

import (
	"fmt"
	"sort"
)

var (
	_IntegerAssertion = IntegerAssertion("")

	_ IntegerValidator = _IntegerAssertion.NonNegativeInteger
	_ IntegerValidator = _IntegerAssertion.NonZero
)

type IntegerAssertion string

func (IntegerAssertion) Assert(v int64, name string, validators ...IntegerValidator) error {
	for _, validate := range validators {
		if err := validate(v, name); err != nil {
			return err
		}
	}
	return nil
}

func (IntegerAssertion) Assertor(v int64, name string) *IntegerAssertor {
	return &IntegerAssertor{v, name}
}

func (IntegerAssertion) NonNegativeInteger(v int64, name string) error {
	if v < 0 {
		return &InvalidArgumentError{
			Name:   name,
			Reason: ERR_NON_NEGATIVE_INTENGER,
		}
	}
	return nil
}

func (IntegerAssertion) NonZero(v int64, name string) error {
	if v == 0 {
		return &InvalidArgumentError{
			Name:   name,
			Reason: ERR_NON_ZERO,
		}
	}
	return nil
}

func (IntegerAssertion) NotIn(values ...int64) IntegerValidator {
	sort.Slice(values, func(i, j int) bool { return values[i] < values[j] })
	return func(v int64, name string) error {
		i := sort.Search(len(values), func(i int) bool { return values[i] >= v })
		if i < len(values) && values[i] == v {
			return &InvalidArgumentError{
				Name:   name,
				Reason: fmt.Sprintf(ERR_INVALID_INTEGER, v),
			}
		}
		return nil
	}
}

func (IntegerAssertion) BetweenRange(min, max int64) IntegerValidator {
	return func(v int64, name string) error {
		if min > v || v > max {
			return &InvalidArgumentError{
				Name:   name,
				Reason: ERR_OUT_OF_RANGE,
			}
		}
		return nil
	}
}
