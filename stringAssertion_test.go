package arg

import "testing"

func TestStringAssertion(t *testing.T) {
	var (
		emprtyString string = ""
		fooString    string = "foo"
		barString    string = "bar"
		bazString    string = "baz"
	)

	{
		err := _StringAssertion.Assert(emprtyString, "emprtyString",
			_StringAssertion.NonEmpty,
			_StringAssertion.In("foo", "bar"),
		)
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument 'emprtyString'; cannot be an empty string"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
	{
		err := _StringAssertion.Assert(fooString, "fooString",
			_StringAssertion.NonEmpty,
			_StringAssertion.In("foo", "bar"),
		)
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		err := _StringAssertion.Assert(barString, "barString",
			_StringAssertion.NonEmpty,
			_StringAssertion.In("foo", "bar"),
		)
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		err := _StringAssertion.Assert(bazString, "bazString",
			_StringAssertion.NonEmpty,
			_StringAssertion.In("foo", "bar"),
		)
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument 'bazString'; specified string 'baz' is invalid"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
}

func TestStringAssertion_WithAssertor(t *testing.T) {
	var (
		emprtyString string = ""
		fooString    string = "foo"
		barString    string = "bar"
		bazString    string = "baz"
	)

	{
		err := _StringAssertion.Assertor(emprtyString, "emprtyString").
			Assert(
				_StringAssertion.NonEmpty,
				_StringAssertion.In("foo", "bar"),
			)
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument 'emprtyString'; cannot be an empty string"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
	{
		err := _StringAssertion.Assertor(fooString, "fooString").
			Assert(
				_StringAssertion.NonEmpty,
				_StringAssertion.In("foo", "bar"),
			)
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		err := _StringAssertion.Assertor(barString, "barString").
			Assert(
				_StringAssertion.NonEmpty,
				_StringAssertion.In("foo", "bar"),
			)
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		err := _StringAssertion.Assertor(bazString, "bazString").
			Assert(
				_StringAssertion.NonEmpty,
				_StringAssertion.In("foo", "bar"),
			)
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument 'bazString'; specified string 'baz' is invalid"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
}

func TestStringAssertion_NonEmpty(t *testing.T) {
	{
		var arg string = "foo"
		err := _StringAssertion.NonEmpty(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg string = ""
		err := _StringAssertion.NonEmpty(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
	}
}

func TestStringAssertion_In(t *testing.T) {
	var validate StringValidator = _StringAssertion.In("foo", "bar")
	{
		var arg string = "foo"
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg string = "baz"
		err := validate(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
	}
}
