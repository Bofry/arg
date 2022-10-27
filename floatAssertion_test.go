package arg

import (
	"math"
	"testing"
)

func TestFloatAssertion(t *testing.T) {
	var (
		nanFloat       float64 = math.NaN()
		infFloat       float64 = math.Inf(0)
		zeroFloat      float64 = 0
		negateFloat    float64 = -0.001
		four_nineFloat float64 = 4.9
		five_sixFloat  float64 = 5.6
	)

	{
		err := _FloatAssertion.Assert(nanFloat, "nanFloat",
			_FloatAssertion.NonNanNorInf,
			_FloatAssertion.NonNegativeNumber,
			_FloatAssertion.NonZero,
			_FloatAssertion.BetweenRange(4.899, 5.001),
		)
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument \"nanFloat\"; cannot be -inf, +inf or NaN"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
	{
		err := _FloatAssertion.Assert(infFloat, "infFloat",
			_FloatAssertion.NonNanNorInf,
			_FloatAssertion.NonNegativeNumber,
			_FloatAssertion.NonZero,
			_FloatAssertion.BetweenRange(4.899, 5.001),
		)
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument \"infFloat\"; cannot be -inf, +inf or NaN"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
	{
		err := _FloatAssertion.Assert(zeroFloat, "zeroFloat",
			_FloatAssertion.NonNanNorInf,
			_FloatAssertion.NonNegativeNumber,
			_FloatAssertion.NonZero,
			_FloatAssertion.BetweenRange(4.899, 5.001),
		)
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument \"zeroFloat\"; should not be zero"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
	{
		err := _FloatAssertion.Assert(negateFloat, "negateFloat",
			_FloatAssertion.NonNanNorInf,
			_FloatAssertion.NonNegativeNumber,
			_FloatAssertion.NonZero,
			_FloatAssertion.BetweenRange(4.899, 5.001),
		)
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument \"negateFloat\"; should be a non-negative number"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
	{
		err := _FloatAssertion.Assert(five_sixFloat, "five_sixFloat",
			_FloatAssertion.NonNanNorInf,
			_FloatAssertion.NonNegativeNumber,
			_FloatAssertion.NonZero,
			_FloatAssertion.BetweenRange(4.899, 5.001),
		)
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument \"five_sixFloat\"; out of range"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
	{
		err := _FloatAssertion.Assert(four_nineFloat, "four_nineFloat",
			_FloatAssertion.NonNanNorInf,
			_FloatAssertion.NonNegativeNumber,
			_FloatAssertion.NonZero,
			_FloatAssertion.BetweenRange(4.899, 5.001),
		)
		if err != nil {
			t.Errorf("should not error")
		}
	}
}

func TestFloatAssertor(t *testing.T) {
	var (
		nanFloat float64 = math.NaN()
	)

	{
		err := _FloatAssertion.Assertor(nanFloat, "nanFloat").
			Assert(
				_FloatAssertion.NonNanNorInf,
				_FloatAssertion.NonNegativeNumber,
				_FloatAssertion.NonZero,
				_FloatAssertion.BetweenRange(4.899, 5.001),
			)
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument \"nanFloat\"; cannot be -inf, +inf or NaN"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
}

func TestFloatAssertion_NonNanNorInf(t *testing.T) {
	{
		var arg float64 = -0.0001
		err := _FloatAssertion.NonNanNorInf(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg float64 = math.Inf(0)
		err := _FloatAssertion.NonNanNorInf(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument \"arg\"; cannot be -inf, +inf or NaN"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
	{
		var arg float64 = math.NaN()
		err := _FloatAssertion.NonNanNorInf(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument \"arg\"; cannot be -inf, +inf or NaN"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
}

func TestFloatAssertion_NonNegativeNumber(t *testing.T) {
	{
		var arg float64 = 0
		err := _FloatAssertion.NonNegativeNumber(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg float64 = 0.001
		err := _FloatAssertion.NonNegativeNumber(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg float64 = -0.0001
		err := _FloatAssertion.NonNegativeNumber(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument \"arg\"; should be a non-negative number"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
}

func TestFloatAssertion_NonZero(t *testing.T) {
	{
		var arg float64 = -0.0001
		err := _FloatAssertion.NonZero(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg float64 = 0.001
		err := _FloatAssertion.NonZero(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg float64 = 0
		err := _FloatAssertion.NonZero(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument \"arg\"; should not be zero"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
}

func TestFloatAssertion_Must(t *testing.T) {
	var validate FloatValidator = _FloatAssertion.Must(
		func(v float64) bool {
			return math.Round(v) == 5
		})

	{
		var arg float64 = -0.0001
		err := validate(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument \"arg\"; specified number -0.0001 is invalid"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
	{
		var arg float64 = 4.9001
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg float64 = 5.0001
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
}

func TestFloatAssertion_Less(t *testing.T) {
	var validate FloatValidator = _FloatAssertion.Less(5.0001)

	{
		var arg float64 = -0.0001
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg float64 = 5.0001
		err := validate(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument \"arg\"; out of range"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
	{
		var arg float64 = 5.0002
		err := validate(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument \"arg\"; out of range"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
}

func TestFloatAssertion_LessOrEqual(t *testing.T) {
	var validate FloatValidator = _FloatAssertion.LessOrEqual(5.0001)

	{
		var arg float64 = -0.0001
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg float64 = 5.0001
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg float64 = 5.0002
		err := validate(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument \"arg\"; out of range"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
}

func TestFloatAssertion_Greater(t *testing.T) {
	var validate FloatValidator = _FloatAssertion.Greater(5.0001)

	{
		var arg float64 = 5.0002
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg float64 = 5.0001
		err := validate(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument \"arg\"; out of range"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
	{
		var arg float64 = -0.0001
		err := validate(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument \"arg\"; out of range"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
}

func TestFloatAssertion_GreaterOrEqual(t *testing.T) {
	var validate FloatValidator = _FloatAssertion.GreaterOrEqual(5.0001)

	{
		var arg float64 = 5.0002
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg float64 = 5.0001
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg float64 = -0.0001
		err := validate(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument \"arg\"; out of range"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
}

func TestFloatAssertion_BetweenRange(t *testing.T) {
	var validate FloatValidator = _FloatAssertion.BetweenRange(4.8999, 5.0001)

	{
		var arg float64 = 4.8999
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg float64 = 4.89991
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg float64 = 5.0001
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg float64 = 5.0000999999
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg float64 = 4.8998999999
		err := validate(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument \"arg\"; out of range"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
	{
		var arg float64 = 5.0001000001
		err := validate(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument \"arg\"; out of range"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
	{
		var arg float64 = -0.0001
		err := validate(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument \"arg\"; out of range"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
}
