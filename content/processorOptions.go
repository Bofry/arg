package content

var _ ProcessorOption = ProcessorOptionFunc(nil)

type ProcessorOptionFunc func(processor *Processor)

func (f ProcessorOptionFunc) apply(processor *Processor) {
	f(processor)
}

func WithErrorHandler(proc ErrorHandler) ProcessorOption {
	return ProcessorOptionFunc(func(processor *Processor) {
		processor.errorHandler = proc
	})
}
