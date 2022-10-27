// Package arg implements functions to validate arguments.
//
// Synopsis:
//  var (
//  	host     string = "127.0.0.1"
//  	status   string = "connecting"
//  	requests int64  = 99
//  )
//
//  err := arg.Assert(
//  	arg.Strings.Assert(host, "host",
//  		arg.Strings.NonEmpty,
//  	),
//  	arg.Strings.Assert(status, "status",
//  		arg.Strings.NonEmpty,
//  		arg.Strings.In("connecting", "closed", "aborted"),
//  	),
//  	arg.Ints.Assert(requests, "requrests",
//  		arg.Ints.NonNegativeInteger,
//  	),
//  )
//  fmt.Println(err)  // Output: <nil>
//
// The Strings, Floats, Ints, Numbers, Values also offer `Assertor()`:
//
//  err := arg.Strings.Assertor(host, "host").
//  	Assert(
//  		arg.Strings.NonEmpty,
//  	)
//
// is preferable to
//
//  err := arg.Strings.Assert(host, "host",
//  	arg.Strings.NonEmpty,
//  )
//
package arg

func Assert(errs ...error) error {
	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}

func ThrowError(name, reason string) error {
	return &InvalidArgumentError{
		Name:   name,
		Reason: reason,
	}
}
