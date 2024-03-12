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

func New(writer io.Writer) *Driver {
	return &Driver{
		writer: writer,
	}
}

func (d *Driver) Send(_ context.Context, req *sms.Request) (*sms.Response, error) {
	if _, err := fmt.Fprintf(d.writer,
		"writer: send sms, the phone: %s, the content: %s\n", req.Phone, req.Content); err != nil {
		return nil, err
	}

	return &sms.Response{
		Status:  sms.Success,
		Message: "writer: send sms success",
	}, nil
}
