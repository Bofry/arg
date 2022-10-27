package arg

import (
	"fmt"

	"github.com/Bofry/arg/internal"
)

var (
	_ValueAssertion = ValueAssertion("")

	_ ValueValidator = _ValueAssertion.NotNil
)

type ValueAssertion string

func (ValueAssertion) Assert(v interface{}, name string, validators ...ValueValidator) error {
	for _, validate := range validators {
		if err := validate(v, name); err != nil {
			return err
		}
	}
	return nil
}

func (ValueAssertion) Assertor(v interface{}, name string) *ValueAssertor {
	return &ValueAssertor{v, name}
}

func (ValueAssertion) NotNil(v interface{}, name string) error {
	if v == nil {
		return &InvalidArgumentError{
			Name:   name,
			Reason: internal.ERR_NIL,
		}
	}
	return nil
}

func (ValueAssertion) Must(fn ValuePredicate) ValueValidator {
	return func(v interface{}, name string) error {
		if !fn(v) {
			return &InvalidArgumentError{
				Name:   name,
				Reason: fmt.Sprintf(internal.ERR_INVALID_VALUE, v),
			}
		}
		return nil
	}
}
