arg
========
Package implements functions to validate arguments.

## Synopsis
```go
func main() {
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
```

