package arg

import (
	"testing"
	"time"
)

func TestTimeAssertion(t *testing.T) {
	var (
		zeroTime     time.Time = time.Time{}
		date1205Time time.Time = time.Date(2022, 12, 5, 10, 57, 34, 0, time.UTC)
		date1211Time time.Time = time.Date(2022, 12, 11, 10, 57, 34, 0, time.UTC)
	)

	{
		err := _TimeAssertion.Assert(zeroTime, "zeroTime",
			_TimeAssertion.IsUTC,
			_TimeAssertion.NonZero,
		)
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"zeroTime\"; should not be zero"
		if err.Error() != expectedErrorMsg {
			t.Errorf("except: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
	{
		err := _TimeAssertion.Assert(date1205Time, "date1205Time",
			_TimeAssertion.IsUTC,
			_TimeAssertion.NonZero,
			_TimeAssertion.AfterOrEqual(
				time.Date(2022, 12, 11, 0, 0, 0, 0, time.UTC),
			),
		)
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"date1205Time\"; out of range"
		if err.Error() != expectedErrorMsg {
			t.Errorf("except: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
	{
		err := _TimeAssertion.Assert(date1211Time, "date1211Time",
			_TimeAssertion.IsUTC,
			_TimeAssertion.NonZero,
			_TimeAssertion.AfterOrEqual(
				time.Date(2022, 12, 11, 0, 0, 0, 0, time.UTC),
			),
		)
		if err != nil {
			t.Errorf("should not error")
		}
	}
}

func TestTimeAssertor(t *testing.T) {
	var (
		zeroTime time.Time = time.Time{}
	)

	{
		err := _TimeAssertion.Assertor(zeroTime, "zeroTime").
			Assert(
				_TimeAssertion.IsUTC,
				_TimeAssertion.NonZero,
			)
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"zeroTime\"; should not be zero"
		if err.Error() != expectedErrorMsg {
			t.Errorf("except: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
}

func TestTimeAssertor_IsUTC(t *testing.T) {
	{
		var arg time.Time = time.Date(2022, 12, 11, 10, 57, 34, 0, time.UTC)
		err := _TimeAssertion.IsUTC(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg time.Time = time.Date(2022, 12, 11, 18, 57, 34, 0, &time.Location{})
		err := _TimeAssertion.IsUTC(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"arg\"; should be a UTC time"
		if err.Error() != expectedErrorMsg {
			t.Errorf("except: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
}

func TestTimeAssertor_NonZero(t *testing.T) {
	{
		var arg time.Time = time.Date(2022, 12, 11, 10, 57, 34, 0, time.UTC)
		err := _TimeAssertion.NonZero(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg time.Time
		err := _TimeAssertion.NonZero(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"arg\"; should not be zero"
		if err.Error() != expectedErrorMsg {
			t.Errorf("except: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
}

func TestTimeAssertor_Must(t *testing.T) {
	var validate TimeValidator = _TimeAssertion.Must(
		func(v time.Time) bool {
			return v.Day() == 5
		})

	{
		var arg time.Time = time.Date(2022, 12, 5, 10, 57, 34, 0, time.UTC)
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg time.Time = time.Date(2022, 12, 11, 10, 57, 34, 0, time.UTC)
		err := validate(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"arg\"; specified time 2022-12-11 10:57:34 +0000 UTC is invalid"
		if err.Error() != expectedErrorMsg {
			t.Errorf("except: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
}

func TestTimeAssertor_InWeekday(t *testing.T) {
	var validate TimeValidator = _TimeAssertion.InWeekday(
		time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday,
	)

	{
		var arg time.Time = time.Date(2022, 12, 5, 10, 57, 34, 0, time.UTC)
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg time.Time = time.Date(2022, 12, 11, 10, 57, 34, 0, time.UTC)
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

func TestTimeAssertor_NotInWeekday(t *testing.T) {
	var validate TimeValidator = _TimeAssertion.NotInWeekday(
		time.Saturday, time.Sunday,
	)

	{
		var arg time.Time = time.Date(2022, 12, 5, 10, 57, 34, 0, time.UTC)
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg time.Time = time.Date(2022, 12, 11, 10, 57, 34, 0, time.UTC)
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

func TestTimeAssertor_AfterOrEqual(t *testing.T) {
	var validate TimeValidator = _TimeAssertion.AfterOrEqual(
		time.Date(2022, 12, 11, 0, 0, 0, 0, time.UTC),
	)

	{
		var arg time.Time = time.Date(2022, 12, 11, 10, 57, 34, 0, time.UTC)
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg time.Time = time.Date(2022, 12, 5, 10, 57, 34, 0, time.UTC)
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

func TestTimeAssertor_BeforeOrEqual(t *testing.T) {
	var validate TimeValidator = _TimeAssertion.BeforeOrEqual(
		time.Date(2022, 12, 11, 0, 0, 0, 0, time.UTC),
	)

	{
		var arg time.Time = time.Date(2022, 12, 5, 10, 57, 34, 0, time.UTC)
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg time.Time = time.Date(2022, 12, 11, 10, 57, 34, 0, time.UTC)
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

func TestTimeAssertor_BetweenRange(t *testing.T) {
	var validate TimeValidator = _TimeAssertion.BetweenRange(
		time.Date(2022, 12, 5, 0, 0, 0, 0, time.UTC),
		time.Date(2022, 12, 11, 0, 0, 0, 0, time.UTC),
	)

	{
		var arg time.Time = time.Date(2022, 12, 5, 10, 57, 34, 0, time.UTC)
		err := validate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg time.Time = time.Date(2022, 12, 11, 10, 57, 34, 0, time.UTC)
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
