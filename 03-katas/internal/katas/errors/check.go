package errors

import "errors"

var (
	ErrUninitiliazed   = errors.New("processor is not initialized")
	ErrPortAlreadyUsed = errors.New("port is already used")
)

// Processor is a dummy component.
type Processor struct{}

func NewProcessor() *Processor {
	return &Processor{}
}

func (p *Processor) Validate() error {
	return ErrUninitiliazed
}

func (p *Processor) Initialize() {
}

func (p *Processor) Run() error {
	return ErrUninitiliazed
}

func (p *Processor) RunOnRandomPort() error {
	return nil
}
