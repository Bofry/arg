package internal

import (
	"net/url"
	"reflect"

	"github.com/Bofry/structproto"
	"github.com/Bofry/structproto/valuebinder"
)

var _ structproto.StructBinder = new(UrlQueryContentBinder)

type UrlQueryContentBinder struct {
	values url.Values
}

func NewUrlQueryContentBinder(values url.Values) *UrlQueryContentBinder {
	instance := &UrlQueryContentBinder{
		values: values,
	}
	return instance
}

func (binder *UrlQueryContentBinder) Init(context *structproto.StructProtoContext) error {
	return nil
}

func (binder *UrlQueryContentBinder) Bind(field structproto.FieldInfo, rv reflect.Value) error {
	if v, ok := binder.values[field.Name()]; ok {
		switch rv.Kind() {
		case reflect.Bool:
			if len(v[0]) == 0 {
				v[0] = True
			}
		}
		return valuebinder.BuildStringBinder(rv).Bind(v[0])
	}
	return nil
}

func (binder *UrlQueryContentBinder) Deinit(context *structproto.StructProtoContext) error {
	return context.CheckIfMissingRequiredFields(func() <-chan string {
		c := make(chan string, 1)
		go func() {
			for k, _ := range binder.values {
				c <- k
			}
			close(c)
		}()
		return c
	})
}
