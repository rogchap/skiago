package skia

import (
	"io"
	"testing"
)

type reader struct {
	called bool
}

func (r *reader) Read(p []byte) (n int, err error) {
	defer func() {
		r.called = true
	}()

	p[0] = 0x1
	return 1, io.EOF
}

func TestGoStream(t *testing.T) {
	t.Parallel()

	r := reader{}
	s := newGoStream(&r)
	s.readTest()

	if !r.called {
		t.Fatal("expected reader to be read")
	}
}
