package plato

import (
	"encoding/json"
	"log"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pbCore "github.com/cynxees/janus-gateway/api/proto/gen/core"
	pb "github.com/cynxees/janus-gateway/api/proto/gen/plato"
	"github.com/cynxees/janus-gateway/internal/dependencies/config"
	"github.com/cynxees/janus-gateway/internal/gateway/handlers"
)

type TopicHandler struct {
	client pb.PlatoTopicServiceClient
}

func NewTopicHandler() *TopicHandler {
	conn, err := grpc.NewClient(config.Config.Plato.Url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("Failed to connect to Plato gRPC server: " + err.Error())
	}
	client := pb.NewPlatoTopicServiceClient(conn)
	return &TopicHandler{client: client}
}

func (h *TopicHandler) PaginateTopic(w http.ResponseWriter, r *http.Request) {
	var req pb.PaginateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.PaginateTopic(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = handlers.HandleResponse(w, resp)
	if err != nil {
		log.Printf("Failed to handle response: %v", err)
		return
	}
}

func (h *TopicHandler) GetTopicById(w http.ResponseWriter, r *http.Request) {
	req := pb.TopicIdRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.GetTopicById(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = handlers.HandleResponse(w, resp)
}

func (h *TopicHandler) GetTopicBySlug(w http.ResponseWriter, r *http.Request) {
	req := pb.SlugRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.GetTopicBySlug(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = handlers.HandleResponse(w, resp)
}

func (h *TopicHandler) InsertTopic(w http.ResponseWriter, r *http.Request) {
	var req pb.InsertTopicRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.InsertTopic(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = handlers.HandleResponse(w, resp)
}

func (h *TopicHandler) UpdateTopic(w http.ResponseWriter, r *http.Request) {
	var req pb.UpdateTopicRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.UpdateTopic(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = handlers.HandleResponse(w, resp)
}

func (h *TopicHandler) DeleteTopic(w http.ResponseWriter, r *http.Request) {

	req := pb.TopicIdRequest{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.DeleteTopic(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = handlers.HandleResponse(w, resp)
}

func (h *TopicHandler) ListTopicsByUserId(w http.ResponseWriter, r *http.Request) {
	req := pbCore.GenericRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.ListTopicsByUserId(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = handlers.HandleResponse(w, resp)
}
