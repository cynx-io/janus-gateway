package plato

import (
	"encoding/json"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "janus/api/proto/gen/go/plato/plato"
	"janus/internal/dependencies/config"
	"janus/internal/gateway/handlers"
	"janus/internal/helper"
)

type ModeHandler struct {
	client pb.PlatoModeServiceClient
}

func NewModeHandler() *ModeHandler {
	conn, err := grpc.NewClient(config.Config.Plato.Url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("Failed to connect to Plato gRPC server: " + err.Error())
	}
	client := pb.NewPlatoModeServiceClient(conn)
	return &ModeHandler{client: client}
}

func (h *ModeHandler) GetModeById(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("mode_id")

	req := pb.ModeIdRequest{}
	var err error
	req.ModeId, err = helper.StringToUint64(id)
	if err != nil {
		http.Error(w, "Invalid mode ID", http.StatusBadRequest)
		return
	}

	resp, err := h.client.GetModeById(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = handlers.HandleResponse(w, resp)
}

func (h *ModeHandler) InsertMode(w http.ResponseWriter, r *http.Request) {
	var req pb.InsertModeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.InsertMode(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = handlers.HandleResponse(w, resp)
}

func (h *ModeHandler) UpdateMode(w http.ResponseWriter, r *http.Request) {
	var req pb.UpdateModeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.UpdateMode(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = handlers.HandleResponse(w, resp)
}

func (h *ModeHandler) DeleteMode(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("mode_id")

	req := pb.ModeIdRequest{}
	var err error
	req.ModeId, err = helper.StringToUint64(id)
	if err != nil {
		http.Error(w, "Invalid mode ID", http.StatusBadRequest)
		return
	}

	resp, err := h.client.DeleteMode(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = handlers.HandleResponse(w, resp)
}

func (h *ModeHandler) ListModesByTopicId(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("topic_id")

	req := pb.TopicIdRequest{}
	var err error
	req.TopicId, err = helper.StringToUint64(id)
	if err != nil {
		http.Error(w, "Invalid topic ID", http.StatusBadRequest)
		return
	}

	resp, err := h.client.ListModesByTopicId(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = handlers.HandleResponse(w, resp)
}
