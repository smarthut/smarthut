package api

type contextKey string

func (c contextKey) String() string {
	return "smarthut api context key " + string(c)
}

const (
	deviceKey = contextKey("device")
	userKey   = contextKey("user")
	bucketKey = contextKey("bucket")
)
