package mitake

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/flc1125/go-sms"
)

var ctx = context.Background()

func TestNew(t *testing.T) {
	driver := New(&Config{
		Addr:     "http://localhost",
		Username: "flc",
		Password: "123456",
	}, WithHTTPClient(http.DefaultClient))

	assert.IsType(t, &Driver{}, driver)
	assert.IsType(t, &Config{}, driver.config)
	assert.IsType(t, &http.Client{}, driver.httpClient)
	assert.Equal(t, "http://localhost", driver.config.Addr)
	assert.Equal(t, "flc", driver.config.Username)
	assert.Equal(t, "123456", driver.config.Password)
}

func TestParseRequest(t *testing.T) {
	driver := &Driver{}

	r1 := driver.parseRequest(&sms.Request{
		Phone:   "1234567890",
		Content: "test",
	})
	assert.Equal(t, "1234567890", r1.Dstaddr)
	assert.Equal(t, "test", r1.Smbody)
	assert.Equal(t, "now", r1.Type)
	assert.Equal(t, "big5", r1.Encoding)

	r2 := driver.parseRequest(&sms.Request{
		Phone:   "1234567890",
		Content: "test",
		Extra:   &Request{},
	})
	assert.Equal(t, "1234567890", r2.Dstaddr)
	assert.Equal(t, "test", r2.Smbody)
	assert.Equal(t, "now", r2.Type)
	assert.Equal(t, "big5", r2.Encoding)
}

func TestSend_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "flc", r.URL.Query().Get("username"))
		assert.Equal(t, "123456", r.URL.Query().Get("password"))
		assert.Equal(t, "1234567890", r.URL.Query().Get("dstaddr"))
		assert.Equal(t, "test", r.URL.Query().Get("smbody"))

		_, _ = fmt.Fprintln(w, "[1]\nmsgid=123123123\nstatuscode=1\nAccountPoint=456789")
	}))
	defer srv.Close()

	driver := New(&Config{
		Addr:     srv.URL,
		Username: "flc",
		Password: "123456",
	})

	resp, err := driver.Send(ctx, &sms.Request{
		Phone:   "1234567890",
		Content: "test",
	})
	assert.NoError(t, err)
	assert.Equal(t, sms.Success, resp.Status)
	assert.Equal(t, "success", resp.Message)
	assert.IsType(t, &Response{}, resp.Extra)
	assert.Equal(t, "123123123", resp.Extra.(*Response).MsgID)
	assert.Equal(t, "1", resp.Extra.(*Response).StatusCode)
	assert.Equal(t, "456789", resp.Extra.(*Response).AccountPoint)
	assert.Empty(t, resp.Extra.(*Response).Error)
}

func TestSend_Fail(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, "[1]\nstatuscode=k\nError=custom error")
	}))
	defer srv.Close()

	driver := New(&Config{
		Addr:     srv.URL,
		Username: "flc",
		Password: "123456",
	})

	resp, err := driver.Send(ctx, &sms.Request{
		Phone:   "1234567890",
		Content: "test",
	})
	assert.Error(t, err)
	assert.Equal(t, sms.Failed, resp.Status)
	assert.Equal(t, "mitake: custom error", resp.Message)
	assert.IsType(t, &Response{}, resp.Extra)
	assert.Equal(t, "k", resp.Extra.(*Response).StatusCode)
	assert.Equal(t, "custom error", resp.Extra.(*Response).Error)
	assert.Empty(t, resp.Extra.(*Response).MsgID)
	assert.Empty(t, resp.Extra.(*Response).AccountPoint)
}
