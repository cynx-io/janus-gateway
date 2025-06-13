package handlers

import (
	"encoding/json"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "janus/api/proto/gen/mercury"
	"janus/internal/dependencies/config"
	"net/http"
)

type CryptoHandler struct {
	client pb.MercuryCryptoServiceClient
}

func NewCryptoHandler() *CryptoHandler {
	conn, err := grpc.Dial(config.Config.Mercury.Url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("Failed to connect to Mercury gRPC server: " + err.Error())
	}

	client := pb.NewMercuryCryptoServiceClient(conn)
	return &CryptoHandler{client: client}
}

func (h *CryptoHandler) SearchCoin(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Query string `json:"query"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.SearchCoin(r.Context(), &pb.SearchCoinRequest{
		Query: req.Query,
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

func (h *CryptoHandler) GetCoinRisk(w http.ResponseWriter, r *http.Request) {
	var req struct {
		CoinId string `json:"coin_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.GetCoinRisk(r.Context(), &pb.GetCoinRiskRequest{
		CoinId: req.CoinId,
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
