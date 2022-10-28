package arg_test

import (
	"fmt"
	"math"
	"strings"

	"github.com/Bofry/arg"
)

func Example() {
	var (
		host     string = "127.0.0.1"
		status   string = "connecting"
		requests int64  = 99
	)

	err := arg.Assert(
		arg.Strings.Assert(host, "host",
			arg.Strings.NonEmpty,
		),
		arg.Strings.Assert(status, "status",
			arg.Strings.NonEmpty,
			arg.Strings.In("connecting", "closed", "aborted"),
		),
		arg.Ints.Assert(requests, "requrests",
			arg.Ints.NonNegativeInteger,
		),
	)
	fmt.Println(err)
	// Output: <nil>
}

func ExampleFloatAssertion_Must() {
	var v float64 = 4.9001
	err := arg.Floats.Assert(v, "v",
		arg.Floats.Must(
			func(v float64) bool {
				return math.Round(v) == 5
			}),
	)

	fmt.Println(err)
	// Output: <nil>
}

func ExampleIntAssertion_Must() {
	var v int64 = 6
	err := arg.Ints.Assert(v, "v",
		arg.Ints.Must(
			func(v int64) bool {
				return (v & 0x01) == 0
			}),
	)

	fmt.Println(err)
	// Output: <nil>
}

func ExampleNumberAssertion_Must() {
	var v arg.Number = "6"
	err := arg.Numbers.Assert(v, "v",
		arg.Numbers.Must(
			func(v arg.Number) bool {
				argv, err := v.Int64()
				if err != nil {
					return false
				}
				return (argv & 0x01) == 0
			}),
	)

	fmt.Println(err)
	// Output: <nil>
}

func ExampleStringAssertion_Must() {
	var v string = "ENV_FOO"
	err := arg.Strings.Assert(v, "v",
		arg.Strings.Must(
			func(v string) bool {
				return strings.HasPrefix(v, "ENV_")
			}),
	)

	fmt.Println(err)
	// Output: <nil>
}

func ExampleStringAssertion_MatchAny() {
	var v string = "demo@mail.com"
	err := arg.Strings.Assert(v, "v",
		arg.Strings.MatchAny(
			arg.EmailPattern,
		),
	)

	fmt.Println(err)
	// Output: <nil>
}
