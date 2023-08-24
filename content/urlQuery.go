package content

import (
	"net/url"

	"github.com/Bofry/arg/content/internal"
	"github.com/Bofry/structproto"
)

var (
	_ BinaryContentProcessor = UrlQuery
	_ TextContentProcessor   = UrlQueryString
)

func UrlQuery(source []byte, target interface{}) error {
	return UrlQueryString(string(source), target)
}

func UrlQueryString(source string, target interface{}) error {
	values, err := url.ParseQuery(source)
	if err != nil {
		return err
	}

	provider := internal.NewUrlQueryContentBinder(values)

	prototype, err := structproto.Prototypify(target,
		&structproto.StructProtoResolveOption{
			TagName: internal.TagName_Query,
		})
	if err != nil {
		return err
	}

	return prototype.Bind(provider)
}
