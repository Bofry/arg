package arg

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestDecimalAssertion(t *testing.T) {
	var (
		zeroDecimal      Decimal = decimal.New(0, 0)   // 0
		negateDecimal    Decimal = decimal.New(-1, -3) // -0.001
		four_nineDecimal Decimal = decimal.New(49, -1) // 4.9
		five_sixDecimal  Decimal = decimal.New(56, -1) // 5.6
	)

	{
		err := _DecimalAssertion.Assert(zeroDecimal, "zeroDecimal",
			_DecimalAssertion.NonZero,
			_DecimalAssertion.BetweenRange(decimal.New(4899, -3) /* 4.899 */, decimal.New(5001, -3) /* 5.001 */),
		)
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"zeroDecimal\"; should not be zero"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
	{
		err := _DecimalAssertion.Assert(negateDecimal, "negateDecimal",
			_DecimalAssertion.NonZero,
			_DecimalAssertion.BetweenRange(decimal.New(4899, -3) /* 4.899 */, decimal.New(5001, -3) /* 5.001 */),
		)
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"negateDecimal\"; out of range"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
	{
		err := _DecimalAssertion.Assert(five_sixDecimal, "five_sixDecimal",
			_DecimalAssertion.NonZero,
			_DecimalAssertion.BetweenRange(decimal.New(4899, -3) /* 4.899 */, decimal.New(5001, -3) /* 5.001 */),
		)
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"five_sixDecimal\"; out of range"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
	{
		err := _DecimalAssertion.Assert(four_nineDecimal, "four_nineDecimal",
			_DecimalAssertion.NonZero,
			_DecimalAssertion.BetweenRange(decimal.New(4899, -3), decimal.New(5001, -3)),
		)
		if err != nil {
			t.Errorf("should not error")
		}
	}
}

func TestDecimalAssertor(t *testing.T) {
	var (
		zeroDecimal Decimal = decimal.New(0, 0)
	)

	{
		err := _DecimalAssertion.Assertor(zeroDecimal, "zeroDecimal").
			Assert(
				_DecimalAssertion.NonZero,
				_DecimalAssertion.BetweenRange(decimal.New(4899, -3) /* 4.899 */, decimal.New(5001, -3) /* 5.001 */),
			)
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"zeroDecimal\"; should not be zero"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
}

func TestDecimalAssertion_NonZero(t *testing.T) {
	{
		var arg Decimal = decimal.New(-1, -4) // -0.0001
		err := _DecimalAssertion.NonZero(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg Decimal = decimal.New(1, -3) // 0.001
		err := _DecimalAssertion.NonZero(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg Decimal = decimal.New(0, 0) // 0
		err := _DecimalAssertion.NonZero(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"arg\"; should not be zero"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
}

func TestDecimalAssertion_Must(t *testing.T) {
	var validate DecimalValidator = _DecimalAssertion.Must(
		func(v Decimal) bool {
			return v.Round(0).Equal(decimal.New(5, 0))
		})

	{
		var arg Decimal = decimal.New(-1, -4) // -0.0001
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
		var arg Decimal = decimal.New(49001, -4) // 4.9001
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg Decimal = decimal.New(50001, -4) // 5.0001
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
}

func TestDecimalAssertion_Less(t *testing.T) {
	var validate DecimalValidator = _DecimalAssertion.Less(decimal.New(50001, -4))

	{
		var arg Decimal = decimal.New(-1, -4) // -0.0001
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg Decimal = decimal.New(50001, -4) // 5.0001
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
		var arg Decimal = decimal.New(50002, -4) // 5.0002
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

func TestDecimalAssertion_LessOrEqual(t *testing.T) {
	var validate DecimalValidator = _DecimalAssertion.LessOrEqual(decimal.New(50001, -4))

	{
		var arg Decimal = decimal.New(-1, -4) // -0.0001
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg Decimal = decimal.New(50001, -4) // 5.0001
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg Decimal = decimal.New(50002, -4) // 5.0002
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

func TestDecimalAssertion_Greater(t *testing.T) {
	var validate DecimalValidator = _DecimalAssertion.Greater(decimal.New(50001, -4))

	{
		var arg Decimal = decimal.New(50002, -4) // 5.0002
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg Decimal = decimal.New(50001, -4) // 5.0001
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
		var arg Decimal = decimal.New(-1, -4) // -0.0001
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

func TestDecimalAssertion_GreaterOrEqual(t *testing.T) {
	var validate DecimalValidator = _DecimalAssertion.GreaterOrEqual(decimal.New(50001, -4))

	{
		var arg Decimal = decimal.New(50002, -4) // 5.0002
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg Decimal = decimal.New(50001, -4) // 5.0001
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg Decimal = decimal.New(-1, -4) // -0.0001
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

func TestDecimalAssertion_BetweenRange(t *testing.T) {
	var validate DecimalValidator = _DecimalAssertion.BetweenRange(decimal.New(48999, -4), decimal.New(50001, -4))

	{
		var arg Decimal = decimal.New(48999, -4) // 4.8999
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg Decimal = decimal.New(489991, -5) // 4.89991
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg Decimal = decimal.New(50001, -4) // 5.0001
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg Decimal = decimal.New(50000999999, -10) // 5.0000999999
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg Decimal = decimal.New(48998999999, -10) // 4.8998999999
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
		var arg Decimal = decimal.New(50001000001, -10) // 5.0001000001
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
		var arg Decimal = decimal.New(-1, -4) // -0.0001
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
