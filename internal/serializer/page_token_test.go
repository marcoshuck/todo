package serializer

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestEncodePageToken(t *testing.T) {
	now := time.Now()
	token := EncodePageToken(now)
	assert.NotNil(t, token)
	var expected string
	assert.IsType(t, expected, token)
	assert.NotEmpty(t, token)
}

func TestDecodePageToken(t *testing.T) {
	now := time.Now()
	token := EncodePageToken(now)

	out, err := DecodePageToken(token)
	assert.NoError(t, err)
	assert.Equal(t, now.Year(), out.Year())
	assert.Equal(t, now.Month(), out.Month())
	assert.Equal(t, now.Day(), out.Day())
	assert.Equal(t, now.Hour(), out.Hour())
	assert.Equal(t, now.Minute(), out.Minute())
	assert.Equal(t, now.Second(), out.Second())
}

func TestDecodePageToken_InvalidToken(t *testing.T) {
	out, err := DecodePageToken("1234")
	assert.Error(t, err)
	assert.Zero(t, out)
}
