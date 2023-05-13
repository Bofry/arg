package arg

import "testing"

func TestStringPtrAssertion(t *testing.T) {
	var NewString = func(v string) *string {
		return &v
	}

	var (
		nilString    *string = nil
		emprtyString *string = NewString("")
		fooString    *string = NewString("foo")
		barString    *string = NewString("bar")
		bazString    *string = NewString("baz")
	)

	{
		err := _StringPtrAssertion.Assert(nilString, "nilString",
			_StringPtrAssertion.NotNil,
			_StringPtrAssertion.NonEmpty,
			_StringPtrAssertion.In("foo", "bar"),
		)
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"nilString\"; cannot be nil"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
	{
		err := _StringPtrAssertion.Assert(emprtyString, "emprtyString",
			_StringPtrAssertion.NonEmpty,
			_StringPtrAssertion.In("foo", "bar"),
		)
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"emprtyString\"; cannot be an empty string"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
	{
		err := _StringPtrAssertion.Assert(emprtyString, "emprtyString",
			_StringPtrAssertion.NonEmpty,
			_StringPtrAssertion.In("foo", "bar"),
		)
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"emprtyString\"; cannot be an empty string"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
	{
		err := _StringPtrAssertion.Assert(nilString, "nilString",
			_StringPtrAssertion.NonEmpty,
			_StringPtrAssertion.In("foo", "bar"),
		)
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		err := _StringPtrAssertion.Assert(fooString, "fooString",
			_StringPtrAssertion.NonEmpty,
			_StringPtrAssertion.In("foo", "bar"),
		)
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		err := _StringPtrAssertion.Assert(barString, "barString",
			_StringPtrAssertion.NonEmpty,
			_StringPtrAssertion.In("foo", "bar"),
		)
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		err := _StringPtrAssertion.Assert(bazString, "bazString",
			_StringPtrAssertion.NonEmpty,
			_StringPtrAssertion.In("foo", "bar"),
		)
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"bazString\"; specified string \"baz\" is invalid"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
}

func TestStringPtrAssertor(t *testing.T) {
	var NewString = func(v string) *string {
		return &v
	}

	var (
		nilString    *string = nil
		emprtyString *string = NewString("")
		fooString    *string = NewString("foo")
		barString    *string = NewString("bar")
		bazString    *string = NewString("baz")
	)

	{
		err := _StringPtrAssertion.Assertor(nilString, "nilString").
			Assert(
				_StringPtrAssertion.NotNil,
				_StringPtrAssertion.NonEmpty,
				_StringPtrAssertion.In("foo", "bar"),
			)
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"nilString\"; cannot be nil"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
	{
		err := _StringPtrAssertion.Assertor(emprtyString, "emprtyString").
			Assert(
				_StringPtrAssertion.NonEmpty,
				_StringPtrAssertion.In("foo", "bar"),
			)
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"emprtyString\"; cannot be an empty string"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
	{
		err := _StringPtrAssertion.Assertor(nilString, "nilString").
			Assert(
				_StringPtrAssertion.NonEmpty,
				_StringPtrAssertion.In("foo", "bar"),
			)
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		err := _StringPtrAssertion.Assertor(fooString, "fooString").
			Assert(
				_StringPtrAssertion.NonEmpty,
				_StringPtrAssertion.In("foo", "bar"),
			)
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		err := _StringPtrAssertion.Assertor(barString, "barString").
			Assert(
				_StringPtrAssertion.NonEmpty,
				_StringPtrAssertion.In("foo", "bar"),
			)
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		err := _StringPtrAssertion.Assertor(bazString, "bazString").
			Assert(
				_StringPtrAssertion.NonEmpty,
				_StringPtrAssertion.In("foo", "bar"),
			)
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"bazString\"; specified string \"baz\" is invalid"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
}

func TestStringPtrAssertion_NotNil(t *testing.T) {
	var NewString = func(v string) *string {
		return &v
	}

	{
		var arg *string = NewString("foo")
		err := _StringPtrAssertion.NotNil(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg *string = nil
		err := _StringPtrAssertion.NotNil(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"arg\"; cannot be nil"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
}
