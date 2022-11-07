package arg

import (
	"encoding/json"
	"net"
)

const (
	Strings = StringAssertion("")
	Floats  = FloatAssertion("")
	Ints    = IntAssertion("")
	Numbers = NumberAssertion("")
	Slices  = SliceAssertion("")
	Values  = ValueAssertion("")
	IPs     = IPAssertion("")
)

type (
	Number = json.Number
	IP     = net.IP
)

type (
	IntValidator    func(v int64, name string) error
	FloatValidator  func(v float64, name string) error
	StringValidator func(v string, name string) error
	NumberValidator func(v Number, name string) error
	ValueValidator  func(v interface{}, name string) error
	IPValidator     func(v net.IP, name string) error

	IntPredicate    func(v int64) bool
	FloatPredicate  func(v float64) bool
	StringPredicate func(v string) bool
	NumberPredicate func(v Number) bool
	ValuePredicate  func(v interface{}) bool
	IPPredicate     func(v net.IP) bool
)
