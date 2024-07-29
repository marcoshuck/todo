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
	assert.True(t, now.Equal(out))
}

func TestDecodePageToken_InvalidToken(t *testing.T) {
	out, err := DecodePageToken("1234")
	assert.Error(t, err)
	assert.Zero(t, out)
}
