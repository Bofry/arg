package arg

import (
	"math"
	"testing"
)

func TestNumberAssertion(t *testing.T) {
	var (
		nonNumber       Number = "unknown"
		nanNumber       Number = "nan"
		zeroNumber      Number = "0"
		negateNumber    Number = "-0.001"
		four_nineNumber Number = "4.9"
	)

	{
		err := _NumberAssertion.Assert(nonNumber, "nonNumber",
			_NumberAssertion.NonNanNorInf,
			_NumberAssertion.NonNegativeNumber,
			_NumberAssertion.NonZero,
			_NumberAssertion.BetweenRange(4.899, 5.001),
		)
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"nonNumber\"; specified number \"unknown\" is invalid. strconv.ParseFloat: parsing \"unknown\": invalid syntax"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
	{
		err := _NumberAssertion.Assert(nanNumber, "nanNumber",
			_NumberAssertion.NonNanNorInf,
			_NumberAssertion.NonNegativeNumber,
			_NumberAssertion.NonZero,
			_NumberAssertion.BetweenRange(4.899, 5.001),
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
		err := _NumberAssertion.Assert(zeroNumber, "zeroNumber",
			_NumberAssertion.NonNanNorInf,
			_NumberAssertion.NonNegativeNumber,
			_NumberAssertion.NonZero,
			_NumberAssertion.BetweenRange(4.899, 5.001),
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
		err := _NumberAssertion.Assert(negateNumber, "negateNumber",
			_NumberAssertion.NonNanNorInf,
			_NumberAssertion.NonNegativeNumber,
			_NumberAssertion.NonZero,
			_NumberAssertion.BetweenRange(4.899, 5.001),
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
		err := _NumberAssertion.Assert(four_nineNumber, "four_nineNumber",
			_NumberAssertion.NonNanNorInf,
			_NumberAssertion.NonNegativeNumber,
			_NumberAssertion.NonZero,
			_NumberAssertion.BetweenRange(4.899, 5.001),
		)
		if err != nil {
			t.Errorf("should not error")
		}
	}
}

func TestNumberAssertor(t *testing.T) {
	var (
		nanNumber Number = "nan"
	)

	{
		err := _NumberAssertion.Assertor(nanNumber, "nanNumber").
			Assert(
				_NumberAssertion.NonNanNorInf,
				_NumberAssertion.NonNegativeNumber,
				_NumberAssertion.NonZero,
				_NumberAssertion.BetweenRange(4.899, 5.001),
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

func TestNumberAssertion_IsNumber(t *testing.T) {
	{
		var arg Number = "0"
		err := _NumberAssertion.IsNumber(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg Number = "0.001"
		err := _NumberAssertion.IsNumber(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg Number = "-0.001"
		err := _NumberAssertion.IsNumber(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg Number = "-inf"
		err := _NumberAssertion.IsNumber(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg Number = "+inf"
		err := _NumberAssertion.IsNumber(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg Number = "nan"
		err := _NumberAssertion.IsNumber(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg Number = "unknown"
		err := _NumberAssertion.IsNumber(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"arg\"; specified number \"unknown\" is invalid. strconv.ParseInt: parsing \"unknown\": invalid syntax"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
}

func TestNumberAssertion_NonNegativeNumber(t *testing.T) {
	{
		var arg Number = "0"
		err := _NumberAssertion.NonNegativeNumber(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg Number = "0.001"
		err := _NumberAssertion.NonNegativeNumber(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg Number = "+inf"
		err := _NumberAssertion.NonNegativeNumber(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg Number = "inf"
		err := _NumberAssertion.NonNegativeNumber(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg Number = "-0.001"
		err := _NumberAssertion.NonNegativeNumber(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"arg\"; should be a non-negative number"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
	{
		var arg Number = "-inf"
		err := _NumberAssertion.NonNegativeNumber(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"arg\"; should be a non-negative number"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
}

func TestNumberAssertion_NonZero(t *testing.T) {
	{
		var arg Number = "-0.0001"
		err := _NumberAssertion.NonZero(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg Number = "0.001"
		err := _NumberAssertion.NonZero(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg Number = "unknown"
		err := _NumberAssertion.NonZero(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"arg\"; specified number \"unknown\" is invalid. strconv.ParseFloat: parsing \"unknown\": invalid syntax"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}

	{
		var arg Number = "0"
		err := _NumberAssertion.NonZero(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"arg\"; should not be zero"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
}

func TestNumberAssertion_NonNanNorInf(t *testing.T) {
	{
		var arg Number = "-0.0001"
		err := _NumberAssertion.NonNanNorInf(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg Number = "0.001"
		err := _NumberAssertion.NonNanNorInf(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg Number = "unknown"
		err := _NumberAssertion.NonNanNorInf(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"arg\"; specified number \"unknown\" is invalid. strconv.ParseFloat: parsing \"unknown\": invalid syntax"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
	{
		var arg Number = "-inf"
		err := _NumberAssertion.NonNanNorInf(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"arg\"; cannot be -inf, +inf or NaN"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
	{
		var arg Number = "+inf"
		err := _NumberAssertion.NonNanNorInf(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"arg\"; cannot be -inf, +inf or NaN"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
	{
		var arg Number = "nan"
		err := _NumberAssertion.NonNanNorInf(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"arg\"; cannot be -inf, +inf or NaN"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
}

func TestNumberAssertion_NotIn(t *testing.T) {
	var validate NumberValidator = _NumberAssertion.NotIn(3, 6)
	{
		var arg Number = "-1"
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg Number = "-1.003"
		err := validate(arg, "arg")
		if err != nil {
			t.Log(err)
			t.Errorf("should not error")
		}
	}
	{
		var arg Number = "6"
		err := validate(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"arg\"; specified number \"6\" is invalid"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
	{
		var arg Number = "3"
		err := validate(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"arg\"; specified number \"3\" is invalid"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
}

func TestNumberAssertion_Must(t *testing.T) {
	var validate NumberValidator = _NumberAssertion.Must(
		func(v Number) bool {
			argv, err := v.Int64()
			if err != nil {
				return false
			}
			return (argv & 0x01) == 0
		})
	{
		var arg Number = "0"
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg Number = "6"
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg Number = "-1"
		err := validate(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"arg\"; specified number \"-1\" is invalid"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
	{
		var arg Number = "-1.003"
		err := validate(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"arg\"; specified number \"-1.003\" is invalid"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
}

func TestNumberAssertion_MustInt(t *testing.T) {
	var validate NumberValidator = _NumberAssertion.MustInt(
		func(v int64) bool {
			return (v & 0x01) == 0
		})

	{
		var arg Number = "1"
		err := validate(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"arg\"; specified number \"1\" is invalid"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
	{
		var arg Number = "6"
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg Number = "3"
		err := validate(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"arg\"; specified number \"3\" is invalid"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
}

func TestNumberAssertion_MustFloat(t *testing.T) {
	var validate NumberValidator = _NumberAssertion.MustFloat(
		func(v float64) bool {
			return math.Round(v) == 5
		})

	{
		var arg Number = "-0.0001"
		err := validate(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"arg\"; specified number \"-0.0001\" is invalid"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
	{
		var arg Number = "4.9001"
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg Number = "5.0001"
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
}

func TestNumberAssertion_Less(t *testing.T) {
	var validate NumberValidator = _NumberAssertion.Less(5.0001)

	{
		var arg Number = "-0.0001"
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg Number = "5.0001"
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
		var arg Number = "5.0002"
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

func TestNumberAssertion_LessOrEqual(t *testing.T) {
	var validate NumberValidator = _NumberAssertion.LessOrEqual(5.0001)

	{
		var arg Number = "-0.0001"
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg Number = "5.0001"
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg Number = "5.0002"
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

func TestNumberAssertion_Greater(t *testing.T) {
	var validate NumberValidator = _NumberAssertion.Greater(5.0001)

	{
		var arg Number = "5.0002"
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg Number = "5.0001"
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
		var arg Number = "-0.0001"
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

func TestNumberAssertion_GreaterOrEqual(t *testing.T) {
	var validate NumberValidator = _NumberAssertion.GreaterOrEqual(5.0001)

	{
		var arg Number = "5.0002"
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg Number = "5.0001"
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg Number = "-0.0001"
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

func TestNumberAssertion_BetweenRange(t *testing.T) {
	var validate NumberValidator = _NumberAssertion.BetweenRange(4.8999, 5.0001)

	{
		var arg Number = "4.8999"
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg Number = "4.89991"
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg Number = "5.0001"
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg Number = "5.0000999999"
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg Number = "4.8998999999"
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
		var arg Number = "5.0001000001"
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
		var arg Number = "-0.0001"
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
