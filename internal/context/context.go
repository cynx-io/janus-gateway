package context

import (
	"context"
	pb "github.com/cynxees/janus-gateway/api/proto/gen/core"
)

func GetBaseRequest(ctx context.Context) *pb.BaseRequest {
	val := ctx.Value(KeyBaseRequest)
	if val == nil {
		return nil
	}
	if req, ok := val.(*pb.BaseRequest); ok {
		return req
	}
	return nil
}

func SetBaseRequest(ctx context.Context, req *pb.BaseRequest) (context.Context, error) {
	return context.WithValue(ctx, KeyBaseRequest, req), nil
}

func SetKey(ctx context.Context, key Key, value string) context.Context {
	if value == "" {
		return ctx
	}
	return context.WithValue(ctx, key, value)
}

func SetUserId(ctx context.Context, userId uint64) context.Context {
	return context.WithValue(ctx, KeyUserId, userId)
}

func GetKeyOrEmpty(ctx context.Context, key Key) string {
	value := ctx.Value(key)
	if value == nil {
		return ""
	}

	strValue, ok := value.(string)
	if !ok {
		return ""
	}

	return strValue
}

func GetKey(ctx context.Context, key Key) *string {
	value := ctx.Value(key)
	if value == nil {
		return nil
	}

	strValue, ok := value.(string)
	if !ok {
		return nil
	}

	return &strValue
}

func GetUserId(ctx context.Context) *uint64 {
	value := ctx.Value(KeyUserId)
	if value == nil {
		return nil
	}

	userId, ok := value.(uint64)
	if !ok {
		return nil
	}
	return &userId
}
