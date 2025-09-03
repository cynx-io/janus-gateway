package ananke

import (
	"encoding/json"
	pb "github.com/cynx-io/janus-gateway/api/proto/gen/ananke"
	"github.com/cynx-io/janus-gateway/internal/dependencies/config"
	"github.com/cynx-io/janus-gateway/internal/gateway/handlers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
)

type WaitlistHandler struct {
	client pb.WaitlistServiceClient
}

func NewWaitlistHandler() *WaitlistHandler {
	conn, err := grpc.NewClient(config.Config.Ananke.Url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("Failed to connect to Ananke gRPC server: " + err.Error())
	}
	client := pb.NewWaitlistServiceClient(conn)
	return &WaitlistHandler{client: client}
}

func (h *WaitlistHandler) RegisterWaitlist(w http.ResponseWriter, r *http.Request) {
	var req pb.RegisterWaitlistRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.RegisterWaitlist(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = handlers.HandleResponse(w, resp)
	if err != nil {
		http.Error(w, "Failed to handle response", http.StatusInternalServerError)
		return
	}
}
