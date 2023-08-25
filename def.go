package arg

import (
	"encoding/json"
	"net"
	"time"

	"github.com/Bofry/arg/content"
)

const (
	Ints      = IntAssertion("")
	UInts     = UIntAssertion("")
	Floats    = FloatAssertion("")
	Numbers   = NumberAssertion("")
	Strings   = StringAssertion("")
	Slices    = SliceAssertion("")
	Values    = ValueAssertion("")
	IPs       = IPAssertion("")
	Times     = TimeAssertion("")
	Durations = DurationAssertion("")

	IntPtr      = IntPtrAssertion("")
	UIntPtr     = UIntPtrAssertion("")
	FloatPtr    = FloatPtrAssertion("")
	NumberPtr   = NumberPtrAssertion("")
	StringPtr   = StringPtrAssertion("")
	TimePtr     = TimePtrAssertion("")
	DurationPtr = DurationPtrAssertion("")
)

type (
	Number = json.Number
	IP     = net.IP
)

type (
	IntValidator      func(v int64, name string) error
	UIntValidator     func(v uint64, name string) error
	FloatValidator    func(v float64, name string) error
	StringValidator   func(v string, name string) error
	NumberValidator   func(v Number, name string) error
	ValueValidator    func(v interface{}, name string) error
	IPValidator       func(v IP, name string) error
	TimeValidator     func(v time.Time, name string) error
	DurationValidator func(v time.Duration, name string) error

	IntPtrValidator      func(v *int64, name string) error
	UIntPtrValidator     func(v *uint64, name string) error
	FloatPtrValidator    func(v *float64, name string) error
	StringPtrValidator   func(v *string, name string) error
	NumberPtrValidator   func(v *Number, name string) error
	TimePtrValidator     func(v *time.Time, name string) error
	DurationPtrValidator func(v *time.Duration, name string) error

	IntPredicate      func(v int64) bool
	UIntPredicate     func(v uint64) bool
	FloatPredicate    func(v float64) bool
	StringPredicate   func(v string) bool
	NumberPredicate   func(v Number) bool
	ValuePredicate    func(v interface{}) bool
	IPPredicate       func(v IP) bool
	TimePredicate     func(v time.Time) bool
	DurationPredicate func(v time.Duration) bool

	IntPtrPredicate      func(v *int64) bool
	UIntPtrPredicate     func(v *uint64) bool
	FloatPtrPredicate    func(v *float64) bool
	StringPtrPredicate   func(v *string) bool
	NumberPtrPredicate   func(v *Number) bool
	TimePtrPredicate     func(v *time.Time) bool
	DurationPtrPredicate func(v *time.Duration) bool
)

type (
	Validatable = content.Validatable
)
