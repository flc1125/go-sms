package writer

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/flc1125/go-sms"
)

type writer struct {
	t *testing.T
}

func newWriter(t *testing.T) *writer {
	return &writer{
		t: t,
	}
}

func (w *writer) Write(p []byte) (n int, err error) {
	assert.Contains(w.t, string(p), "writer: send sms")
	assert.Contains(w.t, string(p), "1234567890")
	assert.Contains(w.t, string(p), "test")

	return len(p), nil
}

func TestDriver(t *testing.T) {
	w := newWriter(t)

	d := New(w)
	resp, err := d.Send(context.Background(), &sms.Request{
		Phone:   "1234567890",
		Content: "test",
	})
	assert.NoError(t, err)
	assert.Equal(t, sms.Success, resp.Status)
	assert.Equal(t, "writer: send sms success", resp.Message)
}
