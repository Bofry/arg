package assert

import (
	"github.com/Bofry/arg"
)

// Reason
const (
	ERR_OUT_OF_RANGE              = "out of range"
	ERR_NON_NEGATIVE_INTENGER     = "should be a non-negative integer"
	ERR_NON_NEGATIVE_NUMBER       = "should be a non-negative number"
	ERR_NAN_OR_INFINITY           = "cannot be -inf, +inf or NaN"
	ERR_NIL                       = "cannot be nil"
	ERR_EMPTY_ARRAY               = "cannot be an empty array"
	ERR_EMPTY_STRING              = "cannot be an empty string"
	ERR_NOT_ARRAY                 = "should be an array"
	ERR_INVALID_INTEGER_ASSERTION = "specified integer %d is invalid"
)

type (
	InvalidArgumentError = arg.InvalidArgumentError
)

type (
	IntegerValidator func(v int64, name string) error
	FloatValidator   func(v float64, name string) error
)
