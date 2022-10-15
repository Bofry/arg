package internal

import "encoding/json"

type JsonNumberAssertor struct {
	v    json.Number
	name string
}

func (arg *JsonNumberAssertor) Assert(validators ...JsonNumberValidator) error {
	return JsonNumber.Assert(arg.v, arg.name, validators...)
}
