package internal

import "encoding/json"

const (
	String     = StringAssertion("")
	Float      = FloatAssertion("")
	Int        = IntAssertion("")
	JsonNumber = JsonNumberAssertion("")
	Slice      = SliceAssertion("")
	Value      = ValueAssertion("")
)

type (
	IntValidator        func(v int64, name string) error
	FloatValidator      func(v float64, name string) error
	StringValidator     func(v string, name string) error
	JsonNumberValidator func(v json.Number, name string) error
	ValueValidator      func(v interface{}, name string) error
)

// Reason
const (
	ERR_OUT_OF_RANGE          = "out of range"
	ERR_NON_NEGATIVE_INTENGER = "should be a non-negative integer"
	ERR_NON_NEGATIVE_NUMBER   = "should be a non-negative number"
	ERR_NON_ZERO              = "should not be zero"
	ERR_NAN_OR_INFINITY       = "cannot be -inf, +inf or NaN"
	ERR_NIL                   = "cannot be nil"
	ERR_EMPTY_ARRAY           = "cannot be an empty array"
	ERR_EMPTY_STRING          = "cannot be an empty string"
	ERR_NOT_ARRAY             = "should be an array"
	ERR_INVALID_INTEGER       = "specified integer %d is invalid"
	ERR_INVALID_STRING        = "specified string '%s' is invalid"
)
