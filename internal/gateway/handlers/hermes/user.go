package hermes

import (
	"encoding/json"
	core "github.com/cynx-io/cynx-core/proto/gen"
	"github.com/cynx-io/cynx-core/src/context"
	"github.com/cynx-io/cynx-core/src/types/usertype"
	pb "github.com/cynx-io/janus-gateway/api/proto/gen/hermes"
	"github.com/cynx-io/janus-gateway/internal/dependencies/config"
	"github.com/cynx-io/janus-gateway/internal/gateway/handlers"
	"github.com/cynx-io/janus-gateway/internal/gateway/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
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
	var req pb.UsernameRequest

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

	resp, err := h.client.GetUser(r.Context(), &pb.UsernameRequest{
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
	req := pb.UsernamePasswordRequest{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.CreateUser(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if resp == nil || resp.Base == nil {
		http.Error(w, "Invalid response from server", http.StatusInternalServerError)
		return
	}

	if resp.Base.Code != "00" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		}
		return
	}

	// Generate JWT token
	claims := &middleware.Claims{
		Username: req.Username,
		UserId:   resp.User.Id,
		UserType: usertype.UserType(resp.User.UserType),
	}
	token, err := middleware.GenerateToken(claims)
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
		Expires:  claims.ExpiresAt.Time,
	})

	err = handlers.HandleResponse(w, resp)
	if err != nil {
		http.Error(w, "Failed to handle response", http.StatusInternalServerError)
		return
	}
}

func (h *UserHandler) CreateUserFromGuest(w http.ResponseWriter, r *http.Request) {
	req := pb.UsernamePasswordRequest{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.CreateUserFromGuest(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if resp == nil || resp.Base == nil {
		http.Error(w, "Invalid response from server", http.StatusInternalServerError)
		return
	}

	if resp.Base.Code != "00" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		}
		return
	}

	// Generate JWT token
	claims := &middleware.Claims{
		Username: req.Username,
		UserId:   resp.User.Id,
		UserType: usertype.UserType(resp.User.UserType),
	}
	token, err := middleware.GenerateToken(claims)
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
		Expires:  claims.ExpiresAt.Time,
	})

	err = handlers.HandleResponse(w, resp)
	if err != nil {
		http.Error(w, "Failed to handle response", http.StatusInternalServerError)
		return
	}
}

func (h *UserHandler) UpsertGuestUser(w http.ResponseWriter, r *http.Request) {
	req := core.GenericRequest{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.UpsertGuestUser(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if resp == nil || resp.Base == nil {
		http.Error(w, "Invalid response from server", http.StatusInternalServerError)
		return
	}

	if resp.Base.Code != "00" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		}
		return
	}

	// Generate JWT token
	claims := &middleware.Claims{
		Username: resp.User.Username,
		UserId:   resp.User.Id,
		UserType: usertype.UserType(resp.User.UserType),
	}
	token, err := middleware.GenerateToken(claims)
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
		Expires:  claims.ExpiresAt.Time,
	})

	err = handlers.HandleResponse(w, resp)
	if err != nil {
		http.Error(w, "Failed to handle response", http.StatusInternalServerError)
		return
	}
}

func (h *UserHandler) PaginateUsers(w http.ResponseWriter, r *http.Request) {
	req := pb.PaginateRequest{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.PaginateUsers(r.Context(), &req)

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
	req := pb.UsernamePasswordRequest{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.ValidatePassword(r.Context(), &req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if resp.Base.Code == "00" {
		// Generate JWT token
		claims := &middleware.Claims{
			Username: req.Username,
			UserId:   resp.User.Id,
			UserType: usertype.UserType(resp.User.UserType),
		}
		token, err := middleware.GenerateToken(claims)
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
			Expires:  claims.ExpiresAt.Time,
		})
	}

	err = handlers.HandleResponse(w, resp)
	if err != nil {
		http.Error(w, "Failed to handle response", http.StatusInternalServerError)
		return
	}
}
