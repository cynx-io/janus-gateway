package handlers

import (
	"encoding/json"
	"janus/internal/context"
	"janus/internal/dependencies/config"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "janus/api/proto/gen/hermes"
	"janus/internal/gateway/middleware"
)

type UserHandler struct {
	client pb.HermesUserServiceClient
}

func NewUserHandler() *UserHandler {
	conn, err := grpc.Dial(config.Config.Hermes.Url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("Failed to connect to Hermes gRPC server: " + err.Error())
	}

	client := pb.NewHermesUserServiceClient(conn)
	return &UserHandler{client: client}
}

func (h *UserHandler) CheckUsername(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.CheckUsername(r.Context(), &pb.CheckUsernameRequest{
		Username: req.Username,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = handleResponse(w, resp)
	if err != nil {
		http.Error(w, "Failed to handle response", http.StatusInternalServerError)
		return
	}
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {

	username, err := context.GetUsername(r.Context())
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	resp, err := h.client.GetUser(r.Context(), &pb.GetUserRequest{
		Username: username,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = handleResponse(w, resp)
	if err != nil {
		http.Error(w, "Failed to handle response", http.StatusInternalServerError)
		return
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.CreateUser(r.Context(), &pb.CreateUserRequest{
		Username: req.Username,
		Password: req.Password,
	})

	if resp.Base.Code != "00" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Generate JWT token
	token, err := middleware.GenerateToken(req.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   false, // Set to true in production
		SameSite: http.SameSiteLaxMode,
	})

	err = handleResponse(w, resp)
	if err != nil {
		http.Error(w, "Failed to handle response", http.StatusInternalServerError)
		return
	}
}

func (h *UserHandler) PaginateUsers(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Page      int    `json:"page"`
		Limit     int    `json:"limit"`
		SortBy    string `json:"sort_by"`
		SortOrder string `json:"sort_order"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.PaginateUsers(r.Context(), &pb.PaginateRequest{
		Page:      int32(req.Page),
		Limit:     int32(req.Limit),
		SortBy:    req.SortBy,
		SortOrder: req.SortOrder,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = handleResponse(w, resp)
	if err != nil {
		http.Error(w, "Failed to handle response", http.StatusInternalServerError)
		return
	}
}

func (h *UserHandler) ValidatePassword(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.ValidatePassword(r.Context(), &pb.ValidatePasswordRequest{
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if resp.Base.Code == "200" {
		// Generate JWT token
		token, err := middleware.GenerateToken(req.Username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Set cookie
		http.SetCookie(w, &http.Cookie{
			Name:     "token",
			Value:    token,
			Path:     "/",
			HttpOnly: true,
			Secure:   false, // Set to true in production
			SameSite: http.SameSiteLaxMode,
		})
	}

	err = handleResponse(w, resp)
	if err != nil {
		http.Error(w, "Failed to handle response", http.StatusInternalServerError)
		return
	}
}
