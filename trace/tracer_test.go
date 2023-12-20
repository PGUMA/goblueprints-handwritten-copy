package trace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Error("return value is nil.")
	} else {
		tracer.Trace("Hello Golang")
		if buf.String() != "Hello Golang\n" {
			t.Errorf("'%s' is incorrect message", buf.String())
		}
	}
}

func TestOff(t *testing.T) {
	var silentTracer Trascer = Off()
	silentTracer.Trace("sample")
}
