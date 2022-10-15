package internal

import (
	"encoding/json"
)

var (
	_JsonNumberAssertion = JsonNumberAssertion("")

	_ JsonNumberValidator = _JsonNumberAssertion.NonNegativeNumber
)

type JsonNumberAssertion string

func (JsonNumberAssertion) Assert(v json.Number, name string, validators ...JsonNumberValidator) error {
	for _, validate := range validators {
		if err := validate(v, name); err != nil {
			return err
		}
	}
	return nil
}

func (JsonNumberAssertion) Assertor(v json.Number, name string) *JsonNumberAssertor {
	return &JsonNumberAssertor{v, name}
}

func (JsonNumberAssertion) NonNegativeNumber(v json.Number, name string) error {
	return FloatValidator(Float.NonNegativeNumber).AssertJsonNumber(v, name)
}

func (JsonNumberAssertion) NonZero(v json.Number, name string) error {
	return FloatValidator(Float.NonZero).AssertJsonNumber(v, name)
}

func (JsonNumberAssertion) NotIn(values ...int64) JsonNumberValidator {
	return IntValidator(Int.NotIn(values...)).AssertJsonNumber
}

func (JsonNumberAssertion) Less(comparand float64) JsonNumberValidator {
	return FloatValidator(Float.Less(comparand)).AssertJsonNumber
}

func (JsonNumberAssertion) LessOrEqual(comparand float64) JsonNumberValidator {
	return FloatValidator(Float.LessOrEqual(comparand)).AssertJsonNumber
}

func (JsonNumberAssertion) Greater(comparand float64) JsonNumberValidator {
	return FloatValidator(Float.Greater(comparand)).AssertJsonNumber
}

func (JsonNumberAssertion) GreaterOrEqual(comparand float64) JsonNumberValidator {
	return FloatValidator(Float.GreaterOrEqual(comparand)).AssertJsonNumber
}

// check if given value is between the specified minimum and maximum values (both inclusive).
func (JsonNumberAssertion) BetweenRange(min, max float64) JsonNumberValidator {
	return FloatValidator(Float.BetweenRange(min, max)).AssertJsonNumber
}
