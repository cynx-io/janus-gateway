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

type AutoFillHandler struct {
	client pb.AutoFillServiceClient
}

func NewAutoFillHandler() *AutoFillHandler {
	conn, err := grpc.NewClient(config.Config.Philyra.Url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("Failed to connect to Philyra gRPC server: " + err.Error())
	}
	client := pb.NewAutoFillServiceClient(conn)
	return &AutoFillHandler{client: client}
}

func (h *AutoFillHandler) AnalyzeForm(w http.ResponseWriter, r *http.Request) {
	var req pb.AnalyzeFormRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.AnalyzeForm(r.Context(), &req)
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
