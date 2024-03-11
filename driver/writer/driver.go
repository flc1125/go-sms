package writer

import (
	"context"
	"fmt"
	"io"

	"github.com/flc1125/go-sms"
)

type Driver struct {
	writer io.Writer
}

func NewDriver(writer io.Writer) *Driver {
	return &Driver{
		writer: writer,
	}
}

func (d *Driver) Send(_ context.Context, req *sms.Request) (*sms.Response, error) {
	if _, err := fmt.Fprintln(d.writer, "writer: send sms", req); err != nil {
		return nil, err
	}

	return &sms.Response{
		Status:  sms.Success,
		Message: "writer: send sms success",
	}, nil
}
