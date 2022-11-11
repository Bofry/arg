package arg

import "testing"

func TestSliceAssertion(t *testing.T) {
	var (
		emptyStringSlice []string   = []string{}
		emptyIntSlice    []int      = []int{}
		emptyStructSlice []struct{} = []struct{}{}
	)

	{
		err := _ValueAssertion.Assert(emptyStringSlice, "emptyStringSlice",
			_SliceAssertion.NonEmpty,
		)
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"emptyStringSlice\"; cannot be an empty array"
		if err.Error() != expectedErrorMsg {
			t.Errorf("except: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
	{
		err := _ValueAssertion.Assert(emptyIntSlice, "emptyIntSlice",
			_SliceAssertion.NonEmpty,
		)
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"emptyIntSlice\"; cannot be an empty array"
		if err.Error() != expectedErrorMsg {
			t.Errorf("except: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
	{
		err := _ValueAssertion.Assert(emptyStructSlice, "emptyStructSlice",
			_SliceAssertion.NonEmpty,
		)
		if err == nil {
			t.Errorf("should get error")
		}
		expectedErrorMsg := "invalid argument \"emptyStructSlice\"; cannot be an empty array"
		if err.Error() != expectedErrorMsg {
			t.Errorf("except: %v\ngot: %v", expectedErrorMsg, err.Error())
		}
	}
}
