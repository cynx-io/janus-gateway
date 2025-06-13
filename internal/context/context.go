package context

import (
	"context"
	"errors"
	"google.golang.org/protobuf/types/known/timestamppb"
	pb "janus/api/proto/gen/go/core/api/proto"
	"time"
)

func GetBaseRequest(ctx context.Context) *pb.BaseRequest {

	timestamp := ctx.Value(KeyTimestamp).(time.Time)

	req := &pb.BaseRequest{
		RequestId:     GetKeyOrEmpty(ctx, KeyRequestId),
		RequestOrigin: GetKeyOrEmpty(ctx, KeyRequestOrigin),
		RequestPath:   GetKeyOrEmpty(ctx, KeyRequestPath),
		Timestamp:     timestamppb.New(timestamp),
		UserId:        GetUserId(ctx),
		Username:      GetKey(ctx, KeyUsername),
	}
	return req
}

func SetBaseRequest(ctx context.Context, req *pb.BaseRequest) (context.Context, error) {
	if req == nil {
		return ctx, errors.New("baseRequest cannot be nil")
	}

	if req.UserId != nil {
		ctx = SetUserId(ctx, *req.UserId)
	}

	if req.RequestId != "" {
		ctx = SetKey(ctx, KeyRequestId, req.RequestId)
	}

	if req.RequestOrigin != "" {
		ctx = SetKey(ctx, KeyRequestOrigin, req.RequestOrigin)
	}

	if req.RequestPath != "" {
		ctx = SetKey(ctx, KeyRequestPath, req.RequestPath)
	}

	if req.Username != nil {
		ctx = SetKey(ctx, KeyUsername, *req.Username)
	}

	if req.Timestamp != nil {
		timestamp := req.Timestamp.AsTime()
		if !timestamp.IsZero() {
			ctx = context.WithValue(ctx, KeyTimestamp, timestamp)
		}
	} else {
		ctx = context.WithValue(ctx, KeyTimestamp, time.Now())
	}

	return ctx, nil
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
