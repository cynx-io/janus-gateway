package hermes

import (
	"encoding/json"
	"github.com/cynxees/cynx-core/src/context"
	"github.com/cynxees/janus-gateway/internal/dependencies/config"
	"github.com/cynxees/janus-gateway/internal/gateway/handlers"
	"net/http"

	pb "github.com/cynxees/janus-gateway/api/proto/gen/hermes"
	"github.com/cynxees/janus-gateway/internal/gateway/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserHandler struct {
	client pb.HermesUserServiceClient
}

func NewUserHandler() *UserHandler {
	conn, err := grpc.NewClient(config.Config.Hermes.Url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("Failed to connect to Hermes gRPC server: " + err.Error())
	}

	client := pb.NewHermesUserServiceClient(conn)
	return &UserHandler{client: client}
}

func (h *UserHandler) CheckUsername(w http.ResponseWriter, r *http.Request) {
	var req pb.CheckUsernameRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.CheckUsername(r.Context(), &req)
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

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {

	username := context.GetKey(r.Context(), context.KeyUsername)
	if username == nil {
		http.Error(w, "Username not provided", http.StatusBadRequest)
		return
	}

	resp, err := h.client.GetUser(r.Context(), &pb.GetUserRequest{
		Username: *username,
	})

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
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		}
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Generate JWT token
	token, err := middleware.GenerateToken(req.Username, resp.User.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set cookie
	http.SetCookie(w, &http.Cookie{
		Name:     config.Config.Cookie.Name,
		Value:    token,
		Path:     config.Config.Cookie.Path,
		HttpOnly: config.Config.Cookie.HttpOnly,
		Domain:   config.Config.Cookie.Domain,
		Secure:   config.Config.Cookie.Secure,
		SameSite: http.SameSiteNoneMode,
	})

	err = handlers.HandleResponse(w, resp)
	if err != nil {
		http.Error(w, "Failed to handle response", http.StatusInternalServerError)
		return
	}
}

func (h *UserHandler) PaginateUsers(w http.ResponseWriter, r *http.Request) {
	var req struct {
		SortBy    string `json:"sort_by"`
		SortOrder string `json:"sort_order"`
		Page      int    `json:"page"`
		Limit     int    `json:"limit"`
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

	err = handlers.HandleResponse(w, resp)
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

	if resp.Base.Code == "00" {
		// Generate JWT token
		token, err := middleware.GenerateToken(req.Username, resp.User.Id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Set cookie
		http.SetCookie(w, &http.Cookie{
			Name:     config.Config.Cookie.Name,
			Value:    token,
			Path:     config.Config.Cookie.Path,
			HttpOnly: config.Config.Cookie.HttpOnly,
			Domain:   config.Config.Cookie.Domain,
			Secure:   config.Config.Cookie.Secure,
			SameSite: http.SameSiteNoneMode,
		})
	}

	err = handlers.HandleResponse(w, resp)
	if err != nil {
		http.Error(w, "Failed to handle response", http.StatusInternalServerError)
		return
	}
}
