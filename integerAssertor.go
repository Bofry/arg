package arg

import (
	"fmt"
	"sort"
)

var (
	_IntegerAssertor = IntegerAssertor("")

	_ IntegerValidator = _IntegerAssertor.NonNegativeInteger
	_ IntegerValidator = _IntegerAssertor.NonZero
)

type IntegerAssertor string

func (IntegerAssertor) Assert(v int64, name string, validators ...IntegerValidator) error {
	for _, validate := range validators {
		if err := validate(v, name); err != nil {
			return err
		}
	}
	return nil
}

func (IntegerAssertor) NonNegativeInteger(v int64, name string) error {
	if v < 0 {
		return &InvalidArgumentError{
			Name:   name,
			Reason: ERR_NON_NEGATIVE_INTENGER,
		}
	}
	return nil
}

func (IntegerAssertor) NonZero(v int64, name string) error {
	if v == 0 {
		return &InvalidArgumentError{
			Name:   name,
			Reason: ERR_NON_ZERO,
		}
	}
	return nil
}

func (IntegerAssertor) NotIn(values ...int64) IntegerValidator {
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

func (IntegerAssertor) BetweenRange(min, max int64) IntegerValidator {
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
