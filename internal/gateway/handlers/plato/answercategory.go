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

type AnswerCategoryHandler struct {
	client pb.PlatoAnswerCategoryServiceClient
}

func NewAnswerCategoryHandler() *AnswerCategoryHandler {
	conn, err := grpc.NewClient(config.Config.Plato.Url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("Failed to connect to Plato gRPC server: " + err.Error())
	}

	client := pb.NewPlatoAnswerCategoryServiceClient(conn)
	return &AnswerCategoryHandler{client: client}
}

func (h *AnswerCategoryHandler) GetAnswerCategoryById(w http.ResponseWriter, r *http.Request) {
	req := pb.AnswerCategoryIdRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.GetAnswerCategoryById(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = handlers.HandleResponse(w, resp)
}

func (h *AnswerCategoryHandler) ListAnswerCategoriesByAnswerId(w http.ResponseWriter, r *http.Request) {
	req := pb.AnswerIdRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.ListAnswerCategoriesByAnswerId(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = handlers.HandleResponse(w, resp)
}

func (h *AnswerCategoryHandler) InsertAnswerCategory(w http.ResponseWriter, r *http.Request) {
	var req pb.InsertAnswerCategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.InsertAnswerCategory(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = handlers.HandleResponse(w, resp)
}

func (h *AnswerCategoryHandler) UpdateAnswerCategory(w http.ResponseWriter, r *http.Request) {
	var req pb.UpdateAnswerCategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.UpdateAnswerCategory(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = handlers.HandleResponse(w, resp)
}

func (h *AnswerCategoryHandler) DeleteAnswerCategory(w http.ResponseWriter, r *http.Request) {
	req := pb.AnswerCategoryIdRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.DeleteAnswerCategory(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = handlers.HandleResponse(w, resp)
}
