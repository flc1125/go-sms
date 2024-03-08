package sms

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type driver struct {
}

type driverRequestExtra struct {
	Code string
}

func (d *driverRequestExtra) IsRequestExtra() {}

type driverResponseExtra struct {
	ReplyCode string
}

func (d *driverResponseExtra) IsResponseExtra() {}

func (d *driver) Send(ctx context.Context, req *Request) (*Response, error) {
	extra, ok := req.Extra.(*driverRequestExtra)
	if !ok {
		return nil, fmt.Errorf("invalid extra")
	}

	return &Response{
		Status:  Success,
		Message: fmt.Sprintf("send sms to %s-%s and content: %s", req.IDDCode, req.Phone, req.Content),
		Extra: &driverResponseExtra{
			ReplyCode: fmt.Sprintf("reply code: %s", extra.Code),
		},
	}, nil
}

func TestClient(t *testing.T) {
	client := New(&driver{})
	resp, err := client.Send(context.Background(), &Request{
		IDDCode: "886",
		Phone:   "123456789",
		Content: "hello",
		Extra:   &driverRequestExtra{Code: "123"},
	})
	assert.Nil(t, err)
	assert.Equal(t, Success, resp.Status)
	assert.Equal(t, "send sms to 886-123456789 and content: hello", resp.Message)
	assert.Equal(t, "reply code: 123", resp.Extra.(*driverResponseExtra).ReplyCode)
}
