package mitake

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/text/encoding/traditionalchinese"

	"github.com/flc1125/go-sms"
)

type Request struct {
	Dstaddr  string `json:"dstaddr"`
	Smbody   string `json:"smbody"`
	Type     string `json:"type"`
	Encoding string `json:"encoding"`
}

func (r *Request) IsRequestExtra() {}

type Response struct {
	MsgID        string `json:"msgid"`
	StatusCode   string `json:"statuscode"`
	AccountPoint string `json:"AccountPoint"`
	Error        string `json:"Error"`
}

func (r *Response) IsResponseExtra() {}

type Config struct {
	Addr     string
	Username string
	Password string
}

type Driver struct {
	config *Config

	httpClient *http.Client
}

type Option func(*Driver)

func WithHTTPClient(httpClient *http.Client) Option {
	return func(d *Driver) {
		d.httpClient = httpClient
	}
}

func New(config *Config, opts ...Option) *Driver {
	d := &Driver{
		config: config,
	}

	for _, o := range opts {
		o(d)
	}

	if d.httpClient == nil {
		d.httpClient = http.DefaultClient
	}

	return d
}

func (d *Driver) Send(ctx context.Context, req *sms.Request) (*sms.Response, error) {
	response, err := d.send(ctx, d.parseRequest(req))
	if err != nil {
		return &sms.Response{
			Status:  sms.Failed,
			Message: err.Error(),
			Extra:   response,
		}, err
	}

	return &sms.Response{
		Status:  sms.Success,
		Message: "success",
		Extra:   response,
	}, nil
}

func (d *Driver) parseRequest(req *sms.Request) *Request {
	request, ok := req.Extra.(*Request)
	if !ok {
		request = &Request{
			Dstaddr: req.Phone,
			Smbody:  req.Content,
		}
	}

	if request.Dstaddr == "" && req.Phone != "" {
		request.Dstaddr = req.Phone
	}

	if request.Smbody == "" && req.Content != "" {
		request.Smbody = req.Content
	}

	if request.Type == "" {
		request.Type = "now"
	}

	if request.Encoding == "" {
		request.Encoding = "big5"
	}

	return request
}

func (d *Driver) newRequest(ctx context.Context, req *Request) (*http.Request, error) {
	var (
		smbody = req.Smbody
		err    error
	)
	if req.Encoding == "big5" {
		if smbody, err = traditionalchinese.Big5.NewEncoder().String(req.Smbody); err != nil {
			return nil, err
		}
	}

	values := url.Values{}
	values.Set("username", d.config.Username)
	values.Set("password", d.config.Password)
	values.Set("dstaddr", req.Dstaddr)
	values.Set("smbody", smbody)
	values.Set("type", req.Type)
	values.Set("encoding", req.Encoding)

	return http.NewRequestWithContext(ctx, http.MethodGet, d.config.Addr+"?"+values.Encode(), nil)
}

func (d *Driver) doRequest(request *http.Request) (*Response, error) {
	resp, err := d.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("mitake: invalid status code: %d", resp.StatusCode)
	}

	return d.decodeResponse(resp)
}

func (d *Driver) decodeResponse(resp *http.Response) (*Response, error) {
	var buffer bytes.Buffer
	if _, err := buffer.ReadFrom(resp.Body); err != nil {
		return nil, err
	}

	var response Response
	for _, str := range strings.Split(buffer.String(), "\n") {
		if str == "" {
			continue
		}
		kv := strings.Split(str, "=")
		if len(kv) != 2 { // nolint:gomnd
			continue
		}

		switch kv[0] {
		case "msgid":
			response.MsgID = kv[1]
		case "statuscode":
			response.StatusCode = kv[1]
		case "AccountPoint":
			response.AccountPoint = kv[1]
		case "Error":
			// big5 to utf-8
			if kv[1] != "" {
				if s, err := traditionalchinese.Big5.NewDecoder().String(kv[1]); err == nil {
					kv[1] = s
				}
			}
			response.Error = kv[1]
		}
	}

	if response.StatusCode != "1" {
		return &response, fmt.Errorf("mitake: %s", response.Error)
	}

	return &response, nil
}

func (d *Driver) send(ctx context.Context, req *Request) (*Response, error) {
	request, err := d.newRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	return d.doRequest(request)
}
