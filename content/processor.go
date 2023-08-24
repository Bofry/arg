package content

type Processor struct {
	target interface{}

	hasError     bool
	errorHandler ErrorHandler
}

func NewProcessor(target interface{}, opts ...ProcessorOption) *Processor {
	p := &Processor{
		target: target,
	}

	for _, opt := range opts {
		opt.apply(p)
	}

	return p
}

func (p *Processor) ProcessBytes(source []byte, proc BinaryContentProcessor) *Processor {
	if p.hasError {
		return p
	}

	var err error
	defer func() {
		if err != nil {
			p.throwError(err)
		}
	}()

	err = proc(source, p.target)
	return p
}

func (p *Processor) ProcessString(source string, proc TextContentProcessor) *Processor {
	if p.hasError {
		return p
	}

	var err error
	defer func() {
		if err != nil {
			p.throwError(err)
		}
	}()

	err = proc(source, p.target)
	return p
}

func (p *Processor) ProcessMap(source map[string]interface{}, proc MapContentProcessor) *Processor {
	if p.hasError {
		return p
	}

	var err error
	defer func() {
		if err != nil {
			p.throwError(err)
		}
	}()

	err = proc(source, p.target)
	return p
}

func (p *Processor) ProcessContent(source interface{}, proc ContentProcessor) *Processor {
	if p.hasError {
		return p
	}

	var err error
	defer func() {
		if err != nil {
			p.throwError(err)
		}
	}()

	err = proc(source, p.target)
	return p
}

func (p *Processor) Validate() {
	if p.hasError {
		return
	}

	var err error
	defer func() {
		if err != nil {
			p.throwError(err)
		}
	}()

	if v, ok := p.target.(Validatable); ok {
		err = v.Validate()
	}
}

func (p *Processor) throwError(err error) {
	// set the hasError flag
	p.hasError = true

	errHandler := p.errorHandler
	if errHandler == nil {
		errHandler = currentErrorHandler()
	}
	if errHandler != nil {
		errHandler(err)
	}
}
