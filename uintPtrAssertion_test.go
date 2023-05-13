package arg

import "testing"

func TestUIntPtrAssertion(t *testing.T) {
	var NewUInt64 = func(v uint64) *uint64 {
		return &v
	}

	var (
		nilInt       *uint64 = nil
		zeroInt      *uint64 = NewUInt64(0)
		negateOneInt *uint64 = NewUInt64(1)
		sixInt       *uint64 = NewUInt64(6)
	)

	{
		err := _UIntPtrAssertion.Assert(nilInt, "nilInt",
			_UIntPtrAssertion.NotNil,
			_UIntPtrAssertion.NonZero,
			_UIntPtrAssertion.LessOrEqual(3),
			_UIntPtrAssertion.BetweenRange(2, 4),
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
		err := _UIntPtrAssertion.Assert(zeroInt, "zeroInt",
			_UIntPtrAssertion.NonZero,
			_UIntPtrAssertion.LessOrEqual(3),
			_UIntPtrAssertion.BetweenRange(2, 4),
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
		err := _UIntPtrAssertion.Assert(negateOneInt, "negateOneInt",
			_UIntPtrAssertion.NonZero,
			_UIntPtrAssertion.LessOrEqual(3),
			_UIntPtrAssertion.BetweenRange(2, 4),
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
		err := _UIntPtrAssertion.Assert(sixInt, "sixInt",
			_UIntPtrAssertion.NonZero,
			_UIntPtrAssertion.LessOrEqual(3),
			_UIntPtrAssertion.BetweenRange(2, 4),
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
		err := _UIntPtrAssertion.Assert(nilInt, "nilInt",
			_UIntPtrAssertion.NonZero,
			_UIntPtrAssertion.LessOrEqual(3),
			_UIntPtrAssertion.BetweenRange(2, 4),
		)
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		err := _UIntPtrAssertion.Assert(negateOneInt, "negateOneInt",
			_UIntPtrAssertion.NonZero,
			_UIntPtrAssertion.LessOrEqual(3),
		)
		if err != nil {
			t.Errorf("should not error")
		}
	}

}

func TestUIntPtrAssertor(t *testing.T) {
	var NewUInt64 = func(v uint64) *uint64 {
		return &v
	}

	var (
		zeroInt *uint64 = NewUInt64(0)
	)

	{
		err := _UIntPtrAssertion.Assertor(zeroInt, "zeroInt").
			Assert(
				_UIntPtrAssertion.NonZero,
				_UIntPtrAssertion.LessOrEqual(3),
				_UIntPtrAssertion.BetweenRange(2, 4),
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

func TestUIntPtrAssertion_NotNil(t *testing.T) {
	var NewUInt64 = func(v uint64) *uint64 {
		return &v
	}

	{
		var arg *uint64 = NewUInt64(0)
		err := _UIntPtrAssertion.NotNil(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg *uint64 = nil
		err := _UIntPtrAssertion.NotNil(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"arg\"; cannot be nil"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
}
