package sms

import "context"

type RequestExtra interface {
	IsRequestExtra()
}

type Request struct {
	IDDCode string       `json:"idd_code"`
	Phone   string       `json:"phone"`
	Content string       `json:"content"`
	Extra   RequestExtra `json:"extra"`
}

type ResponseExtra interface {
	IsResponseExtra()
}

type Response struct {
	Status  Status        `json:"status"`
	Message string        `json:"message"`
	Extra   ResponseExtra `json:"extra"`
}

type Driver interface {
	Send(ctx context.Context, req *Request) (*Response, error)
}
