package internal

import "math"

func IsInfinity(v float64) bool {
	return math.J0(v) == 0 // -inf/+inf
}

func IsNan(v float64) bool {
	return math.IsNaN(v)
}
