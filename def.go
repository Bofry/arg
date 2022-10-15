package arg

import (
	"github.com/Bofry/arg/internal"
)

const (
	String     = internal.String
	Float      = internal.Float
	Int        = internal.Int
	JsonNumber = internal.JsonNumber
	Slice      = internal.Slice
	Value      = internal.Value
)

type (
	IntValidator        = internal.IntValidator
	FloatValidator      = internal.FloatValidator
	StringValidator     = internal.StringValidator
	JsonNumberValidator = internal.JsonNumberValidator
	ValueValidator      = internal.ValueValidator

	InvalidArgumentError = internal.InvalidArgumentError
)
