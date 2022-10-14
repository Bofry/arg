package arg

import "testing"

func TestString(t *testing.T) {
	var (
		emprtyString string = ""
		fooString    string = "foo"
	)
	{
		err := String.Assert(emprtyString, "emprtyString",
			String.NonEmpty,
			String.In("foo", "bar"),
		)
		if err == nil {
			t.Errorf("should get error")
		}
	}
	{
		err := String.Assert(fooString, "fooString",
			String.NonEmpty,
			String.In("foo", "bar"),
		)
		if err != nil {
			t.Errorf("should no error")
		}
	}

	// use Assertor
	{
		err := String.Assertor(emprtyString, "emprtyString").
			Assert(
				String.NonEmpty,
				String.In("foo", "bar"),
			)
		if err == nil {
			t.Errorf("should get error")
		}
	}
}
