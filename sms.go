package sms

import "context"

type Client struct {
	driver Driver
}

func New(driver Driver) *Client {
	return &Client{
		driver: driver,
	}
}

func (c *Client) Send(ctx context.Context, req *Request) (*Response, error) {
	return c.driver.Send(ctx, req)
}
