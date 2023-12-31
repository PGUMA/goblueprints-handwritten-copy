package trace

import (
	"fmt"
	"io"
)

type Tracer interface {
	Trace(...interface{}) // any type, any number of, zero can
}

type tracer struct {
	out io.Writer
}

func (t *tracer) Trace(a ...interface{}) {
	t.out.Write([]byte(fmt.Sprint(a...)))
	t.out.Write([]byte("\n"))
}

func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

type nilTrascer struct{}

func (t *nilTrascer) Trace(a ...interface{}) {}

func Off() Tracer {
	return &nilTrascer{}
}
