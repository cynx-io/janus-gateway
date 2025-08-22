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

type CareerProfileHandler struct {
	client pb.CareerProfileServiceClient
}

func NewCareerProfileHandler() *CareerProfileHandler {
	conn, err := grpc.NewClient(config.Config.Philyra.Url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("Failed to connect to Philyra gRPC server: " + err.Error())
	}
	client := pb.NewCareerProfileServiceClient(conn)
	return &CareerProfileHandler{client: client}
}

func (h *CareerProfileHandler) GetCareerProfile(w http.ResponseWriter, r *http.Request) {
	req := pb.GetCareerProfileRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.GetCareerProfile(r.Context(), &req)
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

func (h *CareerProfileHandler) SyncCareerProfile(w http.ResponseWriter, r *http.Request) {
	var req pb.SyncCareerProfileRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.SyncCareerProfile(r.Context(), &req)
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

func (h *CareerProfileHandler) UpdatePersonalInfo(w http.ResponseWriter, r *http.Request) {
	var req pb.UpdateCareerPersonalInfoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.UpdatePersonalInfo(r.Context(), &req)
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

func (h *CareerProfileHandler) UpdateProfessionalInfo(w http.ResponseWriter, r *http.Request) {
	var req pb.UpdateCareerProfessionalInfoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.UpdateProfessionalInfo(r.Context(), &req)
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

func (h *CareerProfileHandler) UpdateJobPreferences(w http.ResponseWriter, r *http.Request) {
	var req pb.UpdateCareerJobPreferencesRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.UpdateJobPreferences(r.Context(), &req)
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

func (h *CareerProfileHandler) UpdateUserDocuments(w http.ResponseWriter, r *http.Request) {
	var req pb.UpdateCareerDocumentsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.UpdateUserDocuments(r.Context(), &req)
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

func (h *CareerProfileHandler) UpdateCustomResponses(w http.ResponseWriter, r *http.Request) {
	var req pb.UpdateCareerCustomResponsesRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.UpdateCustomResponses(r.Context(), &req)
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
