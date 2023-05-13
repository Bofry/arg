package arg

import "testing"

func TestUIntAssertion(t *testing.T) {
	var (
		zeroInt      uint64 = 0
		negateOneInt uint64 = 1
		sixInt       uint64 = 6
	)

	{
		err := _UIntAssertion.Assert(zeroInt, "zeroInt",
			_UIntAssertion.NonZero,
			_UIntAssertion.LessOrEqual(3),
			_UIntAssertion.BetweenRange(2, 4),
		)
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"zeroInt\"; should not be zero"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
	{
		err := _UIntAssertion.Assert(negateOneInt, "negateOneInt",
			_UIntAssertion.NonZero,
			_UIntAssertion.LessOrEqual(3),
			_UIntAssertion.BetweenRange(2, 4),
		)
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"negateOneInt\"; out of range"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
	{
		err := _UIntAssertion.Assert(sixInt, "sixInt",
			_UIntAssertion.NonZero,
			_UIntAssertion.LessOrEqual(3),
			_UIntAssertion.BetweenRange(2, 4),
		)
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"sixInt\"; out of range"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
}

func TestUIntAssertor(t *testing.T) {
	var (
		zeroInt uint64 = 0
	)

	{
		err := _UIntAssertion.Assertor(zeroInt, "zeroInt").
			Assert(
				_UIntAssertion.NonZero,
				_UIntAssertion.LessOrEqual(3),
				_UIntAssertion.BetweenRange(2, 4),
			)
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"zeroInt\"; should not be zero"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
}

func TestUIntAssertion_NonZero(t *testing.T) {
	{
		var arg uint64 = 5
		err := _UIntAssertion.NonZero(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg uint64 = 1
		err := _UIntAssertion.NonZero(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg uint64 = 0
		err := _UIntAssertion.NonZero(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"arg\"; should not be zero"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
}

func TestUIntAssertion_NotIn(t *testing.T) {
	var validate UIntValidator = _UIntAssertion.NotIn(3, 6)
	{
		var arg uint64 = 1
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg uint64 = 6
		err := validate(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"arg\"; specified integer 6 is invalid"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
	{
		var arg uint64 = 3
		err := validate(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"arg\"; specified integer 3 is invalid"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
}

func TestUIntAssertion_In(t *testing.T) {
	var validate UIntValidator = _UIntAssertion.In(3, 6)
	{
		var arg uint64 = 1
		err := validate(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"arg\"; specified integer 1 is invalid"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
	{
		var arg uint64 = 6
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg uint64 = 3
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
}

func TestUIntAssertion_Must(t *testing.T) {
	var validate UIntValidator = _UIntAssertion.Must(
		func(v uint64) bool {
			return (v & 0x01) == 0
		})

	{
		var arg uint64 = 1
		err := validate(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"arg\"; specified integer 1 is invalid"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
	{
		var arg uint64 = 6
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg uint64 = 3
		err := validate(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"arg\"; specified integer 3 is invalid"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
}

func TestUIntAssertion_LessOrEqual(t *testing.T) {
	var validate UIntValidator = _UIntAssertion.LessOrEqual(5)

	{
		var arg uint64 = 5
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg uint64 = 4
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg uint64 = 6
		err := validate(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"arg\"; out of range"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
}

func TestUIntAssertion_GreaterOrEqual(t *testing.T) {
	var validate UIntValidator = _UIntAssertion.GreaterOrEqual(5)

	{
		var arg uint64 = 5
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg uint64 = 6
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg uint64 = 4
		err := validate(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"arg\"; out of range"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
}

func TestUIntAssertion_BetweenRange(t *testing.T) {
	var validate UIntValidator = _UIntAssertion.BetweenRange(2, 4)

	{
		var arg uint64 = 4
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg uint64 = 2
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg uint64 = 1
		err := validate(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"arg\"; out of range"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
	{
		var arg uint64 = 5
		err := validate(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"arg\"; out of range"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
}
