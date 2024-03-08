package sms

import "context"

type Driver interface {
	Send(ctx context.Context, req *Request) (*Response, error)
}
