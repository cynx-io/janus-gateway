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

type DailyGameHandler struct {
	client pb.PlatoDailyGameServiceClient
}

func NewDailyGameHandler() *DailyGameHandler {
	conn, err := grpc.NewClient(config.Config.Plato.Url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("Failed to connect to Plato gRPC server: " + err.Error())
	}

	client := pb.NewPlatoDailyGameServiceClient(conn)
	return &DailyGameHandler{client: client}
}

func (h *DailyGameHandler) GetDetailDailyGameById(w http.ResponseWriter, r *http.Request) {

	req := pb.DailyGameIdRequest{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.GetDetailDailyGameById(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = handlers.HandleResponse(w, resp)
}

func (h *DailyGameHandler) GetModeDailyGameById(w http.ResponseWriter, r *http.Request) {

	req := pb.ModeIdRequest{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.GetModeDailyGameById(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = handlers.HandleResponse(w, resp)
}

func (h *DailyGameHandler) GetPublicDailyGame(w http.ResponseWriter, r *http.Request) {

	req := pb.ModeIdRequest{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.GetPublicDailyGame(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = handlers.HandleResponse(w, resp)
}

func (h *DailyGameHandler) AttemptAnswer(w http.ResponseWriter, r *http.Request) {
	var req pb.AttemptAnswerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.AttemptAnswer(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = handlers.HandleResponse(w, resp)
}
