package serializer

import (
	"encoding/base64"
	"time"
)

// EncodePageToken converts the given timestamp into a base64-encoded page token.
func EncodePageToken(t time.Time) string {
	return base64.URLEncoding.EncodeToString([]byte(t.Format(time.RFC3339)))
}

// DecodePageToken converts the given token into a valid timestamp.
// It returns an error if the token is not a valid representation of a base64-encoded
// timestamp.
func DecodePageToken(token string) (time.Time, error) {
	b, err := base64.URLEncoding.DecodeString(token)
	if err != nil {
		return time.Time{}, err
	}
	return time.Parse(time.RFC3339, string(b))
}
