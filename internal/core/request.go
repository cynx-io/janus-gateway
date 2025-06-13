package core

import pb "janus/api/proto/gen/go/core/api/proto"

func InitBaseRequest() *pb.BaseRequest {

	userId := uint64(1)
	return &pb.BaseRequest{
		RequestId:     "1231",
		RequestOrigin: "1",
		UserId:        &userId,
	}

}
