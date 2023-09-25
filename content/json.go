package content

import (
	"encoding/json"

	"github.com/Bofry/arg/content/internal"
)

var _ BinaryContentProcessor = Json

func Json(source []byte, target interface{}) error {
	switch target.(type) {
	case json.Unmarshaler:
		return json.Unmarshal(source, target)
	}

	rv, err := internal.Indirect(target)
	if err != nil {
		return err
	}

	binder := internal.BuildJsonContentBinder(rv)
	err = binder.Bind(source)
	if err != nil {
		return err
	}
	return nil
}
