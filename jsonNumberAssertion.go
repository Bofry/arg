package arg

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

func (JsonNumberAssertion) NotIn(values ...int64) JsonNumberValidator {
	return IntegerValidator(Int.NotIn(values...)).AssertJsonNumber
}

func (JsonNumberAssertion) BetweenRange(min, max float64) JsonNumberValidator {
	return FloatValidator(Float.BetweenRange(min, max)).AssertJsonNumber
}
