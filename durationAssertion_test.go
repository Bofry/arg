package arg

import (
	"testing"
	"time"
)

func TestDurationAssertion(t *testing.T) {
	var (
		zeroDuration      time.Duration = 0
		negateOneDuration time.Duration = -1
	)

	{
		err := _DurationAssertion.Assert(zeroDuration, "zeroDuration",
			_DurationAssertion.NonNegative,
			_DurationAssertion.NonZero,
		)
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"zeroDuration\"; should not be zero"
		if err.Error() != expectedErrorMsg {
			t.Errorf("except: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
	{
		err := _DurationAssertion.Assert(negateOneDuration, "negateOneDuration",
			_DurationAssertion.NonNegative,
			_DurationAssertion.NonZero,
		)
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"negateOneDuration\"; should be a non-negative time.Duation"
		if err.Error() != expectedErrorMsg {
			t.Errorf("except: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
	{
		err := _DurationAssertion.Assert(negateOneDuration, "negateOneDuration",
			_DurationAssertion.NonZero,
		)
		if err != nil {
			t.Errorf("should not error")
		}
	}
}

func TestDurationAssertor(t *testing.T) {
	var (
		zeroDuration time.Duration = 0
	)

	{
		err := _DurationAssertion.Assertor(zeroDuration, "zeroDuration").
			Assert(
				_DurationAssertion.NonNegative,
				_DurationAssertion.NonZero,
			)
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"zeroDuration\"; should not be zero"
		if err.Error() != expectedErrorMsg {
			t.Errorf("except: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
}

func TestDurationAssertion_NonNegative(t *testing.T) {
	{
		var arg time.Duration = 0
		err := _DurationAssertion.NonNegative(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg time.Duration = 1
		err := _DurationAssertion.NonNegative(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg time.Duration = -1
		err := _DurationAssertion.NonNegative(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"arg\"; should be a non-negative time.Duation"
		if err.Error() != expectedErrorMsg {
			t.Errorf("except: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
}

func TestDurationAssertion_NonZero(t *testing.T) {
	{
		var arg time.Duration = -1
		err := _DurationAssertion.NonZero(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg time.Duration = 1
		err := _DurationAssertion.NonZero(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg time.Duration = 0
		err := _DurationAssertion.NonZero(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"arg\"; should not be zero"
		if err.Error() != expectedErrorMsg {
			t.Errorf("except: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
}

func TestDurationAssertion_Must(t *testing.T) {
	var validate DurationValidator = _DurationAssertion.Must(
		func(v time.Duration) bool {
			return v.Truncate(5*time.Minute) == v
		})

	{
		var arg, _ = time.ParseDuration("1h5m30s")
		err := validate(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"arg\"; specified time.Duration 1h5m30s is invalid"
		if err.Error() != expectedErrorMsg {
			t.Errorf("except: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
	{
		var arg, _ = time.ParseDuration("1h5m00s")
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg, _ = time.ParseDuration("1h15m30.918273645s")
		err := validate(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"arg\"; specified time.Duration 1h15m30.918273645s is invalid"
		if err.Error() != expectedErrorMsg {
			t.Errorf("except: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
}

func TestDurationAssertion_LessOrEqualByString(t *testing.T) {
	var validate DurationValidator = _DurationAssertion.LessOrEqualByString("2h")

	{
		var arg, _ = time.ParseDuration("1h59m59.999s")
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg, _ = time.ParseDuration("-3h30m")
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg, _ = time.ParseDuration("2h30m")
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

func TestDurationAssertion_GreaterOrEqualByString(t *testing.T) {
	var validate DurationValidator = _DurationAssertion.GreaterOrEqualByString("30m")

	{
		var arg, _ = time.ParseDuration("2h30m")
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg, _ = time.ParseDuration("30m00.999s")
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg, _ = time.ParseDuration("-2h30m")
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

func TestDurationAssertion_BetweenRange(t *testing.T) {
	var validate DurationValidator = _DurationAssertion.BetweenRangeByString("-30m", "2h")

	{
		var arg, _ = time.ParseDuration("-30m")
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg, _ = time.ParseDuration("2h")
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg, _ = time.ParseDuration("1h59m59s")
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg, _ = time.ParseDuration("2h30m")
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
		var arg, _ = time.ParseDuration("-30m00.999s")
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
