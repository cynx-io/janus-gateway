package core

import pb "github.com/cynxees/janus-gateway/api/proto/gen/core"

func InitBaseRequest() *pb.BaseRequest {

	userId := int32(1)
	return &pb.BaseRequest{
		RequestId:     "1231",
		RequestOrigin: "1",
		UserId:        &userId,
	}

}
