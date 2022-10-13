package arg

import "testing"

func TestString(t *testing.T) {
	var (
		emprtyString string = ""
	)

	err := String.Assert(emprtyString, "emprtyString",
		String.NonEmpty,
	)
	if err == nil {
		t.Errorf("should get error")
	}
}
