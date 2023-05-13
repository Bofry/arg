package arg

import "testing"

func TestNumberPtrAssertion(t *testing.T) {
	var NewNumber = func(v Number) *Number {
		return &v
	}

	var (
		nilNumber       *Number = nil
		nonNumber       *Number = NewNumber("unknown")
		nanNumber       *Number = NewNumber("nan")
		zeroNumber      *Number = NewNumber("0")
		negateNumber    *Number = NewNumber("-0.001")
		four_nineNumber *Number = NewNumber("4.9")
	)

	{
		err := _NumberPtrAssertion.Assert(nilNumber, "nilNumber",
			_NumberPtrAssertion.NotNil,
			_NumberPtrAssertion.NonNanNorInf,
			_NumberPtrAssertion.NonNegativeNumber,
			_NumberPtrAssertion.NonZero,
			_NumberPtrAssertion.BetweenRange(4.899, 5.001),
		)
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"nilNumber\"; cannot be nil"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
	{
		err := _NumberPtrAssertion.Assert(nonNumber, "nonNumber",
			_NumberPtrAssertion.NonNanNorInf,
			_NumberPtrAssertion.NonNegativeNumber,
			_NumberPtrAssertion.NonZero,
			_NumberPtrAssertion.BetweenRange(4.899, 5.001),
		)
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"nonNumber\"; specified number \"unknown\" is invalid"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
	{
		err := _NumberPtrAssertion.Assert(nanNumber, "nanNumber",
			_NumberPtrAssertion.NonNanNorInf,
			_NumberPtrAssertion.NonNegativeNumber,
			_NumberPtrAssertion.NonZero,
			_NumberPtrAssertion.BetweenRange(4.899, 5.001),
		)
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"nanNumber\"; cannot be -inf, +inf or NaN"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
	{
		err := _NumberPtrAssertion.Assert(zeroNumber, "zeroNumber",
			_NumberPtrAssertion.NonNanNorInf,
			_NumberPtrAssertion.NonNegativeNumber,
			_NumberPtrAssertion.NonZero,
			_NumberPtrAssertion.BetweenRange(4.899, 5.001),
		)
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"zeroNumber\"; should not be zero"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
	{
		err := _NumberPtrAssertion.Assert(negateNumber, "negateNumber",
			_NumberPtrAssertion.NonNanNorInf,
			_NumberPtrAssertion.NonNegativeNumber,
			_NumberPtrAssertion.NonZero,
			_NumberPtrAssertion.BetweenRange(4.899, 5.001),
		)
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"negateNumber\"; should be a non-negative number"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
	{
		err := _NumberPtrAssertion.Assert(nilNumber, "nilNumber",
			_NumberPtrAssertion.NonNanNorInf,
			_NumberPtrAssertion.NonNegativeNumber,
			_NumberPtrAssertion.NonZero,
			_NumberPtrAssertion.BetweenRange(4.899, 5.001),
		)
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		err := _NumberPtrAssertion.Assert(four_nineNumber, "four_nineNumber",
			_NumberPtrAssertion.NonNanNorInf,
			_NumberPtrAssertion.NonNegativeNumber,
			_NumberPtrAssertion.NonZero,
			_NumberPtrAssertion.BetweenRange(4.899, 5.001),
		)
		if err != nil {
			t.Errorf("should not error")
		}
	}
}

func TestNumberPtrAssertor(t *testing.T) {
	var NewNumber = func(v Number) *Number {
		return &v
	}

	var (
		nanNumber *Number = NewNumber("nan")
	)

	{
		err := _NumberPtrAssertion.Assertor(nanNumber, "nanNumber").
			Assert(
				_NumberPtrAssertion.NonNanNorInf,
				_NumberPtrAssertion.NonNegativeNumber,
				_NumberPtrAssertion.NonZero,
				_NumberPtrAssertion.BetweenRange(4.899, 5.001),
			)
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"nanNumber\"; cannot be -inf, +inf or NaN"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
}

func TestNumberPtrAssertor_NotNil(t *testing.T) {
	var NewNumber = func(v Number) *Number {
		return &v
	}

	{
		var arg *Number = NewNumber("nan")
		err := _NumberPtrAssertion.NotNil(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg *Number = nil
		err := _NumberPtrAssertion.NotNil(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"arg\"; cannot be nil"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
}
