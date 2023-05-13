package arg

import (
	"testing"
	"time"
)

func TestDurationPtrAssertion(t *testing.T) {
	var NewDuration = func(v time.Duration) *time.Duration {
		return &v
	}

	var (
		zeroDuration      *time.Duration = NewDuration(0)
		negateOneDuration *time.Duration = NewDuration(-1)
	)

	{
		err := _DurationPtrAssertion.Assert(zeroDuration, "zeroDuration",
			_DurationPtrAssertion.NonNegative,
			_DurationPtrAssertion.NonZero,
		)
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"zeroDuration\"; should not be zero"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
	{
		err := _DurationPtrAssertion.Assert(negateOneDuration, "negateOneDuration",
			_DurationPtrAssertion.NonNegative,
			_DurationPtrAssertion.NonZero,
		)
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"negateOneDuration\"; should be a non-negative time.Duation"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
	{
		err := _DurationPtrAssertion.Assert(negateOneDuration, "negateOneDuration",
			_DurationPtrAssertion.NonZero,
		)
		if err != nil {
			t.Errorf("should not error")
		}
	}
}

func TestDurationPtrAssertor(t *testing.T) {
	var NewDuration = func(v time.Duration) *time.Duration {
		return &v
	}

	var (
		zeroDuration *time.Duration = NewDuration(0)
	)

	{
		err := _DurationPtrAssertion.Assertor(zeroDuration, "zeroDuration").
			Assert(
				_DurationPtrAssertion.NonNegative,
				_DurationPtrAssertion.NonZero,
			)
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"zeroDuration\"; should not be zero"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
}

func TestDurationPtrAssertion_NotNil(t *testing.T) {
	var NewDuration = func(v time.Duration) *time.Duration {
		return &v
	}

	{
		var arg *time.Duration = NewDuration(0)
		err := _DurationPtrAssertion.NotNil(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg *time.Duration = nil
		err := _DurationPtrAssertion.NotNil(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"arg\"; cannot be nil"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
}
