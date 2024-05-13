package arg

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestDecimalPtrAssertion(t *testing.T) {
	var NewDecimal = func(v Decimal) *Decimal {
		return &v
	}

	var (
		nilDecimal       *Decimal = nil
		zeroDecimal      *Decimal = NewDecimal(decimal.New(0, 0))   // 0
		negateDecimal    *Decimal = NewDecimal(decimal.New(-1, -3)) // -0.001
		four_nineDecimal *Decimal = NewDecimal(decimal.New(49, -1)) // 4.9
		five_sixDecimal  *Decimal = NewDecimal(decimal.New(56, -1)) // 5.6
	)

	{
		err := _DecimalPtrAssertion.Assert(nilDecimal, "nilDecimal",
			_DecimalPtrAssertion.NotNil,
			_DecimalPtrAssertion.NonNegativeNumber,
			_DecimalPtrAssertion.NonZero,
			_DecimalPtrAssertion.BetweenRange(decimal.New(4899, -3) /* 4.899 */, decimal.New(5001, -3) /* 5.001 */),
		)
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"nilDecimal\"; cannot be nil"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
	{
		err := _DecimalPtrAssertion.Assert(zeroDecimal, "zeroDecimal",
			_DecimalPtrAssertion.NonNegativeNumber,
			_DecimalPtrAssertion.NonZero,
			_DecimalPtrAssertion.BetweenRange(decimal.New(4899, -3) /* 4.899 */, decimal.New(5001, -3) /* 5.001 */),
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
		err := _DecimalPtrAssertion.Assert(negateDecimal, "negateDecimal",
			_DecimalPtrAssertion.NonNegativeNumber,
			_DecimalPtrAssertion.NonZero,
			_DecimalPtrAssertion.BetweenRange(decimal.New(4899, -3) /* 4.899 */, decimal.New(5001, -3) /* 5.001 */),
		)
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"negateDecimal\"; should be a non-negative number"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
	{
		err := _DecimalPtrAssertion.Assert(five_sixDecimal, "five_sixDecimal",
			_DecimalPtrAssertion.NonNegativeNumber,
			_DecimalPtrAssertion.NonZero,
			_DecimalPtrAssertion.BetweenRange(decimal.New(4899, -3) /* 4.899 */, decimal.New(5001, -3) /* 5.001 */),
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
		err := _DecimalPtrAssertion.Assert(nilDecimal, "nilDecimal",
			_DecimalPtrAssertion.NonNegativeNumber,
			_DecimalPtrAssertion.NonZero,
			_DecimalPtrAssertion.BetweenRange(decimal.New(4899, -3) /* 4.899 */, decimal.New(5001, -3) /* 5.001 */),
		)
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		err := _DecimalPtrAssertion.Assert(four_nineDecimal, "four_nineDecimal",
			_DecimalPtrAssertion.NonNegativeNumber,
			_DecimalPtrAssertion.NonZero,
			_DecimalPtrAssertion.BetweenRange(decimal.New(4899, -3) /* 4.899 */, decimal.New(5001, -3) /* 5.001 */),
		)
		if err != nil {
			t.Errorf("should not error")
		}
	}
}

func TestDecimalPtrAssertor(t *testing.T) {
	var NewDecimal = func(v Decimal) *Decimal {
		return &v
	}

	var (
		zeroDecimal *Decimal = NewDecimal(decimal.New(0, 0))
	)

	{
		err := _DecimalPtrAssertion.Assertor(zeroDecimal, "zeroDecimal").
			Assert(
				_DecimalPtrAssertion.NonNegativeNumber,
				_DecimalPtrAssertion.NonZero,
				_DecimalPtrAssertion.BetweenRange(decimal.New(4899, -3) /* 4.899 */, decimal.New(5001, -3) /* 5.001 */),
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

func TestDecimalPtrAssertion_NotNil(t *testing.T) {
	var NewDecimal = func(v Decimal) *Decimal {
		return &v
	}

	{
		var arg *Decimal = NewDecimal(decimal.New(0, 0))
		err := _DecimalPtrAssertion.NotNil(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg *Decimal = nil
		err := _DecimalPtrAssertion.NotNil(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"arg\"; cannot be nil"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
}
