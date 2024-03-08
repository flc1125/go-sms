package sms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatus(t *testing.T) {
	assert.Equal(t, Success, Status(0))
	assert.Equal(t, Failed, Status(-1))
	assert.Equal(t, 0, Success.Int())
	assert.Equal(t, -1, Failed.Int())
}
