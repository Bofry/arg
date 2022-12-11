package arg

import "testing"

func TestIntAssertion(t *testing.T) {
	var (
		zeroInt      int64 = 0
		negateOneInt int64 = -1
		sixInt       int64 = 6
	)

	{
		err := _IntAssertion.Assert(zeroInt, "zeroInt",
			_IntAssertion.NonNegativeInteger,
			_IntAssertion.NonZero,
			_IntAssertion.LessOrEqual(3),
			_IntAssertion.BetweenRange(-2, 4),
		)
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"zeroInt\"; should not be zero"
		if err.Error() != expectedErrorMsg {
			t.Errorf("except: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
	{
		err := _IntAssertion.Assert(negateOneInt, "negateOneInt",
			_IntAssertion.NonNegativeInteger,
			_IntAssertion.NonZero,
			_IntAssertion.LessOrEqual(3),
			_IntAssertion.BetweenRange(-2, 4),
		)
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"negateOneInt\"; should be a non-negative integer"
		if err.Error() != expectedErrorMsg {
			t.Errorf("except: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
	{
		err := _IntAssertion.Assert(sixInt, "sixInt",
			_IntAssertion.NonNegativeInteger,
			_IntAssertion.NonZero,
			_IntAssertion.LessOrEqual(3),
			_IntAssertion.BetweenRange(-2, 4),
		)
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"sixInt\"; out of range"
		if err.Error() != expectedErrorMsg {
			t.Errorf("except: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
	{
		err := _IntAssertion.Assert(negateOneInt, "negateOneInt",
			_IntAssertion.NonZero,
			_IntAssertion.LessOrEqual(3),
			_IntAssertion.BetweenRange(-2, 4),
		)
		if err != nil {
			t.Errorf("should not error")
		}
	}
}

func TestIntAssertor(t *testing.T) {
	var (
		zeroInt int64 = 0
	)

	{
		err := _IntAssertion.Assertor(zeroInt, "zeroInt").
			Assert(
				_IntAssertion.NonNegativeInteger,
				_IntAssertion.NonZero,
				_IntAssertion.LessOrEqual(3),
				_IntAssertion.BetweenRange(-2, 4),
			)
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"zeroInt\"; should not be zero"
		if err.Error() != expectedErrorMsg {
			t.Errorf("except: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
}

func TestIntAssertion_NonNegativeInteger(t *testing.T) {
	{
		var arg int64 = 0
		err := _IntAssertion.NonNegativeInteger(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg int64 = 1
		err := _IntAssertion.NonNegativeInteger(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg int64 = -1
		err := _IntAssertion.NonNegativeInteger(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"arg\"; should be a non-negative integer"
		if err.Error() != expectedErrorMsg {
			t.Errorf("except: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
}

func TestIntAssertion_NonZero(t *testing.T) {
	{
		var arg int64 = -1
		err := _IntAssertion.NonZero(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg int64 = 1
		err := _IntAssertion.NonZero(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg int64 = 0
		err := _IntAssertion.NonZero(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"arg\"; should not be zero"
		if err.Error() != expectedErrorMsg {
			t.Errorf("except: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
}

func TestIntAssertion_NotIn(t *testing.T) {
	var validate IntValidator = _IntAssertion.NotIn(3, 6)
	{
		var arg int64 = -1
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg int64 = 6
		err := validate(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"arg\"; specified integer 6 is invalid"
		if err.Error() != expectedErrorMsg {
			t.Errorf("except: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
	{
		var arg int64 = 3
		err := validate(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"arg\"; specified integer 3 is invalid"
		if err.Error() != expectedErrorMsg {
			t.Errorf("except: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
}

func TestIntAssertion_In(t *testing.T) {
	var validate IntValidator = _IntAssertion.In(3, 6)
	{
		var arg int64 = 1
		err := validate(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"arg\"; specified integer 1 is invalid"
		if err.Error() != expectedErrorMsg {
			t.Errorf("except: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
	{
		var arg int64 = 6
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg int64 = 3
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
}

func TestIntAssertion_Must(t *testing.T) {
	var validate IntValidator = _IntAssertion.Must(
		func(v int64) bool {
			return (v & 0x01) == 0
		})

	{
		var arg int64 = 1
		err := validate(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"arg\"; specified integer 1 is invalid"
		if err.Error() != expectedErrorMsg {
			t.Errorf("except: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
	{
		var arg int64 = 6
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg int64 = 3
		err := validate(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"arg\"; specified integer 3 is invalid"
		if err.Error() != expectedErrorMsg {
			t.Errorf("except: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
}

func TestIntAssertion_LessOrEqual(t *testing.T) {
	var validate IntValidator = _IntAssertion.LessOrEqual(5)

	{
		var arg int64 = 5
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg int64 = 4
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg int64 = 6
		err := validate(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"arg\"; out of range"
		if err.Error() != expectedErrorMsg {
			t.Errorf("except: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
}

func TestIntAssertion_GreaterOrEqual(t *testing.T) {
	var validate IntValidator = _IntAssertion.GreaterOrEqual(5)

	{
		var arg int64 = 5
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg int64 = 6
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg int64 = 4
		err := validate(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"arg\"; out of range"
		if err.Error() != expectedErrorMsg {
			t.Errorf("except: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
}

func TestIntAssertion_BetweenRange(t *testing.T) {
	var validate IntValidator = _IntAssertion.BetweenRange(-2, 4)

	{
		var arg int64 = 4
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg int64 = -2
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg int64 = 0
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg int64 = -3
		err := validate(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"arg\"; out of range"
		if err.Error() != expectedErrorMsg {
			t.Errorf("except: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
	{
		var arg int64 = 5
		err := validate(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"arg\"; out of range"
		if err.Error() != expectedErrorMsg {
			t.Errorf("except: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
}
