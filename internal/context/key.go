package context

type Key string

const (
	KeyRequestId     Key = "request_id"
	KeyRequestOrigin Key = "request_origin"
	KeyRequestPath   Key = "request_path"

	KeyUsername Key = "username"

	KeyUserId      Key = "user_id"      // int32
	KeyTimestamp   Key = "timestamp"    // time.Time
	KeyBaseRequest Key = "base_request" // *pb.BaseRequest (protobuf message for base request info

)
