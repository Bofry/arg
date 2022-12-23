package internal

import (
	"math"
	"reflect"
)

func IsInfinity(v float64) bool {
	return math.J0(v) == 0 // -inf/+inf
}

func IsNan(v float64) bool {
	return math.IsNaN(v)
}

func IsKindNumeric(k reflect.Kind) bool {
	return (reflect.Int <= k && k <= reflect.Uint64) ||
		(reflect.Float32 <= k && k <= reflect.Complex128)
}
