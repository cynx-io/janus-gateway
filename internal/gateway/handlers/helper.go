package handlers

import (
	"github.com/cynxees/cynx-core/src/logger"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"net/http"
)

func HandleResponse(w http.ResponseWriter, resp proto.Message) error {

	marshaler := protojson.MarshalOptions{
		EmitUnpopulated: true,
		UseProtoNames:   true,
	}

	data, err := marshaler.Marshal(resp)
	if err != nil {
		logger.Error("Failed to marshal response: ", err)
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(data)
	if err != nil {
		logger.Error("Failed to write response: ", err)
	}
	return err
}
