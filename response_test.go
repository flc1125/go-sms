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
