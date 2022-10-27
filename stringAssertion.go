package arg

import (
	"fmt"
	"regexp"
	"sort"

	"github.com/Bofry/arg/internal"
)

var (
	_StringAssertion = StringAssertion("")

	_ StringValidator = _StringAssertion.NonEmpty
)

type StringAssertion string

func (StringAssertion) Assert(v string, name string, validators ...StringValidator) error {
	for _, validate := range validators {
		if err := validate(v, name); err != nil {
			return err
		}
	}
	return nil
}

func (StringAssertion) Assertor(v string, name string) *StringAssertor {
	return &StringAssertor{
		v:    v,
		name: name,
	}
}

func (StringAssertion) NonEmpty(v string, name string) error {
	if len(v) == 0 {
		return &InvalidArgumentError{
			Name:   name,
			Reason: internal.ERR_EMPTY_STRING,
		}
	}
	return nil
}

func (StringAssertion) In(values ...string) StringValidator {
	sort.Slice(values, func(i, j int) bool { return values[i] < values[j] })
	return func(v, name string) error {
		i := sort.Search(len(values), func(i int) bool { return values[i] >= v })
		if i < len(values) && values[i] != v {
			return &InvalidArgumentError{
				Name:   name,
				Reason: fmt.Sprintf(internal.ERR_INVALID_STRING, v),
			}
		}
		return nil
	}
}

func (StringAssertion) Must(fn StringPredicate) StringValidator {
	return func(v, name string) error {
		if !fn(v) {
			return &InvalidArgumentError{
				Name:   name,
				Reason: fmt.Sprintf(internal.ERR_INVALID_STRING, v),
			}
		}
		return nil
	}
}

func (StringAssertion) MaxLength(size int) StringValidator {
	return func(v, name string) error {
		if len(v) > size {
			return &InvalidArgumentError{
				Name:   name,
				Reason: internal.ERR_STRING_TOO_LONG,
			}
		}
		return nil
	}
}

func (StringAssertion) MinLength(size int) StringValidator {
	return func(v, name string) error {
		if len(v) < size {
			return &InvalidArgumentError{
				Name:   name,
				Reason: internal.ERR_STRING_TOO_SHORT,
			}
		}
		return nil
	}
}

// MatchAny check if given string match any one from specified patterns.
func (StringAssertion) MatchAny(patterns ...string) StringValidator {
	var regex []*regexp.Regexp
	for _, pattern := range patterns {
		r, err := regexp.Compile(pattern)
		if err != nil {
			panic(fmt.Sprintf(internal.ERR_INVALID_REGEX_PATTERN, pattern))
		}
		regex = append(regex, r)
	}
	return func(v, name string) error {
		for _, pattern := range regex {
			if pattern.MatchString(v) {
				return nil
			}
		}
		return &InvalidArgumentError{
			Name:   name,
			Reason: fmt.Sprintf(internal.ERR_INVALID_STRING, v),
		}
	}
}
