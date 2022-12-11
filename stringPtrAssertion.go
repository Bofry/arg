package arg

import (
	"fmt"

	"github.com/Bofry/arg/internal"
)

var (
	_StringPtrAssertion = StringPtrAssertion("")

	_ StringPtrValidator = _StringPtrAssertion.NotNil
	_ StringPtrValidator = _StringPtrAssertion.NonEmpty
)

type StringPtrAssertion string

func (StringPtrAssertion) Assert(v *string, name string, validators ...StringPtrValidator) error {
	for _, validate := range validators {
		if err := validate(v, name); err != nil {
			return err
		}
	}
	return nil
}

func (StringPtrAssertion) Assertor(v *string, name string) *StringPtrAssertor {
	return &StringPtrAssertor{
		v:    v,
		name: name,
	}
}

func (StringPtrAssertion) NotNil(v *string, name string) error {
	if v == nil {
		return &InvalidArgumentError{
			Name:   name,
			Reason: internal.ERR_NIL,
		}
	}
	return nil
}

func (StringPtrAssertion) NonEmpty(ptr *string, name string) error {
	if ptr != nil {
		var (
			v string = *ptr
		)
		return _StringAssertion.NonEmpty(v, name)
	}
	return nil
}

// Must checks if the given string is evaluated to true by specified predicate.
func (StringPtrAssertion) Must(fn StringPtrPredicate) StringPtrValidator {
	return func(v *string, name string) error {
		if !fn(v) {
			return &InvalidArgumentError{
				Name:   name,
				Reason: fmt.Sprintf(internal.ERR_INVALID_STRING, *v),
			}
		}
		return nil
	}
}

func (StringPtrAssertion) In(values ...string) StringPtrValidator {
	return _StringAssertion.In(values...).AssertPtr
}

func (StringPtrAssertion) MaxLength(size int) StringPtrValidator {
	return _StringAssertion.MaxLength(size).AssertPtr
}

func (StringPtrAssertion) MinLength(size int) StringPtrValidator {
	return _StringAssertion.MinLength(size).AssertPtr
}

// MatchAny checks if given string match any one from specified patterns.
func (StringPtrAssertion) MatchAny(patterns ...string) StringPtrValidator {
	return _StringAssertion.MatchAny(patterns...).AssertPtr
}
