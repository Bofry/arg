package content_test

import (
	"reflect"
	"testing"

	"github.com/Bofry/arg/content"
	"go.openly.dev/pointy"
)

func TestKeyValueMap(t *testing.T) {
	type (
		Argv struct {
			String      string   `key:"*string"`
			StringPtr   *string  `key:"*string_ptr"`
			Int64       int64    `key:"int64"`
			True        bool     `key:"true"`
			False       bool     `key:"false"`
			StringSlice []string `key:"string_slice"`
		}
	)

	input := map[string]interface{}{
		"string":       "F0003452",
		"string_ptr":   "KNNS",
		"int64":        280123412341234123,
		"true":         true,
		"string_slice": []string{"T", "ER", "XVV"},
	}

	arg := Argv{}
	err := content.KeyValueMap(input, &arg)
	if err != nil {
		t.Fatal(err)
	}

	expected := Argv{
		String:      "F0003452",
		StringPtr:   pointy.String("KNNS"),
		Int64:       280123412341234123,
		True:        true,
		False:       false,
		StringSlice: []string{"T", "ER", "XVV"},
	}
	if !reflect.DeepEqual(expected, arg) {
		t.Errorf("assert:: expected '%#+v', got '%#+v'", expected, arg)
	}
}

func TestKeyValueMap_WithMapStringString(t *testing.T) {
	type (
		Argv struct {
			String    string  `key:"*string"`
			StringPtr *string `key:"*string_ptr"`
			Int64     int64   `key:"int64"`
			True      bool    `key:"true"`
			False     bool    `key:"false"`
		}
	)

	input := map[string]interface{}{
		"string":     "F0003452",
		"string_ptr": "KNNS",
		"int64":      "280123412341234123",
		"true":       "true",
	}

	arg := Argv{}
	err := content.KeyValueMap(input, &arg)
	if err != nil {
		t.Fatal(err)
	}

	expected := Argv{
		String:    "F0003452",
		StringPtr: pointy.String("KNNS"),
		Int64:     280123412341234123,
		True:      true,
		False:     false,
	}
	if !reflect.DeepEqual(expected, arg) {
		t.Errorf("assert:: expected '%#+v', got '%#+v'", expected, arg)
	}
}
