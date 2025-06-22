package plato

import (
	"encoding/json"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/cynxees/janus-gateway/api/proto/gen/plato"
	"github.com/cynxees/janus-gateway/internal/dependencies/config"
	"github.com/cynxees/janus-gateway/internal/gateway/handlers"
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

	req := pb.ModeIdRequest{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.GetModeById(r.Context(), &req)
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
	err = handlers.HandleResponse(w, resp)
	if err != nil {
		http.Error(w, "Failed to handle response", http.StatusInternalServerError)
		return
	}
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
	err = handlers.HandleResponse(w, resp)
	if err != nil {
		http.Error(w, "Failed to handle response", http.StatusInternalServerError)
		return
	}
}

func (h *ModeHandler) DeleteMode(w http.ResponseWriter, r *http.Request) {

	req := pb.ModeIdRequest{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.DeleteMode(r.Context(), &req)
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

func (h *ModeHandler) ListModesByTopicId(w http.ResponseWriter, r *http.Request) {

	req := pb.TopicIdRequest{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.ListModesByTopicId(r.Context(), &req)
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
