package arg

import "testing"

func TestIntPtrAssertion(t *testing.T) {
	var NewInt64 = func(v int64) *int64 {
		return &v
	}

	var (
		nilInt       *int64 = nil
		zeroInt      *int64 = NewInt64(0)
		negateOneInt *int64 = NewInt64(-1)
		sixInt       *int64 = NewInt64(6)
	)

	{
		err := _IntPtrAssertion.Assert(nilInt, "nilInt",
			_IntPtrAssertion.NotNil,
			_IntPtrAssertion.NonNegativeInteger,
			_IntPtrAssertion.NonZero,
			_IntPtrAssertion.LessOrEqual(3),
			_IntPtrAssertion.BetweenRange(-2, 4),
		)
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"nilInt\"; cannot be nil"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
	{
		err := _IntPtrAssertion.Assert(zeroInt, "zeroInt",
			_IntPtrAssertion.NonNegativeInteger,
			_IntPtrAssertion.NonZero,
			_IntPtrAssertion.LessOrEqual(3),
			_IntPtrAssertion.BetweenRange(-2, 4),
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
		err := _IntPtrAssertion.Assert(negateOneInt, "negateOneInt",
			_IntPtrAssertion.NonNegativeInteger,
			_IntPtrAssertion.NonZero,
			_IntPtrAssertion.LessOrEqual(3),
			_IntPtrAssertion.BetweenRange(-2, 4),
		)
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"negateOneInt\"; should be a non-negative integer"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
	{
		err := _IntPtrAssertion.Assert(sixInt, "sixInt",
			_IntPtrAssertion.NonNegativeInteger,
			_IntPtrAssertion.NonZero,
			_IntPtrAssertion.LessOrEqual(3),
			_IntPtrAssertion.BetweenRange(-2, 4),
		)
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"sixInt\"; out of range"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
	{
		err := _IntPtrAssertion.Assert(nilInt, "nilInt",
			_IntPtrAssertion.NonNegativeInteger,
			_IntPtrAssertion.NonZero,
			_IntPtrAssertion.LessOrEqual(3),
			_IntPtrAssertion.BetweenRange(-2, 4),
		)
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		err := _IntPtrAssertion.Assert(negateOneInt, "negateOneInt",
			_IntPtrAssertion.NonZero,
			_IntPtrAssertion.LessOrEqual(3),
			_IntPtrAssertion.BetweenRange(-2, 4),
		)
		if err != nil {
			t.Errorf("should not error")
		}
	}
}

func TestIntPtrAssertor(t *testing.T) {
	var NewInt64 = func(v int64) *int64 {
		return &v
	}

	var (
		zeroInt *int64 = NewInt64(0)
	)

	{
		err := _IntPtrAssertion.Assertor(zeroInt, "zeroInt").
			Assert(
				_IntPtrAssertion.NonNegativeInteger,
				_IntPtrAssertion.NonZero,
				_IntPtrAssertion.LessOrEqual(3),
				_IntPtrAssertion.BetweenRange(-2, 4),
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

func TestIntPtrAssertion_NotNil(t *testing.T) {
	var NewInt64 = func(v int64) *int64 {
		return &v
	}

	{
		var arg *int64 = NewInt64(0)
		err := _IntPtrAssertion.NotNil(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg *int64 = nil
		err := _IntPtrAssertion.NotNil(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"arg\"; cannot be nil"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
}
