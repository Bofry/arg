package content

import (
	"github.com/Bofry/arg/content/internal"
	"github.com/Bofry/structproto"
	"github.com/Bofry/structproto/valuebinder"
)

var _ MapContentProcessor = KeyValueMap

func KeyValueMap(source map[string]interface{}, target interface{}) error {
	prototype, err := structproto.Prototypify(target,
		&structproto.StructProtoResolveOption{
			TagName: internal.TagName_Key,
		})
	if err != nil {
		return err
	}

	return prototype.BindMap(source, valuebinder.BuildScalarBinder)
}
