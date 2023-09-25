package content

import (
	"encoding/json"
	"log"
	"reflect"
	"sync"
	"sync/atomic"
)

const (
	_LOGGER_PREFIX string = "[httparg] "
)

var (
	logger *log.Logger = log.New(log.Default().Writer(), _LOGGER_PREFIX, log.Default().Flags()|log.Lmsgprefix)

	globalErrorHandler = defaultErrorHandler()

	panicErrorHandler = func(err error) { panic(err) }

	typeOfUnmarshaler = reflect.TypeOf((*json.Unmarshaler)(nil)).Elem()
)

type (
	MapContentProcessor    func(source map[string]interface{}, target interface{}) error
	TextContentProcessor   func(source string, target interface{}) error
	BinaryContentProcessor func(source []byte, target interface{}) error
	ContentProcessor       func(source interface{}, target interface{}) error

	ErrorHandler func(err error)

	Validatable interface {
		Validate() error
	}

	errorHandlerHolder struct {
		once         sync.Once
		ErrorHandler ErrorHandler
	}

	ProcessorOption interface {
		apply(*Processor)
	}
)

func SetErrorHandler(proc ErrorHandler) {
	current := currentErrorHandler()
	if current == nil {
		holder := globalErrorHandler.Load().(*errorHandlerHolder)
		holder.once.Do(func() {
			globalErrorHandler.Store(errorHandlerHolder{
				ErrorHandler: proc,
			})
		})
	}
}

func defaultErrorHandler() *atomic.Value {
	v := &atomic.Value{}
	v.Store(&errorHandlerHolder{})
	return v
}

func currentErrorHandler() ErrorHandler {
	v := globalErrorHandler.Load().(*errorHandlerHolder).ErrorHandler
	if v == nil {
		return panicErrorHandler
	}
	return v
}
