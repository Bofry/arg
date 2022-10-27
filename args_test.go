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
		err := Strings.Assert(emprtyString, "emprtyString",
			Strings.NonEmpty,
			Strings.In("foo", "bar"),
		)
		if err == nil {
			t.Errorf("should get error")
		}
	}
	{
		err := Strings.Assert(fooString, "fooString",
			Strings.NonEmpty,
			Strings.In("foo", "bar"),
		)
		if err != nil {
			t.Errorf("should no error")
		}
	}

	// use Assertor
	{
		err := Strings.Assertor(emprtyString, "emprtyString").
			Assert(
				Strings.NonEmpty,
				Strings.In("foo", "bar"),
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
		err := Floats.Assert(nanFloat, "nanFloat",
			Floats.NonNanNorInf,
			Floats.NonNegativeNumber,
			Floats.NonZero,
			Floats.BetweenRange(4.899, 5.001),
		)
		if err == nil {
			t.Errorf("should get error")
		}
	}

	// use Assertor
	{
		err := Floats.Assert(nanFloat, "nanFloat",
			Floats.NonNanNorInf,
			Floats.NonNegativeNumber,
			Floats.NonZero,
			Floats.BetweenRange(4.899, 5.001),
		)
		if err == nil {
			t.Errorf("should get error")
		}
	}
}
