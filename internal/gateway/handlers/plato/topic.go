package plato

import (
	"encoding/json"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "janus/api/proto/gen/plato/plato"
	"janus/internal/dependencies/config"
	"janus/internal/gateway/handlers"
	"janus/internal/helper"
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
	_ = handlers.HandleResponse(w, resp)
}

func (h *TopicHandler) GetTopicById(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("topic_id")

	req := pb.TopicIdRequest{}
	var err error
	req.TopicId, err = helper.StringToUint64(id)
	if err != nil {
		http.Error(w, "Invalid topic ID", http.StatusBadRequest)
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
	slug := r.URL.Query().Get("slug")
	if slug == "" {
		http.Error(w, "Slug is required", http.StatusBadRequest)
		return
	}

	req := pb.SlugRequest{Slug: slug}
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
	id := r.URL.Query().Get("topic_id")

	req := pb.TopicIdRequest{}
	var err error
	req.TopicId, err = helper.StringToUint64(id)
	if err != nil {
		http.Error(w, "Invalid topic ID", http.StatusBadRequest)
		return
	}

	resp, err := h.client.DeleteTopic(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = handlers.HandleResponse(w, resp)
}
