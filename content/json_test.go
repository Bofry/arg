package content_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/Bofry/arg/content"
	"go.openly.dev/pointy"
)

func TestJson(t *testing.T) {
	type (
		ArgvStruct struct {
			Operator string `json:"operator"`
		}

		Argv struct {
			String      string           `json:"*string"`
			StringPtr   *string          `json:"*string_ptr"`
			Int64       int64            `json:"int64"`
			True        bool             `json:"true"`
			False       bool             `json:"false"`
			StringSlice []string         `json:"string_slice"`
			JsonRaw     *json.RawMessage `json:"json_raw"`

			Struct *ArgvStruct `json:"struct"`
		}
	)

	input := []byte(`{
		"string": "F0003452",
		"string_ptr": "KNNS",
		"int64": 280123412341234123,
		"true": true,
		"string_slice": ["T","ER","XVV"],
		"struct": {
			"operator": "nami"
		},
		"extraInfo": {
			"alias": "Cat Burglar",
			"age"  : 18
		}
	}`)

	arg := Argv{}
	err := content.Json(input, &arg)
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
		Struct: &ArgvStruct{
			Operator: "nami",
		},
	}
	if !reflect.DeepEqual(expected, arg) {
		t.Errorf("assert:: expected '%#+v', got '%#+v'", expected, arg)
	}
}

func TestJson_WithMapStringInterface(t *testing.T) {
	type (
		Argv struct {
			String string                 `json:"*string"`
			Map    map[string]interface{} `json:"map"`
		}
	)

	input := []byte(`{
		"string": "F0003452",
		"map": {
			"alias": "Cat Burglar",
			"age"  : 18
		}
	}`)

	arg := Argv{}
	err := content.Json(input, &arg)
	if err != nil {
		t.Fatal(err)
	}

	expected := Argv{
		String: "F0003452",
		Map: map[string]interface{}{
			"age":   json.Number("18"),
			"alias": "Cat Burglar",
		},
	}
	if !reflect.DeepEqual(expected, arg) {
		t.Errorf("assert:: expected '%#+v', got '%#+v'", expected, arg)
	}
}

func TestJson_WithMapStringString(t *testing.T) {
	type (
		Argv struct {
			String string            `json:"*string"`
			Map    map[string]string `json:"map"`
		}
	)

	input := []byte(`{
		"string": "F0003452",
		"map": {
			"alias": "Cat Burglar",
			"age"  : 18
		}
	}`)

	arg := Argv{}
	err := content.Json(input, &arg)
	if err != nil {
		t.Fatal(err)
	}

	expected := Argv{
		String: "F0003452",
		Map: map[string]string{
			"age":   "18",
			"alias": "Cat Burglar",
		},
	}
	if !reflect.DeepEqual(expected, arg) {
		t.Errorf("assert:: expected '%#+v', got '%#+v'", expected, arg)
	}
}

func TestJson_WithJsonRaw(t *testing.T) {
	type (
		Argv struct {
			String      string           `json:"*string"`
			StringPtr   *string          `json:"*string_ptr"`
			Int64       int64            `json:"int64"`
			True        bool             `json:"true"`
			False       bool             `json:"false"`
			StringSlice []string         `json:"string_slice"`
			JsonRaw     *json.RawMessage `json:"json_raw"`
		}
	)

	input := []byte(`{
		"string": "F0003452",
		"string_ptr": "KNNS",
		"int64": 280123412341234123,
		"true": true,
		"string_slice": ["T","ER","XVV"],
		"json_raw": { "key": "value" }
	}`)

	arg := Argv{}
	err := content.Json(input, &arg)
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
		JsonRaw:     pointy.Pointer[json.RawMessage](json.RawMessage([]byte(`{"key":"value"}`))),
	}
	if !reflect.DeepEqual(expected, arg) {
		t.Errorf("assert:: expected '%#+v', got '%#+v'", expected, arg)
	}
}
