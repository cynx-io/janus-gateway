package athena

import (
	"encoding/json"
	pb "github.com/cynx-io/janus-gateway/api/proto/gen/athena"
	"github.com/cynx-io/janus-gateway/internal/dependencies/config"
	"github.com/cynx-io/janus-gateway/internal/gateway/handlers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
)

type IdeaHandler struct {
	client pb.IdeaServiceClient
}

func NewIdeaHandler() *IdeaHandler {
	conn, err := grpc.NewClient(config.Config.Athena.Url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("Failed to connect to Athena gRPC server: " + err.Error())
	}
	client := pb.NewIdeaServiceClient(conn)
	return &IdeaHandler{client: client}
}

func (h *IdeaHandler) ValidateIdea(w http.ResponseWriter, r *http.Request) {
	var req pb.ValidateIdeaRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.ValidateIdea(r.Context(), &req)
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
