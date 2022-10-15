package arg

import (
	"math"
	"testing"
)

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

func TestFloat(t *testing.T) {
	var (
		nanFloat float64 = math.NaN()
	)

	{
		err := Float.Assert(nanFloat, "nanFloat",
			Float.NonNanNorInf,
			Float.NonNegativeNumber,
			Float.NonZero,
			Float.BetweenRange(4.899, 5.001),
		)
		if err == nil {
			t.Errorf("should get error")
		}
	}

	// use Assertor
	{
		err := Float.Assert(nanFloat, "nanFloat",
			Float.NonNanNorInf,
			Float.NonNegativeNumber,
			Float.NonZero,
			Float.BetweenRange(4.899, 5.001),
		)
		if err == nil {
			t.Errorf("should get error")
		}
	}
}
