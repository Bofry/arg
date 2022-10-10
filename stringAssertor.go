package arg

import (
	"fmt"
	"sort"
)

var (
	_StringAssertor = StringAssertor("")

	_ StringValidator = _StringAssertor.NonEmpty
)

type StringAssertor string

func (StringAssertor) Assert(v string, name string, validators ...StringValidator) error {
	for _, validate := range validators {
		if err := validate(v, name); err != nil {
			return err
		}
	}
	return nil
}

func (StringAssertor) NonEmpty(v string, name string) error {
	if len(v) == 0 {
		return &InvalidArgumentError{
			Name:   name,
			Reason: ERR_EMPTY_STRING,
		}
	}
	return nil
}

func (StringAssertor) In(values ...string) StringValidator {
	sort.Slice(values, func(i, j int) bool { return values[i] < values[j] })
	return func(v string, name string) error {
		i := sort.Search(len(values), func(i int) bool { return values[i] >= v })
		if i < len(values) && values[i] != v {
			return &InvalidArgumentError{
				Name:   name,
				Reason: fmt.Sprintf(ERR_INVALID_STRING, v),
			}
		}
		return nil
	}
}
