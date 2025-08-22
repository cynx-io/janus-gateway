package philyra

import (
	"encoding/json"
	"log"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/cynx-io/janus-gateway/api/proto/gen/philyra"
	"github.com/cynx-io/janus-gateway/internal/dependencies/config"
	"github.com/cynx-io/janus-gateway/internal/gateway/handlers"
)

type ResumeHandler struct {
	client pb.ResumeServiceClient
}

func NewResumeHandler() *ResumeHandler {
	conn, err := grpc.NewClient(config.Config.Philyra.Url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("Failed to connect to Philyra gRPC server: " + err.Error())
	}
	client := pb.NewResumeServiceClient(conn)
	return &ResumeHandler{client: client}
}

func (h *ResumeHandler) CreateResume(w http.ResponseWriter, r *http.Request) {
	var req pb.CreateResumeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.CreateResume(r.Context(), &req)
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

func (h *ResumeHandler) GetResume(w http.ResponseWriter, r *http.Request) {
	req := pb.GetResumeRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Debug logging
	log.Printf("[DEBUG] GetResume - Base request: %+v", req.Base)
	if req.Base != nil {
		log.Printf("[DEBUG] GetResume - Request ID: %s, User ID: %d", req.Base.RequestId, req.Base.UserId)
	} else {
		log.Printf("[DEBUG] GetResume - Base request is nil!")
	}

	resp, err := h.client.GetResume(r.Context(), &req)
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

func (h *ResumeHandler) UpdateResume(w http.ResponseWriter, r *http.Request) {
	var req pb.UpdateResumeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.UpdateResume(r.Context(), &req)
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

func (h *ResumeHandler) DeleteResume(w http.ResponseWriter, r *http.Request) {

	req := pb.DeleteResumeRequest{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.DeleteResume(r.Context(), &req)
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

func (h *ResumeHandler) ListResumes(w http.ResponseWriter, r *http.Request) {
	req := pb.ListResumesRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.ListResumes(r.Context(), &req)
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

func (h *ResumeHandler) GenerateResume(w http.ResponseWriter, r *http.Request) {
	req := pb.GenerateResumeRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.GenerateResume(r.Context(), &req)
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
