package sms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequest_UnimplementedRequestExtra(t *testing.T) {
	r := UnimplementedRequestExtra{}
	r.IsRequestExtra()
	assert.IsType(t, UnimplementedRequestExtra{}, r)
}
