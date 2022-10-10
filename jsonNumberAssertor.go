package arg

import (
	"encoding/json"
)

type JsonNumberAssertor string

func (JsonNumberAssertor) Assert(v json.Number, name string, validators ...JsonNumberValidator) error {
	for _, validate := range validators {
		if err := validate(v, name); err != nil {
			return err
		}
	}
	return nil
}

func (JsonNumberAssertor) NonNegativeNumber(v json.Number, name string) error {
	return FloatValidator(Float.NonNegativeNumber).AssertJsonNumber(v, name)
}

func (JsonNumberAssertor) NotIn(values ...int64) JsonNumberValidator {
	return IntegerValidator(Int.NotIn(values...)).AssertJsonNumber
}

func (JsonNumberAssertor) BetweenRange(min, max float64) JsonNumberValidator {
	return FloatValidator(Float.BetweenRange(min, max)).AssertJsonNumber
}
