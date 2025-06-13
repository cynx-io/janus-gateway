package context

type Key string

const (
	KeyRequestId     Key = "request_id"
	KeyRequestOrigin Key = "request_origin"
	KeyRequestPath   Key = "request_path"

	KeyUsername Key = "username"

	KeyUserId    Key = "user_id"   // uint64
	KeyTimestamp Key = "timestamp" // time.Time

)
