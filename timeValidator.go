package arg

import (
	"fmt"
	"time"

	"github.com/Bofry/arg/internal"
	"github.com/cstockton/go-conv"
)

var (
	timeLayouts []string = []string{
		time.Layout,
		time.RFC822,
		time.RFC822Z,
		time.RFC3339,
		time.RFC3339Nano,
		time.Stamp,
		time.StampMilli,
		time.StampMicro,
		time.StampNano,
		time.RFC850,
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		time.RFC1123,
		time.RFC1123Z,
	}
)

func (fn TimeValidator) AssertPtr(v *time.Time, name string) error {
	if v != nil {
		return fn(*v, name)
	}
	return nil
}

func (fn TimeValidator) AssertValue(v interface{}, name string) error {
	var ts time.Time

	switch v.(type) {
	case time.Time:
		{
			ts = v.(time.Time)
		}
	case string:
		{
			var err error

			str := v.(string)
			for _, layout := range timeLayouts {
				ts, err = time.Parse(layout, str)
				// found
				if err == nil {
					break
				}
			}
			if err != nil {
				err = nil
				ts, err = conv.Time(str)
			}
			if err != nil {
				return &InvalidArgumentError{
					Name:   name,
					Reason: fmt.Sprintf(internal.ERR_UNSUPPORTED_CAST_TIME, v),
				}
			}
		}
	}
	return fn(ts, name)
}
