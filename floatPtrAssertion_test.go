package arg

import (
	"math"
	"testing"
)

func TestFloatPtrAssertion(t *testing.T) {
	var NewFloat64 = func(v float64) *float64 {
		return &v
	}

	var (
		nilFloat       *float64 = nil
		nanFloat       *float64 = NewFloat64(math.NaN())
		infFloat       *float64 = NewFloat64(math.Inf(0))
		zeroFloat      *float64 = NewFloat64(0)
		negateFloat    *float64 = NewFloat64(-0.001)
		four_nineFloat *float64 = NewFloat64(4.9)
		five_sixFloat  *float64 = NewFloat64(5.6)
	)

	{
		err := _FloatPtrAssertion.Assert(nilFloat, "nilFloat",
			_FloatPtrAssertion.NotNil,
			_FloatPtrAssertion.NonNanNorInf,
			_FloatPtrAssertion.NonNegativeNumber,
			_FloatPtrAssertion.NonZero,
			_FloatPtrAssertion.BetweenRange(4.899, 5.001),
		)
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"nilFloat\"; cannot be nil"
		if err.Error() != expectedErrorMsg {
			t.Errorf("except: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
	{
		err := _FloatPtrAssertion.Assert(nanFloat, "nanFloat",
			_FloatPtrAssertion.NonNanNorInf,
			_FloatPtrAssertion.NonNegativeNumber,
			_FloatPtrAssertion.NonZero,
			_FloatPtrAssertion.BetweenRange(4.899, 5.001),
		)
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"nanFloat\"; cannot be -inf, +inf or NaN"
		if err.Error() != expectedErrorMsg {
			t.Errorf("except: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
	{
		err := _FloatPtrAssertion.Assert(infFloat, "infFloat",
			_FloatPtrAssertion.NonNanNorInf,
			_FloatPtrAssertion.NonNegativeNumber,
			_FloatPtrAssertion.NonZero,
			_FloatPtrAssertion.BetweenRange(4.899, 5.001),
		)
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"infFloat\"; cannot be -inf, +inf or NaN"
		if err.Error() != expectedErrorMsg {
			t.Errorf("except: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
	{
		err := _FloatPtrAssertion.Assert(zeroFloat, "zeroFloat",
			_FloatPtrAssertion.NonNanNorInf,
			_FloatPtrAssertion.NonNegativeNumber,
			_FloatPtrAssertion.NonZero,
			_FloatPtrAssertion.BetweenRange(4.899, 5.001),
		)
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"zeroFloat\"; should not be zero"
		if err.Error() != expectedErrorMsg {
			t.Errorf("except: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
	{
		err := _FloatPtrAssertion.Assert(negateFloat, "negateFloat",
			_FloatPtrAssertion.NonNanNorInf,
			_FloatPtrAssertion.NonNegativeNumber,
			_FloatPtrAssertion.NonZero,
			_FloatPtrAssertion.BetweenRange(4.899, 5.001),
		)
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"negateFloat\"; should be a non-negative number"
		if err.Error() != expectedErrorMsg {
			t.Errorf("except: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
	{
		err := _FloatPtrAssertion.Assert(five_sixFloat, "five_sixFloat",
			_FloatPtrAssertion.NonNanNorInf,
			_FloatPtrAssertion.NonNegativeNumber,
			_FloatPtrAssertion.NonZero,
			_FloatPtrAssertion.BetweenRange(4.899, 5.001),
		)
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"five_sixFloat\"; out of range"
		if err.Error() != expectedErrorMsg {
			t.Errorf("except: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
	
	{
		err := _FloatPtrAssertion.Assert(nilFloat, "nilFloat",
			_FloatPtrAssertion.NonNanNorInf,
			_FloatPtrAssertion.NonNegativeNumber,
			_FloatPtrAssertion.NonZero,
			_FloatPtrAssertion.BetweenRange(4.899, 5.001),
		)
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		err := _FloatPtrAssertion.Assert(four_nineFloat, "four_nineFloat",
			_FloatPtrAssertion.NonNanNorInf,
			_FloatPtrAssertion.NonNegativeNumber,
			_FloatPtrAssertion.NonZero,
			_FloatPtrAssertion.BetweenRange(4.899, 5.001),
		)
		if err != nil {
			t.Errorf("should not error")
		}
	}
}

func TestFloatPtrAssertor(t *testing.T) {
	var NewFloat64 = func(v float64) *float64 {
		return &v
	}

	var (
		nanFloat *float64 = NewFloat64(math.NaN())
	)

	{
		err := _FloatPtrAssertion.Assertor(nanFloat, "nanFloat").
			Assert(
				_FloatPtrAssertion.NonNanNorInf,
				_FloatPtrAssertion.NonNegativeNumber,
				_FloatPtrAssertion.NonZero,
				_FloatPtrAssertion.BetweenRange(4.899, 5.001),
			)
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"nanFloat\"; cannot be -inf, +inf or NaN"
		if err.Error() != expectedErrorMsg {
			t.Errorf("except: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
}

func TestFloatPtrAssertion_NotNil(t *testing.T) {
	var NewFloat64 = func(v float64) *float64 {
		return &v
	}

	{
		var arg *float64 = NewFloat64(math.NaN())
		err := _FloatPtrAssertion.NotNil(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg *float64 = nil
		err := _FloatPtrAssertion.NotNil(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"arg\"; cannot be nil"
		if err.Error() != expectedErrorMsg {
			t.Errorf("except: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
}
