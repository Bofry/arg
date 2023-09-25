package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/Bofry/structproto"
	"github.com/Bofry/structproto/reflecting"
	"github.com/Bofry/structproto/valuebinder"
)

var (
	typeOfJsonRawMessage = reflect.TypeOf(json.RawMessage(nil))
	typeOfAny            = reflect.TypeOf([]interface{}{nil}).Elem()
)

var _ structproto.ValueBinder = new(JsonContentBinder)

type JsonContentBinder reflect.Value

func (binder JsonContentBinder) Bind(content interface{}) error {
	rv := reflect.Value(binder)
	return binder.bindJsonValue(rv, content)
}

func (binder JsonContentBinder) bindJsonValue(rv reflect.Value, content interface{}) error {
	rv = reflect.Indirect(reflecting.AssignZero(rv))

	// binding for the special types
	switch rv.Type() {
	case typeOfJsonRawMessage:
		// TODO create JsonRawMessageBinder !!!
		rv = reflect.Indirect(reflecting.AssignZero(rv))
		b, err := json.Marshal(content)
		if err != nil {
			return &valuebinder.ValueBindingError{
				Value: content,
				Kind:  rv.Type().Name(),
				Err:   err}
		}
		var raw = json.RawMessage(b)
		rv.Set(reflect.ValueOf(raw))
		return nil
	}

	var err error
	{
		switch rv.Kind() {
		case reflect.Array, reflect.Slice:
			var jsonArray []interface{}
			{
				switch v := content.(type) {
				case []interface{}:
					jsonArray = v
				default:
					result, err := binder.unmarshalContent(content)
					if err != nil {
						return &valuebinder.ValueBindingError{
							Value: content,
							Kind:  rv.Type().Name(),
							Err:   err}
					}
					jsonArray, _ = result.([]interface{})
				}
			}

			if jsonArray == nil {
				return &valuebinder.ValueBindingError{
					Value: content,
					Kind:  rv.Type().Name(),
					Err:   err}
			}
			return binder.bindJsonArray(rv, jsonArray)
		case reflect.Struct:
			var jsonObject map[string]interface{}
			{
				switch v := content.(type) {
				case map[string]interface{}:
					jsonObject = v
				default:
					result, err := binder.unmarshalContent(content)
					if err != nil {
						return &valuebinder.ValueBindingError{
							Value: content,
							Kind:  rv.Type().Name(),
							Err:   err}
					}
					jsonObject, _ = result.(map[string]interface{})
				}
			}

			if jsonObject == nil {
				return &valuebinder.ValueBindingError{
					Value: content,
					Kind:  rv.Type().Name(),
					Err:   err}
			}
			return binder.bindJsonObject(rv, jsonObject)
		}
	}
	if rv.IsZero() {
		// perform normal binding
		scalarValueBinder := valuebinder.ScalarBinder(rv)
		err = scalarValueBinder.Bind(content)
	}
	return err
}

func (binder JsonContentBinder) bindJsonArray(rv reflect.Value, content []interface{}) error {
	if len(content) > 0 {
		size := len(content)
		container := reflect.MakeSlice(rv.Type(), size, size)
		for i, elem := range content {
			err := binder.bindJsonValue(container.Index(i), elem)
			if err != nil {
				return err
			}
		}
		rv.Set(container)
	}
	return nil
}

func (binder JsonContentBinder) bindJsonObject(rv reflect.Value, content map[string]interface{}) error {
	prototype, err := structproto.Prototypify(rv,
		&structproto.StructProtoResolveOption{
			TagName: TagName_Json,
		})
	if err != nil {
		return err
	}
	return prototype.BindMap(content, BuildJsonContentBinder)
}

func (binder JsonContentBinder) bindJsonMap(rv reflect.Value, content map[string]interface{}) error {
	if rv.Type().Elem() == typeOfAny {
		if content != nil && len(content) > 0 {
			rv.Set(reflect.ValueOf(content))
		}
		return nil
	}

	return fmt.Errorf("unsupported type map[string]%s", rv.Type().Elem().String())
}

func (binder JsonContentBinder) unmarshalContent(content interface{}) (interface{}, error) {
	var (
		result interface{}
		err    error
	)

	switch buffer := content.(type) {
	case string:
		decoder := json.NewDecoder(strings.NewReader(buffer))
		decoder.UseNumber()
		err = decoder.Decode(&result)
	case []byte:
		decoder := json.NewDecoder(bytes.NewReader(buffer))
		decoder.UseNumber()
		err = decoder.Decode(&result)
	default:
		err = fmt.Errorf("cannot marshal content with type %T", content)
	}
	return result, err
}

func BuildJsonContentBinder(rv reflect.Value) structproto.ValueBinder {
	return JsonContentBinder(rv)
}
