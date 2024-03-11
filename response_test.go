package sms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponse_UnimplementedResponseExtra(t *testing.T) {
	r := UnimplementedResponseExtra{}
	r.IsResponseExtra()
	assert.IsType(t, UnimplementedResponseExtra{}, r)
}

func TestResponse(t *testing.T) {
	resp := Response{
		Status:  Success,
		Message: "test",
		Extra:   nil,
	}
	assert.Equal(t, Success, resp.Status)
	assert.Equal(t, "test", resp.Message)
	assert.Nil(t, resp.Extra)
}
