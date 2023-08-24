package content_test

import (
	"reflect"
	"testing"

	"github.com/Bofry/arg/content"
	"go.openly.dev/pointy"
)

func TestUrlQueryString(t *testing.T) {
	type (
		Argv struct {
			ID       int64    `query:"*id"`
			Code     string   `query:"*code"`
			Type     *string  `query:"type"`
			Enabled  bool     `query:"enabled"`
			Disabled bool     `query:"disabled"`
			Tags     []string `query:"tags"`
		}
	)

	input := "id=280123412341234123&code=F0003452&type=KNNS&enabled&tags=T,ER,XVV"

	arg := Argv{}
	err := content.UrlQueryString(input, &arg)
	if err != nil {
		t.Fatal(err)
	}

	expected := Argv{
		ID:       280123412341234123,
		Code:     "F0003452",
		Type:     pointy.String("KNNS"),
		Enabled:  true,
		Disabled: false,
		Tags:     []string{"T", "ER", "XVV"},
	}
	if !reflect.DeepEqual(expected, arg) {
		t.Errorf("assert:: expected '%#+v', got '%#+v'", expected, arg)
	}
}
