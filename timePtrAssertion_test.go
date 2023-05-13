package arg

import (
	"testing"
	"time"
)

func TestTimePtrAssertion(t *testing.T) {
	var NewTime = func(v time.Time) *time.Time {
		return &v
	}

	var (
		nilTime      *time.Time = nil
		zeroTime     *time.Time = NewTime(time.Time{})
		date1205Time *time.Time = NewTime(time.Date(2022, 12, 5, 10, 57, 34, 0, time.UTC))
		date1211Time *time.Time = NewTime(time.Date(2022, 12, 11, 10, 57, 34, 0, time.UTC))
	)

	{
		err := _TimePtrAssertion.Assert(nilTime, "nilTime",
			_TimePtrAssertion.NotNil,
			_TimePtrAssertion.IsUTC,
			_TimePtrAssertion.NonZero,
		)
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"nilTime\"; cannot be nil"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
	{
		err := _TimePtrAssertion.Assert(zeroTime, "zeroTime",
			_TimePtrAssertion.IsUTC,
			_TimePtrAssertion.NonZero,
		)
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"zeroTime\"; should not be zero"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
	{
		err := _TimePtrAssertion.Assert(date1205Time, "date1205Time",
			_TimePtrAssertion.IsUTC,
			_TimePtrAssertion.NonZero,
			_TimePtrAssertion.AfterOrEqual(
				time.Date(2022, 12, 11, 0, 0, 0, 0, time.UTC),
			),
		)
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"date1205Time\"; out of range"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
	{
		err := _TimePtrAssertion.Assert(nilTime, "nilTime",
			_TimePtrAssertion.IsUTC,
			_TimePtrAssertion.NonZero,
			_TimePtrAssertion.AfterOrEqual(
				time.Date(2022, 12, 11, 0, 0, 0, 0, time.UTC),
			),
		)
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		err := _TimePtrAssertion.Assert(date1211Time, "date1211Time",
			_TimePtrAssertion.IsUTC,
			_TimePtrAssertion.NonZero,
			_TimePtrAssertion.AfterOrEqual(
				time.Date(2022, 12, 11, 0, 0, 0, 0, time.UTC),
			),
		)
		if err != nil {
			t.Errorf("should not error")
		}
	}
}

func TestTimePtrAssertor(t *testing.T) {
	var (
		zeroTime *time.Time = &time.Time{}
	)

	{
		err := _TimePtrAssertion.Assertor(zeroTime, "zeroTime").
			Assert(
				_TimePtrAssertion.IsUTC,
				_TimePtrAssertion.NonZero,
			)
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"zeroTime\"; should not be zero"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
}

func TestTimePtrAssertion_NotNil(t *testing.T) {
	{
		var arg *time.Time = &time.Time{}
		err := _TimePtrAssertion.NotNil(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg *time.Time = nil
		err := _TimePtrAssertion.NotNil(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"arg\"; cannot be nil"
		if err.Error() != expectedErrorMsg {
			t.Errorf("expect: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
}
